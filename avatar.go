package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Avatar struct {
	user     *User
	img      *image.RGBA
	fileName string
}

func NewAvatar(user *User, fileName string) *Avatar {
	avatar := &Avatar{user: user, fileName: fileName}
	avatar.img = image.NewRGBA(image.Rect(0, 0, 800, 800))
	return avatar
}

func (a *Avatar) Create() error {
	file, err := os.Create(a.fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	png.Encode(file, a.img)
	return nil
}

func (a *Avatar) Draw() error {
	hash := a.user.ToUserHash()
	lenHash := len(hash)
	var idx int = 0
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			rect := image.Rect(x*50, y*50, x*50+50, y*50+50)
			color := color.RGBA{hash[idx%lenHash], hash[(idx+1)%lenHash], hash[(idx+2)%lenHash], 255}
			draw.Draw(a.img, rect, &image.Uniform{color}, image.ZP, draw.Src)
			idx++
		}
	}
	return nil
}
