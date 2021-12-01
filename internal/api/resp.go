package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
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

type DepartmentBase struct {
	DepartmentId       int64             `json:"departmentId"`
	ParentDepartmentId int64             `gorm:"column:parent_department_id;default:0;NOT NULL"` // 父级部门
	Level              int               `gorm:"-" json:"level"`                                 // 部门层级
	Children           []*DepartmentBase `json:"children"`
}

func ListConvertTree() {
	list := make([]*DepartmentBase, 10)
	str := "[{\"departmentId\":23,\"DepartmentName\":\"部门名称\",\"ParentDepartmentId\":0,\"level\":0,\"Status\":1,\"Sort\":1,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"2578981388\",\"CreateTime\":1637155547,\"UpdateTime\":1637155547,\"Children\":null,\"Type\":0},{\"departmentId\":27,\"DepartmentName\":\"中台\",\"ParentDepartmentId\":23,\"level\":0,\"Status\":1,\"Sort\":2,\"Remark\":\"ass\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637220612,\"UpdateTime\":1637289864,\"Children\":null,\"Type\":0},{\"departmentId\":33,\"DepartmentName\":\"售卖\",\"ParentDepartmentId\":27,\"level\":0,\"Status\":1,\"Sort\":10,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637234563,\"UpdateTime\":1637287893,\"Children\":null,\"Type\":0},{\"departmentId\":34,\"DepartmentName\":\"交易\",\"ParentDepartmentId\":55,\"level\":0,\"Status\":1,\"Sort\":1,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637238814,\"UpdateTime\":1637309229,\"Children\":null,\"Type\":0},{\"departmentId\":35,\"DepartmentName\":\"商品\",\"ParentDepartmentId\":33,\"level\":0,\"Status\":1,\"Sort\":2,\"Remark\":\"分到公司\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637238830,\"UpdateTime\":1637309557,\"Children\":null,\"Type\":0},{\"departmentId\":36,\"DepartmentName\":\"营销\",\"ParentDepartmentId\":33,\"level\":0,\"Status\":1,\"Sort\":2,\"Remark\":\"上地佛啊飞\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637238843,\"UpdateTime\":1637309540,\"Children\":null,\"Type\":0},{\"departmentId\":55,\"DepartmentName\":\"正向交易\",\"ParentDepartmentId\":34,\"level\":0,\"Status\":1,\"Sort\":1,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637288269,\"UpdateTime\":1637308507,\"Children\":null,\"Type\":0},{\"departmentId\":56,\"DepartmentName\":\"逆向交易\",\"ParentDepartmentId\":34,\"level\":0,\"Status\":1,\"Sort\":2,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637288289,\"UpdateTime\":1637308543,\"Children\":null,\"Type\":0},{\"departmentId\":59,\"DepartmentName\":\"二级部门3\",\"ParentDepartmentId\":23,\"level\":0,\"Status\":1,\"Sort\":1,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637290456,\"UpdateTime\":1637290546,\"Children\":null,\"Type\":0},{\"departmentId\":60,\"DepartmentName\":\"二级部门\",\"ParentDepartmentId\":27,\"level\":0,\"Status\":1,\"Sort\":10,\"Remark\":\"\",\"UnitId\":\"2001\",\"OpUid\":\"100000023728\",\"CreateTime\":1637308315,\"UpdateTime\":1637308343,\"Children\":null,\"Type\":0}]"
	err := json.Unmarshal([]byte(str), &list)
	fmt.Println(err)

	treeResp := new(DepartmentBase)

	// 列表转树，时间复杂度：O(n^2)，空间复杂度：O(1)
	for _, base := range list {
		if base.ParentDepartmentId == 0 {
			base.Level = 1
			treeResp = base
		}
		for _, p := range list {
			if p.ParentDepartmentId == base.DepartmentId {
				p.Level = base.Level + 1
				base.Children = append(base.Children, p)
			}
		}
	}

	// 列表转树，时间复杂度：O(n)，空间复杂度：O(n)
	m := make(map[int64]*DepartmentBase)
	for _, base := range list {
		m[base.DepartmentId] = base
	}
	for _, base := range list {
		if base.ParentDepartmentId == 0 {
			base.Level = 1
			treeResp = base
			continue
		}
		parent := m[base.ParentDepartmentId]
		base.Level = parent.Level + 1
		parent.Children = append(parent.Children, base)
	}

	fmt.Println(treeResp)
	fmt.Println("xxxxx")
	fmt.Println("xxxxx")
	fmt.Println("xxxxx")
	fmt.Println("xxxxx")

	// 树转列表，时间复杂度：O(结点数量)，空间复杂度：O(n)
	l := make([]*DepartmentBase, 10)
	l = append(l, &DepartmentBase{
		DepartmentId:       treeResp.DepartmentId,
		ParentDepartmentId: treeResp.ParentDepartmentId,
		Level:              treeResp.Level,
		Children:           nil,
	})
	children := treeResp.Children
	for children != nil {
		var tmp []*DepartmentBase
		for _, child := range children {
			l = append(l, &DepartmentBase{
				DepartmentId:       child.DepartmentId,
				ParentDepartmentId: child.ParentDepartmentId,
				Level:              child.Level,
				Children:           nil,
			})
			tmp = append(tmp, child)
		}
		children = tmp
	}
	fmt.Println(l)
	fmt.Println("xxxxx")
	fmt.Println("xxxxx")
	fmt.Println("xxxxx")
	fmt.Println("xxxxx")
}

func X() {
	fmt.Println(3300*12 + 330*10.8 + 150*12 + 12*3000)
	fmt.Println(3300*12 + 330*10.8 + 150*12 + 12*5000)

	tmp := make([]float64, 7)
	tmp[0] = 23
	salary := 23.0
	fmt.Printf("%.2f,", float64(salary))
	for i := 0; i < 6; i++ {
		salary = salary * (1 + 0.08)
		fmt.Printf("%.2f,", float64(salary))
	}
	fmt.Println()
	salary = 23.0
	yearCost := 10.0
	fmt.Printf("%.2f,", float64(yearCost))
	for i := 0; i < 6; i++ {
		salary = salary * (1 + 0.08)
		yearCost = yearCost * (1 + 0.05)
		fmt.Printf("%.2f,", float64(yearCost))
	}
	fmt.Println()

	salary = 23.0
	yearCost = 10.0
	saveMoney := 13.0
	fmt.Printf("%.2f,", float64(saveMoney))
	for i := 0; i < 6; i++ {
		salary = salary * (1 + 0.08)
		yearCost = yearCost * (1 + 0.05)
		saveMoney = (saveMoney + (salary - yearCost)) * 1.08
		tmp[i+1] = salary
		fmt.Printf("%.2f,", float64(saveMoney))
	}

	fmt.Println()
	for i := 0; i < 7; i++ {
		fmt.Printf("%.2f,", float64(tmp[i]/0.70))
	}
	fmt.Println()
	for i := 0; i < 7; i++ {
		fmt.Printf("%.2f,", float64(tmp[i]/0.70/14))
	}
	fmt.Println()
	for i := 0; i < 7; i++ {
		fmt.Printf("%.2f,", float64(tmp[i]/14))
	}
	fmt.Println()
	fmt.Println(163 / math.Pow(1.04, 7))
}
