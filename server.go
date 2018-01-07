package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micln/go-utils"
	"github.com/micln/gomoku-AI"
)

type Server struct {
	Clients map[string]*Client
}

func NewServer() *Server {
	serv := &Server{
		Clients: make(map[string]*Client),
	}

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

	router := gin.Default()
	{
		h5 := router.Group("/h5")
		{
			h5.GET("/", func(ctx *gin.Context) {
				ctx.File(`views/index.html`)
			})
		}

		api := router.Group("/api")
		{
			api.GET("/start", func(ctx *gin.Context) {
				client := s.NewClient()
				ctx.JSON(200, AjaxSuccess(client))
			})
			api.GET("/fire-human-go", func(ctx *gin.Context) {

				id, _ := ctx.GetQuery("clientId")
				ix, _ := ctx.GetQuery(`x`)
				iy, _ := ctx.GetQuery(`y`)
				x := go_utils.Intval(ix)
				y := go_utils.Intval(iy)

				if id == "" {
					ctx.JSON(401, AjaxError(nil, "参数缺失"))
					return
				}

				clt := s.Clients[id]

				robot := clt.Robot
				board := robot.Board

				board.GoChess(gomoku_AI.NewPoint(x, y), gomoku_AI.C_Player)
				maps := board.CalcScoreMaps(gomoku_AI.C_Robot)

				p := robot.BestPoint(1)
				robot.GoChess(p, gomoku_AI.C_Robot)

				ctx.JSON(200, AjaxSuccess(map[string]interface{}{
					`scoreMaps`: maps,
					`point`:     p,
				}))
			})
		}
	}
	router.Run(":8080")
}

func responseFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte(err.Error())
	} else {
		return b
	}
}
