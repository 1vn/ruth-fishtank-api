package state

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ivanzhangio/ruth-fishtank-api/app/drivers"
	"github.com/ivanzhangio/ruth-fishtank-api/app/models"
)

func Insert(params *models.StateParams) error {
	_, err := drivers.DB.Exec(`
			insert into states 
			(tempangle, statetime, foodmeter, needfood, request)
			values
			($1, $2, $3, $4, $5)`,
		params.GetTempAngle(), params.GetStateTime(), params.GetFoodMeter(),
		params.GetNeedFood(), params.GetRequest())

	if err != nil {
		return err
	}

	return nil
}

func ScanRows(rows *sql.Rows) ([]*models.State, error) {
	ls := make([]*models.State, 0)
	for rows.Next() {
		var c models.State
		err := rows.Scan(&c.StateId, &c.TempAngle, &c.StateTime,
			&c.FoodMeter, &c.NeedFood, &c.Request)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Error: invalid")
		}

		ls = append(ls, &c)

	}
	return ls, nil
}

func GetLatest() (*models.State, error) {
	rows, err := drivers.DB.Query(`select * from states order by statetime desc limit 1`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	states, err := ScanRows(rows)
	if err != nil {
		return nil, err
	}
	var s *models.State
	for _, r := range states {
		s = r
	}

	return s, nil
}
