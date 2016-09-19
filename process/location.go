package process

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/will14smith/agb-calendar/model"

	"googlemaps.github.io/maps"
)

type PlaceApi struct {
	client *maps.Client
}

func NewPlaceApi() (*PlaceApi, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("Please provide a google API key using the GOOGLE_API_KEY environment variable")
	}

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &placeApi{
		client: c,
	}, nil
}

func (api *placeApi) Lookup(seed *model.Location) (*model.Location, error) {
	r := &maps.GeocodingRequest{
		Address: seed.Name,
		Region:  "uk",
	}

	resp, err := api.client.Geocode(context.Background(), r)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return nil, fmt.Errorf("Failed to find location: %s", seed.Name)
	}

	loc := resp[0]

	return &model.Location{
		Name: seed.Name,

		Lat:  loc.Geometry.Location.Lat,
		Long: loc.Geometry.Location.Lng,
	}, nil
}

func (api *placeApi) Directions(mode maps.Mode, from, to *model.Location) (*model.Directions, error) {
	r := &maps.DirectionsRequest{
		Origin:      toDirectionsLocation(from),
		Destination: toDirectionsLocation(to),

		Mode: mode,

		DepartureTime: strconv.FormatInt(getNextSaturday().Unix(), 10),
	}

	routes, _, err := api.client.Directions(context.Background(), r)
	if err != nil {
		return nil, err
	}
	if len(routes) == 0 {
		return nil, fmt.Errorf("Failed to find route")
	}

	return getDirections(routes[0])
}

func getDirections(route maps.Route) (*model.Directions, error) {
	directions := model.Directions{
		Distance: 0,
		Duration: 0,
	}

	for _, leg := range route.Legs {
		directions.Distance = directions.Distance + int64(leg.Meters)
		directions.Duration = directions.Duration + leg.Duration
	}

	return &directions, nil
}

func toDirectionsLocation(location *model.Location) string {
	if location.Lat != 0 || location.Long != 0 {
		return fmt.Sprintf("%f, %f", location.Lat, location.Long)
	}

	return location.Name
}

func getNextSaturday() time.Time {
	n := time.Now()
	t := time.Date(n.Year(), n.Month(), n.Day(), 9, 0, 0, 0, time.UTC)

	t = t.AddDate(0, 0, 1)
	for t.Weekday() != time.Saturday {
		t = t.AddDate(0, 0, 1)
	}

	return t
}
