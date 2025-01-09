package db

import (
	"database/sql"
	"dragon_battle/config"
	"dragon_battle/models"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(cfg config.DBConfig) {
	//dsn := fmt.Sprintf("%s:%s@%s(%s)/%s",
	//	cfg.User, cfg.Passwd, cfg.Net, cfg.Addr, cfg.DBName)
	dsn := "root:123456@tcp(127.0.0.1:3307)/dragon_battle"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection is not alive: %v", err)
	}

	log.Println("Connected to MySQL database!")
}

func Message(from_id int, to_id int, text string) error {
	query := `INSERT INTO Messanges (ID_Player1,ID_Player2,Messanges) VALUES(?,?,?) `
	text = strconv.Itoa(from_id) + text
	_, err := DB.Exec(query, from_id, to_id, text)
	return err
}

func Cre_Player(player models.Player) error {
	// Câu lệnh SQL để chèn dữ liệu vào bảng Information và Achievement
	queryInfo := `INSERT INTO Information (ID, Name, Balance, Level, Avatar) VALUES (?, ?, ?, ?, ?)`
	queryAchieve := `INSERT INTO Achievement (ID, Win, Lose, Num_Dragon, Num_Token) VALUES (?, ?, ?, ?, ?)`

	// Chèn vào bảng Information
	_, err := DB.Exec(queryInfo, player.Information.ID, player.Name, player.Balance, player.Level, player.Avatar)
	if err != nil {
		return fmt.Errorf("failed to insert into Information: %v", err)
	}

	// Chèn vào bảng Achievement
	_, err = DB.Exec(queryAchieve, player.Achievement.ID, 0, 0, 0, player.Num_token)
	if err != nil {
		return fmt.Errorf("failed to insert into Achievement: %v", err)
	}

	return nil
}
func GetPOnePlayer(playerID int) (models.Player, error) {
	query := `
		SELECT 
			i.ID, i.Name, i.Balance, i.Level, i.Avatar,
			a.Win, a.Lose, a.Num_Dragon, a.Num_token
		FROM Information i
		JOIN Achievement a ON i.ID = a.ID
		WHERE i.ID = ?`

	row := DB.QueryRow(query, playerID)

	var player models.Player
	err := row.Scan(
		&player.ID, &player.Name, &player.Balance, &player.Level, &player.Avatar,
		&player.Win, &player.Lose, &player.Num_Dragon, &player.Num_token,
	)
	if err != nil {
		return models.Player{}, err
	}

	return player, nil
}

func AddFriend(playerID, friendID int) error {
	// Kiểm tra xem bạn bè đã tồn tại chưa
	queryCheck := `SELECT COUNT(*) FROM Friends WHERE PlayerID = ? AND FriendID = ?`
	var count int
	err := DB.QueryRow(queryCheck, playerID, friendID).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check friendship: %v", err)
	}
	if count > 0 {
		return fmt.Errorf("friendship already exists")
	}

	// Thêm bạn bè
	query := `INSERT INTO Friends (PlayerID, FriendID) VALUES (?, ?)`
	_, err = DB.Exec(query, playerID, friendID)
	if err != nil {
		return fmt.Errorf("failed to add friend: %v", err)
	}

	// Thêm quan hệ ngược lại (nếu cần)
	_, err = DB.Exec(query, friendID, playerID)
	if err != nil {
		return fmt.Errorf("failed to add reverse friendship: %v", err)
	}

	return nil
}

func GetFriends(playerID int) ([]models.Player, error) {
	query := `
		SELECT 
			i.ID, i.Name, i.Balance, i.Level, i.Avatar,
			a.Win, a.Lose, a.Num_Dragon, a.Num_token
		FROM Information i
		JOIN Achievement a ON i.ID = a.ID
		JOIN Friends f ON f.FriendID = i.ID
		WHERE f.PlayerID = ?
		ORDER BY i.ID ASC
		`

	rows, err := DB.Query(query, playerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get friends: %v", err)
	}
	defer rows.Close()

	var friends []models.Player
	for rows.Next() {
		var friend models.Player
		if err := rows.Scan(&friend.ID, &friend.Name, &friend.Balance, &friend.Level, &friend.Avatar, &friend.Win, &friend.Lose, &friend.Num_Dragon, &friend.Num_token); err != nil {
			return nil, fmt.Errorf("failed to scan friend: %v", err)
		}
		friends = append(friends, friend)
	}

	return friends, nil
}

func DeleteFriend(playerID, friendID int) error {
	// Câu truy vấn SQL để xóa mối quan hệ bạn bè
	query := `
        DELETE FROM Friends
        WHERE PlayerID = ? AND FriendID = ?`

	// Thực hiện câu truy vấn
	_, err := DB.Exec(query, playerID, friendID)
	if err != nil {
		return fmt.Errorf("failed to delete friend: %v", err)
	}

	return nil
}
