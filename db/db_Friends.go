package db

import (
	"dragon_battle/models"
	"fmt"
)

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
