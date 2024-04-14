package dictionary_item

import (
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/jinzhu/copier"
)

type Service struct {
	client       *ent.Client
	dictItemRepo *Repository
}

func NewService() *Service {
	return &Service{
		client:       global.Client,
		dictItemRepo: NewRepository(),
	}
}

func (s *Service) Create(req *CreateInput) error {
	e := &ent.Dictionary_Item{}
	_ = copier.Copy(e, req)
	err := s.dictItemRepo.Create(e, s.client)
	return err
}

func (s *Service) Update(req *UpdateInput) error {
	e := &ent.Dictionary_Item{}
	_ = copier.Copy(e, req)
	err := s.dictItemRepo.Update(e, s.client)
	return err
}

func (s *Service) Delete(req *DeleteInput) error {
	e := &ent.Dictionary_Item{ID: req.ID}
	err := s.dictItemRepo.Delete(e, s.client)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*DictionaryItemOutput, int, error) {
	es, total, err := s.dictItemRepo.QueryPage(w, s.client)
	output := make([]*DictionaryItemOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, total, err
}
