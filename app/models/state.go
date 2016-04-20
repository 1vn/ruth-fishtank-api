package models

import (
	"time"
)

type State struct {
	StateId   int64     `json:"state_id"`
	TempAngle int       `json:"temp_angle"`
	StateTime time.Time `json:"state_time"`
	FoodMeter int       `json:"food_meter"`
	NeedFood  bool      `json:"need_food"`
	Request   bool      `json:"request"`
}

type StateParams struct {
	TempAngle NullInt64 `json:"temp_angle"`
	StateTime time.Time `json:"state_time"`
	FoodMeter NullInt64 `json:"food_meter"`
	NeedFood  NullBool  `json:"need_food"`
	Request   NullBool  `json:"request"`
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
