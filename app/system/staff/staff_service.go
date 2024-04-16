package staff

import (
	"github.com/casbin/casbin/v2"
	"github.com/golden-ocean/fiber-ocean/app/system/staff_position"
	"github.com/golden-ocean/fiber-ocean/app/system/staff_role"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/golden-ocean/fiber-ocean/platform/database"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type Service struct {
	client            *ent.Client
	staffRepo         *Repository
	staffPositionRepo *staff_position.Repository
	staffRoleRepo     *staff_role.Repository
	casbinEnforcer    *casbin.Enforcer
}

func NewService() *Service {
	return &Service{
		client:            global.Client,
		staffRepo:         NewRepository(),
		staffPositionRepo: staff_position.NewRepository(),
		staffRoleRepo:     staff_role.NewRepository(),
		casbinEnforcer:    global.Enforcer,
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &ent.Staff{}
	_ = copier.Copy(e, r)
	err := database.WithTx(s.client, func(tx *ent.Tx) error {
		// 创建员工
		re, err := s.staffRepo.Create(e, tx.Client())
		if err != nil {
			return err
		}
		// 创建员工职位关联
		err = s.staffPositionRepo.Create(&staff_position.CreateInput{
			StaffID:     re.ID,
			PositionIDs: r.PositionIDs}, tx.Client())
		if err != nil {
			return err
		}
		// 创建员工角色关联
		err = s.staffRoleRepo.Create(&staff_role.CreateInput{
			StaffID: re.ID,
			RoleIDs: r.RoleIDs}, tx.Client())
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	staff_position_list, err := s.staffPositionRepo.Query(&staff_position.WhereParams{StaffID: r.ID}, s.client)
	if err != nil {
		return err
	}
	db_p_ids := lo.Map(staff_position_list, func(e *ent.Staff_Position, _ int) string { return e.PositionID })
	remove_p_ids, add_p_ids := lo.Difference(db_p_ids, r.PositionIDs)

	staff_role_list, err := s.staffRoleRepo.Query(&staff_role.WhereParams{StaffID: r.ID}, s.client)
	if err != nil {
		return err
	}
	db_r_ids := lo.Map(staff_role_list, func(e *ent.Staff_Role, _ int) string { return e.RoleID })
	remove_r_ids, add_r_ids := lo.Difference(db_r_ids, r.RoleIDs)

	e := &ent.Staff{}
	_ = copier.Copy(e, r)
	err = database.WithTx(s.client, func(tx *ent.Tx) error {
		err := s.staffRepo.Update(e, s.client)
		if err != nil {
			return err
		}
		if len(remove_p_ids) > 0 {
			if err := s.staffPositionRepo.Delete(&staff_position.DeleteInput{StaffID: r.ID, PositionIDs: remove_p_ids}, s.client); err != nil {
				return err
			}
		}
		if len(add_p_ids) > 0 {
			if err := s.staffPositionRepo.Create(&staff_position.CreateInput{StaffID: r.ID, PositionIDs: add_p_ids}, s.client); err != nil {
				return err
			}
		}
		if len(remove_r_ids) > 0 {
			if err := s.staffRoleRepo.Delete(&staff_role.DeleteInput{StaffID: r.ID, RoleIDs: remove_r_ids}, s.client); err != nil {
				return err
			}
		}
		if len(add_r_ids) > 0 {
			if err := s.staffRoleRepo.Create(&staff_role.CreateInput{StaffID: r.ID, RoleIDs: add_r_ids}, s.client); err != nil {
				return err
			}
		}
		return nil
	})
	if err == nil {
		s.casbinEnforcer.LoadPolicy()
	}
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	err := database.WithTx(s.client, func(tx *ent.Tx) error {
		// 删除 staff_position 关联
		if err := s.staffPositionRepo.Delete(&staff_position.DeleteInput{StaffID: r.ID}, s.client); err != nil {
			return err
		}
		// 删除 staff_role 关联
		if err := s.staffRoleRepo.Delete(&staff_role.DeleteInput{StaffID: r.ID}, s.client); err != nil {
			return err
		}
		// 删除 staff
		if err := s.staffRepo.Delete(&ent.Staff{ID: r.ID}, s.client); err != nil {
			return err
		}
		return nil
	})
	if err == nil {
		s.casbinEnforcer.LoadPolicy()
	}
	return err
}

func (s *Service) QueryByUniqueField(w *WhereParams) (*ent.Staff, error) {
	e, err := s.staffRepo.QueryByUniqueField(w, s.client)
	return e, err
}

// func (s *Service) QueryPage(w *WhereParams) ([]*StaffOutput, int, error) {
// 	es, total, err := s.staffRepo.QueryPage(w, s.client)
// 	output := make([]*StaffOutput, 0)
// 	_ = copier.Copy(&output, es)
// 	return output, total, err
// }

func (s *Service) QueryPage(w *WhereParams) ([]*StaffOutput, int, error) {
	es, total, err := s.staffRepo.QueryPage(w, s.client)
	if err != nil {
		return nil, 0, err
	}
	s_ids := lo.Map(es, func(e *ent.Staff, _ int) string {
		return e.ID
	})
	staffs_positions, err := s.staffPositionRepo.Query(&staff_position.WhereParams{StaffIDs: s_ids}, s.client)
	if err != nil {
		return nil, 0, err
	}
	staffs_roles, err := s.staffRoleRepo.Query(&staff_role.WhereParams{StaffIDs: s_ids}, s.client)
	if err != nil {
		return nil, 0, err
	}
	output := make([]*StaffOutput, 0)
	_ = copier.Copy(&output, es)
	for _, entity := range output {
		position_ids := lo.FilterMap(staffs_positions, func(e *ent.Staff_Position, _ int) (string, bool) {
			if entity.ID == e.StaffID {
				return e.PositionID, true
			}
			return "", false
		})
		role_ids := lo.FilterMap(staffs_roles, func(e *ent.Staff_Role, _ int) (string, bool) {
			if entity.ID == e.StaffID {
				return e.RoleID, true
			}
			return "", false
		})
		entity.PositionIDs = position_ids
		entity.RoleIDs = role_ids
	}
	return output, total, err
}
