package config

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func AccessibleRoles() map[string]map[string]bool {
	const listingsServicePath = "/listings.Listings/"
	const applicationServicePath = "/application.Application/"

	return map[string]map[string]bool{
		// Listing Protected routes
		listingsServicePath + "CreateApartment": {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "CreateBuilding":  {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "CreateRealtor":   {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "DeleteApartment": {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "DeleteBuilding":  {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "DeleteRealtor":   {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "UpdateApartment": {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "UpdateBuilding":  {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "UpdateRealtor":   {"admin": true, "realtor": true, "owner": true},
		listingsServicePath + "Upload":          {"admin": true, "realtor": true, "owner": true},
		// Application Protected routes
		applicationServicePath + "Apply":             {"admin": true, "renter": true},
		applicationServicePath + "CreateApplication": {"admin": true, "renter": true},
		applicationServicePath + "DeleteApplication": {"admin": true, "realtor": true, "owner": true, "renter": true},
		applicationServicePath + "DeleteResponse":    {"admin": true, "realtor": true, "owner": true},
		applicationServicePath + "GetApplication":    {"admin": true, "realtor": true, "owner": true, "renter": true},
		applicationServicePath + "GetResponse":       {"admin": true, "realtor": true, "owner": true, "renter": true},
		applicationServicePath + "ListApplications":  {"admin": true, "realtor": true, "owner": true, "renter": true},
		applicationServicePath + "ListResponses":     {"admin": true, "realtor": true, "owner": true, "renter": true},
		applicationServicePath + "UpdateApplication": {"admin": true, "realtor": true, "owner": true, "renter": true},
		applicationServicePath + "UpdateResponse":    {"admin": true, "realtor": true, "owner": true},
		applicationServicePath + "Upload":            {"admin": true, "realtor": true, "owner": true, "renter": true},
	}
}

// ListGRPCResources is a helper function that lists all URLs that are registered on gRPC server.
// This makes it easy to register all the relevant routes in your HTTP router of choice.
func ListGRPCResources(server *grpc.Server) []string {
	var ret []string
	for serviceName, serviceInfo := range server.GetServiceInfo() {
		for _, methodInfo := range serviceInfo.Methods {
			fullResource := fmt.Sprintf("/%s/%s", serviceName, methodInfo.Name)
			ret = append(ret, fullResource)
		}
	}
	return ret
}

// NewDBContext returns a new Context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*Performance/100)
}
