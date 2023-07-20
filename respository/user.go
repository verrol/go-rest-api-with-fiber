package respository

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/verrol/go-rest-api-with-fiber/database"
	"github.com/verrol/go-rest-api-with-fiber/model"
	"xorm.io/xorm"
)

// handle persistence database persistence by abtracting it away the sql DB details
type UserRepository interface {
	Get(id string) (model.UserInfo, error)
	GetByUsername(username string) (model.UserInfo, error)
	GetPassword(username string) (string, error)
	Create(username, password, roles string) (string, error)
	Delete(id string) error
}

type userRepoImpl struct {
	db *xorm.Engine
}

func NewUserRepository(db *xorm.Engine) UserRepository {
	return &userRepoImpl{db}
}

func (pr *userRepoImpl) Get(id string) (model.UserInfo, error) {
	user := model.UserInfo{}

	if has, err := pr.db.ID(id).Get(&user); err != nil {
		log.Info("error while getting user by id", "id", id, "err", err)
		return user, err
	} else if !has {
		log.Info("No record found for user by id", "id", id)
		return user, nil
	}

	return user, nil
}

func (pr *userRepoImpl) GetByUsername(username string) (model.UserInfo, error) {
	user := model.UserInfo{
		Username: username,
	}

	if has, err := pr.db.Get(&user); err != nil {
		log.Info("error while getting user by username", "username", username, "err", err)
		return user, err
	} else if !has {
		log.Info("No record found!", "username", username)
		return user, nil
	}

	return user, nil
}

func (pr *userRepoImpl) GetPassword(username string) (string, error) {
	user := database.UserInfo{
		Username: username,
	}

	if has, err := pr.db.Get(&user); err != nil {
		log.Info("error while getting user password", "username", username, "err", err)
		return "", err
	} else if !has {
		log.Info("No record found!", "username", username)
		return "", nil
	}

	return user.Password, nil
}

func (pr *userRepoImpl) Create(username, password, roles string) (string, error) {
	user := database.UserInfo{
		Id:       uuid.NewString(),
		Username: username,
		Password: password,
		Roles:    roles,
	}

	if _, err := pr.db.Insert(&user); err != nil {
		log.Error("Error while inserting user", "err", err)
		return "", err
	}

	return user.Id, nil
}

func (pr *userRepoImpl) Delete(id string) error {
	if _, err := pr.db.ID(id).Delete(new(model.UserInfo)); err != nil {
		log.Error("Error while deleting user", "err", err, "id", id)
		return err
	}

	return nil
}
