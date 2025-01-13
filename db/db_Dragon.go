package db

import (
	"dragon_battle/models"
	"fmt"
)

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
