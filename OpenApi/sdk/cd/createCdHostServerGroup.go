package cd

type CreateCdHostServerGroupReq struct {
	Action         string   `json:"Action"`         // CreateCdHostServerGroup
	AgentMachineID int64    `json:"AgentMachineId"` // 10
	AuthMethod     string   `json:"AuthMethod"`     // PUBLIC_KEY
	DisplayName    string   `json:"DisplayName"`    // test
	Ips            []string `json:"Ips"`            // 127.0.0.1
	Labels         []struct {
		Key   string `json:"Key"`   // testKey
		Value string `json:"Value"` // testValue
	} `json:"Labels"`
	Port     int64  `json:"Port"`     // 36000
	UserName string `json:"UserName"` // root
}

func (req *CreateCdHostServerGroupReq) SetAction() string {
	req.Action = "CreateCdHostServerGroup"
	return req.Action
}

type CreateCdHostServerGroupResp struct {
	Response struct {
		Data struct {
			HostServerGroup struct {
				Account      string `json:"Account"` // cred-10-41
				AgentMachine struct {
					ID     int64  `json:"Id"`     // 10
					Name   string `json:"Name"`   // master
					Status string `json:"Status"` // SUCCESS
				} `json:"AgentMachine"`
				DisplayName string `json:"DisplayName"` // test
				ID          int64  `json:"Id"`          // 41
			} `json:"HostServerGroup"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 9d8775ae-cd91-d0ca-d0b2-24761bb073c7
	} `json:"Response"`
}

// CreateCdHostServerGroup 添加 CD 主机组
func (c *CdClient) CreateCdHostServerGroup(req CreateCdHostServerGroupReq) (resp CreateCdHostServerGroupResp, err error) {
	err = c.Call(&req, &resp)
	return
}
