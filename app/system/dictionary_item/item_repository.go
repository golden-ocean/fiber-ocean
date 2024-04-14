package dictionary_item

import (
	"context"
	"errors"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/dictionary"
	"github.com/golden-ocean/fiber-ocean/ent/dictionary_item"
)

type Repository struct {
	ctx context.Context
}

func NewRepository() *Repository {
	return &Repository{
		ctx: context.Background(),
	}
}

func (r *Repository) Create(e *ent.Dictionary_Item, c *ent.Client) error {
	err := c.Dictionary_Item.Create().
		SetDictionaryID(e.DictionaryID).
		SetLabel(e.Label).
		SetValue(e.Value).
		SetColor(e.Color).
		SetSort(e.Sort).
		SetStatus(e.Status).
		SetRemark(e.Remark).
		Exec(r.ctx)
	switch {
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "label") {
			return errors.New(ErrorLabelRepeat)
		}
		if strings.Contains(err.Error(), "value") {
			return errors.New(ErrorValueRepeat)
		}
	case err != nil:
		return errors.New(CreatedFail)
	}
	return err
}

func (r *Repository) Update(e *ent.Dictionary_Item, c *ent.Client) error {
	b := c.Dictionary_Item.Update().Where(dictionary_item.ID(e.ID))
	if len(e.Label) > 0 {
		b.SetLabel(e.Label)
	}
	if len(e.Value) > 0 {
		b.SetValue(e.Value)
	}
	if len(e.Color) > 0 {
		b.SetColor(e.Color)
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
		if strings.Contains(err.Error(), "label") {
			return errors.New(ErrorLabelRepeat)
		}
		if strings.Contains(err.Error(), "value") {
			return errors.New(ErrorValueRepeat)
		}
	case err != nil:
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *Repository) Delete(e *ent.Dictionary_Item, c *ent.Client) error {
	num, err := c.Dictionary_Item.Delete().Where(dictionary_item.IDEQ(e.ID)).Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(w *WhereParams, c *ent.Client) (bool, error) {
	b := c.Dictionary_Item.Query()
	if len(w.DictionaryID) > 0 {
		b = b.Where(dictionary_item.DictionaryIDEQ(w.DictionaryID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err

}

func (r *Repository) QueryPage(w *WhereParams, c *ent.Client) ([]*ent.Dictionary_Item, int, error) {
	b := c.Dictionary_Item.Query().Where(dictionary_item.DictionaryIDEQ(w.DictionaryID))
	if len(strings.TrimSpace(w.Label)) > 0 {
		b = b.Where(dictionary_item.LabelContainsFold(strings.TrimSpace(w.Label)))
	}
	if len(strings.TrimSpace(w.Value)) > 0 {
		b = b.Where(dictionary_item.ValueContainsFold(strings.TrimSpace(w.Value)))
	}
	if len(strings.TrimSpace(w.Status)) > 0 {
		b = b.Where(dictionary_item.StatusEqualFold(strings.TrimSpace(w.Status)))
	}
	if len(strings.TrimSpace(w.Remark)) > 0 {
		b = b.Where(dictionary_item.RemarkContainsFold(strings.TrimSpace(w.Remark)))
	}
	total, err := b.Count(r.ctx)
	if err != nil {
		return nil, 0, err
	}
	b.Select(SelectFields...)
	entities, err := b.Order(dictionary_item.BySort()).
		Limit(w.PageSize).
		Offset((w.Current - 1) * w.PageSize).
		All(r.ctx)
	return entities, total, err
}

func (r *Repository) QueryByCode(code string, c *ent.Client) ([]*ent.Dictionary_Item, error) {
	es, err := c.Dictionary_Item.Query().
		Select(SelectFields...).
		Where(func(s *sql.Selector) {
			d := sql.Table(dictionary.Table)
			s.Join(d).On(s.C(dictionary_item.FieldDictionaryID), d.C(dictionary.FieldID))
			s.Where(sql.EQ(d.C(dictionary.FieldCode), code))
		}).
		Order(dictionary_item.BySort()).
		All(r.ctx)
	return es, err
}
