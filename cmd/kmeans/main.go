package main

import (
	"flag"
	"image"
	"os"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/kmeans"

	"github.com/gen95mis/golog"
)

var (
	pathOpen       = flag.String("open", "dataset/images/00_000200.png", "path to file with dataset")
	pathSave       = flag.String("save", "save/img.png", "path to image save")
	pathImageLabel = flag.String("label", "", "path to image label")
	pathLog        = flag.String("log", "tmp/kmeans.log", "path to log file")
	centroid       = flag.Int("centroid", 16, "count centroids")
	isDraw         = flag.Bool("draw", false, "draw segmentation image")
	lvl            = flag.String("lvl", "Warn", "log level")
	logPath        *os.File
	logger         *golog.Logger
)

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

	var (
		imgLabel image.Image
		err      error
	)

	if *pathImageLabel != "" {
		imgLabel, err = gil.OpenImage(*pathImageLabel)
		if err != nil {
			logger.Fatalln(err)
		}
	}

	img, err := gil.OpenImage(*pathOpen)
	if err != nil {
		logger.Fatalln(err)
	}

	kmeans := kmeans.NewKMeans(img, *centroid, imgLabel)
	newImg, err := kmeans.Predict(*isDraw)
	if err != nil {
		logger.Fatalln(err)
	}

	if err := gil.SaveImage(*pathSave, newImg); err != nil {
		logger.Fatalln(err)
	}
}
