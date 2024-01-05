package repository

type Port struct {
	Id          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Timezone    string    `json:"timezone"`
	Province    *string   `json:"province"`
	Coordinates []float32 `json:"coordinates"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Unlocs      []string  `json:"unlocs"`
}

type PortSummary struct {
	ReceivedTotal uint64
}
