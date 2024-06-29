package pkg

import (
	"dir_to_json/model"
	"encoding/json"
	"os"
	"path/filepath"
)

// Function DirToJson converts a directory structure to a JSON string.
// It takes a path to a directory, a level of depth to explore and a boolean to show hidden files.
// It returns a JSON string and an error.
func DirToJson(path string, level int, showHide bool) (string, error) {

	// Create file explore model
	fileExplore := model.FileExplore{
		Path:   path,
		Level:  level,
		Hidden: showHide,
	}

	// Build file structure
	rootFileStructure, err := buildFileStructure(fileExplore)
	if err != nil {
		return "", err
	}

	// Convert to json
	jsonData, err := json.Marshal(rootFileStructure)
	if err != nil {
		return "", err
	}

	// Return json string
	return string(jsonData), nil
}

// build File Structure recursively.
func buildFileStructure(fileExploremodel model.FileExplore) (model.FileStructure, error) {

	// Get info path root
	fileInfo, err := os.Stat(fileExploremodel.Path)
	if err != nil {
		return model.FileStructure{}, err
	}

	// Create file structure
	fileStruct := model.FileStructure{
		Name:  fileInfo.Name(),
		IsDir: fileInfo.IsDir(),
		Path:  fileExploremodel.Path,
	}

	// If level is 0, return
	if fileExploremodel.Level == 0 {
		return fileStruct, nil
	}

	// If is a directory, get children
	if fileInfo.IsDir() {

		// Get children files
		files, err := os.ReadDir(fileExploremodel.Path)
		if err != nil {
			return model.FileStructure{}, err
		}

		// Iterate over children files
		for _, file := range files {

			// hidden files
			if !fileExploremodel.Hidden {
				// Ignore hide files
				if file.Name()[:1] == "." {
					continue
				}
			}

			// Build child file structure
			childPath := filepath.Join(fileExploremodel.Path, file.Name())

			// Create child file explore model
			childFileExplore := model.FileExplore{
				Path:   childPath,
				Level:  fileExploremodel.Level - 1,
				Hidden: fileExploremodel.Hidden,
			}

			// Recursively build child file structure
			childStruct, err := buildFileStructure(childFileExplore)
			if err != nil {
				return model.FileStructure{}, err
			}

			// Append child file structure to parent
			fileStruct.Children = append(fileStruct.Children, childStruct)
		}
	} else {
		// If is a file, get size and set in file structure
		fileStruct.Size = fileInfo.Size()
	}

	// Return file structure
	return fileStruct, nil
}
