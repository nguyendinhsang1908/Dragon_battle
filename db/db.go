package db

import (
	"database/sql"
	"dragon_battle/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

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
