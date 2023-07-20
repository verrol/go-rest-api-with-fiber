package middleware

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func AuthReq() fiber.Handler {
	users := map[string]string{
		"admin": "admin",
		"user":  "user",
	}
	auth := basicauth.New(basicauth.Config{
		Users: users,
		Realm: "Authorization Required",
		Authorizer: func(user, pass string) bool {
			log.Info("verifying auth", "user", user)
			
			p, ok := users[user]
			if !ok {
				return false
			}
			return p == pass
		},
	})
	return auth
}
