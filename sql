package test

type TblSchoolArchive struct {
	Id                int    `gorm:"id" json:"id"`
	SchoolId          int    `gorm:"school_id" json:"school_id"`                     // 学校id
	SchoolInfo        string `gorm:"school_info" json:"school_info"`                 // 学校信息
	MenuInfo          string `gorm:"menu_info" json:"menu_info"`                     // 菜单信息
	RoleInfo          string `gorm:"role_info" json:"role_info"`                     // 角色信息
	DataAuthorityInfo string `gorm:"data_authority_info" json:"data_authority_info"` // 数据权限信息
	MenuAuthInfo      string `gorm:"menu_auth_info" json:"menu_auth_info"`           // 用户菜单权限
}

func (*TblSchoolArchive) TableName() string {
	return "tblSchoolArchive"
}
