package geocensus

type Result struct {
	Input          Input          `json:"input"`
	AddressMatches AddressMatches `json:"addressMatches"`
}

type Input struct {
	Benchmark Benchmark `json:"benchmark"`
	Address   Address   `json:"address"`
}

type Benchmark struct {
	Id                   string `json:"id"`
	BenchmarkName        string `json:"benchmarkName"`
	BenchmarkDescription string `json:"benchmarkDescription"`
	IsDefault            bool   `json:"isDefault"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type AddressMatches []struct {
	MatchedAddress    string            `json:"matchedAddress"`
	Coordinates       Coordinates       `json:"coordinates"`
	TigerLine         TigerLine         `json:"tigerLine"`
	AddressComponents AddressComponents `json:"addressComponents"`
}

type TigerLine struct {
	TigerLineId string `json:"tigerLineId"`
	Side        string `json:"side"`
}

type AddressComponents struct {
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
}

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
