package cd

type ModifyCdHostServerGroupReq struct {
	Action         string   `json:"Action"`         // ModifyCdHostServerGroup
	AgentMachineID int64    `json:"AgentMachineId"` // 10
	AuthMethod     string   `json:"AuthMethod"`     // PUBLIC_KEY
	DisplayName    string   `json:"DisplayName"`    // test-1
	ID             int64    `json:"Id"`             // 41
	Ips            []string `json:"Ips"`            // 127.0.0.1
	Labels         []struct {
		Key   string `json:"Key"`   // testKey
		Value string `json:"Value"` // testValue
	} `json:"Labels"`
	Port     int64  `json:"Port"`     // 36000
	UserName string `json:"UserName"` // root
}

func (req *ModifyCdHostServerGroupReq) SetAction() string {
	req.Action = "ModifyCdHostServerGroup"
	return req.Action
}

type ModifyCdHostServerGroupResp struct {
	Response struct {
		Data struct {
			HostServerGroup struct {
				Account      string `json:"Account"` // cred-10-41
				AgentMachine struct {
					ID     int64  `json:"Id"`     // 10
					Name   string `json:"Name"`   // master
					Status string `json:"Status"` // SUCCESS
				} `json:"AgentMachine"`
				DisplayName string `json:"DisplayName"` // test-1
				ID          int64  `json:"Id"`          // 41
			} `json:"HostServerGroup"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 10906500-927d-82ab-7e95-3f5242712f12
	} `json:"Response"`
}

// ModifyCdHostServerGroup 修改 CD 主机组
func (c *CdClient) ModifyCdHostServerGroup(req ModifyCdHostServerGroupReq) (resp ModifyCdHostServerGroupResp, err error) {
	err = c.Call(&req, &resp)
	return
}
