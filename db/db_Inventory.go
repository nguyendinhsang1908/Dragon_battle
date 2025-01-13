package db

import (
	"dragon_battle/models"
	"fmt"
)

/// Inventory

// Add_items inventory
func Db_Add_Inventory(playerID int, items int, quantity int) error {
	query := `
	INSERT INTO Inventory(PlayerID,ItemID,Quantity) VALUES (?,?,?)
	`

	_, err := DB.Query(query, playerID, items, quantity)
	if err != nil {
		return fmt.Errorf("failed to create inventory: %v", err)
	}
	return nil
}

// Get All Items Inventory
func Db_GetAllItems_inven(playerid int) ([]models.Inventory, error) {
	query := `
	SELECT ItemID,Quantity
	FROM Inventory
	WHERE PlayerID = ?
	ORDER BY ItemID
	`
	rows, err := DB.Query(query, playerid)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}

	defer rows.Close()

	var items []models.Inventory
	for rows.Next() {
		var item models.Inventory
		err := rows.Scan(&item.InventoryID, &item.PlayerID, &item.ItemID, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("scan err: %v", err)
		}
		items = append(items, item)
	}
	return items, nil
}

// Delete items inventory
func Db_Delete_item_inventory(inve_id int) error {
	query := `
	DELETE FROM Inventory
	WHERE InventoryID = ?
	`

	// Thực hiện câu truy vấn
	_, err := DB.Exec(query, inve_id)
	if err != nil {
		return fmt.Errorf("failed to delete item: %v", err)
	}

	return nil
}

// Update number items in inventory
func Db_Update_item_inventory(inve_id int, item_id int, quantity int) error {
	query := `
	UPDATE (ItemID,Quantity) VALUES (?,?) 
	FROM Inventory 
	WHERE InventoryID = ?
	`

	// Thực hiện câu truy vấn
	_, err := DB.Exec(query, inve_id, item_id, quantity)
	if err != nil {
		return fmt.Errorf("failed to update quantity item: %v", err)
	}

	return nil
}
