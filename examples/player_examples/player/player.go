package player

// Player represents a person in an online game.
type Player struct {
	Nick  string
	Score int
}

// Players is a wrapper for the []Player slice.
// go:generate hasgo -T=Player -S=Players
type Players []Player
