package repository

import (
	"database/sql"
	"strings"

	"aah-form-based-auth/app/models"

	aah "aahframe.work"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var TestLoginFailUserName = "fail"

func (repo *SormRepository) GetAllUsers() ([]*models.User, error) {
	ul, err := models.Users(
		qm.Load("Roles"),
		qm.Load("Permissions"),
		qm.OrderBy("email"),
	).All(repo.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return ul, nil
}

func (repo *SormRepository) GetUserByID(id string) (*models.User, error) {
	u, err := models.Users(models.UserWhere.ID.EQ(id)).One(repo.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *SormRepository) GetUserByEmail(email string) (*models.User, error) {
	u, err := models.Users(
		models.UserWhere.Email.EQ(strings.TrimSpace(email)),
		qm.Load(models.UserRels.Roles),
		qm.Load(models.UserRels.Permissions),
	).One(repo.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *SormRepository) CreateUser(u *models.User) (err error) {
	u.ID = uuid.New().String()
	err = u.Insert(repo.db, boil.Infer())
	return
}

func (repo *SormRepository) UpdateUser(email string, args UpdateUserArgs) error {
	var (
		changed bool
	)

	u, err := models.Users(
		models.UserWhere.Email.EQ(strings.TrimSpace(email)),
		qm.Load(models.UserRels.Roles),
		qm.Load(models.UserRels.Permissions),
	).One(repo.db)

	if err != nil {
		return err
	}

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	changes := []string{}
	if args.FirstName != u.FirstName {
		u.FirstName = args.FirstName
		changes = append(changes, "first_name")
	}
	if args.LastName != u.LastName {
		u.LastName = args.LastName
		changes = append(changes, "last_name")
	}
	if args.Email != u.Email {
		u.Email = args.Email
		changes = append(changes, "email")
	}
	if args.IsLocked != u.IsLocked.Bool {
		u.IsLocked = null.BoolFrom(args.IsLocked)
		changes = append(changes, "is_locked")
	}

	if len(changes) > 0 {
		if _, err = u.Update(repo.db, boil.Whitelist(changes...)); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		changed = true
	}

	// clear roles
	if len(args.Roles) == 0 {
		for _, r := range u.R.Roles {
			err = u.RemoveRoles(repo.db, r)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
		}
	}

	// clear permissions
	if len(args.Permissions) == 0 {
		for _, p := range u.R.Permissions {
			err = u.RemovePermissions(repo.db, p)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
		}
	}

	if len(args.Roles) > 0 {
		for _, r := range args.Roles {
			// existed in user's relative roles?
			find := false
			for _, v := range u.R.Roles {
				if v.Name.String == r {
					find = true
				} else {
					err = u.RemoveRoles(repo.db, v)
					if err != nil {
						if err := tx.Rollback(); err != nil {
							return err
						}
						return err
					}
				}
			}
			if find {
				continue
			}
			// role name existed in role table?
			r1, err := models.Roles(qm.Where("name = ?", r)).One(repo.db)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
			if r1 != nil {
				err = u.AddRoles(repo.db, false, r1)
				if err != nil {
					if err := tx.Rollback(); err != nil {
						return err
					}
					return err
				}
				changed = true
			}
			// common user has no permission to add a role
			aah.App().Log().Debugf("user %s tries to add a role %s", u.Email, r)
		}
	}

	if len(args.Permissions) > 0 {
		for _, p := range args.Permissions {
			// existed in user's relative permission?
			find := false
			for _, v := range u.R.Permissions {
				if v.Name.String == p {
					find = true
				}
			}
			if find {
				continue
			}
			// permission name existed in permission table?
			p1, err := models.Permissions(qm.Where("name = ?", p)).One(repo.db)
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
			if p1 != nil {
				err = u.AddPermissions(repo.db, false, p1)
				if err != nil {
					if err := tx.Rollback(); err != nil {
						return err
					}
					return err
				}
				changed = true
			}
			// common user has no permission to add a permission
			aah.App().Log().Debugf("user %s tries to add a role %s", u.Email, p)
		}
	}

	u.Reload(repo.db)
	// Commit
	if err := tx.Commit(); err != nil {
		return err
	}

	if changed {
		aah.App().PublishEvent("UserEdited", &UserEvent{
			UserEmail: u.Email,
		})
	}

	return nil
}

func (repo *SormRepository) Exist(email string) (exist bool, err error) {
	exist, err = models.Users(qm.Where("email = ?", email)).Exists(repo.db)
	return
}
