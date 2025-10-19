package pokeapi

// RespShallowLocations represents the response from the location-area endpoint
type RespShallowLocations struct {
    Count    int     `json:"count"`
    Next     *string `json:"next"`     // Note: pointer to handle null values
    Previous *string `json:"previous"` // Note: pointer to handle null values
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

// RespSpecificLocation represents the response for the specific location-area endpoint
type RespSpecificLocation struct{
    EncounterMethodRates []struct{
        EncounterMethod struct{
            Name string `json:"name"`
            URL  string `json:"url"`
        }`json:"encounter_method"`
        VersionDetails []struct{
            Rate int `json:"rate"`
            Version struct{
                Name string `json:"name"`
                URL  string `json:"url"`
            } `json:"version"`
        }`json:"version_details"`
    } `json:"encounter_method_rates"`
    GameIndex int `json:"game_index"`
    ID int `json:"id"`
    Name string `json:"name"`
    Location struct{            
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"location"`
    Names []struct{
        Language struct{
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"language"`
        Name string `json:"name"`
    }  `json:"names"`
    PokemonEncounters[]struct{
        Pokemon struct{
            Name string `json:"name"`
            URL  string `json:"url"`
        }`json:"pokemon"`
        VersionDetails []struct{
            EncounterDetails []struct{
                Chance int `json:"chnce"`
                ConditionValues []struct{
                    Name string `json:"name"`
                    URL  string `json:"url"`
                }`json:"condition_values"`
                MaxLevel int `json:"max_level"`
                MinLevel int `json:"min_level"`
                Method struct{
                    Name string `json:"name"`
                    URL  string `json:"url"`
                }`json:"method"`
            }`json:"encounter_details"`
            MaxChance int `json:"max_chance"`
            Version struct{
                Name string `json:"name"`
                URL  string `json:"url"`
            }`json:"version"`
        }`json:"version_details"`
    }`json:"pokemon_encounters"`
}