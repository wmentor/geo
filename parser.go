package geo

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidCoords error = errors.New("Invalid coords")

	reDec   *regexp.Regexp = regexp.MustCompile(`^\s*(\-?\d+(?:\.\d+)?)[˚°]?\s*,\s*(\-?\d+(?:\.\d+)?)[˚°]?\s*$`)
	reDecNS *regexp.Regexp = regexp.MustCompile(`^\s*(\d+(?:\.\d+)?)[˚°]\s*(N|S)\s*,\s*(\d+(?:\.\d+)?)[˚°]\s*(E|W)\s*$`)
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

	if list := reDecNS.FindStringSubmatch(coords); len(list) == 5 {
		lat, _ := strconv.ParseFloat(list[1], 64)
		long, _ := strconv.ParseFloat(list[3], 64)
		if lat <= 90 && long <= 180 {
			sLat := float64(1)
			if list[2] == "S" {
				sLat = -1
			}
			sLong := float64(1)
			if list[4] == "W" {
				sLong = -1
			}
			return lat * sLat, long * sLong, nil
		}
		return 0, 0, ErrInvalidCoords
	}

	return 0, 0, ErrInvalidCoords
}
