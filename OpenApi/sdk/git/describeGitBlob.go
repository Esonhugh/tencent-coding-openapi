package git

type DescribeGitBlobReq struct {
	Action  string `json:"Action"`  // DescribeGitBlob
	BlobSha string `json:"BlobSha"` // a7954ee405100719a59c1e03f4ac98ce92a2a257
	DepotID int64  `json:"DepotId"` // 2
}

func (req *DescribeGitBlobReq) SetAction() string {
	req.Action = "DescribeGitBlob"
	return req.Action
}

type DescribeGitBlobResp struct {
	Response struct {
		Data struct {
			BlobSha  string `json:"BlobSha"`  // a7954ee405100719a59c1e03f4ac98ce92a2a257
			Content  string `json:"Content"`  // bmFtZTogZXh0ZW5kZWRfdGV4dApkZXNjcmlwdGlvbjogRXh0ZW5kZWQgb2ZmaWNpYWwgdGV4dCB0byBidWlsZCBzcGVjaWFsIHRleHQgbGlrZSBpbmxpbmUgaW1hZ2Ugb3IgQHNvbWVib2R5IHF1aWNrbHksaXQgYWxzbyBzdXBwb3J0IGN1c3RvbSBiYWNrZ3JvdW5kLGN1c3RvbSBvdmVyIGZsb3cgYW5kIGN1c3RvbSBzZWxlY3Rpb24gdG9vbGJhciBhbmQgaGFuZGxlcy4KdmVyc2lvbjogNS4wLjQKaG9tZXBhZ2U6IGh0dHBzOi8vZ2l0aHViLmNvbS9mbHV0dGVyY2FuZGllcy9leHRlbmRlZF90ZXh0CgplbnZpcm9ubWVudDoKICBzZGs6ICc+PTIuMTIuMCA8My4wLjAnCiAgZmx1dHRlcjogIj49Mi4wLjAiCgpkZXBlbmRlbmNpZXM6CiAgZXh0ZW5kZWRfdGV4dF9saWJyYXJ5OiBeNS4wLjAKICBmbHV0dGVyOgogICAgc2RrOiBmbHV0dGVyCgpkZXZfZGVwZW5kZW5jaWVzOgogIGZsdXR0ZXJfdGVzdDoKICAgIHNkazogZmx1dHRlcg==
			Encoding string `json:"Encoding"` // base64
			Size     int64  `json:"Size"`     // 469
		} `json:"Data"`
		RequestID string `json:"RequestId"` // acba0c73-8dc2-c30d-7753-eac38d9b6bf6
	} `json:"Response"`
}

// DescribeGitBlob 查询 Git Blob 信息
func (g *GitClient) DescribeGitBlob(req DescribeGitBlobReq) (resp DescribeGitBlobResp, err error) {
	err = g.Call(&req, &resp)
	return
}
