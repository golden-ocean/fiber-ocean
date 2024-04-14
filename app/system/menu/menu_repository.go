package menu

import (
	"context"
	"errors"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/menu"
	"github.com/golden-ocean/fiber-ocean/ent/role_menu"
	"github.com/golden-ocean/fiber-ocean/ent/staff_role"
)

type Repository struct {
	ctx context.Context
}

func NewRepository() *Repository {
	return &Repository{
		ctx: context.Background(),
	}
}

func (r *Repository) Create(e *ent.Menu, c *ent.Client) error {
	err := c.Menu.Create().
		SetName(e.Name).
		SetNillableParentID(&e.ParentID).
		SetIcon(e.Icon).
		SetPath(e.Path).
		SetPermission(e.Permission).
		SetComponent(e.Component).
		SetType(e.Type).
		SetMethod(e.Method).
		SetVisible(e.Visible).
		SetSort(e.Sort).
		SetStatus(e.Status).
		SetRemark(e.Remark).
		Exec(r.ctx)
	switch {
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "name") {
			return errors.New(ErrorNameRepeat)
		}
	case err != nil:
		return errors.New(CreatedFail)
	}
	return err
}

func (r *Repository) Update(e *ent.Menu, c *ent.Client) error {
	b := c.Menu.Update().Where(menu.IDEQ(e.ID))
	if len(e.Name) > 0 {
		b.SetName(e.Name)
	}
	if len(e.Icon) > 0 {
		b.SetIcon(e.Icon)
	}
	if len(e.Path) > 0 {
		b.SetPath(e.Path)
	}
	if len(e.Permission) > 0 {
		b.SetPermission(e.Permission)
	}
	if len(e.Component) > 0 {
		b.SetComponent(e.Component)
	}
	if len(e.Type) > 0 {
		b.SetType(e.Type)
	}
	if e.Sort > 0 {
		b.SetSort(e.Sort)
	}
	if len(e.Status) > 0 {
		b.SetStatus(e.Status)
	}
	if len(e.Remark) > 0 {
		b.SetRemark(e.Remark)
	}
	b.SetVisible(e.Visible)
	b.SetParentID(e.ParentID)
	err := b.Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "name") {
			return errors.New(ErrorNameRepeat)
		}
	case err != nil:
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *Repository) Delete(e *ent.Menu, c *ent.Client) error {
	num, err := c.Menu.Delete().Where(menu.IDEQ(e.ID)).Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(w *WhereParams, c *ent.Client) (bool, error) {
	b := c.Menu.Query()
	if len(w.ParentID) > 0 {
		b = b.Where(menu.ParentIDEQ(w.ParentID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err
}

func (r *Repository) QueryAll(w *WhereParams, c *ent.Client) ([]*ent.Menu, error) {
	b := c.Menu.Query()
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(menu.NameContains(strings.TrimSpace(w.Name)))
	}
	if len(strings.TrimSpace(w.Status)) > 0 {
		b = b.Where(menu.StatusEQ(strings.TrimSpace(w.Status)))
	}
	if len(strings.TrimSpace(w.Remark)) > 0 {
		b = b.Where(menu.StatusEQ(strings.TrimSpace(w.Remark)))
	}
	b.Select(SelectFields...)
	es, err := b.Order(menu.BySort()).All(r.ctx)
	return es, err
}

func (r *Repository) QueryByStaffID(id string, c *ent.Client) ([]*ent.Menu, error) {
	es, err := c.Menu.Query().
		Select(SelectFields...).
		Where(func(s *sql.Selector) {
			rm := sql.Table(role_menu.Table)
			sr := sql.Table(staff_role.Table)
			s.Join(rm).On(s.C(menu.FieldID), rm.C(role_menu.FieldMenuID)).
				Join(sr).On(rm.C(role_menu.FieldRoleID), sr.C(staff_role.FieldRoleID)).
				Where(sql.EQ(sr.C(staff_role.FieldStaffID), id))
		}).
		Order(menu.BySort()).
		All(r.ctx)
	return es, err
}
