package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c *App) Index() revel.Result {
	return c.Render()
}

func (c *App) Hello() revel.Result {
	return c.RenderText(fmt.Sprintf("%d", 1))
}

func (c *App) GetActions() revel.Result {

	return c.RenderText(fmt.Sprintf("%s", 1))
}
