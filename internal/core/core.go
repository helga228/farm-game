package core

import (
	"ferma/internal/entity"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
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

func getPerson(user entity.User) (pixel.Picture, error) {
	path := "pictures/person/" + user.Pers + ".png"
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}
	return pic, nil
}

func StartGame(user entity.User, win *pixelgl.Window) {

	person, err := getPerson(user)
	if err != nil {
		panic(err)
	}

	stockImg, err := loadPicture("pictures/menu/stock.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(person, person.Bounds())

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
