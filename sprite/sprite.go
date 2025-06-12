package sprite

import (
	"fmt"
	"image"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func LoadSprite() (*Sprite, error) {
	tmpDir := os.TempDir()
	outFile := filepath.Join(tmpDir, "out.png")

	cmd := exec.Command(
		os.ExpandEnv("$HOME/.cargo/bin/lpcg-build"),
		"./spritesheets",
		"body::bodies::male::light",
		"head::heads::human::male::light",

		//fmt.Sprintf("head::heads::%s", head),
		//fmt.Sprintf("arms::armour::plate::male::iron"),
		outFile,
	)

	// Use shell to expand the ~ (home dir)
	//cmd.Env = append(os.Environ(), "HOME="+os.Getenv("HOME"))
	//cmd.Path = "/bin/bash"
	//cmd.Args = []string{"bash", "-c", cmd.String()}

	log.Printf("Executing command: %s", cmd.String())

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("command failed: %w", err)
	}
	defer os.Remove(outFile)

	file, err := os.Open(outFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open output file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return &Sprite{
		img: img,
	}, nil

	//spriteWidth := 64
	//spriteHeight := 64
	//
	//bounds := img.Bounds()
	//sheetWidth := bounds.Dx()
	//sheetHeight := bounds.Dy()
	//
	//cols := sheetWidth / spriteWidth
	//rows := sheetHeight / spriteHeight
	//
	//fmt.Printf("Found %d cols and %d rows\n", cols, rows)
	//
	//// Output directory
	//outputDir := "output"
	//os.Mkdir(outputDir, os.ModePerm)
	//
	//// Crop each sprite
	//count := 0
	//for y := 0; y < rows; y++ {
	//	for x := 0; x < cols; x++ {
	//		rect := image.Rect(
	//			x*spriteWidth,
	//			y*spriteHeight,
	//			(x+1)*spriteWidth,
	//			(y+1)*spriteHeight,
	//		)
	//
	//		sprite := image.NewRGBA(rect)
	//		draw.Draw(sprite, rect, img, rect.Min, draw.Src)
	//
	//		// Adjust rectangle to start from (0,0)
	//		normalizedSprite := image.NewRGBA(image.Rect(0, 0, spriteWidth, spriteHeight))
	//		draw.Draw(normalizedSprite, normalizedSprite.Bounds(), img, rect.Min, draw.Src)
	//
	//		// Save the sprite
	//		outputPath := filepath.Join(outputDir, fmt.Sprintf("sprite_%d_%d.png", y, x))
	//		outFile, err := os.Create(outputPath)
	//		if err != nil {
	//			panic(err)
	//		}
	//		defer outFile.Close()
	//
	//		err = png.Encode(outFile, normalizedSprite)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		count++
	//	}
	//}
	//
	//fmt.Printf("Extracted %d sprites\n", count)
	//return nil, nil
}
