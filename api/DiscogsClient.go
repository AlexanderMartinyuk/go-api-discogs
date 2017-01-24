package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	apiURL      string = "https://api.discogs.com/"
	artistURL   string = apiURL + "artists/%d"
	releasesURL string = apiURL + "artists/%d/releases"
	mastersURL  string = apiURL + "masters/%d"
	searchURL   string = apiURL + "database/search?"
)

// DiscogsClient provides the client and associated elements for interacting with the www.discogs.com
type DiscogsClient struct {
	client *http.Client
	token  string
}

// NewDiscogsClient generates a new client for the Discogs API
func NewDiscogsClient(httpClient *http.Client, token string) *DiscogsClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &DiscogsClient{client: httpClient, token: token}
}

// SearchRelease search for artists by name
func (dc *DiscogsClient) SearchRelease(releaseName string) (r *DSearch, err error) {
	params := url.Values{}
	params.Add("q", releaseName)
	params.Add("type", "release")

	return dc.search(params)
}

// SearchArtist search for artists by name
func (dc *DiscogsClient) SearchArtist(artistName string) (r *DSearch, err error) {
	params := url.Values{}
	params.Add("q", artistName)
	params.Add("type", "artist")

	return dc.search(params)
}

// Search search for artists by name
func (dc *DiscogsClient) Search(query string) (r *DSearch, err error) {
	params := url.Values{}
	params.Add("q", query)

	return dc.search(params)
}

// GetReleasesByArtistID search for artists by id
func (dc *DiscogsClient) GetReleasesByArtistID(artistID int) (r *DReleases, err error) {
	r = &DReleases{}
	url := fmt.Sprintf(releasesURL, artistID)

	err = dc.fetchApiJson(url, nil, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// GetArtistByID search for artists by id
func (dc *DiscogsClient) GetArtistByID(artistID int) (r *DArtist, err error) {
	r = &DArtist{}
	url := fmt.Sprintf(artistURL, artistID)

	err = dc.fetchApiJson(url, nil, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// GetReleaseDetailsByID search for album by id
func (dc *DiscogsClient) GetReleaseDetailsByID(releaseID int) (r *DReleaseDetails, err error) {
	r = &DReleaseDetails{}
	url := fmt.Sprintf(mastersURL, releaseID)

	err = dc.fetchApiJson(url, nil, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (dc *DiscogsClient) search(values url.Values) (r *DSearch, err error) {
	r = &DSearch{}
	values.Add("token", dc.token)

	err = dc.fetchApiJson(searchURL, values, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// fetchApiJson makes a request to the API and decodes the response.
// `actionUrl` is the final path component that specifies the API call
// `parameters` include the API key
// `result` is modified as an output parameter. It must be a pointer to a ZC JSON structure.
func (dc *DiscogsClient) fetchApiJson(actionUrl string, values url.Values, result interface{}) (err error) {

	resp, err := dc.makeApiGetRequest(actionUrl, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(result)
}

// makeApiGetRequest fetches a URL with querystring via HTTP GET and
//  returns the response if the status code is HTTP 200
// `parameters` should not include the apikey.
// The caller must call `resp.Body.Close()`.
func (dc *DiscogsClient) makeApiGetRequest(fullUrl string, values url.Values) (resp *http.Response, err error) {

	req, err := http.NewRequest("GET", fullUrl+values.Encode(), nil)
	if err != nil {
		return resp, err
	}

	resp, err = dc.client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 200 {
		msg := &DError{}
		dec := json.NewDecoder(resp.Body)

		err = dec.Decode(msg)
		if err != nil {
			return resp, err
		}

		return resp, ClientError{msg: msg.Message}
	}

	return resp, nil
}
