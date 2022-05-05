package models

import "time"

type Community struct {
	CommunityId   int64  `json:"community_id" `
	CommunityName string `json:"community_name"`
}

type CommunityDetail struct {
	CommunityId   int64     `json:"community_id"`
	CommunityName string    `json:"community_name" `
	Introduction  string    `json:"introduction,omitempty"`
	CreateTime    time.Time `json:"create_time"`
}
