package main

import (
	"flag"
	"log"
	"os"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/knn"
	"urban-image-segmentation/internal/gil/knn/storage"

	"github.com/gen95mis/golog"
)

var (
	pathOpen  = flag.String("open", "dataset/images/00_000200.png", "path to file with dataset")
	pathSave  = flag.String("save", "save/img.png", "path to image save")
	pathLabel = flag.String("label", "dataset/knn-dataset/labels.csv", "path to labels knn")
	pathLog   = flag.String("log", "tmp/knn.log", "path to log file")
	lvl       = flag.String("lvl", "Warn", "log level")
	logger    *golog.Logger
)

func init() {
	flag.Parse()

	file, err := os.OpenFile(*pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	logger, err = golog.NewLogger(file, "", *lvl, log.LstdFlags)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: Раскрасить отсегментированное изображение
func main() {
	s := storage.NewStorage(logger)
	if err := s.Read(*pathLabel); err != nil {
		logger.Fatalln(err)
	}

	img, err := gil.OpenImage(*pathOpen)
	if err != nil {
		logger.Fatalln(err)
	}

	knn := knn.NewKNN(img, &s.Labels)
	newImg, _ := knn.Predict()

	if err := gil.SaveImage(*pathSave, newImg); err != nil {
		logger.Fatalln(err)
	}
}
