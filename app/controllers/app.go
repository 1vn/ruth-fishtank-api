package controllers

import (
	"fmt"
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

func (c *App) RecordState() revel.Result {
	return c.RenderText(fmt.Sprintf("%d,%d,%t,%t", 10, 10, true, false))
}

func (c *App) GetActions() revel.Result {
	return c.RenderText(fmt.Sprintf("%s", 1))
}
