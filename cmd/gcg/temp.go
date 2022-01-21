package main

/**
func AddEmploymentDocument2(
	ctx context.Context,
	employmentID ID,
	fileID FileID,
	repositoryID ID) (result Employment, err error) {
	var m struct {
		AddEmploymentDocument struct {
			employmentID ID
			fileID       FileID
			repositoryID ID
		} `graphql:"addEmploymentDocument(employmentID: $employmentID, fileID: $fileID, repositoryID: $repositoryID)"`
	}
	vars := map[string]interface{}{
		"employmentID": employmentID,
		"fileID":       fileID,
		"repositoryID": repositoryID,
	}
	// err := client.Mutate(ctx, &m, vars)
	log.Println(vars, m)
	return
}

**/
