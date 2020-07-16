package geo

import (
	"testing"
)

func TestParseCoords(t *testing.T) {

	tPC := func(coords string, lat float64, long float64, err error) {
		rlat, rlong, rerr := ParseCoords(coords)
		if rlat != lat || rlong != long || err != rerr {
			t.Fatalf("ParseCoords failed for: %s", coords)
		}
	}

	tPC("1, 2", 1, 2, nil)
	tPC("12˚, 32˚", 12, 32, nil)
	tPC("-71, 35", -71, 35, nil)
	tPC("-71, -92", -71, -92, nil)
	tPC("-71˚, -92˚", -71, -92, nil)
	tPC("-171˚, -92˚", 0, 0, ErrInvalidCoords)
	tPC("# -171˚, -92˚", 0, 0, ErrInvalidCoords)
}
