package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pschlafley/fileFunctions"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	var title pterm.BigTextPrinter = pterm.DefaultBigText

	title.WithLetters(
		putils.LettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightGreen)),
		putils.LettersFromStringWithStyle("_", pterm.NewStyle(pterm.FgLightYellow)),
		putils.LettersFromStringWithStyle("TERMINAL", pterm.NewStyle(pterm.FgLightGreen))).
		Render()

	showMainMenu()
}

func getConfirmAnswer(answer bool) string {
	if answer {
		return "Yes"
	}
	return "No"
}

func confirmCreateDirectoryAndFile(fileName, path string, searchForFileSpinner pterm.SpinnerPrinter, returnToMainMenu pterm.InteractiveSelectPrinter) {
	searchForFileSpinner.Fail("No such file or directory found.\n")
	confirmCreateDirectory, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("There was no directory found at the given path. Would you like to create the directory and file?").WithTextStyle(pterm.FgYellow.ToStyle()).Show()

	if getConfirmAnswer(confirmCreateDirectory) == "Yes" {
		createDirectorySpinner, _ := pterm.DefaultSpinner.Start("Creating the directory at the given path.")
		time.Sleep(time.Second * 2)

		fileFunctions.CreateDirectory(path)

		createDirectorySpinner.Success("Directory created!\n")

		showConfirmCreateFileSpinner(fileName, path)

		showMainMenu()
	} else if getConfirmAnswer(confirmCreateDirectory) == "No" {
		showMainMenu()
	}

}

func fileWasNotFoundFunc(fileName, path string, searchForFileSpinner pterm.SpinnerPrinter) {
	searchForFileSpinner.Fail("File Not Found.\n")
	createFileIfNotFound, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("The file was not found at the given path! Would you like to create it?\n").WithTextStyle(pterm.FgYellow.ToStyle()).Show()

	if getConfirmAnswer(createFileIfNotFound) == "Yes" {
		createFileSpinnerIfFileNotFound, _ := pterm.DefaultSpinner.Start("Creating the file at the given path.")
		time.Sleep(time.Second * 2)
		fileFunctions.CreateFile(fileName, path)
		createFileSpinnerIfFileNotFound.Success("File created!\n")
	} else {
		os.Exit(1)
	}
}

func showConfirmCreateFileSpinner(fileName, path string) {
	confirmCreateFile, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Would you also like to create the file in the directory?").WithTextStyle(pterm.FgCyan.ToStyle()).Show()

	if getConfirmAnswer(confirmCreateFile) == "Yes" {
		createFileSpinner, _ := pterm.DefaultSpinner.Start("Creating the file at the given path.")
		time.Sleep(time.Second * 2)

		createFileSpinner.Success("File created!")

		fileFunctions.CreateFile(fileName, path)
	}
}

func showMainMenu() {
	var selectOptions []string = []string{
		"Create a file",
		"Delete a file",
		"Find a file",
		"Create a Directory",
		"Create Golang Package",
		"Exit",
	}

	menuOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(selectOptions).Show()
	var returnToMainMenu pterm.InteractiveSelectPrinter = pterm.DefaultInteractiveSelect
	returnToMainMenu.Options = selectOptions

	if menuOptions == "Create a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()
		_, _, createFileError := fileFunctions.CreateFile(fileName, path)
		var errors []string
		var spinner pterm.SpinnerPrinter

		if createFileError != nil {
			_, after, _ := strings.Cut(createFileError.Error(), ":")
			errors = append(errors, after)
		}

		if len(errors) > 0 && errors[0] == " no such file or directory" {
			confirmCreateDirectoryAndFile(fileName, path, spinner, returnToMainMenu)
		}

	} else if menuOptions == "Find a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()
		fileWasFound, fileName, path, errorArr := fileFunctions.FindFile(fileName, path)

		searchForFileSpinner, _ := pterm.DefaultSpinner.Start("Searching for the given file...")
		time.Sleep(time.Second * 2)

		if len(errorArr) > 0 && errorArr[0] == " no such file or directory" {
			searchForFileSpinner.Stop()
			confirmCreateDirectoryAndFile(fileName, path, *searchForFileSpinner, returnToMainMenu)
			return
		}

		if !fileWasFound {
			searchForFileSpinner.Stop()
			fileWasNotFoundFunc(fileName, path, *searchForFileSpinner)
			return
		}

		searchForFileSpinner.Success("The file was found!")

	} else if menuOptions == "Create a Directory" {
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		fileFunctions.CreateDirectory(path)
	} else if menuOptions == "Delete a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		fileFunctions.DeleteFile(fileName, path)
	} else if menuOptions == "Create Golang Package" {
		var projectDetails []string

		pterm.DefaultBasicText.Println("Ok. Lets get some information first...")

		goFileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		projectDirectory, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Where would you like the Project to be located?").Show()
		projectURL, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Please enter the package URL").Show()

		projectDetails = append(projectDetails, goFileName, projectDirectory, projectURL)

		fmt.Print(projectDetails)

	} else if menuOptions == "Exit" {
		exitSpinner, _ := pterm.DefaultSpinner.Start("One moment please ...")
		exitSpinner.Info("Goodbye!")
		os.Exit(1)
	}
}
