package models

type Information struct {
	ID      int
	Name    string
	Balance float32
	Level   int
	Avatar  string
}
type Achievement struct {
	ID         int
	Win        int
	Lose       int
	Num_Dragon int
	Num_token  int
}

type Player struct {
	ID int
	Information
	Achievement
	Friends map[int]Friend
}

type Friend struct {
	ID     int
	Name   string
	Status string // e.g., "online", "offline"
}

type Message struct {
	ID_player1 int
	ID_player2 int
	Message    string
}
