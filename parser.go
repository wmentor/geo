package geo

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidCoords error = errors.New("Invalid coords")

	reDec    *regexp.Regexp = regexp.MustCompile(`^\s*(\-?\d+(?:\.\d+)?)[˚°]?\s*,\s*(\-?\d+(?:\.\d+)?)[˚°]?\s*$`)
	reDecNS  *regexp.Regexp = regexp.MustCompile(`^\s*(\d+(?:\.\d+)?)[˚°]\s*(N|S)\s*,\s*(\d+(?:\.\d+)?)[˚°]\s*(E|W)\s*$`)
	reDecGH1 *regexp.Regexp = regexp.MustCompile(`^\s*(\d+(?:\.\d+)?)_(N|S)_(\d+(?:\.\d+)?)_(E|W)(?:_.+)?$`)
	reDecGH2 *regexp.Regexp = regexp.MustCompile(`^\s*(\d+)_(\d+)_(N|S)_(\d+)_(\d+)_(E|W)(?:_.+)?$`)
	reDecGH3 *regexp.Regexp = regexp.MustCompile(`^\s*(\d+)_(\d+)_(\d+(?:\.\d+)?)_(N|S)_(\d+)_(\d+)_(\d+(?:\.\d+)?)_(E|W)(?:_.+)?$`)
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

	if list := reDecGH1.FindStringSubmatch(coords); len(list) == 5 {
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

	if list := reDecGH2.FindStringSubmatch(coords); len(list) == 7 {
		latG, _ := strconv.ParseFloat(list[1], 64)
		latM, _ := strconv.ParseFloat(list[2], 64)
		longG, _ := strconv.ParseFloat(list[4], 64)
		longM, _ := strconv.ParseFloat(list[5], 64)
		if latG <= 90 && longG <= 180 && latM < 60 && longM < 60 {
			sLat := float64(1)
			if list[3] == "S" {
				sLat = -1
			}
			sLong := float64(1)
			if list[6] == "W" {
				sLong = -1
			}
			return (latG + latM/60) * sLat, (longG + longM/60) * sLong, nil
		}
		return 0, 0, ErrInvalidCoords
	}

	if list := reDecGH3.FindStringSubmatch(coords); len(list) == 9 {
		latG, _ := strconv.ParseFloat(list[1], 64)
		latM, _ := strconv.ParseFloat(list[2], 64)
		latS, _ := strconv.ParseFloat(list[3], 64)
		longG, _ := strconv.ParseFloat(list[5], 64)
		longM, _ := strconv.ParseFloat(list[6], 64)
		longS, _ := strconv.ParseFloat(list[7], 64)
		if latG <= 90 && longG <= 180 && latM < 60 && longM < 60 && latS < 60 && longS < 60 {
			sLat := float64(1)
			if list[4] == "S" {
				sLat = -1
			}
			sLong := float64(1)
			if list[8] == "W" {
				sLong = -1
			}
			return (latG + latM/60 + latS/3600) * sLat, (longG + longM/60 + longS/3600) * sLong, nil
		}
		return 0, 0, ErrInvalidCoords
	}

	return 0, 0, ErrInvalidCoords
}
