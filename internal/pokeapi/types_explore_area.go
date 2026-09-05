package pokeapi

type ExploreAreaResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}


type PokemonEncounter struct {
	Pokemon PokemonResource `json:"pokemon"`
}

type PokemonResource struct {
	Name string `json:"name"`  
	URL string `json:"url"`
}


