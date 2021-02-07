package main

import (
	"flag"
	"os"
	"urban-image-segmentation/internal/dataset/softdataset"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/kmeans"

	"github.com/gen95mis/golog"
)

var (
	pathOpen  = flag.String("open", "dataset/images/00_000200.png", "path to file with dataset")
	pathSave  = flag.String("save", "save/img.png", "path to image save")
	pathLabel = flag.String("label", "dataset/saft-dataset/labels.csv", "path to labels kmeans")
	pathLog   = flag.String("log", "tmp/kmeans.log", "path to log file")
	lvl       = flag.String("lvl", "Warn", "log level")
	logPath   *os.File
	logger    *golog.Logger
)

const centroids = 16

func init() {
	var err error

	flag.Parse()

	logPath, err = os.OpenFile(*pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		golog.Fatalln(err)
	}

	logger, err = golog.NewLogger(logPath, "", *lvl, golog.LstdFlags)
	if err != nil {
		golog.Fatalln(err)
	}
}

func main() {
	defer logPath.Close()

	s := softdataset.NewStorage(logger)
	if err := s.Read(*pathLabel); err != nil {
		logger.Fatalln(err)
	}

	img, err := gil.OpenImage(*pathOpen)
	if err != nil {
		logger.Fatalln(err)
	}

	kmeans := kmeans.NewKMeans(img, centroids, s.Labels)
	if _, err := kmeans.Predict(); err != nil {
		logger.Fatalln(err)
	}

	newImg, err := kmeans.Predict()
	if err != nil {
		logger.Fatalln(err)
	}

	if err := gil.SaveImage(*pathSave, newImg); err != nil {
		logger.Fatalln(err)
	}
}
