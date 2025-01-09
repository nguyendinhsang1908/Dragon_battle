package untils

import (
	"dragon_battle/models"
)

func Crea_new_data() models.Player {
	info := models.Information{
		ID:      2,
		Name:    "sang",
		Balance: 500,
		Level:   1,
		Avatar:  "NULL",
	}
	ach := models.Achievement{
		ID:         2,
		Win:        0,
		Lose:       0,
		Num_Dragon: 0,
		Num_token:  0,
	}
	pl := models.Player{
		Information: info,
		Achievement: ach,
	}
	return pl
}
