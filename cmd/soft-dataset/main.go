package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"os"
	"strings"
	"urban-image-segmentation/internal/dataset/label"
	"urban-image-segmentation/internal/dataset/softdataset"
	"urban-image-segmentation/internal/gil"

	"github.com/gen95mis/golog"
)

var (
	pathImages = flag.String("pathImages", "", "path to images")
	pathLabels = flag.String("pathLabels", "", "path to classified images")
	splits     = flag.String("splits", "", "path to dataset")
	pathSave   = flag.String("save", "", "save path to dataset")
	pathLog    = flag.String("log", "tmp/soft-dataset.log", "path to log file")
	lvl        = flag.String("lvl", "Warn", "log level")
	logger     *golog.Logger
)

func init() {
	flag.Parse()

	file, err := os.OpenFile(*pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		golog.Fatal(err)
	}

	logger, err = golog.NewLogger(file, "", *lvl, golog.LstdFlags)
	if err != nil {
		golog.Fatal(err)
	}
}

func main() {
	s := softdataset.NewStorage(logger)

	images, err := openDataset()
	if err != nil {
		logger.Fatalln(err)
	}

	for _, name := range images {
		pathImg := fmt.Sprintf("%s/%s", *pathImages, name)
		pathLbl := fmt.Sprintf("%s/%s", *pathLabels, name)

		img, err := gil.OpenImage(pathImg)
		if err != nil {
			logger.Errorln(err)
			return
		}

		lbl, err := gil.OpenImage(pathLbl)
		if err != nil {
			logger.Errorln(err)
			return
		}

		lblCount := make([]int, 16)
		for i := 0; i < img.Bounds().Max.X; i++ {
			for j := 0; j < img.Bounds().Max.Y; j++ {
				imgR, imgG, imgB, imgA := img.At(i, j).RGBA()
				imgRGBA := color.RGBA{
					uint8(imgR >> 8),
					uint8(imgG >> 8),
					uint8(imgB >> 8),
					uint8(imgA >> 8),
				}

				lblR, lblG, lblB, lblA := lbl.At(i, j).RGBA()
				lblRGBA := color.RGBA{
					uint8(lblR >> 8),
					uint8(lblG >> 8),
					uint8(lblB >> 8),
					uint8(lblA >> 8),
				}

				l := label.Labels[lblRGBA]
				if lblCount[l] < 250 {
					s.Add(imgRGBA, l)
					lblCount[l]++
				}
			}
		}
		fmt.Printf("%+v\n", lblCount)
	}

	if err := s.Save(*pathSave); err != nil {
		logger.Errorln(err)
	}
}

func openDataset() ([]string, error) {
	nameImages := make([]string, 0)

	file, err := os.Open(*splits)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
		str = strings.Trim(str, "\n")
		nameImages = append(nameImages, str)
	}

	return nameImages, nil
}
