package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

var _ persist.Adapter = (*CasbinAdapter)(nil)

// CasbinAdapter casbin适配器
type CasbinAdapter struct {
	//RoleModel         *model.Role
	//RoleMenuModel     *model.RoleMenu
	////MenuResourceModel *model.MenuActionResource
	//UserModel         *model.User
	//UserRoleModel     *model.UserRole
}

func (c CasbinAdapter) LoadPolicy(model model.Model) error {
	fmt.Println("dsdsds")
	panic("implement me")
}

func (c CasbinAdapter) SavePolicy(model model.Model) error {
	panic("implement me")
}

func (c CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	panic("implement me")
}

func (c CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	panic("implement me")
}

func (c CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	panic("implement me")
}
