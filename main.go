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
		"Edit a file",
		"Delete a file",
		"Find a file",
		"Create a Directory",
	}

	result, _ := pterm.DefaultInteractiveSelect.WithOptions(selectOptions).Show()

	area.Update(
		result,
	)

	if result == "Create a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()
		fileNameText := pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprint(fileName)
		pathText := pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprint(path)

		area.Update(pathText, fileNameText)
		fileFunctions.CreateFile(fileName, path)

	} else if result == "Find a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()
		fileNameText := pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprint(fileName)
		pathText := pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprint(path)
		fileWasFound, fileName, path, errorArr := fileFunctions.FindFile(fileName, path)

		pterm.Info.Printf("\nSearching for: %v in the path: %v\n", fileNameText, pathText)

		if fileWasFound {
			area.Update(pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprintf("%v was found at %v", fileName, path))
		}

		if len(errorArr) > 0 && errorArr[0] == " no such file or directory" {
			confirmCreateDirectory, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("There was no directory found at the given path. Would you like to create the directory and file?").WithTextStyle(pterm.FgRed.ToStyle()).Show()

			createDirectorySpinner, _ := pterm.DefaultSpinner.Start("Creating the directory at the given path.")
			time.Sleep(time.Second * 2)
			createDirectorySpinner.Success("Directory created!")

			if getConfirmAnswer(confirmCreateDirectory) == "Yes" {
				fileFunctions.CreateDirectory(path)
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

		if !fileWasFound {
			confirmCreateFile, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("The file was not found at the given path! Would you like to create it?").WithTextStyle(pterm.FgRed.ToStyle()).Show()

			createFileSpinner, _ := pterm.DefaultSpinner.Start("Creating the file at the given path.")
			time.Sleep(time.Second * 2)
			createFileSpinner.Success("File created!")

			if getConfirmAnswer(confirmCreateFile) == "Yes" {
				fileFunctions.CreateFile(fileName, path)
			}
		}

	} else if result == "Create a Directory" {
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		fileFunctions.CreateDirectory(path)
	}

	area.Stop()
}

func getConfirmAnswer(answer bool) string {
	if answer {
		return "Yes"
	}
	return "No"
}
