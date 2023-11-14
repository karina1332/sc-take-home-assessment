package folders_test

import (
	"testing"
	"github.com/gofrs/uuid"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	// Test case 1: There is a matching OrgID in the Folders
	t.Run("Test 1: Folders match OrgID", func(t *testing.T) {
		// Create test data that exists in sample.json
		testOrgID := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		testReq := &folders.FetchFolderRequest {
			OrgID: testOrgID,
		}

		result, err := folders.GetAllFolders(testReq)

		assert.NoError(t, err, "Unexpected error") 								 // Checks that no error is returned
		assert.NotNil(t, result, "Expected a non-nil result for matching OrgID") // Checks that the result is not nil
		assert.NotEmpty(t, result.Folders, "Expected non-empty list of folders") // Check that the result is not empty
	})

	// Test Case 2: There is no matching OrgID in the folders
	t.Run("Test 2: Folders do not match OrgID", func(t *testing.T) {
		// Create test data that does not exist in sample.json
		testOrgID := uuid.FromStringOrNil("7ee73e98-b5a7-4ff5-a710-bfd8077ac0a9")
		testReq := &folders.FetchFolderRequest {
			OrgID: testOrgID,
		}

		result, err := folders.GetAllFolders(testReq)

		assert.Error(t, err, "Expected an error for non-matching OrgID") 	  // Checks that an error has occured
		assert.Nil(t, result, "Expected a nil result for non-matching OrgID") // Checks that the result is nil
	})

	// Test Case 3: The provided OrgID is invalid
	t.Run("Test 3: Invalid OrgID", func(t *testing.T) {
		// Create test data
		testOrgID := uuid.FromStringOrNil("")
		testReq := &folders.FetchFolderRequest {
			OrgID: testOrgID,
		}

		result, err := folders.GetAllFolders(testReq)

		assert.Error(t, err, "Expected an error for invalid OrgID")		 // Checks that an error has occured
		assert.Nil(t, result, "Expected a nil result for invalid OrgID") // Checks that the result is nil
	})


}
