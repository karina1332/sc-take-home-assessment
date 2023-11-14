package folders

import (
	"errors"
	"math/rand"
	"math"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
const defaultPageSize = 100

// Generates and returns three maps for pagination: token to index, index to token, and index to folder pointer
// Takes a slice of Folder pointers as input
// Returns the three maps
func CreateMaps(folders []*Folder) (map[string]int, map[int]string, map[int]*Folder, error) {
	// Error handling if folders is nil
	if folders == nil {
		return nil, nil, nil, errors.New("Input slice of folders is invalid")
	}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Calculate number of tokens required based on data size and page number
	noOfTokens := int(math.Ceil(float64(len(folders)) / float64(defaultPageSize)))

	// Declare the maps
	tokenToIndexMap := make(map[string]int)
	indexToTokenMap := make(map[int]string)
	indexToFolderMap := make(map[int]*Folder)

	// Create the maps with the length as number of tokens
	for i := 0; i < noOfTokens; i++ {
		// Generate a random token of length 5 using characters from the charset
		token := make([]byte, 5)
		for i := range token {
			token[i] = charset[rand.Intn(len(charset))]
		}

		// If it is on the last page, set token to "" since there is no more data left
		if i == noOfTokens - 1 {
			tokenToIndexMap[""] = i * defaultPageSize
			indexToTokenMap[i * defaultPageSize] = ""

		} else {
		// Assign token to its corresponding index
			tokenToIndexMap[string(token)] = i * defaultPageSize
			indexToTokenMap[i * defaultPageSize] = string(token)
		}
	}

	// Create the map which maps the index to its corresponding folder
	for key, value := range folders {
		indexToFolderMap[key] = value
	}

	return tokenToIndexMap, indexToTokenMap, indexToFolderMap, nil
}


// Retrieves a page of folders based on the provided token
// Takes three maps and a token as input
// Returns the page which is a slice of Folders, the new token, and an error
func GetFoldersWithPagination(tokenIndexMap map[string]int, indexTokenMap map[int]string, indexFolderMap map[int]*Folder, token string) ([]Folder, string, error) {
	var fs []Folder
	var startIndex int

	// Find the starting index based on the token
	startIndex, exists := tokenIndexMap[token]
	if !exists {
		return nil, "", errors.New("Invalid token provided")
	}
	
	// Iterate until the pageSize is reached and append the folder to result slice
	for i := startIndex; i < startIndex + defaultPageSize; i++ {
		// Break if the index exceeds the actual length of the data set as a whole
		if i >= len(indexFolderMap) {
			break
		}

		fs = append(fs, *indexFolderMap[i])
	}

	// Set new token for the next pagination based on the next index
	token = indexTokenMap[startIndex + defaultPageSize]

	return fs, token, nil
}
