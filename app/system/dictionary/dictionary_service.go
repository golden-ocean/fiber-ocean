package dictionary

import (
	"errors"

	"github.com/golden-ocean/fiber-ocean/app/system/dictionary_item"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/jinzhu/copier"
)

type Service struct {
	client             *ent.Client
	dictionaryRepo     *Repository
	dictioanryItemRepo *dictionary_item.Repository
}

func NewService() *Service {
	return &Service{
		client:             global.Client,
		dictionaryRepo:     NewRepository(),
		dictioanryItemRepo: dictionary_item.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &ent.Dictionary{}
	_ = copier.Copy(e, r)
	err := s.dictionaryRepo.Create(e, s.client)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	e := &ent.Dictionary{}
	_ = copier.Copy(e, r)
	err := s.dictionaryRepo.Update(e, s.client)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	if exist, _ := s.dictioanryItemRepo.Exist(&dictionary_item.WhereParams{DictionaryID: r.ID}, s.client); exist {
		return errors.New(ErrorExistChildren)
	}
	e := &ent.Dictionary{ID: r.ID}
	err := s.dictionaryRepo.Delete(e, s.client)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*DictionaryOutput, int, error) {
	es, total, err := s.dictionaryRepo.QueryPage(w, s.client)
	output := make([]*DictionaryOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) QueryByCode(code string) ([]*dictionary_item.DictionaryItemOutput, error) {
	es, err := s.dictioanryItemRepo.QueryByCode(code, s.client)
	output := make([]*dictionary_item.DictionaryItemOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, err
}
