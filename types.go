package main

type Weather struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Dt_txt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Country  string `json:"country"`
		Timezone int    `json:"timezone"`
		Sunrise  int    `json:"sunrise"`
		Sunset   int    `json:"sunset"`
	} `json:"city"`
}
