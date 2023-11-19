package main

import (
	"image"
	"os"

	_ "image/png"

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

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("per.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Greenyellow)
	spritePos := pixel.V(400, 300)
	sprite.Draw(win, pixel.IM.Moved(spritePos))
	for !win.Closed() {
		// Обработка событий клавиатуры
		if win.Pressed(pixelgl.KeyLeft) {
			spritePos.X -= 5
		}
		if win.Pressed(pixelgl.KeyRight) {
			spritePos.X += 5
		}
		if win.Pressed(pixelgl.KeyUp) {
			spritePos.Y += 5
		}
		if win.Pressed(pixelgl.KeyDown) {
			spritePos.Y -= 5
		}

		// Очистка экрана
		win.Clear(colornames.Greenyellow)

		// Рисование спрайта на новой позиции
		sprite.Draw(win, pixel.IM.Moved(spritePos))

		// Обновление экрана
		win.Update()
	}
	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
