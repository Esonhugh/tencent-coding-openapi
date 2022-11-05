package define

import (
	"fmt"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
)

type MemberObject struct {
	Avatar          string `json:"Avatar"`          // http://e.coding.net/static/fruit_avatar/Fruit-4.png
	Email           string `json:"Email"`           // blockuser@gmail.com
	EmailValidation int64  `json:"EmailValidation"` // 1
	ID              int64  `json:"Id"`              // 6
	Name            string `json:"Name"`            // blockuser
	NamePinYin      string `json:"NamePinYin"`      // blockuser
	Phone           int64  `json:"Phone,string"`    // 13800138006
	PhoneValidation int64  `json:"PhoneValidation"` // 1
	Status          int64  `json:"Status"`          // -1
	TeamID          int64  `json:"TeamId"`          // 1
}
type MemberObjectList []MemberObject

func (p MemberObjectList) PrintSelf() {
	var t Print.Table
	t.Header = []string{"ID", "Name", "Email", "Phone", "Avatar"}
	for _, each := range p {
		t.Body = append(t.Body,
			[]string{fmt.Sprintf("%v", each.ID),
				each.Name, each.Email,
				fmt.Sprintf("%v", each.Phone),
				each.Avatar})
	}
	t.Print("")
}

func ConvertToMemberObjectList(p []struct {
	Avatar          string `json:"Avatar"`          // http://e.coding.net/static/fruit_avatar/Fruit-4.png
	Email           string `json:"Email"`           // blockuser@gmail.com
	EmailValidation int64  `json:"EmailValidation"` // 1
	ID              int64  `json:"Id"`              // 6
	Name            string `json:"Name"`            // blockuser
	NamePinYin      string `json:"NamePinYin"`      // blockuser
	Phone           int64  `json:"Phone,string"`    // 13800138006
	PhoneValidation int64  `json:"PhoneValidation"` // 1
	Status          int64  `json:"Status"`          // -1
	TeamID          int64  `json:"TeamId"`          // 1
}) MemberObjectList {
	var r MemberObjectList
	for _, v := range p {
		r = append(r, v)
	}
	return r
}
