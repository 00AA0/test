package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CrmTagGroup struct {
	Id         int64     `gorm:"id" json:"id"`                  // ID
	Name       string    `gorm:"name" json:"name"`              // 标签组名
	CreateTime time.Time `gorm:"create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"update_time" json:"updateTime"` // 更新时间
	Role       int8      `gorm:"role" json:"role"`              // 角色
	OpId       string    `gorm:"op_id" json:"opId"`             // 操作人accountID
	Deleted    int8      `gorm:"deleted" json:"deleted"`        // 软删状态
}
type CrmTags struct {
	Id         int64     `gorm:"id" json:"id"` // ID
	Name       string    `gorm:"name" json:"name"`
	CreateTime time.Time `gorm:"create_time" json:"createTime"`  // 创建时间
	UpdateTime time.Time `gorm:"update_time" json:"updateTime"`  // 更新时间
	TagGroupId int64     `gorm:"tag_group_id" json:"tagGroupId"` // 标签组id
	Deleted    int64     `gorm:"deleted" json:"deleted"`         // 软删状态
	OpId       string    `gorm:"op_id" json:"opId"`              // 操作人accountID
}

func (CrmTags) TableName() string {
	return "tblCrmTags"
}
func (CrmTagGroup) TableName() string {
	return "tblCrmTagGroup"
}

