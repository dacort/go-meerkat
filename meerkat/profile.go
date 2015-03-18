package meerkat

import (
	"fmt"
)

const (
	profileURL   = "https://resources.meerkatapp.co/users/%s/profile?v=2"
	followersURL = "https://social-cdn.meerkatapp.co/users/%s/followers?v=2"
)

type ProfileService struct {
	Client *Client
}

type Profile struct {
	Info  UserInfo  `json:"info"`
	Stats UserStats `json:"stats"`
}

type UserInfo struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	TwitterID   string `json:"twitterId"`
	Privacy     string `json:"privacy"`
	Bio         string `json:"bio"`
}

type UserStats struct {
	Streams        []UserStream `json:"streams"`
	StreamsCount   int          `json:"streamsCount"`
	FollowingCount int          `json:"followingCount"`
	FollowersCount int          `json:"followersCount"`
	Score          int          `json:"score"`
}

type UserStream struct {
	ID           string `json:"id"`
	EndTimestamp int    `json:"endTime"`
}

// A slimed down and flattened version of a user profile
type FollowerProfile struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	Username     string `json:"username"`
	ProfileImage string `json:"profileImage"`
	Score        int    `json:"score"`
}

func (p *ProfileService) Get(userID string) (*Profile, error) {
	u := fmt.Sprintf(profileURL, userID)

	req, err := p.Client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	profile := new(Profile)
	_, err = p.Client.Do(req, profile)
	return profile, err
}

func (p *ProfileService) Followers(userID string) (*[]FollowerProfile, error) {
	url := fmt.Sprintf(followersURL, userID)

	req, err := p.Client.NewRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	followers := new([]FollowerProfile)
	_, err = p.Client.Do(req, followers)
	return followers, err
}
