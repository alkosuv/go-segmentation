package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strings"
	"urban-image-segmentation/internal/dataset/label"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/knn/storage"
)

var (
	pathImages = flag.String("pathImages", "", "path to image")
	pathLabels = flag.String("pathLabels", "", "path to classified image")
	splits     = flag.String("splits", "", "path to dataset")
	pathSave   = flag.String("save", "", "save path to dataset")
	pathLog    = flag.String("log", "tmp/knn_selection.log", "path to log file")
	logger     *log.Logger
)

func init() {
	flag.Parse()

	file, err := os.OpenFile(*pathLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger = log.New(file, "", log.LstdFlags)
}

func main() {
	s := storage.NewStorage(logger)

	images, err := openDataset()
	if err != nil {
		logger.Fatalln(err)
	}

	for _, name := range images {
		pathImg := fmt.Sprintf("%s/%s", *pathImages, name)
		pathLbl := fmt.Sprintf("%s/%s", *pathLabels, name)

		img, err := gil.OpenImage(pathImg)
		if err != nil {
			logger.Println(err)
			return
		}

		lbl, err := gil.OpenImage(pathLbl)
		if err != nil {
			logger.Println(err)
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
		logger.Println(err)
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
