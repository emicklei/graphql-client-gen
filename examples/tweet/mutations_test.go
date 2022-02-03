package tweet

import "testing"

/**
mutation createTweet($body:String)
        {	createTweet(body: $body) {
        		id
        	}

        }
**/
func TestCreateTweet(t *testing.T) {
	m := CreateTweetMutation{}
	m.Data.ID = "?"
	req := m.Build("hello")
	t.Log("\n", req.Query)
}
