// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        int32
	Email     string
	Password  pgtype.Text
	Timezone  string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}