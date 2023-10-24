package banklocsrv

type Atm struct {
	Address   string                `json:"address" bson:"address"`
	Latitude  float64               `json:"latitude" bson:"latitude"`
	Longitude float64               `json:"longitude" bson:"longitude"`
	AllDay    bool                  `json:"allDay" bson:"allDay"`
	Services  map[string]AtmService `json:"services" bson:"services"`
}

type AtmService struct {
	ServiceCapability string `json:"serviceCapability" bson:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity" bson:"serviceActivity"`
}

type AtmsData struct {
	Atms []Atm `json:"atms"`
}
