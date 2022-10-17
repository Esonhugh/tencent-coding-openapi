package TeamSDK

type DescribeCodingProjectsRequest struct {
	Action      string `json:"Action"`
	ProjectName string `json:"ProjectName"`
	PageNumber  int    `json:"PageNumber"`
	PageSize    int    `json:"PageSize"`
}

func (req *DescribeCodingProjectsRequest) SetAction() string {
	req.Action = "DescribeCodingProjects"
	return req.Action
}

type DescribeCodingProjectsResponse struct {
	Response struct {
		RequestId string `json:"RequestId"`
		Data      struct {
			PageNumber  int `json:"PageNumber"`
			PageSize    int `json:"PageSize"`
			TotalCount  int `json:"TotalCount"`
			ProjectList []struct {
				Id          int    `json:"Id"`
				CreatedAt   int64  `json:"CreatedAt"`
				UpdatedAt   int64  `json:"UpdatedAt"`
				Status      int    `json:"Status"`
				Type        int    `json:"Type"`
				MaxMember   int    `json:"MaxMember"`
				Name        string `json:"Name"`
				DisplayName string `json:"DisplayName"`
				Description string `json:"Description"`
				Icon        string `json:"Icon"`
				TeamOwnerId int    `json:"TeamOwnerId"`
				UserOwnerId int    `json:"UserOwnerId"`
				StartDate   int    `json:"StartDate"`
				EndDate     int    `json:"EndDate"`
				TeamId      int    `json:"TeamId"`
				IsDemo      bool   `json:"IsDemo"`
				Archived    bool   `json:"Archived"`
			} `json:"ProjectList"`
		} `json:"Data"`
	} `json:"Response"`
}

func (c *TeamSdkClient) DescribeCodingProjects(req DescribeCodingProjectsRequest) (rsp DescribeCodingProjectsResponse, e error) {
	e = c.Call(&req, &rsp)
	return
}
