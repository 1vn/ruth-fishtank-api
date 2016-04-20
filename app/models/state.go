package models

import (
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

type StateParams struct {
	TempAngle sql.NullInt64
	StateTime time.Time
	FoodMeter sql.NullInt64
	NeedFood  sql.NullBool
	Request   sql.NullBool
}

func (s *StateParams) GetTempAngle() *int64 {
	if !s.TempAngle.Valid {
		return nil
	}
	return &s.TempAngle.Int64
}

func (s *StateParams) GetFoodMeter() *int64 {
	if !s.FoodMeter.Valid {
		return nil
	}
	return &s.FoodMeter.Int64
}

func (s *StateParams) GetRequest() *bool {
	if !s.Request.Valid {
		return nil
	}
	return &s.Request.Bool
}

func (s *StateParams) GetNeedFood() *bool {
	if !s.NeedFood.Valid {
		return nil
	}
	return &s.NeedFood.Bool
}

func (s *StateParams) GetStateTime() *time.Time {
	if s.StateTime.IsZero() {
		return nil
	}
	return &s.StateTime
}
