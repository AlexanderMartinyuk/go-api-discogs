package api

// DSearch ...
type DSearch struct {
	Pagination struct {
		PerPage                 int    `json:"per_page"`
		Items                   int    `json:"items"`
		Page                    int    `json:"page"`
		Urls                    string `json:"name"`
		Pages                   int    `json:"pages"`
	} `json:"pagination"`
	Results      []struct {
		Thumb                   string  `json:"thumb"`
		Title                   string  `json:"title"`
		URI                     string  `json:"uri"`
		ResourceURL             string  `json:"resource_url"`
		Type                    string  `json:"type"`
		ID                      int     `json:"id"`
	} `json:"results"`
}