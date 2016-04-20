package state

import (
	"database/sql"
	"github.com/ivanzhangio/ruth-fishtank-api/drivers"
	"github.com/ivanzhangio/ruth-fishtank-api/models"
)

func Insert(params *models.StateParams) error {

	_, err := drivers.DB.Exec(`
			insert into states 
			(tempangle, statetime, foodmeter, needfood, request)
			values
			($1, $2, $3, $4, $5)`,
		a.GetTempAngle(), a.GetStateTime(), a.GetFoodMeter(), a.GetNeedfood(), a.GetRequest())

	if err != nil {
		return err
	}

	return nil
}
