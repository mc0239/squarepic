package main

import (
	"hash/fnv"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func hashInput(value string) (int64, error) {
	fnvhash := fnv.New64a()
	_, err := fnvhash.Write([]byte(value))
	// logger.Log(logm.LvlNotice, "fnv: %d", int64(fnvhash.Sum64()))
	return int64(fnvhash.Sum64()), err

}

func generateImage(filepath string, params generationParams, rand *rand.Rand) error {

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, params.size, params.size))

	randomColor := color.NRGBA{
		R: uint8(rand.Int()), G: uint8(rand.Int()), B: uint8(rand.Int()), A: 255,
	}

	unit := params.size / params.squaresCount

	for yu := 0; yu < params.squaresCount; yu++ {
		for xu := 0; xu < params.squaresCount; xu++ {
			isColoredUnit := rand.Intn(10) > 6
			isShadowdUnit := rand.Intn(10) > 4
			var shadow uint8
			if isShadowdUnit {
				shadow = uint8(rand.Intn(6) * 3)
			} else {
				shadow = 0
			}
			for y := unit * yu; y < unit*(yu+1); y++ {
				for x := unit * xu; x < unit*(xu+1); x++ {
					if isColoredUnit {
						img.Set(x, y, randomColor)
					} else {
						img.Set(x, y, color.NRGBA{R: 255 - shadow, G: 255 - shadow, B: 255 - shadow, A: 255})
					}
				}
			}
		}
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
