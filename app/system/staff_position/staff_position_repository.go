package staff_position

import (
	"context"
	"errors"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
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
	p_ids := req.PositionIDs
	err := c.Staff_Position.MapCreateBulk(p_ids, func(c *ent.StaffPositionCreate, i int) {
		c.SetPositionsID(p_ids[i]).SetStaffsID(s_id)
	}).Exec(r.ctx)
	return err
}

func (r *Repository) Delete(req *DeleteInput, c *ent.Client) error {
	b := c.Staff_Position.Delete()
	if len(req.StaffID) > 0 {
		b = b.Where(staff_position.StaffIDEQ(req.StaffID))
	}
	if len(req.PositionIDs) > 0 {
		b = b.Where(staff_position.PositionIDIn(req.PositionIDs...))
	}
	num, err := b.Exec(r.ctx)
	switch {
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(e *ent.Staff_Position, c *ent.Client) (bool, error) {
	b := c.Staff_Position.Query()
	if len(e.StaffID) > 0 {
		b = b.Where(staff_position.StaffIDEQ(e.StaffID))
	}
	if len(e.PositionID) > 0 {
		b = b.Where(staff_position.PositionIDEQ(e.PositionID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err
}

func (r *Repository) Query(w *WhereParams, c *ent.Client) ([]*ent.Staff_Position, error) {
	b := c.Staff_Position.Query()
	if len(w.StaffID) > 0 {
		b = b.Where(staff_position.StaffIDEQ(w.StaffID))
	}
	if len(w.StaffIDs) > 0 {
		b = b.Where(staff_position.StaffIDIn(w.StaffIDs...))
	}
	if len(w.PositionID) > 0 {
		b = b.Where(staff_position.PositionIDEQ(w.PositionID))
	}
	es, err := b.Select(staff_position.FieldStaffID, staff_position.FieldPositionID).All(r.ctx)
	if err != nil {
		return nil, err
	}
	return es, err
}
