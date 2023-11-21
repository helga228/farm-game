package core

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func Run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("pictures/person/farmer_f.png")
	if err != nil {
		panic(err)
	}

	stockImg, err := loadPicture("pictures/menu/stock.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	stock := pixel.NewSprite(stockImg, stockImg.Bounds())

	win.Clear(colornames.Green)

	farmPos := pixel.V(400, 300)
	stockPos := pixel.V(970, 710)

	sprite.Draw(win, pixel.IM.Moved(farmPos))
	stock.Draw(win, pixel.IM.Moved(stockPos))

	for !win.Closed() {
		// Обработка событий клавиатуры
		if win.Pressed(pixelgl.KeyA) {
			farmPos.X -= 5
		}
		if win.Pressed(pixelgl.KeyD) {
			farmPos.X += 5
		}
		if win.Pressed(pixelgl.KeyW) {
			farmPos.Y += 5
		}
		if win.Pressed(pixelgl.KeyS) {
			farmPos.Y -= 5
		}

		// Очистка экрана
		win.Clear(colornames.Green)

		// Рисование спрайта на новой позиции
		sprite.Draw(win, pixel.IM.Moved(farmPos))
		stock.Draw(win, pixel.IM.Moved(stockPos))

		// Обновление экрана
		win.Update()
	}
}
