package match

type Coordinator struct {
	queue []GameRequest
}

func InitCoordinator() Coordinator {
	co := Coordinator{queue: []GameRequest{}}
	return co
}

func (c *Coordinator) Add(gr GameRequest) {
	c.queue = append(c.queue, gr)
}
