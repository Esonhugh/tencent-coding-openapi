package git

type CreateSshKeyReq struct {
	Action         string `json:"Action"`         // CreateSshKey
	Content        string `json:"Content"`        // ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEPWDzF9gULVngb8Bv/SEGXVa6i5I1mFQK9nuUib0fbR4rFAvfBhbJtHoed4O+/KOc3Sxw/A0dQB6uZB7jfzw8RY5dtOvEkI1WlZ8/+7cC7Q14DkzElfCNodha/kP5dS23cIYZdESfuorWCRJZCq5CWoaV9io2DCdD1cvKbAr1CcdBNJ2UODvqxVllLnSw7rZW6D4hdJFjcLAP+ZaoZb1tqpkuS8I5QiF5smuMeLx8uGdM0/C9Pe0EXcsLk7N/YwqO4F6KnPj1TltLqs7O5hXN8BuMsLT90O2EHynq6Z1/ErSaCIszXQJQsYTJr5bJsaRZAkucDnldkX0YocnerNfH0PH+NwB0JT8Q/v9ao2FFDGG+u0fSYQ4fnOsANhNMnwr3hbQjg/S01DZBS6XryiL+MyDWB2vUgNxXQzCUGhwY2Vn9A8J9sfydTDDIQL/RDkDaBe7e4wIMqXvSP9UB/jXV2IZsbzq+QHcXW3SZJ8ueGcwI6krzlsXHZtw7xkUZhrlDHlP0q5nrrm7NZzkduwDFDZ8Pm4vig98+Wj/KDPKcig3b/y/q8XJKTK0D46x+85rS7Ti7157wltUuyqyTgoGnMZuGboU3zRvbbhEWF3PbDTCWbx+4E/YVubzq1MWQjmtglBsc7IHx79+5+D8MNrYb086WUUVIjHfzCoyAHcu99w== lichen@coding.net
	ExpirationDate string `json:"ExpirationDate"` // 2022-09-09
	Title          string `json:"Title"`          // Hello
}

func (req *CreateSshKeyReq) SetAction() string {
	req.Action = "CreateSshKey"
	return req.Action
}

type CreateSshKeyResp struct {
	Response struct {
		RequestID  string `json:"RequestId"` // 8643a755-559e-604e-6420-fe5cc74e0ec2
		SshKeyInfo struct {
			CreatedAt      int64  `json:"CreatedAt"`      // 1.634808370882e+12
			ExpirationDate string `json:"ExpirationDate"` // 2022-09-09
			FingerPrint    string `json:"FingerPrint"`    // 0c:36:f5:60:fd:da:3b:be:22:49:38:96:64:50:49:9f
			HasExpired     bool   `json:"HasExpired"`     // false
			ID             int64  `json:"Id"`             // 16
			OwnerID        int64  `json:"OwnerId"`        // 44
			Title          string `json:"Title"`          // Hello
		} `json:"SshKeyInfo"`
	} `json:"Response"`
}

// CreateSshKey 导入用户 SSH 公钥
func (g *GitClient) CreateSshKey(req CreateSshKeyReq) (resp CreateSshKeyResp, err error) {
	err = g.Call(&req, &resp)
	return
}
