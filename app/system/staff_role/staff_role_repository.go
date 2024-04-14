package staff_role

import (
	"context"
	"errors"

	"github.com/golden-ocean/fiber-ocean/ent"
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

func (r *Repository) Create(req *CreateInput, c *ent.Client) error {
	s_id := req.StaffID
	r_ids := req.RoleIDs
	err := c.Staff_Role.MapCreateBulk(r_ids, func(c *ent.StaffRoleCreate, i int) {
		c.SetRolesID(r_ids[i]).SetStaffsID(s_id)
	}).Exec(r.ctx)
	return err
}

func (r *Repository) Delete(req *DeleteInput, c *ent.Client) error {
	s_id := req.StaffID
	r_ids := req.RoleIDs
	b := c.Staff_Role.Delete()
	if len(req.StaffID) > 0 {
		b = b.Where(staff_role.StaffIDEQ(s_id))
	}
	if len(req.RoleIDs) > 0 {
		b = b.Where(staff_role.RoleIDIn(r_ids...))
	}
	num, err := b.Exec(r.ctx)
	switch {
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(e *ent.Staff_Role, c *ent.Client) (bool, error) {
	b := c.Staff_Role.Query()
	if len(e.StaffID) > 0 {
		b = b.Where(staff_role.StaffIDEQ(e.StaffID))
	}
	if len(e.RoleID) > 0 {
		b = b.Where(staff_role.RoleIDEQ(e.RoleID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err
}

func (r *Repository) Query(w *WhereParams, c *ent.Client) ([]*ent.Staff_Role, error) {
	b := c.Staff_Role.Query()
	if len(w.StaffID) > 0 {
		b = b.Where(staff_role.StaffIDEQ(w.StaffID))
	}
	if len(w.StaffIDs) > 0 {
		b = b.Where(staff_role.StaffIDIn(w.StaffIDs...))
	}
	if len(w.RoleID) > 0 {
		b = b.Where(staff_role.RoleIDEQ(w.RoleID))
	}
	es, err := b.Select(staff_role.FieldStaffID, staff_role.FieldRoleID).All(r.ctx)
	if err != nil {
		return nil, err
	}
	return es, err
}
