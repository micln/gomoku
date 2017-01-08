package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
	"github.com/micln/gomoku-AI"
)

type Server struct {
	Clients map[string]*Client

	beeApp *beego.App
}

func NewServer() *Server {
	serv := &Server{
		Clients: make(map[string]*Client),
		beeApp:  beego.NewApp(),
	}

	//beego.BConfig.WebConfig.Session.SessionOn = true
	//beego.BConfig.WebConfig.Session.SessionProvider = `file`
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = `./tmp`

	return serv
}

func (s *Server) NewClient() *Client {
	clt := &Client{
		Id:    genClientId(),
		Robot: gomoku_AI.NewRobot(),
	}

	s.Clients[clt.Id] = clt

	return clt
}

var (
	clientIds map[int]bool
)

func init() {
	clientIds = make(map[int]bool)
}

func genClientId() string {
	id := time.Now().Nanosecond() % 100007
	if clientIds[id] == true {
		return genClientId()
	}
	return fmt.Sprintf(`%d`, id)
}

func (s *Server) Start() {
	s.beeApp.Handlers.AddAuto(&GameController{})
	s.beeApp.Run()

	return
	//
	//response := func(success bool, content interface{}, message string) interface{} {
	//	m := make(map[string]interface{})
	//	m[`success`] = success
	//	m[`content`] = content
	//	m[`message`] = message
	//
	//	return m
	//}
	//
	//s.beeApp.Handlers.Get(`/`, func(ctx *context.Context) {
	//	ctx.ResponseWriter.Write(responseFile(`views/index.html`))
	//})
	//
	//s.beeApp.Handlers.Get(`/start`, func(ctx *context.Context) {
	//	m := make(map[string]int)
	//	m[`clientId`] = s.NewClient()
	//
	//	ctx.Output.JSON(response(true, m, ``), true, false)
	//})
	//
	//s.beeApp.Handlers.Get(`/human-go`, func(ctx *context.Context) {
	//
	//	clientId := go_utils.Intval(ctx.Input.Query(`clientId`))
	//	x := go_utils.Intval(ctx.Input.Query(`x`))
	//	y := go_utils.Intval(ctx.Input.Query(`y`))
	//
	//	rb := s.Clients[clientId]
	//	p := rb.HumanGo(x, y)
	//
	//	ctx.Output.JSON(response(true, p, ``), true, false)
	//})
	//
	//s.beeApp.Run()
}

func responseFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte(err.Error())
	} else {
		return b
	}
}
