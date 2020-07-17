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
	tPC("12˚ N, 45˚ E", 12, 45, nil)
	tPC("55.7558° N, 37.6173° E", 55.7558, 37.6173, nil)
	tPC("34.6037° S, 58.3816° W", -34.6037, -58.3816, nil)
	tPC("94.6037° S, 58.3816° W", 0, 0, ErrInvalidCoords)
	tPC("55.7558_N_37.6173_E", 55.7558, 37.6173, nil)
	tPC("55.7558_S_37.6173_W", -55.7558, -37.6173, nil)
	tPC("55.7558_N_37.6173_E_param:123", 55.7558, 37.6173, nil)
	tPC("55_30_S_5_15_W", -55.5, -5.25, nil)
	tPC("55_30_0_S_5_15_0_W", -55.5, -5.25, nil)
}
