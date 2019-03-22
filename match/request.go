package match

type GameRequest struct {
	playerId  int
	gameType  GameType
	createdAt int
	active    bool
	rank      int
}

type Game struct {
	teams     []Group
	createdAt int
}

type Group struct {
	players []int
}

type GameType struct {
	Name      string
	PlayerPot int
}

func NewRequest(id int, t GameType, time int, active bool, rank int) GameRequest {
	gr := GameRequest{playerId: id, gameType: t, active: active, createdAt: time, rank: rank}
	return gr
}

func (gr *GameRequest) Threshold(time int) int {
	tm := time - gr.createdAt
	if tm > 5 {
		return 3
	} else {
		return time
	}
}

func (g *Group) Add(player int) {
	g.players = append(g.players, player)
}

func (g *Group) Distribute(strategy string) []Group {
	if strategy == "1v1" {
		return g.Distribute1v1()
	} else {
		return []Group{}
	}
}

func (g *Group) Distribute1v1() []Group {
	groups := []Group{}
	group1 := Group{}
	group2 := Group{}

	for i := 0; i < len(g.players); i++ {
		if i%2 == 0 {
			group1.Add(g.players[i])
		} else {
			group2.Add(g.players[i])
		}
	}
	groups = append(groups, group1)
	groups = append(groups, group2)
	return groups
}
