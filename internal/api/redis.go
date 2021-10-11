package api

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

// https://www.cnblogs.com/caoshouling/p/12219037.html
const (
	// 设置键的过期时间为 second 秒。SET key value EX second 效果等同于 SETEX key second value 。
	EXSECONDS = "EX"
	// 设置键的过期时间为 millisecond 毫秒。 SET key value PX millisecond 效果等同于 PSETEX key millisecond value。
	PXMILLISSECONDS = "PX"

	// 只在键不存在时，才对键进行设置操作。 SET key value NX 效果等同于 SETNX key value 。
	NOTEXISTS = "NX"
	// 只在键已经存在时，才对键进行设置操作。
	EXISTS = "XX"

	/*
		注意: 由于SET命令加上选项已经可以完全取代SETNX, SETEX, PSETEX的功能，所以在将来的版本中，redis可能会不推荐使用并且最终抛弃这几个命令。
		使用SET代替SETNX，相当于SETNX+EXPIRE实现了原子性，不必担心SETNX成功，EXPIRE失败的问题！有效的避免死锁，解决了Redis2.6.12之前版本存在的问题。
	*/

)

type Redis struct {
	redisClent redis.Conn
}

var RedisClent Redis

func InitRedis() {
	redisClent, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	RedisClent.redisClent = redisClent
	fmt.Println("redis connection success")
}

/*
	key
*/

func (r *Redis) Del(keys ...interface{}) (int64, error) {
	return redis.Int64(r.redisClent.Do("DEL", keys...))
}

/*
	lock
*/
func (r *Redis) Setnx(key string, value interface{}, expire uint64) (bool, error) {
	_, err := redis.String(r.redisClent.Do("SET", key, value, EXSECONDS, expire, NOTEXISTS))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *Redis) Unlock(key string, value interface{}) (flag bool, err error) {
	script := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return -1 end"
	reply, err := r.Lua(script, 1, key, value)
	if err != nil {
		return false, err
	}
	if reply.(int64) == -1 {
		return false, nil
	}
	return true, nil
}

/*
	list
*/
func (r *Redis) LPush(key string, value string) (int, error) {
	return redis.Int(r.redisClent.Do("LPUSH", key, value))
}
func (r *Redis) RPush(key string, value string) (int, error) {
	return redis.Int(r.redisClent.Do("RPUSH", key, value))
}
func (r *Redis) LPop() {}
func (r *Redis) RPop() {}
func (r *Redis) BLPop(key string, timeout int64) (err error) {
	slices, err := redis.ByteSlices(r.redisClent.Do("BRPOP", key, timeout))
	if err == redis.ErrNil {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println(string(slices[0]), string(slices[1]))
	return nil
}
func (r *Redis) BRPop(key string, timeout int64) (slices [][]byte, err error) {
	slices, err = redis.ByteSlices(r.redisClent.Do("BLPOP", key, timeout))
	if err == redis.ErrNil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	fmt.Println(string(slices[0]), string(slices[1]))
	return slices, nil
}

/*
	zset
*/

func (r *Redis) ZAdd(key string, maps map[string]int64) (int64, error) {
	args := []interface{}{key}
	for member, score := range maps {
		args = append(args, score, member)
	}
	return redis.Int64(r.redisClent.Do("ZADD", args...))
}

// 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。
// 有序集成员按 score 值递增(从小到大)次序排列。
// withscores指定是否返回得分，
// limit 是否分页方法，false返回所有的数据
func (r *Redis) ZRangeByScore(key string, min int, max int, withscores, limit bool, offset int, count int) (slices [][]byte, err error) {
	args := []interface{}{key, min, max}
	if withscores {
		args = append(args, "WITHSCORES")
	}
	if limit {
		args = append(args, "LIMIT", offset, count)
	}
	slices, err = redis.ByteSlices(r.redisClent.Do("ZRANGEBYSCORE", args...))
	if err == redis.ErrNil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	for _, slice := range slices {
		fmt.Println(string(slice))
	}
	return slices, nil
}

// 移除有序集 key 中的一个或多个成员，不存在的成员将被忽略
// return: 被成功移除的成员的数量，不包括被忽略的成员
func (r *Redis) ZRem(key string, members ...string) (num int64, err error) {
	args := packArgs(key, members)
	num, err = redis.Int64(r.redisClent.Do("ZREM", args...))
	if err == redis.ErrNil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return num, nil
}
func (r *Redis) ZRangeByScoreAndZRem(key string, min int64, max int64, withscores, limit bool, offset int, count int) (slices [][]byte, err error) {

	script := "local tbl = redis.call('ZRANGEBYSCORE', KEYS[1], KEYS[2], KEYS[3], 'limit', '0', '1') " +
		"if #tbl > 0 then return redis.call('ZREM', KEYS[1], tbl[1]) else return -1 end"
	reply, err := r.Lua(script, 3, key, min, max)
	fmt.Println(reply)
	return slices, nil
}

/*
	lua
*/
func (r *Redis) Lua(script string, keyCount int, keysAndArgs ...interface{}) (interface{}, error) {
	defer r.redisClent.Close()
	lua := redis.NewScript(keyCount, script)
	reply, err := lua.Do(r.redisClent, keysAndArgs...)
	return reply, err
}

func packArgs(items ...interface{}) (args []interface{}) {
	for _, item := range items {
		v := reflect.ValueOf(item)
		switch v.Kind() {
		case reflect.Slice:
			if v.IsNil() {
				continue
			}
			for i := 0; i < v.Len(); i++ {
				args = append(args, v.Index(i).Interface())
			}
		case reflect.Map:
			if v.IsNil() {
				continue
			}
			for _, key := range v.MapKeys() {
				args = append(args, key.Interface(), v.MapIndex(key).Interface())
			}
		default:
			args = append(args, v.Interface())
		}
	}
	return args
}
