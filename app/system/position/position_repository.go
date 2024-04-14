package position

import (
	"context"
	"errors"
	"strings"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/position"
)

type Repository struct {
	ctx context.Context
}

func NewRepository() *Repository {
	return &Repository{
		ctx: context.Background(),
	}
}

func (r *Repository) Create(e *ent.Position, c *ent.Client) error {
	err := c.Position.Create().
		SetName(e.Name).
		SetCode(e.Code).
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

func (r *Repository) Update(e *ent.Position, c *ent.Client) error {
	b := c.Position.Update().Where(position.IDEQ(e.ID))
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

func (r *Repository) Delete(e *ent.Position, c *ent.Client) error {
	num, err := c.Position.Delete().Where(position.IDEQ(e.ID)).Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) QueryPage(w *WhereParams, c *ent.Client) ([]*ent.Position, int, error) {
	b := c.Position.Query()
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(position.NameContainsFold(strings.TrimSpace(w.Name)))
	}
	if len(strings.TrimSpace(w.Code)) > 0 {
		b = b.Where(position.CodeContainsFold(strings.TrimSpace(w.Code)))
	}
	if len(strings.TrimSpace(w.Status)) > 0 {
		b = b.Where(position.StatusEqualFold(strings.TrimSpace(w.Status)))
	}
	if len(strings.TrimSpace(w.Remark)) > 0 {
		b = b.Where(position.RemarkContainsFold(strings.TrimSpace(w.Remark)))
	}
	total, err := b.Count(r.ctx)
	if err != nil {
		return nil, 0, err
	}
	b.Select(SelectFields...)
	es, err := b.Order(position.BySort()).
		Limit(w.PageSize).
		Offset((w.Current - 1) * w.PageSize).
		All(r.ctx)
	return es, total, err
}

func (r *Repository) QueryAll(w *WhereParams, c *ent.Client) ([]*ent.Position, error) {
	b := c.Position.Query()
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(position.NameContains(strings.TrimSpace(w.Name)))
	}
	if len(strings.TrimSpace(w.Code)) > 0 {
		b = b.Where(position.CodeContains(strings.TrimSpace(w.Code)))
	}
	if len(strings.TrimSpace(w.Status)) > 0 {
		b = b.Where(position.StatusEQ(strings.TrimSpace(w.Status)))
	}
	if len(strings.TrimSpace(w.Remark)) > 0 {
		b = b.Where(position.RemarkContains(strings.TrimSpace(w.Remark)))
	}
	b.Select(SelectFields...)
	es, err := b.Order(position.BySort()).All(r.ctx)
	return es, err
}
