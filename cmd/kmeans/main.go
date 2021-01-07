package main

import (
	"flag"
	"log"
	"os"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/kmeans"

	"github.com/gen95mis/golog"
)

var (
	pathOpen = flag.String("open", "dataset/images/00_000200.png", "path to file with dataset")
	// pathSave = flag.String("save", "save/img.png", "path to image save")
	pathLog = flag.String("log", "tmp/kmeans.log", "path to log file")
	lvl     = flag.String("lvl", "Warn", "log level")
	logPath *os.File
	logger  *golog.Logger
)

func init() {
	var err error

	flag.Parse()

	logPath, err = os.OpenFile(*pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	logger, err = golog.NewLogger(logPath, "", *lvl, log.LstdFlags)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	defer logPath.Close()

	img, err := gil.OpenImage(*pathOpen)
	if err != nil {
		logger.Fatalln(err)
	}

	kmeans := kmeans.NewKMeans(img, 16)
	if _, err := kmeans.Predict(); err != nil {
		logger.Fatalln(err)
	}

}
