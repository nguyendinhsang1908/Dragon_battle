package db

import (
	"database/sql"
	"dragon_battle/models"
	"fmt"
)

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
