package main

import (
	"fmt"
)

type TblUserRole struct {
	Id         int64 `gorm:"id"           json:"id"`        // 自增id
	Uid        int64 `gorm:"uid"          json:"uid"`       // 用户id
	SchoolId   int64 `gorm:"school_id"    json:"schoolId"`  // 学校id
	RoleId     int64 `gorm:"role_id"      json:"roleId"`    // 角色id
	Status     int8  `gorm:"status"       json:"status"`    // 1:正常 2:删除
	OpUid      int64 `gorm:"op_uid"       json:"opUid"`     // 操作人uid
	CreateTime int64 `gorm:"create_time" json:"createTime"` // 创建时间
	UpdateTime int64 `gorm:"update_time" json:"updateTime"` // 更新时间
}

type RoleMenu struct {
	Id         int64  `gorm:"id"           json:"id"`        // 自增id
	RoleId     int64  `gorm:"role_id"      json:"roleId"`    // 角色id
	MenuId     int64  `gorm:"menu_id"      json:"menuId"`    // 菜单id
	Opids      string `gorm:"opids"        json:"opids"`     // 菜单操作权限， 空则表示拥有所有操作权限  [1,2,3]
	Status     int8   `gorm:"status"       json:"status"`    // 1:正常 2:删除
	OpUid      int64  `gorm:"op_uid"       json:"opUid"`     // 操作人uid
	CreateTime int64  `gorm:"create_time" json:"createTime"` // 创建时间
	UpdateTime int64  `gorm:"update_time" json:"updateTime"` // 更新时间
}

//_, _ = e.AddPolicy("root", "/auth/*", "get")
//_, _ = e.AddGroupingPolicy("name1", "root")
//_, _ = e.UpdatePolicy([]string{"eve", "data3", "read"}, []string{"eve", "data3", "write"})
//ok, _ = e.UpdateGroupingPolicy([]string{"name1", "root"}, []string{"admin", "data4_admin"})
//_, _ = e.RemovePolicy("alice", "data1", "read")
//ok, _ = e.RemoveGroupingPolicy("name1", "root")

func AddPolicy(roleMenu []RoleMenu) (ok bool, err error) {
	for _, menu := range roleMenu {
		fmt.Println(menu.RoleId, menu.MenuId, menu.Opids)
	}

	//m := make([]map[string]string, 0)

	//m := map[string]string{
	//	"schui/auth/adduser": "100112",
	//	"schui/auth/editrole": "100117",
	//}
	//
	//mGet := map[string]string{
	//	"schui/auth/getmenutree": "schui/auth/editrole#schui/auth/addrole",
	//	"schui/auth/getroleinfo": "schui/auth/editrole",
	//	"schui/auth/adduserinfo": "schui/auth/adduser",
	//	"schui/auth/edituserinfos": "schui/auth/edituserbatch#schui/auth/edituser",
	//	"schui/teacher/list": "schui/class/add#schui/class/update",
	//	"schui/class/detail": "schui/class/update",
	//}
	return
}
func AddGroupingPolicy(userRole []TblUserRole) (ok bool, err error) {
	for _, role := range userRole {
		fmt.Println(role.Uid, role.RoleId)
	}
	return
}
func UpdatePolicy(roleMenuOld []RoleMenu, roleMenuNew []RoleMenu) (ok bool, err error) {

	return
}
func UpdateGroupingPolicy(userRoleOld []TblUserRole, userRoleNew []TblUserRole, roleMenu []RoleMenu) (ok bool, err error) {

	return
}
func RemovePolicy(roleMenu []RoleMenu) (ok bool, err error) {
	return
}
func RemoveGroupingPolicy(userRole []TblUserRole) (ok bool, err error) {
	return
}
