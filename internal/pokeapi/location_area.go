package pokeapi

import(
	"encoding/json"
)

func (c *Client) GetLocationAreas(url string)(LocationAreaResponse, error) {
	var data LocationAreaResponse

	if url == "" {
		url = baseURL + "/location-area/?limit=20"
	}

	// check if response stored in cache
	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &data); err != nil {
			return LocationAreaResponse{}, err
		} 

		return data, nil
	}

	// make API request
	if err := c.get(url, &data); err != nil {
		return LocationAreaResponse{}, err
	}

	val ,err := json.Marshal(data)
	if err != nil {
		return LocationAreaResponse{}, err 
	}

	c.cache.Add(url, val)

	return data, nil
}
