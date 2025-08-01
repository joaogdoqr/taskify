// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package store

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Task struct {
	ID          int32              `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Priority    int32              `json:"priority"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
