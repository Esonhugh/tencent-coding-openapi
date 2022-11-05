package define

import (
	"fmt"
	"github.com/esonhugh/tencent-coding-openapi/utils/Print"
)

// DeoptObject define.
type DeoptObject struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	HttpsUrl  string `json:"HttpsUrl"`
	ProjectId int    `json:"ProjectId"`
	SshUrl    string `json:"SshUrl"`
	WebUrl    string `json:"WebUrl"`
	VscType   string `json:"VscType"`
}
type DeoptObjectList []DeoptObject

func (p DeoptObjectList) PrintSelf() {
	var t Print.Table
	t.Header = []string{"ID", "Name", "HttpsUrl", "ProjectId", "SshUrl", "WebUrl", "VscType"}
	for _, each := range p {
		t.Body = append(t.Body,
			[]string{fmt.Sprintf("%v", each.Id),
				each.Name, each.HttpsUrl,
				fmt.Sprintf("%v", each.ProjectId),
				each.SshUrl,
				each.WebUrl,
				each.VscType})
	}
	t.Print("")
}
func ConvertDeoptObjectList(p []struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	HttpsUrl  string `json:"HttpsUrl"`
	ProjectId int    `json:"ProjectId"`
	SshUrl    string `json:"SshUrl"`
	WebUrl    string `json:"WebUrl"`
	VscType   string `json:"VscType"`
}) DeoptObjectList {
	var r DeoptObjectList
	for _, v := range p {
		r = append(r, v)
	}
	return r
}
