package casbinx

import (
	"errors"
	"fmt"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/golden-ocean/fiber-ocean/app/system/menu"
	"github.com/golden-ocean/fiber-ocean/app/system/role_menu"
	"github.com/golden-ocean/fiber-ocean/app/system/staff_role"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/constants"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/samber/lo"
)

type Adapter struct {
	client        *ent.Client
	staffRoleRepo staff_role.Repository
	roleMenuRepo  role_menu.Repository
	menuRepo      menu.Repository
}

func NewAdapter() *Adapter {
	return &Adapter{
		client:        global.Client,
		staffRoleRepo: *staff_role.NewRepository(),
		roleMenuRepo:  *role_menu.NewRepository(),
		menuRepo:      *menu.NewRepository(),
	}
}

func (a *Adapter) LoadPolicy(model model.Model) error {
	err := a.loadRolePolicy(model)
	if err != nil {
		return err
	}
	err = a.loadUserPolicy(model)
	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) SavePolicy(model model.Model) error {
	return errors.New("not implemented")
}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	//TODO implement me
	panic("implement me")
}

// load role policy (p,role_id,path,method)
func (a *Adapter) loadRolePolicy(model model.Model) error {
	rms, err := a.roleMenuRepo.Query(&role_menu.WhereParams{}, a.client)
	if err != nil {
		return err
	}
	m_ids := lo.Map(rms, func(rm *ent.Role_Menu, _ int) string {
		return rm.MenuID
	})
	menu_ids := lo.Uniq(m_ids)
	menus, err := a.menuRepo.Query(&menu.WhereParams{IDs: menu_ids}, a.client)
	if err != nil {
		return err
	}
	// 构建 p role_id path method
	for _, rm := range rms {
		for _, m := range menus {
			if m.ID == rm.MenuID && m.Type == constants.Button {
				line := fmt.Sprintf("p,%s,%s,%s", rm.RoleID, m.Path, m.Method)
				err := persist.LoadPolicyLine(line, model)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *Adapter) loadUserPolicy(model model.Model) error {
	staffs_roles, err := a.staffRoleRepo.Query(&staff_role.WhereParams{}, a.client)
	if err != nil {
		return err
	}
	// 构建 g staff_id role_id
	for _, staff_role := range staffs_roles {
		line := fmt.Sprintf("g,%s,%s", staff_role.StaffID, staff_role.RoleID)
		err := persist.LoadPolicyLine(line, model)
		if err != nil {
			return err
		}
	}
	return nil
}
