package messari

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

// Client struct is a struct which holds data related to making API requests against Client's API
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	apiKey     string
}

// New func returns an instance of a Messari
func New(apiKey string) *Client {
	if apiKey == "" {
		panic("need apiKey in order to create new MessariClient!")
	}

	m := &Client{
		httpClient: &http.Client{},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "data.messari.io",
		},
		apiKey: apiKey,
	}

	return m
}

func (m *Client) buildURL(path string) string {
	url := m.baseURL.ResolveReference(&url.URL{Path: path})
	return url.String()
}

func (m *Client) setRequestQuery(req *http.Request, query map[string][]string) {
	q := req.URL.Query()
	for key, valueSlice := range query {
		for _, value := range valueSlice {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()
}

func (m *Client) setRequestHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func (m *Client) request(method string, path string, body interface{}, query map[string][]string) (*http.Response, error) {
	if method == http.MethodGet {
		reqURL := m.buildURL(path)
		req, err := http.NewRequest(http.MethodGet, reqURL, nil)
		if err != nil {
			return nil, fmt.Errorf("could not make GET request: %w", err)
		}
		m.setRequestHeaders(req, map[string]string{
			"x-messari-api-key": m.apiKey,
		})
		if query != nil {
			m.setRequestQuery(req, query)
		}
		logrus.Debugf("making GET request to %s", req.URL)
		resp, err := m.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("could not do GET request: %w", err)
		}
		return resp, nil
	}
	if method == http.MethodPost {
		var buf io.ReadWriter
		if body != nil {
			buf = new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, fmt.Errorf("could not encode body to JSON: %w", err)
			}
		}
		reqURL := m.buildURL(path)
		req, err := http.NewRequest(http.MethodGet, reqURL, buf)
		if err != nil {
			return nil, fmt.Errorf("could not make POST request: %w", err)
		}
		m.setRequestHeaders(req, map[string]string{
			"x-messari-api-key": m.apiKey,
		})
		if query != nil {
			m.setRequestQuery(req, query)
		}
		logrus.Debugf("making POST request to %s", req.URL)
		resp, err := m.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("could not do POST request: %w", err)
		}
		return resp, nil
	}
	return nil, fmt.Errorf("request method %s is not supported", method)
}

// GetAllAssetsOptions struct holds options for the GetAllAssets func call
type GetAllAssetsOptions struct {
	Page             *int
	Sort             *string
	Limit            *int
	Fields           []string
	WithMertricsOnly *bool
	WithProfilesOnly *bool
}

// GetAllAssets function gets all assets from Messari's API. Accepts a fields []string which
// indicates which fields to return for each Asset. Pass nil if you need all.
func (m *Client) GetAllAssets(options *GetAllAssetsOptions) (*GetAllAssetsResp, error) {

	// Default Options
	opts := &GetAllAssetsOptions{
		Page:             intPtr(1),
		Sort:             nil,
		Limit:            intPtr(20),
		Fields:           nil,
		WithMertricsOnly: nil,
		WithProfilesOnly: nil,
	}

	if options != nil {
		if err := copier.CopyWithOption(opts, options, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, fmt.Errorf("could not set options: %w", err)
		}
	}

	query := map[string][]string{
		"page":   {strconv.Itoa(*opts.Page)},
		"limit":  {strconv.Itoa(*opts.Limit)},
		"fields": opts.Fields,
	}
	if opts.WithMertricsOnly != nil {
		if *opts.WithMertricsOnly == true {
			query["with-metrics"] = []string{"true"}
		}
	}
	if opts.WithProfilesOnly != nil {
		if *opts.WithProfilesOnly == true {
			query["with-profiles"] = []string{"true"}
		}
	}
	if opts.Sort != nil {
		query["sort"] = []string{*opts.Sort}
	}

	resp, err := m.request(http.MethodGet, "/api/v2/assets", nil, query)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %w", err)
	}
	defer resp.Body.Close()

	// All pages have been found and nexg page doesn't exist
	if resp.StatusCode == 404 && opts.Page != nil {
		return &GetAllAssetsResp{}, nil
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("server returned Status %d instead of 200", resp.StatusCode)
	}

	var assetsResp GetAllAssetsResp
	if err := json.NewDecoder(resp.Body).Decode(&assetsResp); err != nil {
		return nil, fmt.Errorf("could not unmarshal json from response body: %w", err)
	}

	return &assetsResp, nil
}

// GetAssetOptions struct holds option fields for GetAsset func
type GetAssetOptions struct {
	Fields []string
}

// GetAsset func fetches basic metadata of a given asset symbol or slug
func (m *Client) GetAsset(symbolOrSlug string, options *GetAssetOptions) (*GetAssetResp, error) {
	// default options
	opts := &GetAssetOptions{
		Fields: nil,
	}

	if options != nil {
		if err := copier.CopyWithOption(opts, options, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, fmt.Errorf("could not copy options: %w", err)
		}
	}
	query := map[string][]string{}

	if opts.Fields != nil {
		query["fields"] = opts.Fields
	}

	resp, err := m.request(http.MethodGet, fmt.Sprintf("/api/v1/assets/%s", symbolOrSlug), nil, query)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("server returned Status %d instead of 200", resp.StatusCode)
	}

	var assetResp GetAssetResp
	if err := json.NewDecoder(resp.Body).Decode(&assetResp); err != nil {
		return nil, fmt.Errorf("could not unmarshal json from response body: %w", err)
	}

	return &assetResp, nil

}

// GetAssetMetricsOptions struct holds option fields for GetAssetMetrics func
type GetAssetMetricsOptions struct {
	Fields []string
}

// GetAssetMetrics func returns an asset's metrics given a symbol or slug
func (m *Client) GetAssetMetrics(symbolOrSlug string, options *GetAssetMetricsOptions) (*GetAssetMetricsResp, error) {
	// default options
	opts := &GetAssetOptions{
		Fields: nil,
	}

	if options != nil {
		if err := copier.CopyWithOption(opts, options, copier.Option{IgnoreEmpty: true}); err != nil {
			return nil, fmt.Errorf("could not copy options: %w", err)
		}
	}
	query := map[string][]string{}

	if opts.Fields != nil {
		query["fields"] = opts.Fields
	}

	resp, err := m.request(http.MethodGet, fmt.Sprintf("/api/v1/assets/%s/metrics", symbolOrSlug), nil, query)
	if err != nil {
		return nil, fmt.Errorf("could not make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("server returned Status %d instead of 200", resp.StatusCode)
	}

	var assetMetricsResp GetAssetMetricsResp
	if err := json.NewDecoder(resp.Body).Decode(&assetMetricsResp); err != nil {
		return nil, fmt.Errorf("could not unmarshal json from response body: %w", err)
	}

	return &assetMetricsResp, nil
}

// GetAssetMetricsAggregateOptions struct holds option fields for GetAssetMetricsAggregate func
type GetAssetMetricsAggregateOptions struct {
	Tags   []string
	Sector []string
}

func assign(target map[string]interface{}, a ...map[string]interface{}) map[string]interface{} {
	for _, source := range a {
		for key, value := range source {
			target[key] = value
		}
	}
	return target
}

func intPtr(v int) *int {
	return &v
}

func strPtr(v string) *string {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}
