package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

// Scraper collects metrics from a tinode server.
type Scraper struct {
	// Target Tinode server address.
	address string
	// List of simple numeric metrics to scrape.
	simpleMetrics []string
	// List of histogram metrics to scrape.
	histogramMetrics []string
}

// Histogram struct.
type histogram struct {
	count   uint64
	sum     float64
	buckets map[float64]uint64
}

var errKeyNotFound = errors.New("key not found")
var errMalformed = errors.New("input malformed")

// CollectRaw gathers all metrics from the configured Tinode instance,
// and returns them as a map.
func (s *Scraper) CollectRaw() (map[string]any, error) {
	stats, err := s.Scrape()
	if err != nil {
		log.Println("Failed to fetch or parse response", err)
		return nil, err
	}
	metrics, err := s.parseStatsRaw(stats)
	if err != nil {
		return nil, err
	}
	metrics["up"] = 1.0
	return metrics, nil
}

// Scrape fetches the data from Tinode server using HTTP GET then decodes the response.
func (s *Scraper) Scrape() (map[string]any, error) {
	resp, err := http.Get(s.address)
	if err != nil {
		log.Println("Failed to connect to server", err)
		return nil, err
	}
	defer resp.Body.Close()

	var stats map[string]any
	err = json.NewDecoder(resp.Body).Decode(&stats)
	return stats, err
}

func (s *Scraper) parseStatsRaw(stats map[string]any) (map[string]any, error) {
	metrics := make(map[string]any)
	for _, key := range s.simpleMetrics {
		if val, err := parseNumeric(stats, key); err == nil {
			metrics[key] = val
		} else {
			return nil, err
		}
	}
	for _, key := range s.histogramMetrics {
		if val, err := parseHisto(stats, key); err == nil {
			metrics[key] = val
		} else {
			return nil, err
		}
	}
	return metrics, nil
}

// Extracts a simple histogram from `stats` and returns a cumulative histogram
// corresponding to the simple histogram.
// Returns: (count, sum, buckets, error) tuple.
func parseHisto(stats map[string]any, key string) (*histogram, error) {
	// Histogram is presented as a json with the predefined fields: count, sum, count_per_bucket, bounds.
	count, err := parseNumeric(stats, key+".count")
	if err != nil {
		return nil, err
	}
	sum, err := parseNumeric(stats, key+".sum")
	if err != nil {
		return nil, err
	}
	buckets, err := parseList(stats, key+".count_per_bucket")
	if err != nil {
		return nil, err
	}
	bounds, err := parseList(stats, key+".bounds")
	if err != nil {
		return nil, err
	}
	n := len(buckets)
	if n != len(bounds)+1 {
		return nil, errMalformed
	}
	result := make(map[float64]uint64)
	s := uint64(0)
	for i, v := range bounds {
		s += uint64(buckets[i])
		result[v] = s
	}
	return &histogram{count: uint64(count), sum: sum, buckets: result}, nil
}

// Extracts a list of numerics from `stats` for the given path.
func parseList(stats map[string]any, path string) ([]float64, error) {
	value, err := parseMetric(stats, path)
	if err != nil {
		return nil, err
	}
	listval, ok := value.([]any)
	if !ok {
		log.Println("Value at path is not a float64 array:", path, value)
		return nil, errMalformed
	}
	result := []float64{}
	for _, v := range listval {
		result = append(result, v.(float64))
	}
	return result, nil
}

// Extracts a numeric from `stats` for the given path.
func parseNumeric(stats map[string]any, path string) (float64, error) {
	value, err := parseMetric(stats, path)
	if err != nil {
		return 0, err
	}
	floatval, ok := value.(float64)
	if !ok {
		log.Println("Value at path is not a float64:", path, value)
		return 0, errKeyNotFound
	}
	return floatval, nil
}

// Extracts a metric from `stats` for the given path.
func parseMetric(stats map[string]any, path string) (any, error) {
	parts := strings.Split(path, ".")
	var value any
	var found bool
	value = stats
	for i := range parts {
		subset, ok := value.(map[string]any)
		if !ok {
			log.Println("Invalid key path:", path)
			return 0, errKeyNotFound
		}
		value, found = subset[parts[i]]
		if !found {
			log.Println("Invalid key path:", path, "(", parts[i], ")")
			return 0, errKeyNotFound
		}
	}

	return value, nil
}
