package main

import (
	"os"
	"time"

	"github.com/pschlafley/fileFunctions"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	//var progressBarSteps pterm.BasicTextPrinter
	var area pterm.AreaPrinter = *pterm.DefaultArea.WithRemoveWhenDone(true)
	var title pterm.BigTextPrinter = pterm.DefaultBigText

	area.Center = true

	title.WithLetters(
		putils.LettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightGreen)),
		putils.LettersFromStringWithStyle("_", pterm.NewStyle(pterm.FgLightYellow)),
		putils.LettersFromStringWithStyle("TERMINAL", pterm.NewStyle(pterm.FgLightGreen))).
		Render()

	var selectOptions []string = []string{
		"Create a file",
		"Delete a file",
		"Find a file",
		"Create a Directory",
	}

	menuOptions, _ := pterm.DefaultInteractiveSelect.WithOptions(selectOptions).Show()

	area.Update(
		menuOptions,
	)

	if menuOptions == "Create a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()
		fileNameText := pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprint(fileName)
		pathText := pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprint(path)

		area.Update(pathText, fileNameText)
		fileFunctions.CreateFile(fileName, path)

	} else if menuOptions == "Find a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()
		fileWasFound, fileName, path, errorArr := fileFunctions.FindFile(fileName, path)

		searchForFileSpinner, _ := pterm.DefaultSpinner.Start("Searching for the given file...")
		time.Sleep(time.Second * 2)

		if len(errorArr) > 0 && errorArr[0] == " no such file or directory" {
			searchForFileSpinner.Stop()
			confirmCreateDirectoryAndFile(fileName, path, *searchForFileSpinner)
		}

		if !fileWasFound {
			searchForFileSpinner.Stop()
			fileWasNotFoundFunc(fileName, path, *searchForFileSpinner)
		}

		searchForFileSpinner.Success("The file was found!")

	} else if menuOptions == "Create a Directory" {
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		fileFunctions.CreateDirectory(path)
	} else if menuOptions == "Delete a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		fileFunctions.DeleteFile(fileName, path)
	}

	area.Stop()
}

func getConfirmAnswer(answer bool) string {
	if answer {
		return "Yes"
	}
	return "No"
}

func confirmCreateDirectoryAndFile(fileName, path string, searchForFileSpinner pterm.SpinnerPrinter) {
	searchForFileSpinner.Fail("No such file or directory found.\n")
	confirmCreateDirectory, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("There was no directory found at the given path. Would you like to create the directory and file?").WithTextStyle(pterm.FgYellow.ToStyle()).Show()

	if getConfirmAnswer(confirmCreateDirectory) == "Yes" {
		createDirectorySpinner, _ := pterm.DefaultSpinner.Start("Creating the directory at the given path.")
		time.Sleep(time.Second * 2)
		fileFunctions.CreateDirectory(path)
		createDirectorySpinner.Success("Directory created!")
	} else {
		os.Exit(1)
	}

	confirmCreateFile, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Would you also like to create the file in the directory?").WithTextStyle(pterm.FgCyan.ToStyle()).Show()

	if getConfirmAnswer(confirmCreateFile) == "Yes" {
		createFileSpinner, _ := pterm.DefaultSpinner.Start("Creating the file at the given path.")
		time.Sleep(time.Second * 2)

		createFileSpinner.Success("File created!")

		fileFunctions.CreateFile(fileName, path)
		os.Exit(1)
	} else {
		os.Exit(1)
	}

}

func fileWasNotFoundFunc(fileName, path string, searchForFileSpinner pterm.SpinnerPrinter) {
	searchForFileSpinner.Fail("File Not Found.\n")
	createFileIfNotFound, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("The file was not found at the given path! Would you like to create it?").WithTextStyle(pterm.FgYellow.ToStyle()).Show()

	if getConfirmAnswer(createFileIfNotFound) == "Yes" {
		createFileSpinnerIfFileNotFound, _ := pterm.DefaultSpinner.Start("Creating the file at the given path.")
		time.Sleep(time.Second * 2)
		fileFunctions.CreateFile(fileName, path)
		createFileSpinnerIfFileNotFound.Success("File created!")
	} else {
		os.Exit(1)
	}
}
