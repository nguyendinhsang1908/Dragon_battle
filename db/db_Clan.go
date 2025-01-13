package db

import "fmt"

// Create Clan
func Db_Crea_Clan(playerID int, n_clan string) error {

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
func Db_Send_Cl_Mess_Request(senderID int, receiverID int) error {
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
