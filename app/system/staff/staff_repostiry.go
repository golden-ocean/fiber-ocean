package staff

import (
	"context"
	"errors"
	"strings"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/staff"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
	"github.com/golden-ocean/fiber-ocean/ent/staff_role"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
)

type Repository struct {
	ctx context.Context
}

func NewRepository() *Repository {
	return &Repository{
		ctx: context.Background(),
	}
}

func (r *Repository) Create(e *ent.Staff, c *ent.Client) (*ent.Staff, error) {
	re, err := c.Staff.Create().
		SetUsername(e.Username).
		SetPassword(utils.GeneratePassword(e.Password)).
		SetName(e.Name).
		SetEmail(e.Email).
		SetMobile(e.Mobile).
		SetGender(e.Gender).
		SetAvatar(e.Avatar).
		SetWorkStatus(e.WorkStatus).
		SetSort(e.Sort).
		SetStatus(e.Status).
		SetRemark(e.Remark).
		SetNillableOrganizationID(&e.OrganizationID).
		Save(r.ctx)
	switch {
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "username") {
			return nil, errors.New(ErrorUsernameRepeat)
		}
		if strings.Contains(err.Error(), "email") {
			return nil, errors.New(ErrorEmailRepeat)
		}
		if strings.Contains(err.Error(), "mobile") {
			return nil, errors.New(ErrorMobileRepeat)
		}
	case err != nil:
		return nil, errors.New(CreatedFail)
	}
	return re, err
}

func (r *Repository) Update(e *ent.Staff, c *ent.Client) error {
	b := c.Staff.Update().Where(staff.IDEQ(e.ID))
	if len(e.Username) > 0 {
		b.SetUsername(e.Username)
	}
	if len(e.Password) > 0 {
		b.SetPassword(utils.GeneratePassword(e.Password))
	}
	if len(e.Name) > 0 {
		b.SetName(e.Name)
	}
	if len(e.Email) > 0 {
		b.SetEmail(e.Email)
	}
	if len(e.Mobile) > 0 {
		b.SetMobile(e.Mobile)
	}
	if len(e.Gender) > 0 {
		b.SetGender(e.Gender)
	}
	if len(e.Avatar) > 0 {
		b.SetAvatar(e.Avatar)
	}
	if len(e.WorkStatus) > 0 {
		b.SetWorkStatus(e.WorkStatus)
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
	if len(e.OrganizationID) > 0 {
		b.SetOrganizationID(e.OrganizationID)
	}
	err := b.Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case ent.IsConstraintError(err):
		if strings.Contains(err.Error(), "username") {
			return errors.New(ErrorUsernameRepeat)
		}
		if strings.Contains(err.Error(), "email") {
			return errors.New(ErrorEmailRepeat)
		}
		if strings.Contains(err.Error(), "mobile") {
			return errors.New(ErrorMobileRepeat)
		}
	case err != nil:
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *Repository) Delete(e *ent.Staff, c *ent.Client) error {
	num, err := c.Staff.Delete().Where(staff.IDEQ(e.ID)).Exec(r.ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.New(ErrorNotExist)
	case err != nil || num == 0:
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) Exist(w *WhereParams, c *ent.Client) (bool, error) {
	b := c.Staff.Query()
	if len(w.OrganizationID) > 0 {
		b = b.Where(staff.OrganizationIDEQ(w.OrganizationID))
	}
	exist, err := b.Exist(r.ctx)
	return exist, err
}

func (r *Repository) QueryByUniqueField(w *WhereParams, c *ent.Client) (*ent.Staff, error) {
	b := c.Staff.Query()
	if len(strings.TrimSpace(w.Username)) > 0 {
		b = b.Where(staff.UsernameContains(w.Username))
	}
	if len(strings.TrimSpace(w.Email)) > 0 {
		b = b.Where(staff.EmailContains(w.Email))
	}
	if len(strings.TrimSpace(w.Mobile)) > 0 {
		b = b.Where(staff.MobileContains(w.Mobile))
	}
	e, err := b.First(r.ctx)
	return e, err
}

func (r *Repository) QueryPage(w *WhereParams, c *ent.Client) ([]*ent.Staff, int, error) {
	b := c.Staff.Query()
	if len(strings.TrimSpace(w.Username)) > 0 {
		b = b.Where(staff.UsernameContains(w.Username))
	}
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(staff.NameContains(w.Name))
	}
	if len(w.OrganizationID) > 0 {
		b = b.Where(staff.OrganizationID(w.OrganizationID))
	}
	if len(strings.TrimSpace(w.Email)) > 0 {
		b = b.Where(staff.EmailContains(w.Email))
	}
	if len(strings.TrimSpace(w.Mobile)) > 0 {
		b = b.Where(staff.MobileContains(w.Mobile))
	}
	if len(strings.TrimSpace(w.Gender)) > 0 {
		b = b.Where(staff.GenderEqualFold(w.Gender))
	}
	if len(strings.TrimSpace(w.WorkStatus)) > 0 {
		b = b.Where(staff.WorkStatusEqualFold(w.WorkStatus))
	}
	if len(strings.TrimSpace(w.Status)) > 0 {
		b = b.Where(staff.StatusEQ(w.Status))
	}
	if len(strings.TrimSpace(w.Remark)) > 0 {
		b = b.Where(staff.RemarkContains(w.Remark))
	}

	total, err := b.Count(r.ctx)
	if err != nil {
		return nil, 0, err
	}
	es, err := b.Order(staff.BySort()).
		WithPositions(func(pq *ent.StaffPositionQuery) {
			pq.Select(staff_position.FieldPositionID)
		}).
		WithRoles(func(rq *ent.StaffRoleQuery) {
			rq.Select(staff_role.FieldRoleID)
		}).
		Limit(w.PageSize).
		Offset((w.Current - 1) * w.PageSize).
		All(r.ctx)

	return es, total, err
}
