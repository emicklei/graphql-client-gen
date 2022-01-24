package tweet

import (
	"context"
	"encoding/json"
	"testing"
)

func TestTweetQuery(t *testing.T) {
	q := Tweet{
		ID:   Get.ID,
		Body: &Get.String,
		Author: &User{
			Username: &Get.String,
		},
		Stats: &Stat{
			Views: &Get.Int32,
			Likes: &Get.Int32,
		},
		Responders: &TweetRespondersFunction{
			//Limit: 10,
			User: User{
				ID: Get.ID,
			},
		},
	}
	data, _ := json.Marshal(q)
	t.Log(string(data))
	t.Log(BuildQuery(q))
}

func TestTweetMutation(t *testing.T) {
	Mutation.CreateTweet(mockClient{}, context.Background(), "body")
}

var Get = struct {
	String string
	ID     interface{}
	Int32  int32
}{
	"",
	0,
	0,
}

// Tweet(id: ID!): Tweet

type Kuery struct{}

func (q Kuery) Tweet(id interface{}) (Tweet, error) {
	return Tweet{}, nil
}

type mockClient struct{}

func (_ mockClient) Mutate(ctx context.Context, m interface{}, variables map[string]interface{}) error {
	return nil
}
