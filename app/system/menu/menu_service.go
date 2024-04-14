package menu

import (
	"errors"
	"strings"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/constants"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type Service struct {
	client   *ent.Client
	menuRepo *Repository
}

func NewService() *Service {
	return &Service{
		client:   global.Client,
		menuRepo: NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &ent.Menu{}
	_ = copier.Copy(e, r)
	err := s.menuRepo.Create(e, s.client)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	all, err := s.menuRepo.QueryAll(&WhereParams{}, s.client)
	if err != nil {
		return err
	}
	ok := s.isValidParentID(r.ID, r.ParentID, all)
	if !ok {
		return errors.New(ErrorPidCantEqSelfAndChildId)
	}
	e := &ent.Menu{}
	_ = copier.Copy(e, r)
	err = s.menuRepo.Update(e, s.client)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	if exist, _ := s.menuRepo.Exist(&WhereParams{ParentID: r.ID}, s.client); exist {
		return errors.New(ErrorExistChildren)
	}
	e := &ent.Menu{ID: r.ID}
	err := s.menuRepo.Delete(e, s.client)
	return err
}

func (s *Service) QueryTree(w *WhereParams) ([]*MenuOutput, error) {
	entities, err := s.menuRepo.QueryAll(w, s.client)
	output := make([]*MenuOutput, 0)
	_ = copier.Copy(&output, entities)
	tree := s.buildTree(output)
	return tree, err
}

func (s *Service) QueryByStaffID(id string) ([]*MenuOutput, []string, error) {
	es, err := s.menuRepo.QueryByStaffID(id, s.client)
	if err != nil {
		return nil, nil, err
	}
	// 去除button权限菜单
	filterMenus := lo.Filter(es, func(item *ent.Menu, _ int) bool { return item.Type != constants.Button })
	// button 权限
	buttons := lo.Filter(es, func(item *ent.Menu, _ int) bool { return item.Type == constants.Button })
	permissions := lo.Map(buttons, func(item *ent.Menu, _ int) string {
		var build strings.Builder
		build.WriteString(item.Path)
		build.WriteString(":")
		build.WriteString(item.Method)
		return build.String()
	})
	output := make([]*MenuOutput, 0)
	_ = copier.Copy(&output, filterMenus)
	tree := s.buildTree(output)
	return tree, permissions, err
}

func (s *Service) buildTree(menus []*MenuOutput) []*MenuOutput {
	menuMap := make(map[string]*MenuOutput)
	// 映射到一个字典中
	for _, m := range menus {
		menuMap[m.ID] = m
	}
	for _, m := range menus {
		if parent, ok := menuMap[m.ParentID]; ok {
			parent.Children = append(parent.Children, m)
		}
	}
	// 寻找根组织并返回
	var roots []*MenuOutput
	for _, m := range menus {
		if len(m.ParentID) == 0 {
			roots = append(roots, m)
		}
	}
	return roots
}

func (s *Service) isValidParentID(id, parentID string, menus []*ent.Menu) bool {
	if id == parentID {
		return false
	}
	ancestors := make(map[string]struct{})
	menuMap := make(map[string]*ent.Menu)
	for _, m := range menus {
		menuMap[m.ID] = m
	}
	for len(parentID) > 0 {
		if _, ok := ancestors[parentID]; ok {
			return false
		}
		ancestors[parentID] = struct{}{}
		parentMenu, ok := menuMap[parentID]
		if !ok {
			return false
		}
		if parentMenu.ID == id {
			return false
		}
		parentID = parentMenu.ParentID
	}
	return true
}
