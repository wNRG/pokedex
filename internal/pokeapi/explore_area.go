package pokeapi

import(
	"encoding/json"
)

func (c *Client) GetPokemonInArea(name string) (LocationArea, error) {
	var data LocationArea 

	url := baseURL + "/location-area/" + name

	// check if in cache
	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &data); err != nil {
			return LocationArea{}, err
		} 

		return data, nil
	}

	// make API request
	if err := c.get(url, &data); err != nil {
		return LocationArea{}, err
	}

	val, err := json.Marshal(data)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, val)

	return data, nil
}

