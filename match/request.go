package match

import (
	"time"
)

type GameRequest struct {
	playerId  int
	gameType  GameType
	createdAt time.Time
}

type GameType struct {
	Name      string
	PlayerPot int
}
