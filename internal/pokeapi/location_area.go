package pokeapi

type LocationAreaResponse struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`

}

func (c *Client) GetLocationAreas(url string)(LocationAreaResponse, error) {
	var data LocationAreaResponse

	if url == "" {
		url = baseURL + "/location-area/?limit=20"
	}
	if err := c.get(url, &data); err != nil {
		return LocationAreaResponse{}, err
	}

	return data, nil
}
