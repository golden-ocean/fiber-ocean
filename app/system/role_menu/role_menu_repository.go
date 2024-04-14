package role_menu

import (
	"context"
	"errors"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/role_menu"
)

type Repository struct {
	ctx context.Context
}

func NewRepository() *Repository {
	return &Repository{
		ctx: context.Background(),
	}
}

func (r *Repository) Create(req *CreateInput, c *ent.Client) error {
	r_id := req.RoleID
	m_ids := req.MenuIDs
	err := c.Role_Menu.MapCreateBulk(m_ids, func(c *ent.RoleMenuCreate, i int) {
		c.SetMenusID(m_ids[i]).SetRolesID(r_id)
	}).Exec(r.ctx)
	return err
}

func (r *Repository) Delete(req *DeleteInput, c *ent.Client) error {
	r_id := req.RoleID
	m_ids := req.MenuIDs
	num, err := c.Role_Menu.Delete().Where(role_menu.RoleIDEQ(r_id), role_menu.MenuIDIn(m_ids...)).Exec(r.ctx)
	switch {
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(e *ent.Role_Menu, c *ent.Client) (bool, error) {
	b := c.Role_Menu.Query()
	if len(e.RoleID) > 0 {
		b = b.Where(role_menu.RoleIDEQ(e.RoleID))
	}
	if len(e.MenuID) > 0 {
		b = b.Where(role_menu.MenuIDEQ(e.MenuID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err
}

func (r *Repository) Query(w *WhereParams, c *ent.Client) ([]*ent.Role_Menu, error) {
	b := c.Role_Menu.Query()
	if len(w.RoleID) > 0 {
		b = b.Where(role_menu.RoleIDEQ(w.RoleID))
	}
	if len(w.RoleIDs) > 0 {
		b = b.Where(role_menu.RoleIDIn(w.RoleIDs...))
	}
	es, err := b.Select(role_menu.FieldRoleID, role_menu.FieldMenuID).All(r.ctx)
	if err != nil {
		return nil, err
	}
	return es, err
}
