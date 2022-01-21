package main

/**
import (
	"context"
	"log"

	"github.com/shurcooL/graphql"
)

type Client struct {
	graphqlClient *graphql.Client
}

func (c *Client) AddEmploymentDocument2(
	ctx context.Context,
	employmentID ID,
	fileID FileID,
	repositoryID ID) (result Employment, err error) {
	type AddEmploymentDocument struct {
		*Employment  `json:"-"` // result embedded
		employmentID ID         `json:"employmentID"`
		fileID       FileID     `json:"fileID"`
		repositoryID ID         `json:"repositoryID"`
	}
	var m struct {
		X AddEmploymentDocument `graphql:"addEmploymentDocument(employmentID: $employmentID, fileID: $fileID, repositoryID: $repositoryID)"`
	}
	vars := AddEmploymentDocument{
		*Employment,
		employmentID,
		fileID,
		repositoryID,
	}
	err = c.graphqlClient.Mutate(ctx, &m, &vars)
	log.Println(vars, m, err)
	return
}
**/
