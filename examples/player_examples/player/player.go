package player

type Player struct {
	Nick  string
	Score int
}

//go:generate hasgo -T=Player -S=Players
type Players []Player
