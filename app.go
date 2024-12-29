package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {

	// **IMPORTANT:** Update this to your desired root folder path.
	rootFolder := `C:\Users\Yura\Desktop\OneDriveYulia\OneDrive\Зображення\Фотографії з камери`

	fileNames, err := getAllFileNames(rootFolder)
	if err != nil {
		fmt.Printf("Error getting file names: %v\n", err)
		return
	}

	// Extract only file names from paths
	fileNamesOnly := make([]string, len(fileNames))
	for i, filePath := range fileNames {
		fileNamesOnly[i] = filepath.Base(filePath)[9:] // **IMPORTANT:** [9:] removes first 9 characters of the file name.
	}

	// Sort the file names
	sort.Strings(fileNamesOnly)

	// Write the sorted file names to results.txt
	err = writeToFile("results.txt", fileNamesOnly)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Println("File names written to results.txt")
}

// getAllFileNames recursively traverses the folder and retrieves all file names
func getAllFileNames(root string) ([]string, error) {
	var fileNames []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Propagate the error, don't stop execution
		}

		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileNames, nil
}

// writeToFile writes a slice of strings to a file, one line per string
func writeToFile(fileName string, lines []string) error {
	fileContent := strings.Join(lines, "\n")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, fileContent)
	return err
}
