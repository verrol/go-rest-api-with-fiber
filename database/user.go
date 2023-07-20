package database

import (
	"github.com/charmbracelet/log"
	"xorm.io/xorm"
)

type UserInfo struct {
	Id       string `xorm:"pk"`
	Username string
	Password string
	Roles    string // comma separated roles
	Created  string `xorm:"created"`
	Updated  string `xorm:"updated"`
	Deleted  string `xorm:"deleted"`
}

func CreateUserTable(db *xorm.Engine) error {
	if err := db.Sync(new(UserInfo)); err != nil {
		log.Error("error creating user table", "table", "user", "error", err)
	}

	log.Info("user table created")
	return nil
}
