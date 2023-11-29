package main

import (
	"ferma/internal/core"
	"ferma/internal/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
)

func main() { // все конфигурации/репозитории/адапторы и тд будут подключаться тут
	pixelgl.Run(Run)
}
func Run() {
	user := entity.User{ //Пока нет базы используем так, потом будем из базы подтягивать остальные данные (прогресс и прочее)
		Id:   "1",
		Name: "root",
		Pers: "nlo",
	}
	cfg := pixelgl.WindowConfig{
		Title:  "Mars farm",
		Bounds: pixel.R(0, 0, 1224, 968),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	core.StartGame(user, win)
}
