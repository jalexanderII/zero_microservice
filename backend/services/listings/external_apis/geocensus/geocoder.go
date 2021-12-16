package geocensus

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// DefaultTimeout for the request execution
const DefaultTimeout = time.Second * 8

type ResponseParser struct {
	Result struct {
		Input struct {
			Benchmark struct {
				Id                   string `json:"id"`
				BenchmarkName        string `json:"benchmarkName"`
				BenchmarkDescription string `json:"benchmarkDescription"`
				IsDefault            bool   `json:"isDefault"`
			} `json:"benchmark"`
			Address struct {
				Street string `json:"street"`
				City   string `json:"city"`
				State  string `json:"state"`
				Zip    string `json:"zip"`
			} `json:"address"`
		} `json:"input"`
		AddressMatches AddressMatches `json:"addressMatches"`
	} `json:"result"`
}

type AddressMatches []struct {
	MatchedAddress string      `json:"matchedAddress"`
	Coordinates    Coordinates `json:"coordinates"`
	TigerLine      struct {
		TigerLineId string `json:"tigerLineId"`
		Side        string `json:"side"`
	} `json:"tigerLine"`
	AddressComponents struct {
		FromAddress     string `json:"fromAddress"`
		ToAddress       string `json:"toAddress"`
		PreQualifier    string `json:"preQualifier"`
		PreDirection    string `json:"preDirection"`
		PreType         string `json:"preType"`
		StreetName      string `json:"streetName"`
		SuffixType      string `json:"suffixType"`
		SuffixDirection string `json:"suffixDirection"`
		SuffixQualifier string `json:"suffixQualifier"`
		City            string `json:"city"`
		State           string `json:"state"`
		Zip             string `json:"zip"`
	} `json:"addressComponents"`
}

type Matches struct {
	M []Match
}

type Match struct {
	MatchedAddress string      `json:"matchedAddress"`
	Coordinates    Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func GetGeoCode(street, city, state, zip string) (Coordinates, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), DefaultTimeout)
	defer cancel()

	url := fmt.Sprintf(
		"https://geocoding.geo.census.gov/geocoder/locations/address?street=%s&city=%s&state=%s&zip=%s&benchmark=Public_AR_Current&format=json",
		parser(street), parser(city), parser(state), parser(zip),
	)

	m, err := response(ctx, url, &ResponseParser{})
	if err != nil {
		log.Fatalln(err)
	}

	return GetZipFromAddr(m, street, city, state, zip)
}

func parser(s string) string {
	return strings.ReplaceAll(s, " ", "+")
}

// Response gets response from url
func response(ctx context.Context, url string, obj *ResponseParser) (Matches, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Matches{}, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Matches{}, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Matches{}, err
	}

	body := strings.Trim(string(data), " []")
	if body == "" {
		return Matches{}, nil
	}

	if err := json.Unmarshal([]byte(body), obj); err != nil {
		fmt.Printf("payload: %s\n", body)
		return Matches{}, err
	}

	return getMatches(obj.Result.AddressMatches), nil
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

// FastStringConv fast int to string conversion
// https://stackoverflow.com/questions/39442167/convert-int32-to-string-in-golang
func FastStringConv(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
