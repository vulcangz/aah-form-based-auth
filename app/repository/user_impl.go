package repository

import (
	"aah-form-based-auth/app/models"
	"strings"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var TestLoginFailUserName = "fail"

func (repo *SormRepository) GetUserByID(id string) (*models.User, error) {
	u, err := models.Users(models.UserWhere.ID.EQ(id)).One(repo.db)
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
	if err != nil {
		return nil, err
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
