package organization

import (
	"context"
	"errors"
	"strings"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/organization"
)

type Repository struct {
	ctx context.Context
}

func NewRepository() *Repository {
	return &Repository{
		ctx: context.Background(),
	}
}

func (r *Repository) Create(e *ent.Organization, c *ent.Client) error {
	err := c.Organization.Create().
		SetName(e.Name).
		SetCode(e.Code).
		SetNillableParentID(&e.ParentID).
		SetSort(e.Sort).
		SetStatus(e.Status).
		SetRemark(e.Remark).
		Exec(r.ctx)
	switch {
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "name") {
			return errors.New(ErrorNameRepeat)
		}
		if strings.Contains(err.Error(), "code") {
			return errors.New(ErrorCodeRepeat)
		}
	case err != nil:
		return errors.New(CreatedFail)
	}
	return err
}

func (r *Repository) Update(e *ent.Organization, c *ent.Client) error {
	// if !xid.ID.IsNil(e.ParentID) {
	// 	_, err := c.Organization.Query().Select(organization.FieldID).Where(organization.IDEQ(e.ParentID)).First(r.ctx)
	// 	if ent.IsNotFound(err) {
	// 		return errors.New(ErrorPidNotExist)
	// 	}
	// }
	b := c.Organization.Update().Where(organization.IDEQ(e.ID))
	if len(e.Name) > 0 {
		b.SetName(e.Name)
	}
	if len(e.Code) > 0 {
		b.SetCode(e.Code)
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
	b.SetParentID(e.ParentID)
	err := b.Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "name") {
			return errors.New(ErrorNameRepeat)
		}
		if strings.Contains(err.Error(), "code") {
			return errors.New(ErrorCodeRepeat)
		}
	case err != nil:
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *Repository) Delete(e *ent.Organization, c *ent.Client) error {
	num, err := c.Organization.Delete().Where(organization.IDEQ(e.ID)).Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(w *WhereParams, c *ent.Client) (bool, error) {
	b := c.Organization.Query()
	if len(w.ID) > 0 {
		b = b.Where(organization.IDEQ(w.ID))
	}
	if len(w.ParentID) > 0 {
		b = b.Where(organization.ParentIDEQ(w.ParentID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err
}

func (r *Repository) QueryAll(w *WhereParams, c *ent.Client) ([]*ent.Organization, error) {
	b := c.Organization.Query()
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(organization.NameContainsFold(strings.TrimSpace(w.Name)))
	}
	if len(strings.TrimSpace(w.Code)) > 0 {
		b = b.Where(organization.CodeContainsFold(strings.TrimSpace(w.Code)))
	}
	if len(strings.TrimSpace(w.Status)) > 0 {
		b = b.Where(organization.StatusEqualFold(strings.TrimSpace(w.Status)))
	}
	if len(strings.TrimSpace(w.Remark)) > 0 {
		b = b.Where(organization.RemarkContainsFold(strings.TrimSpace(w.Remark)))
	}
	b.Select(SelectFields...)
	es, err := b.Order(organization.BySort()).All(r.ctx)
	return es, err
}
