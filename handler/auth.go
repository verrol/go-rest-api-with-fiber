package handler

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/verrol/go-rest-api-with-fiber/respository"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/xorm"
)

type AuthHandlers struct {
	repo respository.UserRepository
}

func NewAuthHandlers(db *xorm.Engine) *AuthHandlers {
	return &AuthHandlers{repo: respository.NewUserRepository(db)}
}

func (ah *AuthHandlers) Register(c *fiber.Ctx) error {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Roles    string `json:"roles"`
	}

	if err := c.BodyParser(&user); err != nil {
		log.Error("error parsing body")
		return c.SendStatus(http.StatusBadRequest)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error hashing password")
		return c.SendStatus(http.StatusInternalServerError)
	}

	id, err := ah.repo.Create(user.Username, string(hashedPassword), user.Roles)
	if err != nil {
		log.Error("error creating user")
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(id)
}

func (ah *AuthHandlers) Login(c *fiber.Ctx) error {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		log.Error("error parsing body")
		return c.SendStatus(http.StatusBadRequest)
	}

	hashedPassword, err := ah.repo.GetPassword(user.Username)
	if err != nil {
		log.Error("error getting password for user", "username", user.Username)
		return c.SendStatus(http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		log.Error("failed login attempted")
		return c.SendStatus(http.StatusBadGateway)
	}

	return c.JSON(map[string]string{"token": "token", "roles": "roles", "username": user.Username})
}
