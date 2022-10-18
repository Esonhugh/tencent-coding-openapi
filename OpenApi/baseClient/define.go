package baseClient

// ProjectObject define
/*
`
	{
	"Name": "coding-demo",
	"Id": 2,
	"Type": 2,
	"DisplayName": "示例项目",
	"Icon": "https://codehub.cn/7971.png",
	"Description": "CODING 示例项目",
	"CreatedAt": 1572933083682,
	"MaxMember": 0,
	"TeamId": 12,
	"UserOwnerId": 0,
	"IsDemo": true,
	"Archived": false,
	"StartDate": 0,
	"UpdatedAt": 1572933083682,
	"TeamOwnerId": 0,
	"EndDate": 0,
	"Status": 1
}`
*/
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

// DeoptObject define.
/*
`
	{
	"Id": 8521598,
	"Name": "demo",
	"HttpsUrl": "https://e.coding.net/codingcorp/pual_test/demo.git",
	"ProjectId": 8104562,
	"SshUrl": "git@e.coding.net:codingcorp/pual_test/demo.git",
	"WebUrl": "https://codingcorp.coding.net/p/pual_test/d/demo",
	"VcsType": "git"
	}
`
*/
type DeoptObject struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	HttpsUrl  string `json:"HttpsUrl"`
	ProjectId int    `json:"ProjectId"`
	SshUrl    string `json:"SshUrl"`
	WebUrl    string `json:"WebUrl"`
	VscType   string `json:"VscType"`
}
