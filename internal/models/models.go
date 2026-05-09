package models

type Point struct{
	Lat 	 	float64  `json:"lat"`	// создали структуру для координат потом засунули в Location 
	Lon 	 	float64  `json:"lon"`
}

type Location struct {
	ID       	string   `json:"id"`
	Name     	string   `json:"name"`
	Address  	string   `json:"address_name"`
	Point 	 	Point	  `json:"point"` // все координаты
}

type ViewLocation struct{
	ID 			string 		
	Name 		string 	
	Address 	string	
	Lat			float64 // обьявление координат
	Lon 		float64
	MapURL  	string // статитческий URL карты
}

type APIResponse struct {
	Result struct {
		Items []Location `json:"items"`
	} `json:"result"`
}