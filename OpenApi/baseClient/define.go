package baseClient

type ProjectObject struct {
	Id          int    `json:"Id"`
	CreatedAt   int64  `json:"CreatedAt"`
	UpdatedAt   int64  `json:"UpdatedAt"`
	Status      int    `json:"Status"`
	Type        int    `json:"Type"`
	MaxMember   int    `json:"MaxMember"`
	Name        string `json:"Name"`
	DisplayName string `json:"DisplayName"`
	Description string `json:"Description"`
	Icon        string `json:"Icon"`
	TeamOwnerId int    `json:"TeamOwnerId"`
	UserOwnerId int    `json:"UserOwnerId"`
	StartDate   int    `json:"StartDate"`
	EndDate     int    `json:"EndDate"`
	TeamId      int    `json:"TeamId"`
	IsDemo      bool   `json:"IsDemo"`
	Archived    bool   `json:"Archived"`
}
