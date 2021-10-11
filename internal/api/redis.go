package api

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
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

/*
	lua
*/
func (r *Redis) Lua(script string, keyCount int, keysAndArgs ...interface{}) (interface{}, error) {
	defer r.redisClent.Close()
	lua := redis.NewScript(keyCount, script)
	reply, err := lua.Do(r.redisClent, keysAndArgs...)
	return reply, err
}
