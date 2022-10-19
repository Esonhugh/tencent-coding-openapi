package git

type CreateGitDeployKeyReq struct {
	Action         string `json:"Action"`         // CreateGitDeployKey
	AllowWrite     bool   `json:"AllowWrite"`     // true
	Content        string `json:"Content"`        // ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDdNcUV5lLNsPHsHhfEut18aMBEJ4ifzMJ0utoxDOPqgDyh7nhEzGKYQooORN/F8u1WJ4P7Jcn5H7kgkiYn18qlg69p3ZslKw3UmBtY7CEryvZmcwZi/mKlVXHJN4hLvdiPxVEfMNjM4IZM/UMxCd3m4bSBybz4K1p/IiXca1foQkJrdZoOcfXhTMpjRjL2FMf2lYnLXA71DbgVYR3ACQZGTZQ4A6gPjQ1hvfklGpKTdAwgcegKTdqka3Yiv7GG0ekZWGuHAo/ddY+Nhlz8/UNoRSzYBLkhDRYkuhZrIi2a8/wI1BWpmPuCM8cmum0EuPONrxW0criY3wrl5i+JbZqn8N/tFU0GbxQrIk8ID2+Az/T6VEsTzqWPVm4hjHyCJXgBye9BA3YH1nYPAWy1yf6tzwj53ibLmhrxleUHZmJGtiEFtH4bvdIIBEGhIMiSK9tWb3jH/Ncq+FGTYeK871hKZBYjHJzcDRjGyGLvDuPe4uMDGDEQoeQZGOHkPtdKz3QoDKr78yPH1auadq/c1kmH98r3M4yebPdiuYndbQYZEqcbJDeWB/mPNQQxmARpjpqhUXctyMmj580MQPqxvHgxmmJFIp5vf8qALFRF/koN4sVcZxGHa0bjjhjmGmilL8GphqV3qSdn8dHgwHQ97Oi1I5M5qCxUcLcIN7uWYmMwvQ== your.email@example.com
	DepotID        int64  `json:"DepotId"`        // 2
	ExpirationDate string `json:"ExpirationDate"` // 2022-12-12
	Title          string `json:"Title"`          // deploy key 1
}

func (req *CreateGitDeployKeyReq) SetAction() string {
	req.Action = "CreateGitDeployKey"
	return req.Action
}

type CreateGitDeployKeyResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 39a5e264-2bdf-5840-37b1-e9587ea0ff17
	} `json:"Response"`
}

// CreateGitDeployKey 新建部署公钥
func (g *GitClient) CreateGitDeployKey(req CreateGitDeployKeyReq) (resp CreateGitDeployKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
