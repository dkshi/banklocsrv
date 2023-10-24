package banklocsrv

type Office struct {
	SalePointName       string        `json:"salePointName" bson:"salePointName"`
	Address             string        `json:"address" bson:"address"`
	Status              string        `json:"status" bson:"status"`
	OpenHours           []OfficeHours `json:"openHours" bson:"openHours"`
	OpenHoursIndividual []OfficeHours `json:"openHoursIndividual" bson:"openHoursIndividual"`
	Rko                 string        `json:"rko" bson:"rko"`
	OfficeType          string        `json:"officeType" bson:"officeType"`
	SalePointFormat     string        `json:"salePointFormat" bson:"salePointFormat"`
	SuoAvailability     string        `json:"suoAvailability" bson:"suoAvailability"`
	HasRamp             string        `json:"hasRamp" bson:"hasRamp"`
	Latitude            float64       `json:"latitude" bson:"latitude"`
	Longitude           float64       `json:"longitude" bson:"longitude"`
	MetroStation        string        `json:"metroStation" bson:"metroStation"`
	Distance            int           `json:"distance" bson:"distance"`
	Kep                 bool          `json:"kep" bson:"kep"`
	MyBranch            bool          `json:"myBranch" bson:"myBranch"`
	Load                int           `json:"load" bson:"load"`
}

type OfficeHours struct {
	Days  string `json:"days" bson:"days"`
	Hours string `json:"hours" bson:"hours"`
}

type OfficesData struct {
	Offices []Office `json:"offices" bson:"offices"`
}
