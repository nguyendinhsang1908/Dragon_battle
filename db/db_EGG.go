package db

import (
	"dragon_battle/models"
	"fmt"
)

func Db_Add_Egg(egg_name string, rate string) error {
	query := `
		INSERT INTO Eggs (Name,rate)
		VALUES (?,?)
		`
	_, err := DB.Exec(query, egg_name, rate)
	if err != nil {
		return fmt.Errorf("failed to insert egg: %v", err)
	}
	return nil
}

// delete egg by egg_id
func Db_Delete_EggID(egg_id int) error {
	query := `
		DELETE FROM Eggs
		WHERE EGG_id = ?
	`
	//thuc hien query
	_, err := DB.Exec(query, egg_id)
	if err != nil {
		return fmt.Errorf("failed to delete egg: %v", err)
	}

	return nil
}

// delete by name's egg
func Db_Delete_EggName(na_egg string) error {
	query := `
		DELETE FROM Eggs
		WHERE Name = ?
	`
	//thuc hien query
	_, err := DB.Exec(query, na_egg)
	if err != nil {
		return fmt.Errorf("failed to delete egg: %v", err)
	}

	return nil
}

func GetAllEggs() ([]models.Egg, error) {
	query := `
		SELECT * FROM Eggs
	`
	// Thực hiện truy vấn
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query egg: %v", err)
	}
	defer rows.Close()

	var eggs []models.Egg
	for rows.Next() {
		var egg models.Egg
		err := rows.Scan(&egg.Egg_id, &egg.Name, &egg.Rate, &egg.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan egg: %v", err)
		}
		eggs = append(eggs, egg)
	}
	return eggs, nil
}
