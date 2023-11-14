package folders

import (
	"github.com/gofrs/uuid"
	"errors"
)

// Retrieves all folders based on the request and converts it into a FetchFolderResponse
// Takes a pointer to the FetchFolderRequest struct
// Returns a pointer to the FetchFolderResponse struct and an error
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var err error

	// Initialise an empty slice to store Folders fetched by FetchAllFoldersByOrgID
	f := []Folder{}

	// Fetch folders and handle errors
	r, err:= FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// Iterate through the fetched folders and append each Folder to new slice
	for _, v := range r {
		f = append(f, *v)
	}

	var fp []*Folder

	// Iterate through the slice of Folders and append the Folder pointer to fp
	for _, v1 := range f {
		var f1 Folder = v1 // We redeclare this variable at every iteration to give it a unique memory address 
		fp = append(fp, &f1)
	}

	// Create the FetchFolderResponse pointer
	var ffr *FetchFolderResponse

	// Assign ffr the address of FetchFolderResponse struct and initiliase struct member with fp 
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// Retrieves a list of folders associated with a specific orgID
// Takes orgID as a parameter
// Returns a slice of Folder pointers and an error
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// Obtain sample list of Folders
	folders := GetSampleData()

	// Initialise empty slice of Folder pointers to store folders associated with orgID
	resFolder := []*Folder{}

	// Iterate through each folder in sample data
	for _, folder := range folders {
		// Append folder to slice if folder orgID matches specified orgID
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	// Return an error if there is no matching folders with the requested orgID
	if len(resFolder) == 0 {
		return nil, errors.New("Error: No folders found for the specified orgID")
	}

	return resFolder, nil
}
