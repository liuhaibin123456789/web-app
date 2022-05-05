package model

type ReqVoteData struct {
	PostId    int64   `json:"post_id,string" binding:"required" form:"post_id"`
	Direction float64 `json:"direction" binding:"required" form:"direction"`
}

type ReqPost struct {
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	CommunityId int64  `json:"community_id,string"`
}
