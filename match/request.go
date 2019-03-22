package match

type GameRequest struct {
	playerId  int
	gameType  GameType
	createdAt int
	active    bool
}

type GameType struct {
	Name      string
	PlayerPot int
}

func NewRequest(id int, t GameType, time int, active bool) GameRequest {
	gr := GameRequest{playerId: id, gameType: t, active: active, createdAt: time}
	return gr
}
