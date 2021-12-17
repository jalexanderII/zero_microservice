package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNearbySchools(t *testing.T) {
	type args struct {
		lat float64
		lng float64
	}
	tests := []struct {
		name    string
		args    args
		verbose bool
		want    PlacesResult
	}{
		{
			name:    "GetNearbySchool",
			args:    args{-33.8670522, 151.1957362},
			verbose: false,
			want: PlacesResult{
				FormattedAddress:  "",
				Geometry:          Coordinates{Location: LatLng{Lat: -33.8692158, Lng: 151.2094192}},
				Name:              "Style Academy Australia",
				Types:             []string{"university", "school", "point_of_interest", "establishment"},
				PermanentlyClosed: false,
				BusinessStatus:    "OPERATIONAL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj, err := GetNearbySchools(tt.args.lat, tt.args.lng, tt.verbose)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, obj.Results[0])
		})
	}
}
