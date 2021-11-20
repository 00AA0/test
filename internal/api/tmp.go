package api

//根据用户查询顶级机构：登陆

type GetUnitIdReq struct {
	Uid int64 `json:"uid"`
}

type GetUnitIdResp struct {
	UnitId int64 `json:"unitId"`
}

//根据（uid，）机构id 查询机构信息：基础信息
type GetOrganizationInfoReq struct {
	Uid    int64 `json:"uid"`
	UnitId int64 `json:"UnitId"`
}
type GetOrganizationInfoResp struct {
	UnitId           int64  `json:"unitId"`
	OrganizationName string `json:"organizationName"`
	Phone            string `json:"phone"`
	Address          string `json:"address"`
	//appinfo
}

//根据机构ID（机构名称），机构类型查询下辖机构：下辖机构
type GetChildOrganizationListReq struct {
	Uid          int64 `json:"uid"`
	UnitId       int64 `json:"unitId"`
	BusinessType int   `json:"businessType"`
}
type GetChildOrganizationListResp struct {
	ChildOrganization []Organization `json:"childOrganization"`
}

type Organization struct {
	UnitId           int64  `json:"unitId"`
	OrganizationName string `json:"organizationName"`
	BusinessType     int8   `json:"businessType"`
	// 办学性质
}

//分配角色：roleIds，机构ids，uid
type SetRoleReq struct {
	RoleIds []int64 `json:"roleIds" binding:"required"`
	Uid     int64   `json:"uid"`
	UnitIds []int64 `json:"UnitIds"`
}
type SetRoleResp struct {
}

//批量分配角色：[]分配角色

//MQ异步将（用户，用户-角色）写入 sch
//
//分配角色展示：机构id｜全量角色信息，下辖机构
type SetRoleInfoReq struct {
	Uid    int64 `json:"uid"`
	UnitId int64 `json:"UnitId"`
}
type SetRoleInfoResp struct {
	RoleInfo          []RoleInfo     `json:"roleInfo"`
	ChildOrganization []Organization `json:"childOrganization"`
}
type RoleInfo struct {
	RoleName string `json:"roleName"`
	RoleId   int64  `json:"roleId"`
	Selected bool   `json:"selected"`
}

//获取用户列表：顶级机构分配的人员不能操作，包括重置密码
//
//重置密码：
//
//删除用户：
