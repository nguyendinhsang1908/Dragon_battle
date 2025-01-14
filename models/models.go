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

type Stat struct {
	Hp      int
	Mp      int
	Dame    int
	Defense int
	Speed   int
}

type Dragon struct {
	ID    int
	Name  string
	Point int
	Stat
	Rarity  string // Common,epic,legendary,immo
	Element string // Fire,Natural,Terra,Ice,Metal,Dark,Light,Chaos
}

type Item struct {
	//ItemID      int
	Name        string `json:"name_item" binding:"required"`
	Type        string `json:"type" binding:"required"` //
	IsStackable *bool  `json:"IsStackable" binding:"required"`
	MaxStack    int    `json:"MaxStack" binding:"required,gt=0"`
}

type Inventory struct {
	InventoryID int
	PlayerID    int
	ItemID      int
	Quantity    int
}

type Egg struct {
	Egg_id    int
	Name      string `json:"name_egg" bingding:"required"`
	Rate      string `json:"rate" bingding:"required"`
	CreatedAt time.Time
}

// Event model: Logs in-game events and rewards
type Event struct {
	EventID       int       `json:"event_id" db:"event_id"`
	EventName     string    `json:"event_name" db:"event_name"`
	EventType     string    `json:"event_type" db:"event_type"`
	StartTime     time.Time `json:"start_time" db:"start_time"`
	EndTime       time.Time `json:"end_time" db:"end_time"`
	RewardDetails string    `json:"reward_details" db:"reward_details"` // JSON stored as string
}

// Ranking model: Stores leaderboard data
type Ranking struct {
	RankingID int       `json:"ranking_id" db:"ranking_id"`
	PlayerID  int       `json:"player_id" db:"player_id"`
	RankType  string    `json:"rank_type" db:"rank_type"` // 'global', 'clan', etc.
	Score     int64     `json:"score" db:"score"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type ClanMember struct {
	MembershipID int       `json:"membership_id" db:"membership_id"`
	ClanID       int       `json:"clan_id" db:"clan_id"`
	PlayerID     int       `json:"player_id" db:"player_id"`
	Role         string    `json:"role" db:"role"` // 'member', 'admin', 'leader'
	JoinedAt     time.Time `json:"joined_at" db:"joined_at"`
}
