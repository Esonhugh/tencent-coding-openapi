package cd

type DescribeCdHostServerGroupReq struct {
	Action string `json:"Action"` // DescribeCdHostServerGroup
	ID     int64  `json:"Id"`     // 41
}

func (req *DescribeCdHostServerGroupReq) SetAction() string {
	req.Action = "DescribeCdHostServerGroup"
	return req.Action
}

type DescribeCdHostServerGroupResp struct {
	Response struct {
		Data struct {
			HostServerGroup struct {
				Account      string `json:"Account"` // cred-10-41
				AgentMachine struct {
					ID     int64  `json:"Id"`     // 10
					Name   string `json:"Name"`   // master
					Status string `json:"Status"` // SUCCESS
				} `json:"AgentMachine"`
				AuthMethod  string   `json:"AuthMethod"`  // PUBLIC_KEY
				DisplayName string   `json:"DisplayName"` // test-1
				ID          int64    `json:"Id"`          // 41
				Ips         []string `json:"Ips"`         // 127.0.0.1
				Labels      []struct {
					Key   string `json:"Key"`   // testKey
					Value string `json:"Value"` // testValue
				} `json:"Labels"`
				Port     int64  `json:"Port"`     // 36000
				UserName string `json:"UserName"` // root
			} `json:"HostServerGroup"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 175127c8-bc19-5a38-b6c9-8f1a46f15c66
	} `json:"Response"`
}

// DescribeCdHostServerGroup 获取 CD 主机组
func (c *CdClient) DescribeCdHostServerGroup(req DescribeCdHostServerGroupReq) (resp DescribeCdHostServerGroupResp, err error) {
	err = c.Call(&req, &resp)
	return
}
