package auth

import (
	"errors"

	"github.com/golden-ocean/fiber-ocean/app/system/menu"
	"github.com/golden-ocean/fiber-ocean/app/system/staff"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/constants"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
	"github.com/jinzhu/copier"
)

type Service struct {
	client       *ent.Client
	staffService *staff.Service
	menuService  *menu.Service
}

func NewService() *Service {
	return &Service{
		client:       global.Client,
		staffService: staff.NewService(),
		menuService:  menu.NewService(),
	}
}

func (s *Service) Login(req *LoginInput) (*ent.Staff, error) {
	e, err := s.staffService.QueryByUniqueField(&staff.WhereParams{Username: req.Username})
	if err != nil {
		return nil, errors.New(ErrorUsernameOrPassword)
	}
	if e.Status != constants.ENABLE {
		return nil, errors.New(ErrorDisableStatus)
	}
	match := utils.ComparePasswords(e.Password, req.Password)
	if !match {
		return nil, errors.New(ErrorUsernameOrPassword)
	}
	return e, err
}

func (s *Service) Logout(id string) error {
	return nil
}

func (s *Service) QueryInfo(id string) (*InfoOutput, error) {
	re, err := s.staffService.QueryByUniqueField(&staff.WhereParams{ID: id})
	so := &staff.StaffOutput{}
	_ = copier.Copy(so, re)
	if err != nil {
		return nil, err
	}
	mo, po, err := s.menuService.QueryByStaffID(re.ID)
	if err != nil {
		return nil, err
	}
	return &InfoOutput{
		Staff:       so,
		Menus:       mo,
		Permissions: po,
	}, err
}

// func (s *Service) Refresh(id string) (*staff.Staff, error) {
// 	e, err := s.staffService.QueryByUniqueField(&staff.WhereParams{ID: id})
// 	return e, err
// }
