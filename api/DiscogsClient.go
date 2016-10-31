package api

import (
    "net/url"
    "net/http"
    "os"
    "encoding/json"
    "fmt"
)

const (
	apiURL            string = "https://api.discogs.com/database"
	searchURL         string = apiURL + "/search?"
)

// DiscogsClient provides the client and associated elements for interacting with the www.discogs.com
type DiscogsClient struct {
	client      *http.Client 
	token       string       
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

func (dc *DiscogsClient) search(values url.Values) (r *DSearch, err error) {
	r = &DSearch{}
	
	values.Add("token", dc.token)
    
	if err := dc.fetchApiJson(searchURL, values, r); err != nil {
		return r, err
	}
	return r, err
}


// fetchApiJson makes a request to the API and decodes the response.
// `actionUrl` is the final path component that specifies the API call
// `parameters` include the API key
// `result` is modified as an output parameter. It must be a pointer to a ZC JSON structure.
func (dc *DiscogsClient) fetchApiJson(actionUrl string, values url.Values, result interface{}) (err error) {
	var resp *http.Response
	resp, err = dc.makeApiGetRequest(actionUrl, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(result); err != nil {
		return err
	}
	//TODO checkServiceError(body)
	return err
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
		var msg string = fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)
		resp.Write(os.Stdout)
		return resp, ClientError{msg: msg}
	}
	return resp, nil
}