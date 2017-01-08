package main

import (
	"github.com/astaxie/beego"
	"github.com/micln/gomoku-AI"
)

func init() {
}

type GameController struct {
	beego.Controller
}

func (c *GameController) client() *Client {
	id := c.GetString(`clientId`)
	return serv.Clients[id]
}

func (c *GameController) Game() {
	c.TplName = `index.html`
	return
}

func (c *GameController) Start() {
	client := serv.NewClient()
	c.ajaxSuccess(client)
}
func (c *GameController) HumanGo() {
	clt := c.client()
	x, _ := c.GetInt(`x`)
	y, _ := c.GetInt(`y`)

	robot := clt.Robot
	board := robot.Board

	board.GoChess(gomoku_AI.NewPoint(x, y), gomoku_AI.C_Player)

	maps := board.CalcScoreMaps(gomoku_AI.C_Robot)

	p := robot.BestStep(4)
	robot.GoChess(p, gomoku_AI.C_Robot)

	c.ajaxSuccess(map[string]interface{}{
		`scoreMaps`: maps,
		`point`:     p,
	})
}

func (c *GameController) ajaxSuccess(content interface{}, messages ...string) {
	c.ajax(true, content, messages...)
}
func (c *GameController) ajaxError(content interface{}, messages ...string) {
	c.ajax(false, content, messages...)
}

func (c *GameController) ajax(success bool, content interface{}, messages ...string) {
	m := make(map[string]interface{})
	m[`success`] = success
	m[`content`] = content

	message := ``
	if len(messages) > 0 {
		message = messages[0]
	}

	m[`message`] = message

	c.Data[`json`] = m
	c.ServeJSON()
}
