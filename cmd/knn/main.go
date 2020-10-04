package main

import (
	"flag"
	"log"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/knn"
)

func main() {
	var (
		pathOpen = flag.String("open", "", "path to image open")
		pathSave = flag.String("save", "", "path to image save")
	)
	flag.Parse()

	img, err := gil.OpenImage(*pathOpen)
	if err != nil {
		log.Fatal(err)
	}

	knn := knn.NewKNN(img)
	newImg, _ := knn.Predict()

	if err := gil.SaveImage(*pathSave, newImg); err != nil {
		log.Fatal(err)
	}
}
