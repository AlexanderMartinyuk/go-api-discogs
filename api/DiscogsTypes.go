package api


type DError struct {
	Message string `json:"message"`
}

type DSearch struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Items int `json:"items"`
		Page int `json:"page"`
		Urls struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
		Pages int `json:"pages"`
	} `json:"pagination"`
	Results []struct {
		Thumb string `json:"thumb"`
		Title string `json:"title"`
		URI string `json:"uri"`
		ResourceURL string `json:"resource_url"`
		Type string `json:"type"`
		ID int `json:"id"`
		Style []string `json:"style,omitempty"`
		Format []string `json:"format,omitempty"`
		Country string `json:"country,omitempty"`
		Barcode []interface{} `json:"barcode,omitempty"`
		Community struct {
			Have int `json:"have"`
			Want int `json:"want"`
		} `json:"community,omitempty"`
		Label []string `json:"label,omitempty"`
		Catno string `json:"catno,omitempty"`
		Year string `json:"year,omitempty"`
		Genre []string `json:"genre,omitempty"`
	} `json:"results"`
}

type DReleases struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Items int `json:"items"`
		Page int `json:"page"`
		Urls struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
		Pages int `json:"pages"`
	} `json:"pagination"`
	Releases []struct {
		Thumb string `json:"thumb"`
		Artist string `json:"artist"`
		MainRelease int `json:"main_release,omitempty"`
		Title string `json:"title"`
		Role string `json:"role"`
		Year int `json:"year"`
		ResourceURL string `json:"resource_url"`
		Type string `json:"type"`
		ID int `json:"id"`
		Status string `json:"status,omitempty"`
		Format string `json:"format,omitempty"`
		Label string `json:"label,omitempty"`
		Trackinfo string `json:"trackinfo,omitempty"`
	} `json:"releases"`
}

type DArtist struct {
	Profile string `json:"profile"`
	Urls []string `json:"urls"`
	ReleasesURL string `json:"releases_url"`
	Name string `json:"name"`
	URI string `json:"uri"`
	Realname string `json:"realname"`
	Members []struct {
		Active bool `json:"active"`
		ResourceURL string `json:"resource_url"`
		ID int `json:"id"`
		Name string `json:"name"`
	} `json:"members"`
	Images []struct {
		URI string `json:"uri"`
		Height int `json:"height"`
		Width int `json:"width"`
		ResourceURL string `json:"resource_url"`
		Type string `json:"type"`
		URI150 string `json:"uri150"`
	} `json:"images"`
	ResourceURL string `json:"resource_url"`
	ID int `json:"id"`
	DataQuality string `json:"data_quality"`
	Namevariations []string `json:"namevariations"`
}

type DReleaseDetails struct {
	Styles []string `json:"styles"`
	Artists []struct {
		Join string `json:"join"`
		Name string `json:"name"`
		Anv string `json:"anv"`
		Tracks string `json:"tracks"`
		Role string `json:"role"`
		ResourceURL string `json:"resource_url"`
		ID int `json:"id"`
	} `json:"artists"`
	VersionsURL string `json:"versions_url"`
	Year int `json:"year"`
	Images []struct {
		URI string `json:"uri"`
		Height int `json:"height"`
		Width int `json:"width"`
		ResourceURL string `json:"resource_url"`
		Type string `json:"type"`
		URI150 string `json:"uri150"`
	} `json:"images"`
	ID int `json:"id"`
	Tracklist []struct {
		Duration string `json:"duration"`
		Position string `json:"position"`
		Type string `json:"type_"`
		Title string `json:"title"`
	} `json:"tracklist"`
	Genres []string `json:"genres"`
	NumForSale int `json:"num_for_sale"`
	Title string `json:"title"`
	MainRelease int `json:"main_release"`
	MainReleaseURL string `json:"main_release_url"`
	URI string `json:"uri"`
	ResourceURL string `json:"resource_url"`
	LowestPrice float64 `json:"lowest_price"`
	DataQuality string `json:"data_quality"`
}