package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// Create the maps for pagination and handle any errors
	tokenToIndexMap, indexToTokenMap, indexToFolderMap, err := folders.CreateMaps(res.Folders)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// Get token to retrieve first chunk of data
	token := indexToTokenMap[0]

	// Infinite loop that breaks once it reaches the last page
	for {
		// Retrieve the page and new token and handle any errors
		page, newToken, err := folders.GetFoldersWithPagination(tokenToIndexMap, indexToTokenMap, indexToFolderMap, token)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		// Print the page
		folders.PrettyPrint(page)

		// If there is no more pages, break the loop
		if token == "" {
			break
		}

		// Set token that points to next page
		token = newToken

		// Print the token retrieved
		fmt.Printf("\nToken: %v\n", token)

		// Wait for user input to fetch the next page
		fmt.Print("\nPress Enter to fetch the next page...")
		fmt.Scanln()
	}
}
