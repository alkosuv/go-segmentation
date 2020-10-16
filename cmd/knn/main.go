package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/knn"
	"urban-image-segmentation/internal/gil/knn/storage"
)

var (
	pathOpen  = flag.String("open", "", "path to file with dataset")
	pathSave  = flag.String("save", "", "path to image save")
	pathLabel = flag.String("label", "dataset/knn_dataset/labels.csv", "path to labels knn")
	pathLog   = flag.String("log", "tmp/knn.log", "path to log file")
	logger    *log.Logger
)

func init() {
	flag.Parse()

	file, err := os.OpenFile(*pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(file, "", log.LstdFlags)
}

func main() {
	s := storage.NewStorage(logger)
	if err := s.Read(*pathLabel); err != nil {
		logger.Fatalln(err)
	}

	fmt.Println(s.Labels)

	img, err := gil.OpenImage(*pathOpen)
	if err != nil {
		logger.Fatalln(err)
	}

	knn := knn.NewKNN(img)
	newImg, _ := knn.Predict()

	if err := gil.SaveImage(*pathSave, newImg); err != nil {
		logger.Fatalln(err)
	}
}