func initMySql() (client *gorm.DB, err error) {
	client, err = gorm.Open(mysql.Open("root:123456789@(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func main() {
	db, _ := initMySql()
	//var tagGroup1 []CrmTagGroup
	var tagGroup2 []CrmTagGroup
	//db = db.Find(&tagGroup)
	//fmt.Println(len(tagGroup))

	//db.Table("tblCrmTagGroup").Where("name = ?", "test_3").Update("name", "test_33")

	go func() {
		err := db.Transaction(func(tx *gorm.DB) error {
			//tx = tx.Find(&tagGroup1)
			//fmt.Println(len(tagGroup1))
			time.Sleep(2 * time.Second)
			fmt.Println("开始更新")
			tx = tx.Table("tblCrmTagGroup").Where("name = ?", "test_4").Update("name", "44test_44")
			fmt.Println("更新结束")
			db.Table("tblCrmTagGroup").Find(&tagGroup2)
			fmt.Println(len(tagGroup2))
			//fmt.Println(tagGroup2)
			return nil
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	//time.Sleep(2 * time.Second)

	go func() {
		err := db.Transaction(func(tx *gorm.DB) error {
			//tagGroup4 := &CrmTagGroup{
			//	Name:       "test_4",
			//	CreateTime: time.Now(),
			//	UpdateTime: time.Now(),
			//	Role:       1,
			//	Deleted:    0,
			//}
			//time.Sleep(time.Second)
			//tx = tx.Create(tagGroup4)
			//fmt.Println(tagGroup4.Id)
			var tagGroup3 []CrmTagGroup
			tx.Where("name = ?", "test_4").Find(&tagGroup3)
			fmt.Println(len(tagGroup3))
			time.Sleep(5 * time.Second)
			fmt.Println("睡眠结束")
			return nil
		})
		fmt.Println("tijiao")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	time.Sleep(10 * time.Second)

	//fmt.Println("test2_1")
	//fmt.Println("test_1")

	//tmp := &CrmTag{}
	//var tmp CrmTag

	//leadsInfo := LeadsInfo{Id: 1, StageID: 2}

	//fmt.Println(tmp)

	//fmt.Println(tmp == (CrmTag{}))
	//fmt.Println(leadsInfo == (LeadsInfo{}))

	//router := gin.Default()
	//// 处理multipart forms提交文件时默认的内存限制是32 MiB
	//// 可以通过下面的方式修改
	//// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	//router.POST("/upload", func(c *gin.Context) {
	//	// 单个文件
	//	file, err := c.FormFile("upload")
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"message": err.Error(),
	//		})
	//		return
	//	}
	//
	//	log.Println(file.Filename)
	//	//dst := fmt.Sprintf("desktop/%s", file.Filename + "1")
	//	// 上传文件到指定的目录
	//	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
	//		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	//	})
	//})

	//router.GET("/method", func(c *gin.Context) {
	//
	//	for {
	//		now := time.Now()
	//		defer func() {
	//			if i := recover(); i != nil {
	//				fmt.Println(i)
	//
	//			}
	//		}()
	//		time.Sleep(1000*1000*1000)
	//		since := time.Since(now)
	//		fmt.Println(since)
	//		if since > 0 {
	//			panic("aa")
	//		}
	//		c.String(200, "hello hello")
	//	}
	//})
	//
	//router.Run(":8080")

	//fmt.Println(zap.Error(errors.New("sdsd")))
	//log.Printf("aa")
	//fmt.Println(len("😊"))
	//err := errors.New("qqq")
	//1.创建路由
	//默认使用了2个中间件Logger(), Recovery()
	//r := gin.New()
	// 注册中间件
	//r.Use(MiddleWare())
	//r.Use(MiddleWare2())
	//r.Use(MiddleWare3())
	//// {}为了代码规范
	//{
	//	r.GET("/ce", func(c *gin.Context) {
	//		// 取值
	//		req, _ := c.Get("request")
	//		fmt.Println("request:", req)
	//		// 页面接收
	//		c.JSON(200, gin.H{"request": req})
	//		panic("error test")
	//		//if err != nil {
	//		//	panic("error")
	//		//}
	//
	//	})
	//}
	//r.Run()

	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Unix(1623168000, 0).Format("2006-01-02 15:04:05"))
	//parse, _ := time.Parse("01/02/2006", "06/21/2017")
	//
	//fmt.Println(CrmTag{}.TagName)
	//fmt.Println(parse.Unix())
	//fmt.Println(len("a"))
	//res := make(map[string]interface{})
	//fmt.Println(res)
	//tag := make([]CrmTag, 10)
	//
	////fmt.Println(len(da))
	//for i := 0; i < len(tag); i++ {
	//	tag[i] = CrmTag{
	//		//TagId: int64(i),
	//		TagName: "n" + strconv.Itoa(i),
	//	}
	//	//da["a" + strconv.Itoa(i)] = i
	//}

	//for k, v := range da {
	//	fmt.Println(k, v)
	//}

	//marshal, _ := json.Marshal(tag)
	//_ = json.Unmarshal(marshal, &da)
	//fmt.Println(da)
	//var a interface{} = 0
	//i := da[0]["tagId"]
	//i2 := da[0]["tagName"]

	//m := make(map[string]string, 5)
	//fmt.Println(len(m))

	//test(tag, map[string]interface{}{"a": 1, "b": 2})
	//crmTag := CrmTag{TagId: 1, TagName: "a"}
	//of := reflect.TypeOf(crmTag)
	//fmt.Println(of.Field(1).Name)
	//str := "1、极好的是"
	//fmt.Println(str[1:])
	//fmt.Println(10737418240 / 1024 / 1024 / 1024)
	//var a int = 12345678900
	//var b int32 = -2004567890
	//fmt.Println(b)
	//fmt.Println(unsafe.Sizeof(a))
	//fmt.Println(strconv.IntSize)
	//fmt.Println(67108864/1024/1024)

	//da := make([]map[string]interface{}, 5)
	//for j := 0; j < 5; j++ {
	//	da[j] = make(map[string]interface{}, 5)
	//}
	//for i := 0; i < 5; i++ {
	//	da[i]["role"] = "role" + strconv.Itoa(i)
	//	da[i]["phone"] = "phone" + strconv.Itoa(i)
	//	if i == 3 {
	//		da[i]["uid"] = "uid" + strconv.Itoa(i)
	//	}
	//	da[i]["name"] = "name" + strconv.Itoa(i)
	//	da[i]["remark"] = "remark" + strconv.Itoa(i)
	//}
	//err := UpdateBatch(da, "tablename", []string{"role", "phone"}, []string{"uid", "name", "remark"})
	//if err != nil {
	//
	//}

	//t1(errors.New("test"))

}

func t1(err error) { t2(err) }
func t2(err error) { t3(err) }
func t3(err error) {
	t4(err)
}
func t4(err error) {

	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Println(errors.New(file + "-->第" + strconv.Itoa(line) + "行-->err:" + err.Error()))
	}
}

func UpdateBatch( /*ctx *gin.Context, */ data []map[string]interface{}, tableName string, condition []string, fields []string) error {
	res := "update " + tableName + " as t set "
	sqlMap := make(map[string]string, len(fields))
	for i := 0; i < len(fields); i++ {
		sqlMap[fields[i]] = "t." + fields[i] + "= case "
	}

	sql := ""
	for _, m := range data {
		sql = "when t."
		for _, c := range condition {
			v, ok := m[c]
			if ok {
				sql = sql + c + "=" + v.(string) + " and t."
			}
		}
		sql = strings.TrimRight(sql, "and t.")
		sql = sql + " then "
		for _, f := range fields {
			v, ok := m[f]
			if ok {
				sqlMap[f] = sqlMap[f] + sql + v.(string) + " "
			}
		}
	}

	for k, v := range sqlMap {
		v = v + "else t." + k + " end,"
		res = res + v
	}
	res = strings.TrimRight(res, ",")
	fmt.Println(res)

	return nil
}

func test(t interface{}, m interface{}) {
	fieldValue := reflect.ValueOf(t)
	//fieldType := reflect.TypeOf(t)//.Elem().Elem()
	//fmt.Println(fieldValue.Index(1).Field(0))
	da := make([]map[string]interface{}, fieldValue.Len())
	l := fieldValue.Len()
	for i := 0; i < l; i++ {
		da[i] = make(map[string]interface{}, 2)
	}

	marshal, _ := json.Marshal(t)
	_ = json.Unmarshal(marshal, &da)
	//m2 := structs.Map(&fieldValue)
	//fmt.Println(m2)

	//data = make(map[string]interface{})
	//objT := reflect.TypeOf(obj)
	//objV := reflect.ValueOf(obj)
	//for i := 0; i < objT.NumField(); i++ {
	//	data[objT.Field(i).Name] = fieldType.Field(i).Interface()
	//}
	//fmt.Println(fieldValue.Len())
	//fmt.Println(fieldValue.Index(0).Field(1))
	//name := reflect.TypeOf(fieldValue.Index(1)).Field(0).Name
	//fmt.Println(name)
	//
	//for i := 0; i < fieldValue.Len(); i++ {
	//	for j := 0; j < fieldValue.Index(i).NumField(); j++ {
	//		//da[i]["a"] = i
	//		da[i][reflect.TypeOf(fieldValue.Index(i)).Field(j).Name] = fieldValue.Index(i).Field(j).Interface()
	//	}
	//}
	fmt.Println(da)

	//value := reflect.ValueOf(m)
	//fmt.Println(m.(map[string]interface{}))

	//sliceLength := fieldType.Len()
	//fieldNum := fieldType.NumField()
	//tag := t.([]CrmTag)
	//fmt.Println("==============")
	//fmt.Println(tag)
	//fmt.Println(fieldValue.Kind())
	//fmt.Println(fieldType)
	//fmt.Println(sliceLength)
	//fmt.Println(fieldNum)
}

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		//c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		//fmt.Println("中间件执行完毕")
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了1111111")
		// 设置变量到Context的key中，可以通过Get()取
		//c.Set("request", "中间件")
		//c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕1111111111", status)
		//fmt.Println("中间件执行完毕")
		t2 := time.Since(t)
		fmt.Println("11111time:", t2)
	}
}

func MiddleWare3() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer CatchRecoverRpc()
		c.Next()
	}
}

func CatchRecoverRpc() {
	// panic捕获
	if err := recover(); err != nil {
		//fmt.Println("error recover")
		fmt.Println(err)
		t4(errors.New(err.(string)))
		//c := *gin.Context
		////请求url
		//path := c.Request.URL.Path
		//raw := c.Request.URL.RawQuery
		//if raw != "" {
		//	path = path + "?" + raw
		//}
		////请求报文
		//body, _ := ioutil.ReadAll(c.Request.Body)
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		//fields := []zlog.Field{
		//	zlog.String("logId", zlog.GetLogID(c)),
		//	zlog.String("requestId", zlog.GetRequestID(c)),
		//	zlog.String("uri", path),
		//	zlog.String("refer", c.Request.Referer()),
		//	zlog.String("clientIp", c.ClientIP()),
		//	zlog.String("module", env.AppName),
		//	zlog.String("ua", c.Request.UserAgent()),
		//	zlog.String("host", c.Request.Host),
		//}
		//zlog.InfoLogger(c, "Panic[recover]", fields...)
		//
		//base.RenderJsonAbort(c, components.SystemError)
	}
}
