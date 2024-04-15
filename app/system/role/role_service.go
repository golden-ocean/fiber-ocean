package role

import (
	"errors"

	"github.com/golden-ocean/fiber-ocean/app/system/role_menu"
	"github.com/golden-ocean/fiber-ocean/app/system/staff_role"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/golden-ocean/fiber-ocean/platform/database"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type Service struct {
	client        *ent.Client
	roleRepo      *Repository
	staffRoleRepo *staff_role.Repository
	roleMenuRepo  *role_menu.Repository
}

func NewService() *Service {
	return &Service{
		client:        global.Client,
		roleRepo:      NewRepository(),
		staffRoleRepo: staff_role.NewRepository(),
		roleMenuRepo:  role_menu.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &ent.Role{}
	_ = copier.Copy(e, r)
	err := s.roleRepo.Create(e, s.client)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	e := &ent.Role{}
	_ = copier.Copy(e, r)
	err := s.roleRepo.Update(e, s.client)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	// 检查 staff_role 的关联
	if exist, _ := s.staffRoleRepo.Exist(&ent.Staff_Role{RoleID: r.ID}, s.client); exist {
		return errors.New(ErrorExistStaff)
	}
	// 检查 role_menu 的关联
	if exist, _ := s.roleMenuRepo.Exist(&ent.Role_Menu{RoleID: r.ID}, s.client); exist {
		return errors.New(ErrorExistMenu)
	}
	e := &ent.Role{ID: r.ID}
	err := s.roleRepo.Delete(e, s.client)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*RoleOutput, int, error) {
	es, total, err := s.roleRepo.QueryPage(w, s.client)
	output := make([]*RoleOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) QueryAll(w *WhereParams) ([]*RoleOutput, error) {
	es, err := s.roleRepo.QueryAll(w, s.client)
	output := make([]*RoleOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, err
}

func (s *Service) QueryMenus(id string) ([]string, error) {
	es, err := s.roleMenuRepo.Query(&role_menu.WhereParams{RoleID: id}, s.client)
	ids := lo.Map(es, func(e *ent.Role_Menu, _ int) string { return e.MenuID })
	return ids, err
}

func (s *Service) GrantMenus(r *RoleMenuInput) error {
	db_menu_ids, err := s.QueryMenus(r.RoleID)
	if err != nil {
		return err
	}
	remove_ids, add_ids := lo.Difference(db_menu_ids, r.MenuIDs)
	err = database.WithTx(s.client, func(tx *ent.Tx) error {
		if len(remove_ids) > 0 {
			if err := s.roleMenuRepo.Delete(&role_menu.DeleteInput{RoleID: r.RoleID, MenuIDs: remove_ids}, s.client); err != nil {
				return err
			}
		}
		if len(add_ids) > 0 {
			if err := s.roleMenuRepo.Create(&role_menu.CreateInput{RoleID: r.RoleID, MenuIDs: add_ids}, s.client); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
