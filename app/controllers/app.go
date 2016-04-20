package controllers

import (
	"fmt"
	"github.com/ivanzhangio/ruth-fishtank-api/app/models/state"
	"github.com/ivanzhangio/ruth-fishtank-api/conf"
	"github.com/revel/revel"
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

	return nil
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
