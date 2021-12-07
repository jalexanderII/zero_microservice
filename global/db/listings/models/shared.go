package models

import (
	"time"

	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

type Place struct {
	FullAddress  string
	Street       string
	City         string
	State        string
	ZipCode      int32
	Neighborhood string
	Unit         string
	Coordinates  Coordinates
}

type Coordinates struct {
	Lat float64
	Lng float64
}

type Content struct {
	Id       string
	Filename string
	FileId   []byte
	Source   IsContentSource
	Type     listingsPB.Content_Type
}

type IsContentSource interface {
	isContent_Source()
}

type ContentApartmentRef struct {
	ApartmentRef string
}

type ContentBuildingRef struct {
	BuildingRef string
}

func (*ContentApartmentRef) isContent_Source() {}

func (*ContentBuildingRef) isContent_Source() {}

type ListingMetrics struct {
	AvailableOn  time.Time
	DaysOnMarket int32
}
