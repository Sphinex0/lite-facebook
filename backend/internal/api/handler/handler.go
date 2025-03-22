package handler

import (
	"database/sql"

	"social-network/internal/repository"
	"social-network/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(db *sql.DB) *Handler {
	userData := repository.Database{
		Db: db,
	}

	userService := service.Service{
		Database: &userData,
	}

	return &Handler{
		Service: &userService,
	}
}
