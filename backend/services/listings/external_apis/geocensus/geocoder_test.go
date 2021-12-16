package geocensus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGeoCode(t *testing.T) {
	type args struct{ street, city, state, zip string }
	tests := []struct {
		name string
		args args
		want Coordinates
	}{
		{
			name: "GetGeoCodeZip",
			args: args{"22 Eldridge St", "New York", "NY", "10002"},
			want: Coordinates{X: -73.993546, Y: 40.71506},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGeoCodeZip(tt.args.street, tt.args.city, tt.args.state, tt.args.zip)
			assert.Equal(t, tt.want, got)
			assert.NoError(t, err)
		})
	}
}
