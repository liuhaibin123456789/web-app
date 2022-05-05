package model

type ResCommunity struct {
	CommunityId   int64  `json:"community_id,string"`
	CommunityName string `json:"community_name"`
}

// ResPost 获取帖子详情时，需要回复的结构体
type ResPost struct {
	UserName  string  `json:"user_name"`
	VoteScore float64 `json:"vote_score"`
	*Post
	*Community
}

type ResRegister struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
}
