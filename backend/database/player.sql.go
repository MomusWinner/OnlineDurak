// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: player.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPlayer = `-- name: CreatePlayer :one
insert into player (name, age, rating)
values ($1, $2, 0)
returning id
`

type CreatePlayerParams struct {
	Name string
	Age  int16
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createPlayer, arg.Name, arg.Age)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}
