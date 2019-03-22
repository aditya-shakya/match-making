package match

import (
	"fmt"
	"sort"
)

type Coordinator struct {
	queue []GameRequest
	games []Game
}

type byRank []GameRequest

var gameRequests []GameRequest

func (gameRequests byRank) Len() int { return len(gameRequests) }
func (gameRequests byRank) Swap(i, j int) {
	gameRequests[i], gameRequests[j] = gameRequests[j], gameRequests[i]
}
func (gameRequests byRank) Less(i, j int) bool {
	// if gameRequests[i].rank < gameRequests[j].rank {
	// 	return true
	// }
	// if gameRequests[i].rank > gameRequests[j].rank {
	// 	return false
	// }
	return gameRequests[i].rank > gameRequests[j].rank
}

func InitCoordinator() Coordinator {
	co := Coordinator{queue: []GameRequest{}}
	return co
}

func (c *Coordinator) Add(gr GameRequest) {
	c.queue = append(c.queue, gr)
}

func (c *Coordinator) AddGame(g Game) {
	c.games = append(c.games, g)
}

func (c *Coordinator) Filter(n int, time int) []GameRequest {
	fmt.Println("filter Requests")
	fmt.Print("Time: ")
	fmt.Println(time)
	requests := []GameRequest{}
	for i := 0; i < len(c.queue); i++ {
		request := c.queue[i]
		if request.active == true && request.createdAt <= time {
			requests = append(requests, request)
		}
	}
	sort.Sort(byRank(requests))
	fmt.Printf("filtered Queue: %v\n", requests)

	return requests
}

func (c *Coordinator) ChooseGameSet(n int, time int, requests []GameRequest) {
	fmt.Println("Choosing gameSet")
	fmt.Print("Time: ")
	fmt.Println(time)
	requests = c.Filter(n, time)
	if len(requests) < n {
		fmt.Println("smallers size")
		return
	} else {
		fmt.Println("looking for valid")
		i := 0
		for i < len(requests)-n+1 {
			fmt.Println(i)
			valid := true
			for j := 0; j < n; j++ {
				diff1 := requests[i].rank - requests[i+j].rank
				diff2 := requests[i+j].rank - requests[i+n-1].rank
				th := requests[i+j].Threshold(time)
				fmt.Printf("i=%d j=%d diff1=%v diff2=%v th=%v\n", i, j, diff1, diff2, th)
				if diff1 > th || diff2 > th {
					valid = false
					break
				}
			}
			fmt.Println("i: ", valid)
			if valid {
				group := Group{}
				for j := 0; j < n; j++ {
					group.Add(requests[i+j].playerId)
					c.CloseRequest(requests[i+j])
				}
				fmt.Println("Group: ", group)

				teams := group.Distribute("1v1")
				game := Game{teams: teams, createdAt: time}
				c.AddGame(game)
				i = i + n
			} else {
				i = i + 1
			}
		}
	}

	fmt.Print("Games: ")
	fmt.Println(c.games)

}

func (c *Coordinator) SimulateResult(n int) {
	time := 0
	started := false
	for true {
		reqs := c.Filter(n, time)
		if (len(c.queue)-len(c.games)*n) < n && started == true {
			break
		} else {
			started = true
			c.ChooseGameSet(n, time, reqs)
			time = time + 1
		}
	}
}

func (c *Coordinator) CloseRequest(gr GameRequest) {
	for i := 0; i < len(c.queue); i++ {
		if c.queue[i] == gr {
			c.queue[i].active = false
			break
		}
	}
}
