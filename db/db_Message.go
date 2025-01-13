package db

import (
	"database/sql"
	"dragon_battle/models"
	"fmt"
	"strconv"
)

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
	return messages, nil
}
