package cd

type DeleteCdHostServerGroupReq struct {
	Action string `json:"Action"` // DeleteCdHostServerGroup
	ID     int64  `json:"Id"`     // 41
}

func (req *DeleteCdHostServerGroupReq) SetAction() string {
	req.Action = "DeleteCdHostServerGroup"
	return req.Action
}

type DeleteCdHostServerGroupResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // a3d7e086-df2b-02a5-46f8-68bf99e2f514
	} `json:"Response"`
}

// DeleteCdHostServerGroup 删除 CD 主机组
func (c *CdClient) DeleteCdHostServerGroup(req DeleteCdHostServerGroupReq) (resp DeleteCdHostServerGroupResp, err error) {
	err = c.Call(&req, &resp)
	return
}
