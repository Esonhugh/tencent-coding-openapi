package define

import (
	"fmt"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
)

// ProjectObject define
type ProjectObject struct {
	Archived    bool   `json:"Archived"`    // false
	CreatedAt   int64  `json:"CreatedAt"`   // 1.572933083682e+12
	Description string `json:"Description"` // CODING 示例项目
	DisplayName string `json:"DisplayName"` // 示例项目
	EndDate     int64  `json:"EndDate"`     // 0
	Icon        string `json:"Icon"`        // https://codehub.cn/7971.png
	ID          int64  `json:"Id"`          // 2
	IsDemo      bool   `json:"IsDemo"`      // true
	MaxMember   int64  `json:"MaxMember"`   // 0
	Name        string `json:"Name"`        // coding-demo
	StartDate   int64  `json:"StartDate"`   // 0
	Status      int64  `json:"Status"`      // 1
	TeamID      int64  `json:"TeamId"`      // 12
	TeamOwnerID int64  `json:"TeamOwnerId"` // 0
	Type        int64  `json:"Type"`        // 2
	UpdatedAt   int64  `json:"UpdatedAt"`   // 1.572933083682e+12
	UserOwnerID int64  `json:"UserOwnerId"` // 0
}
type ProjectObjectList []ProjectObject

func (p ProjectObjectList) PrintSelf() {
	var t Print.Table
	t.Header = []string{"ID", "Name", "DisplayName", "Description"}
	for _, each := range p {
		t.Body = append(t.Body, []string{fmt.Sprintf("%v", each.ID), each.Name, each.DisplayName, each.Description})
	}
	t.Print("")
}
func ConvertProjectObjectList(p []struct {
	Archived    bool   `json:"Archived"`
	CreatedAt   int64  `json:"CreatedAt"`
	Description string `json:"Description"`
	DisplayName string `json:"DisplayName"`
	EndDate     int64  `json:"EndDate"`
	Icon        string `json:"Icon"`
	ID          int64  `json:"Id"`
	IsDemo      bool   `json:"IsDemo"`
	MaxMember   int64  `json:"MaxMember"`
	Name        string `json:"Name"`
	StartDate   int64  `json:"StartDate"`
	Status      int64  `json:"Status"`
	TeamID      int64  `json:"TeamId"`
	TeamOwnerID int64  `json:"TeamOwnerId"`
	Type        int64  `json:"Type"`
	UpdatedAt   int64  `json:"UpdatedAt"`
	UserOwnerID int64  `json:"UserOwnerId"`
}) ProjectObjectList {
	var r ProjectObjectList
	for _, v := range p {
		r = append(r, v)
	}
	return r
}
