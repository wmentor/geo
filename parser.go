package geo

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidCoords error = errors.New("Invalid coords")

	reDec *regexp.Regexp = regexp.MustCompile(`^\s*(\-?\d+(?:\.\d+)?)˚?\s*,\s*(\-?\d+(?:\.\d+)?)˚?\s*$`)
)

func ParseCoords(coords string) (float64, float64, error) {

	if list := reDec.FindStringSubmatch(coords); len(list) == 3 {
		lat, _ := strconv.ParseFloat(list[1], 64)
		long, _ := strconv.ParseFloat(list[2], 64)
		if lat >= -90 && lat <= 90 && long >= -180 && long <= 180 {
			return lat, long, nil
		}
		return 0, 0, ErrInvalidCoords
	}

	return 0, 0, ErrInvalidCoords
}
