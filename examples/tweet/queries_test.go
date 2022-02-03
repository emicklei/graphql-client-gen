package tweet

import "testing"

/**
query testTweets($sort_order:String,$limit:Int,$skip:Int,$sort_field:String)
        {	Tweets(limit: $limit,skip: $skip,sort_field: $sort_field,sort_order: $sort_order) {
        		id
        	}

        }
**/
func TestTweetsQuery(t *testing.T) {
	q := TweetsQuery{}
	q.Data = []TweetsQueryData{
		{
			Tweet: Tweet{ID: "?"},
		},
	}
	req := q.Build("testTweets", 1, 0, "id", "desc")
	t.Log("\n", req.Query)
}
