package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

const MAX_PRINT_BODY_LEN = 512

type responseBodyWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

// https://zhuanlan.zhihu.com/p/265395083
// 无法接收 string，重写 WriteString
//func (r responseBodyWriter) WriteString(s string) (n int, err error)  {
//	r.bodyBuf.WriteString(s)
//	return r.ResponseWriter.WriteString(s)
//}
func (w responseBodyWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

// 如果需要筛选 response，需要在 c.json() 前进行筛选
func GinResponse(c *gin.Context) {
	w := &responseBodyWriter{bodyBuf: &bytes.Buffer{}, ResponseWriter: c.Writer}
	// 使用 responseBodyWriter 替换 gin 中的 responseWriter,
	// 替换的目的是把 response 返回值缓存起来
	c.Writer = w
	c.Next()
	// 打印 response 的返回值
	m := make(map[string]interface{})
	i := w.bodyBuf.Bytes()
	_ = json.Unmarshal(i, &m)
	m["saasa"] = 1111
	//marshal, _ := json.Marshal(m)
	//_, _ = c.Writer.Write(marshal)
	//w.bodyBuf.Write(marshal)
	fmt.Println("Response body: " + w.bodyBuf.String())
}
