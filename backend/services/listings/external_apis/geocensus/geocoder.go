package geocensus

import (
	"context"
	"fmt"
	"log"
	"strings"

	config "github.com/jalexanderII/zero_microservice"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/external_apis"
)

type ResponseParser struct {
	Result Result `json:"result"`
}

type Matches struct {
	M []Match
}

type Match struct {
	MatchedAddress string      `json:"matchedAddress"`
	Coordinates    Coordinates `json:"coordinates"`
}

func GetGeoCodeZip(street, city, state, zip string, verbose bool) (Coordinates, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultTimeout)
	defer cancel()

	url := fmt.Sprintf(
		"https://geocoding.geo.census.gov/geocoder/locations/address?street=%s&city=%s&state=%s&zip=%s&benchmark=Public_AR_Current&format=json",
		parser(street), parser(city), parser(state), parser(zip),
	)

	obj := &ResponseParser{}

	err := external_apis.MakeGet(ctx, url, obj, verbose)
	if err != nil {
		log.Fatalln(err)
	}

	m := getMatches(obj.Result.AddressMatches)

	return GetZipFromAddr(m, street, city, state, zip)
}

func parser(s string) string {
	return strings.ReplaceAll(s, " ", "+")
}

func getMatches(res AddressMatches) Matches {
	m := make([]Match, len(res))
	for idx, val := range res {
		m[idx] = Match{val.MatchedAddress, val.Coordinates}
	}
	return Matches{m}
}

func GetZipFromAddr(m Matches, street, city, state, zip string) (Coordinates, error) {
	entered := fmt.Sprintf(
		"%s, %s, %s, %s",
		strings.ToLower(street), strings.ToLower(city), strings.ToLower(state), strings.ToLower(zip),
	)
	var f []string
	for _, addr := range m.M {
		fetched := strings.ToLower(addr.MatchedAddress)
		if entered == fetched {
			return addr.Coordinates, nil
		}
		f = append(f, fetched)
	}
	return Coordinates{0, 0}, fmt.Errorf("no addresses matched, original: %v, fetched: %v", entered, f)
}
