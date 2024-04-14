package position

import (
	"errors"

	"github.com/golden-ocean/fiber-ocean/app/system/staff_position"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/jinzhu/copier"
)

type Service struct {
	client            *ent.Client
	positionRepo      *Repository
	staffPositionRepo *staff_position.Repository
}

func NewService() *Service {
	return &Service{
		client:            global.Client,
		positionRepo:      NewRepository(),
		staffPositionRepo: staff_position.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &ent.Position{}
	_ = copier.Copy(e, r)
	err := s.positionRepo.Create(e, s.client)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	e := &ent.Position{}
	_ = copier.Copy(e, r)
	err := s.positionRepo.Update(e, s.client)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	// 检查 staff_position 关联
	if exist, _ := s.staffPositionRepo.Exist(&ent.Staff_Position{PositionID: r.ID}, s.client); exist {
		return errors.New(ErrorExistStaff)
	}
	e := &ent.Position{ID: r.ID}
	err := s.positionRepo.Delete(e, s.client)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*PositionOutput, int, error) {
	es, total, err := s.positionRepo.QueryPage(w, s.client)
	output := make([]*PositionOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) QueryAll(w *WhereParams) ([]*PositionOutput, error) {
	es, err := s.positionRepo.QueryAll(w, s.client)
	output := make([]*PositionOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, err
}
