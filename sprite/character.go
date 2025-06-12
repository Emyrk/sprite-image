package sprite

import (
	"image"
	"image/draw"
)

const (
	spriteWidth  = 64
	spriteHeight = 64
)

type Sprite struct {
	img image.Image
}

func (s *Sprite) Frame(column, row int) (image.Image, error) {
	rect := image.Rect(
		column*spriteWidth,
		row*spriteHeight,
		(column+1)*spriteWidth,
		(row+1)*spriteHeight,
	)

	normalizedSprite := image.NewRGBA(image.Rect(0, 0, spriteWidth, spriteHeight))
	draw.Draw(normalizedSprite, normalizedSprite.Bounds(), s.img, rect.Min, draw.Src)
	return normalizedSprite, nil
}

func (s *Sprite) Back() (image.Image, error)    { return s.Frame(0, 0) }
func (s *Sprite) Forward() (image.Image, error) { return s.Frame(2, 0) }
