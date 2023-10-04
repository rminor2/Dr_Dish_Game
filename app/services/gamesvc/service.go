package gamesvc

import (
	"context"
	"errors"

	gamedb "dish/app/datastore"
)

type gameSvc struct {
	repo *gamedb.Queries
}

var gameInst *gameSvc

func init() {
	gameInst = &gameSvc{
		repo: gamedb.New(gamedb.RWInstance()),
	}
}

func GameSVC() *gameSvc {
	return gameInst
}

func (g *gameSvc) CreateGame(ctx context.Context, gameTitle string) (int32, error) {
	if len(gameTitle) <= 6 {
		return -1, errors.New("game title must be longer than 6 characters")
	}

	return g.repo.InsertGame(ctx, gameTitle)
}
