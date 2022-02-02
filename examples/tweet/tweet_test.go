package tweet

import (
	"testing"
)

/**
	id
	body
	Author {
			username
	}
	Stats {
			views
			likes
	}
	Responders(limit:10) {
		id
	}
**/
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
			Limit: 10,
		},
	}
	tv := map[string]valueAndType{}
	s, v := buildQuery("doit", q, tv)
	t.Log("\n", s)
	t.Log("\n", v)
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
