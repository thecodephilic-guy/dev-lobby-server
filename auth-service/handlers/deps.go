package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// HandlerDeps holds all dependencies for handlers
type HandlerDeps struct {
	DB *pgxpool.Pool
}

// NewHandlerDeps creates a new handler dependencies instance
func NewHandlerDeps(db *pgxpool.Pool) *HandlerDeps {
	return &HandlerDeps{
		DB: db,
	}
}
