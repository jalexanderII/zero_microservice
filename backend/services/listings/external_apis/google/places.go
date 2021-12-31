package google

import (
	"context"
	"fmt"
	"log"

	"github.com/jalexanderII/zero_microservice/backend/services/listings/external_apis"
	"github.com/jalexanderII/zero_microservice/config"
)

type ResponseParser struct {
	HtmlAttributions []string             `json:"html_attributions,omitempty"`
	NextPageToken    string               `json:"next_page_token,omitempty"`
	Results          []PlacesSearchResult `json:"results,omitempty"`
	Status           string               `json:"status,omitempty"`
}

type NearbyPlaces struct {
	Results []PlacesResult
}

type PlacesResult struct {
	FormattedAddress  string
	Geometry          Coordinates
	Name              string
	Types             []string
	PermanentlyClosed bool
	BusinessStatus    string
}

func GetNearbySchools(lat, lng float64, verbose bool) (NearbyPlaces, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultTimeout)
	defer cancel()

	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%v,%v&radius=1500&type=school&key=%s", lat, lng, config.PLACESKEY)
	fmt.Printf("calling with url: %s\n", url)

	obj := &ResponseParser{}

	err := external_apis.MakeGet(ctx, url, obj, verbose)
	if err != nil {
		log.Fatalln(err)
	}
	if obj.Status != "OK" {
		return NearbyPlaces{}, fmt.Errorf("api call failled for reason: %s", obj.Status)
	}
	return toPlaceResult(obj.Results), nil
}

func toPlaceResult(results []PlacesSearchResult) NearbyPlaces {
	res := make([]PlacesResult, len(results))
	for idx, place := range results {
		res[idx] = PlacesResult{
			FormattedAddress:  place.FormattedAddress,
			Geometry:          Coordinates{place.Geometry.Location},
			Name:              place.Name,
			Types:             place.Types,
			PermanentlyClosed: place.PermanentlyClosed,
			BusinessStatus:    place.BusinessStatus,
		}
	}
	return NearbyPlaces{res}
}
