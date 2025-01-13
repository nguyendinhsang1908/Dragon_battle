package db

import (
	"database/sql"
	"dragon_battle/models"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Post Message in BD
func PostMessage(sender_id int, receiver_id int, clan_id int, team_id int, text string, messageType string) error {
	query := `
        INSERT INTO Messages (sender_id, receiver_id, clan_id, team_id, message_type, message)
        VALUES (?, ?, ?, ?, ?, ?)
    `

	var receiverID sql.NullInt64
	var clanID sql.NullInt64
	var teamID sql.NullInt64

	// Xử lý các trường hợp "NULL"
	switch messageType {
	case "world":
		text = strconv.Itoa(sender_id) + " : " + text
		receiverID.Valid = false
		clanID.Valid = false
		teamID.Valid = false
		_, err := DB.Exec(query, sender_id, receiverID, clanID, teamID, messageType, text)
		if err != nil {
			return fmt.Errorf("error inserting world message: %v", err)
		}
		return nil

	case "clan":
		text = strconv.Itoa(sender_id) + " : " + text
		receiverID.Valid = false
		clanID.Int64 = int64(clan_id)
		clanID.Valid = true
		teamID.Valid = false
		_, err := DB.Exec(query, sender_id, receiverID, clanID, teamID, messageType, text)
		if err != nil {
			return fmt.Errorf("error inserting clan message: %v", err)
		}
		return nil

	case "friend":
		text = strconv.Itoa(sender_id) + " : " + text
		receiverID.Int64 = int64(receiver_id)
		receiverID.Valid = true
		clanID.Valid = false
		teamID.Valid = false
		_, err := DB.Exec(query, sender_id, receiverID, clanID, teamID, messageType, text)
		if err != nil {
			return fmt.Errorf("error inserting friend message: %v", err)
		}
		return nil

	case "team":
		text = strconv.Itoa(sender_id) + " : " + text
		receiverID.Valid = false
		clanID.Valid = false
		teamID.Int64 = int64(team_id)
		teamID.Valid = true
		_, err := DB.Exec(query, sender_id, receiverID, clanID, teamID, messageType, text)
		if err != nil {
			return fmt.Errorf("error inserting team message: %v", err)
		}
		return nil

	default:
		return fmt.Errorf("invalid message_type: %s", messageType)
	}
}

func GetMessages(sender_id int, receiver_id int, clan_id int, team_id int, message_type string) ([]models.Message, error) {

	var query string
	var rows *sql.Rows
	var err error

	switch message_type {
	case "world":
		query = `SELECT id, sender_id, receiver_id, clan_id, team_id, message_type, message, created_at 
                 FROM Messages 
                 WHERE message_type = 'world' 
                 ORDER BY created_at ASC`
		rows, err = DB.Query(query)

	case "team":
		query = `SELECT id, sender_id, receiver_id, clan_id, team_id, message_type, message, created_at 
                 FROM Messages 
                 WHERE message_type = 'team' AND team_id = ? 
                 ORDER BY created_at ASC`
		rows, err = DB.Query(query, team_id)

	case "clan":
		query = `SELECT id, sender_id, receiver_id, clan_id, team_id, message_type, message, created_at 
                 FROM Messages 
                 WHERE message_type = 'clan' AND clan_id = ? 
                 ORDER BY created_at ASC`
		rows, err = DB.Query(query, clan_id)

	case "friend":
		query = `SELECT id, sender_id, receiver_id, clan_id, team_id, message_type, message, created_at 
        FROM Messages 
        WHERE ((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)) 
        AND message_type = 'friend' 
        ORDER BY created_at ASC
		`
		rows, err = DB.Query(query, sender_id, receiver_id, receiver_id, sender_id)
	default:
		return nil, fmt.Errorf("invalid message_type: %s", message_type)
	}

	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.ClanID, &msg.TeamID, &msg.MessageType, &msg.Message, &msg.CreatedAt)
		if err != nil {

			return nil, fmt.Errorf("scan error: %v", err)
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	fmt.Println(messages)
	return messages, nil
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

func UpdateName(playerID int, name string) error {
	query := `
        UPDATE Information
        SET name = ?
        WHERE ID = ?
    `

	// Thực hiện truy vấn cập nhật
	_, err := DB.Exec(query, name, playerID)
	if err != nil {
		return fmt.Errorf("failed to update name: %v", err)
	}

	return nil
}

// Update Avatar
func UpdateAvatar(playerID int, avatar string) error {
	query := `
	UPDATE Information
	SET Avatar = ?
	WHERE ID = ?
`

	// Thực hiện truy vấn cập nhật
	_, err := DB.Exec(query, avatar, playerID)
	if err != nil {
		return fmt.Errorf("failed to update name: %v", err)
	}

	return nil
}

// Create Clan
func Crea_Clan(playerID int, n_clan string) error {

	var query string
	var err error

	query = `
		INSERT INTO Clans(name,own) VALUES (?,?)
		`
	_, err = DB.Query(query, n_clan, playerID)
	if err != nil {
		return fmt.Errorf("failed to insert into Clans: %v", err)
	}

	//Update clan_id in Information for player
	query = `
		UPDATE Information
		SET clan_name = ?
		WHERE ID = ?
		`
	_, err = DB.Query(query, n_clan, playerID)
	if err != nil {
		return fmt.Errorf("failed to update clan_id for player: %v", err)
	}
	return nil
}

// Send friend request to databse
func Send_Fr_Mess_Request(senderID int, receiverID int) error {
	query := `
	INSERT INTO FriendRequests (sender_id, receiver_id, status)
	VALUES (?, ?, 'pending')
	`
	_, err := DB.Exec(query, senderID, receiverID)
	if err != nil {
		return fmt.Errorf("failed to send friend request: %v", err)
	}
	return nil
}

// Reponse to friend request
func Accept_Fr_Request(requestID int) error {
	query := `
	UPDATE FriendRequests
	SET status = "accepted"
	WHERE id = ?
`
	_, err := DB.Exec(query, requestID)
	if err != nil {
		return fmt.Errorf("failed to accept friend request: %v", err)
	}
	return nil
}

// if reject friend request
func Reject_Fr_Request(requestID int) error {
	query := `
        UPDATE FriendRequests
        SET status = "rejected"
        WHERE id = ?
    `
	_, err := DB.Exec(query, requestID)
	if err != nil {
		return fmt.Errorf("failed to reject friend request: %v", err)
	}
	return nil
}

// delete Friends request in database
func Delete_Fr_Request(requestID int) error {
	query := `
		DELETE FROM FriendRequests
		WHERE id = ?
	`

	//thuc hien query
	_, err := DB.Exec(query, requestID)
	if err != nil {
		return fmt.Errorf("failed to delete reject friend request: %v", err)
	}

	return nil
}

// Send clan request to databse
func Send_Cl_Mess_Request(senderID int, receiverID int) error {
	query := `
	INSERT INTO ClanRequest (sender_id, receiver_id, status)
	VALUES (?, ?, 'pending')
	`
	_, err := DB.Exec(query, senderID, receiverID)
	if err != nil {
		return fmt.Errorf("failed to send clan request: %v", err)
	}
	return nil
}

// create and ADD Dragon
func Crea_Dragon(Name string, point int, hp int, mp int, dame int, defense int, speed int, rarity string, element string) error {
	query := `
	INSERT INTO Dragons (Name, Point,Hp,Mp,Dame,Defense,Speed,Rarity,Element)
	VALUES (?, ?, ?,?,?,?,?,?,?)
	`
	_, err := DB.Exec(query, Name, point, hp, mp, dame, defense, speed, rarity, element)
	if err != nil {
		return fmt.Errorf("failed to insert dragon: %v", err)
	}
	return nil
}

// Delete_Dragon
func Delete_Dragon(name_dr string) error {
	query := `
	DELETE FROM Dragons
	WHERE Name = ?
`

	//thuc hien query
	_, err := DB.Exec(query, name_dr)
	if err != nil {
		return fmt.Errorf("failed to delete dragon: %v", err)
	}

	return nil
}

func GetAllDragon() ([]models.Dragon, error) {
	query := `SELECT * FROM Dragons`

	// Thực hiện truy vấn
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query dragons: %v", err)
	}
	defer rows.Close()

	var dragons []models.Dragon

	// Lặp qua các hàng kết quả
	for rows.Next() {
		var dragon models.Dragon
		err := rows.Scan(
			&dragon.ID,
			&dragon.Name,
			&dragon.Point,
			&dragon.Hp,
			&dragon.Mp,
			&dragon.Dame,
			&dragon.Defense,
			&dragon.Speed,
			&dragon.Rarity,
			&dragon.Element,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan dragon: %v", err)
		}
		dragons = append(dragons, dragon)
	}

	// Kiểm tra lỗi sau khi lặp qua các hàng
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return dragons, nil
}

func Add_Item(na_item string, is_stack bool, max_stack int) error {
	query := `
	INSERT INTO Items (Name,IsStackable,MaxStack)
	VALUES (?, ?, ?)
	`
	_, err := DB.Exec(query, na_item, is_stack, max_stack)
	if err != nil {
		return fmt.Errorf("failed to insert Item: %v", err)
	}
	return nil
}

func Delete_Item(na_item string) error {
	query := `
	DELETE FROM Items
	WHERE Name = ?
`

	//thuc hien query
	_, err := DB.Exec(query, na_item)
	if err != nil {
		return fmt.Errorf("failed to delete item: %v", err)
	}

	return nil
}

func Get_Item(na_item string) (models.Item, error) {
	query := `
	SELECT Name, IsStackable, MaxStack
	FROM Items
	WHERE Name = ?
	`
	var item models.Item

	err := DB.QueryRow(query, na_item).Scan(&item.Name, &item.IsStackable, &item.MaxStack)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Item{}, fmt.Errorf("item not found")
		}
		return models.Item{}, fmt.Errorf("failed to get item: %v", err)
	}

	return item, nil
}
