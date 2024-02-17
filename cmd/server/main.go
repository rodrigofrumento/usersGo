package main

import (
	"log/slog"

	"github.com/rodrigofrumento/usersGo/config/logger"
)

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func (u user) LodUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "HIDDEN"),
	)
}

func main() {
	logger.InitLogger()

	user := user{
		Name:     "John Doe",
		Age:      33,
		Password: "123456",
	}

	slog.Info("starting API")
	slog.Info("creating user", "user", user.LodUser())
}
