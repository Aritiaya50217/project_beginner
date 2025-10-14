package proxy

type RoundRobin struct {
	Servers []string
	Index   int
}

func (roundRobin *RoundRobin) Next() string {
	if len(roundRobin.Servers) == 0 {
		return ""
	}
	server := roundRobin.Servers[roundRobin.Index]
	roundRobin.Index = (roundRobin.Index + 1) % len(roundRobin.Servers)
	return server
}
