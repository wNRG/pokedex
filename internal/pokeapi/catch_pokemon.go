package pokeapi

import(
	"encoding/json"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	var data Pokemon
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &data); err != nil {
			return Pokemon{}, err
		}

		return data, nil
	}

	if err := c.get(url, &data); err != nil {
		return Pokemon{}, err
	}

	val, err := json.Marshal(data)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, val)

	return data, nil
}
