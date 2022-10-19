package git

type CreateMemberSshKeyReq struct {
	Action         string `json:"Action"`         // CreateMemberSshKey
	Content        string `json:"Content"`        // ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDdNcUV5lLNsPHsHhfEut18aMBEJ4ifzMJ0utoxDOPqgDyh7nhEzGKYQooORN/F8u1WJ4P7Jcn5H7kgkiYn18qlg69p3ZslKw3UmBtY7CEryvZmcwZi/mKlVXHJN4hLvdiPxVEfMNjM4IZM/UMxCd3m4bSBybz4K1p/IiXca1foQkJrdZoOcfXhTMpjRjL2FMf2lYnLXA71DbgVYR3ACQZGTZQ4A6gPjQ1hvfklGpKTdAwgcegKTdqka3Yiv7GG0ekZWGuHAo/ddY+Nhlz8/UNoRSzYBLkhDRYkuhZrIi2a8/wI1BWpmPuCM8cmum0EuPONrxW0criY3wrl5i+JbZqn8N/tFU0GbxQrIk8ID2+Az/T6VEsTzqWPVm4hjHyCJXgBye9BA3YH1nYPAWy1yf6tzwj53ibLmhrxleUHZmJGtiEFtH4bvdIIBEGhIMiSK9tWb3jH/Ncq+FGTYeK871hKZBYjHJzcDRjGyGLvDuPe4uMDGDEQoeQZGOHkPtdKz3QoDKr78yPH1auadq/c1kmH98r3M4yebPdiuYndbQYZEqcbJDeWB/mPNQQxmARpjpqhUXctyMmj580MQPqxvHgxmmJFIp5vf8qALFRF/koN4sVcZxGHa0bjjhjmGmilL8GphqV3qSdn8dHgwHQ97Oi1I5M5qCxUcLcIN7uWYmMwvQ== your.email@example.com
	ExpirationDate string `json:"ExpirationDate"` // 2025-12-31
	MemberUserID   int64  `json:"MemberUserId"`   // 2
	Title          string `json:"Title"`          // create a ssh key for member
}

func (req *CreateMemberSshKeyReq) SetAction() string {
	req.Action = "CreateMemberSshKey"
	return req.Action
}

type CreateMemberSshKeyResp struct {
	Response struct {
		RequestID  string `json:"RequestId"` // 7197fc1a-7493-0756-8da5-123fc3e80b65
		SshKeyInfo struct {
			CreatedAt      int64  `json:"CreatedAt"`      // 1.642398559986e+12
			ExpirationDate string `json:"ExpirationDate"` // 2025-12-31
			FingerPrint    string `json:"FingerPrint"`    // 32:a9:dd:de:4c:3f:df:6f:0a:43:23:48:23:06:d2:31
			HasExpired     bool   `json:"HasExpired"`     // false
			ID             int64  `json:"Id"`             // 2
			OwnerID        int64  `json:"OwnerId"`        // 2
			Title          string `json:"Title"`          // create a ssh key for member
		} `json:"SshKeyInfo"`
	} `json:"Response"`
}

// CreateMemberSshKey 为团队成员添加 SSH 公钥
func (g *GitClient) CreateMemberSshKey(req CreateMemberSshKeyReq) (resp CreateMemberSshKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
