// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: match.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addPlayerPlacement = `-- name: AddPlayerPlacement :one
insert into player_placement (match_result_id, player_id, player_place, rating_change)
values ($1, $2, $3, $4)
returning match_result_id, player_id, player_place, rating_change
`

type AddPlayerPlacementParams struct {
	MatchResultID pgtype.UUID
	PlayerID      pgtype.UUID
	PlayerPlace   int16
	RatingChange  int32
}

func (q *Queries) AddPlayerPlacement(ctx context.Context, arg AddPlayerPlacementParams) (PlayerPlacement, error) {
	row := q.db.QueryRow(ctx, addPlayerPlacement,
		arg.MatchResultID,
		arg.PlayerID,
		arg.PlayerPlace,
		arg.RatingChange,
	)
	var i PlayerPlacement
	err := row.Scan(
		&i.MatchResultID,
		&i.PlayerID,
		&i.PlayerPlace,
		&i.RatingChange,
	)
	return i, err
}

const createMatchResult = `-- name: CreateMatchResult :one
insert into match_result (player_count, game_result)
values ($1, $2)
returning id, player_count, game_result
`

type CreateMatchResultParams struct {
	PlayerCount int16
	GameResult  GameResult
}

func (q *Queries) CreateMatchResult(ctx context.Context, arg CreateMatchResultParams) (MatchResult, error) {
	row := q.db.QueryRow(ctx, createMatchResult, arg.PlayerCount, arg.GameResult)
	var i MatchResult
	err := row.Scan(&i.ID, &i.PlayerCount, &i.GameResult)
	return i, err
}
