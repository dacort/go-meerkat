package meerkat

import "fmt"

const (
	broadcastURL = "http://resources.meerkatapp.co/broadcasts/%s/summary"
)

type BroadcastService struct {
	Client *Client
}

type Broadcast struct {
	Broadcaster BroadcastUser `json:"broadcaster"`
	Location    string        `json:"location"`
	Caption     string        `json:"caption"`
	TweetID     int64         `json:"tweetId"`

	Likes     int `json:"likesCount"`
	Comments  int `json:"commentsCount"`
	Restreams int `json:"restreamsCount"`
	Watchers  int `json:"watchersCount"`

	Status       string `json:"status"`
	EndTimestamp int    `json:"endTime"`
}

type BroadcastUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

func (p *BroadcastService) Get(broadcastID string) (*Broadcast, error) {
	b := fmt.Sprintf(broadcastURL, broadcastID)

	req, err := p.Client.NewRequest("GET", b, "")
	if err != nil {
		return nil, err
	}

	broadcast := new(Broadcast)
	_, err = p.Client.Do(req, broadcast)
	return broadcast, err
}
