package models

import (
	"database/sql"
	"github.com/ivanzhangio/ruth-fishtank-api/app/drivers"
	"time"
)

type State struct {
	StateId   int64
	TempAngle int
	StateTime time.Time
	FoodMeter int
	NeedFood  bool
	Request   bool
}

func (a *Action) Save() int64 {
	var stateId int64
	if a.StateId != 0 {
		_, err := drivers.DB.Exec(`
			update states set 
			tempangle=$2,
			statetime=$3,
			foodmeter=$4
			where
			stateid=$1
			`, a.StateId, a.TempAngle, a.StateTime, a.FoodMeter)
		if err != nil {
			panic(err)
		}
		stateId = a.StateId
	} else {
		_, err := drivers.DB.Exec(`
			insert into states 
			(tempangle, statetime, foodmeter)
			values
			($1, $2, $3)
			returning stateid`,
			a.TempAngle, a.StateTime, a.FoodMeter).Scan(stateId)

		if err != nil {
			panic(err)
		}
	}

	return stateId
}
