package controllers

import (
	"fmt"
	"github.com/ivanzhangio/ruth-fishtank-api/app/models"
	"github.com/ivanzhangio/ruth-fishtank-api/app/models/state"
	"github.com/ivanzhangio/ruth-fishtank-api/conf"
	"github.com/revel/revel"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Init() revel.Result {
	token := c.Params.Get("access_token")
	if token != conf.PASS {
		c.Response.Status = 403
		mp := map[string]interface{}{
			"message": "bad token",
		}
		return c.RenderJson(mp)
	}
	return nil
}

func (c *App) Index() revel.Result {
	return c.Render()
}

func (c *App) Hello() revel.Result {
	return c.RenderText(fmt.Sprintf("%d", 1))
}

func (c *App) UpdateState(temp_angle, food_meter, state_time int64,
	need_food, request bool) revel.Result {

	tm := time.Unix(state_time, 0)

	err := state.Insert(&models.StateParams{
		TempAngle: models.NullInt64{Int64: temp_angle, Valid: true},
		FoodMeter: models.NullInt64{Int64: food_meter, Valid: true},
		StateTime: tm,
		NeedFood:  models.NullBool{Bool: need_food, Valid: true},
		Request:   models.NullBool{Bool: request, Valid: true},
	})

	if err != nil {
		mp := map[string]interface{}{
			"message": fmt.Sprintf("insert error, bro: %s", err.Error()),
		}
		return c.RenderJson(mp)
	}

	mp := map[string]interface{}{
		"message": fmt.Sprintf("insert success, bro"),
	}
	return c.RenderJson(mp)
}

func (c *App) GetState() revel.Result {
	s, err := state.GetLatest()
	if err != nil {
		c.Response.Status = 500

		mp := map[string]interface{}{
			"message": fmt.Sprintf("server error, bro: %s", err.Error()),
		}
		return c.RenderJson(mp)
	}

	if s == nil {
		c.Response.Status = 404
		mp := map[string]interface{}{
			"message": "no states",
		}
		return c.RenderJson(mp)
	}

	return c.RenderJson(s)
}
