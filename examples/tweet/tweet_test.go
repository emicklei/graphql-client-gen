package tweet

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestUserQuery(t *testing.T) {
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
	}
	data, _ := json.Marshal(q)
	t.Log(string(data))
	doReflect(q)
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

func doReflect(q interface{}) {
	rt := reflect.TypeOf(q)
	rv := reflect.ValueOf(q)
	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		if !fv.IsZero() {
			sf := rt.Field(i)
			fmt.Println(sf.Tag.Get("graphql"))
			// is struct or pointer to struct
			k := sf.Type
			if k.Kind() == reflect.Pointer {
				k = k.Elem()
			}
			if k.Kind() == reflect.Struct {
				fmt.Println("struct")
			}
		}
	}
}
