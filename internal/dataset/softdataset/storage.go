package softdataset

import (
	"bufio"
	"image/color"
	"io"
	"os"
	"urban-image-segmentation/internal/sys"

	"github.com/gen95mis/golog"
)

type Storage struct {
	Labels []Label
	logger *golog.Logger
}

func NewStorage(logger *golog.Logger) *Storage {
	s := new(Storage)
	s.logger = logger
	return s
}

func (s *Storage) Add(rgba color.RGBA, index int) {
	l := Label{rgba, index}
	s.Labels = append(s.Labels, l)
}

func (s *Storage) Save(name string) error {
	file, err := sys.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	defer file.Close()

	for _, l := range s.Labels {
		if _, err := file.WriteString(l.String()); err != nil {
			s.logger.Errorln(err)
		}
	}

	return nil
}

func (s *Storage) Read(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			s.logger.Errorln(err)
		}

		l, err := NewLabel(str)
		if err != nil {
			s.logger.Errorln(err)
		}

		s.Labels = append(s.Labels, *l)
	}

	return nil
}
