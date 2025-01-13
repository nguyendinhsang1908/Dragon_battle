package models

import "time"

type Information struct {
	ID        int
	Name      string
	Balance   float32
	Level     int
	Avatar    string
	Team_id   int
	Clan_name string
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
}

type Friend struct {
	ID     int
	Name   string
	Status string // e.g., "online", "offline"
}

type Message struct {
	ID          int       `json:"id"`           // ID tin nhắn
	SenderID    int       `json:"sender_id"`    // ID người gửi
	ReceiverID  *int      `json:"receiver_id"`  // ID người nhận (Friend chat)
	ClanID      *int      `json:"clan_id"`      // ID clan (Clan chat)
	TeamID      *int      `json:"team_id"`      // ID team (Team chat)
	MessageType string    `json:"message_type"` // Loại tin nhắn (world, clan, friend, team)
	Message     string    `json:"message"`      // Nội dung tin nhắn
	CreatedAt   time.Time `json:"created_at"`   // Thời gian tạo
}

type Clan struct {
	ID       int    // ID clan
	Name     string // Tên clan
	Level_cl int    // Cấp clan
	Num_mem  int    // Số lượng thành viên
	Own_id   int    // ID Chủ clan
}

type MessageRequest struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type Dragon struct {
	ID      int
	Name    string
	Point   int
	Hp      int
	Mp      int
	Dame    int
	Defense int
	Speed   int
	Rarity  string // Common,epic,legendary,immo
	Element string // Fire,Natural,Terra,Ice,Metal,Dark,Light,Chaos
}

type Item struct {
	//ItemID      int
	Name        string `json:"name_item" binding:"required"`
	IsStackable *bool  `json:"IsStackable" binding:"required"`
	MaxStack    int    `json:"MaxStack" binding:"required,gt=0"`
}

type Inventory struct {
	InventoryID int
	PlayerID    int
	ItemID      int
	Quantity    int
}
