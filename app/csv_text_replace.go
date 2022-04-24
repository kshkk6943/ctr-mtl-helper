package app

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/kshkk6943/ctr-mtl-helper/app/constants"
	"github.com/kshkk6943/ctr-mtl-helper/app/helper"
	"github.com/kshkk6943/ctr-mtl-helper/app/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type CsvTextReplace struct{}

func NewCsvTextReplace() CsvTextReplace {
	return CsvTextReplace{}
}
func validateTextToReplaceValue(textToReplace models.TextToReplace) bool {
	return textToReplace.Old != nil && textToReplace.New != nil
}

func (c CsvTextReplace) GetLibrary(folderRoot string) []models.TextToReplace {
	files, err := ioutil.ReadDir(folderRoot + constants.LibraryFolderName)
	if err != nil {
		os.Exit(helper.CloseDueToError(err))
	}
	var textToReplace []models.TextToReplace
	var hasError = false
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".csv") {
			fileLocation := folderRoot + constants.LibraryFolderName + file.Name()
			libraryFile, err := os.OpenFile(fileLocation, os.O_RDONLY, os.ModePerm)
			if err != nil {
				os.Exit(helper.CloseDueToError(err))
			}

			var textToReplaceFromFile []models.TextToReplace
			if err := gocsv.UnmarshalFile(libraryFile, &textToReplaceFromFile); err != nil {
				fmt.Println("CSV File has issue/s: " + fileLocation)
				fmt.Println("Error: " + err.Error())
				fmt.Println(constants.CMDLineSplitter)
				hasError = true
			}

			for _, textToReplaceSlice := range textToReplaceFromFile {
				if !validateTextToReplaceValue(textToReplaceSlice) {
					fmt.Println("CSV File has issue/s: " + fileLocation)
					fmt.Println("Error: Missing header/s")
					fmt.Println(constants.CMDLineSplitter)
					hasError = true
				}
			}

			_ = libraryFile.Close()

			if !hasError {
				textToReplace = append(textToReplace, textToReplaceFromFile...)
			}
		}
	}

	if hasError {
		os.Exit(helper.CloseApplication())
	}

	return textToReplace
}

func (c CsvTextReplace) RunTextReplace(folderRoot string) {
	textsToReplace := c.GetLibrary(folderRoot)
	files, err := ioutil.ReadDir(folderRoot + constants.UnprocessedFolderName)
	if err != nil {
		os.Exit(helper.CloseDueToError(err))
	}
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".txt") {
			fileLocation := folderRoot + constants.UnprocessedFolderName + file.Name()
			fileBytes, err := ioutil.ReadFile(fileLocation)
			if err != nil {
				os.Exit(helper.CloseDueToError(err))
			}
			stringContent := string(fileBytes)
			for _, textToReplace := range textsToReplace {
				stringContent = strings.ReplaceAll(
					stringContent,
					*textToReplace.Old,
					*textToReplace.New,
				)
			}
			filenameWithoutExtension := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			newFileLocation := folderRoot + constants.FinishedJobDestinationFolderName +
				filenameWithoutExtension + ".txt"
			newFile, err := os.OpenFile(newFileLocation,
				os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755,
			)
			_, _ = newFile.WriteString(stringContent)
			_ = newFile.Close()
		}
	}
}
