package main

import (
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
		confirmCreateDirectory, confirmCreateDirectoryErr := pterm.DefaultInteractiveConfirm.WithConfirmText("Yes").WithDefaultText("Would you like to create the directory and file at the given path?").WithDefaultValue(true).Show()

		area.Update(pathText, fileNameText)
		fileWasFound, fileName, path, errorArr := fileFunctions.FindFile(fileName, path)

		if fileWasFound {
			area.Update(pterm.DefaultBasicText.WithStyle(pterm.FgLightGreen.ToStyle()).Sprintf("%v was found at %v", fileName, path))
		} else if errorArr[0] == "no such file or directory" {
			area.Update(confirmCreateDirectory)

			if confirmCreateDirectory {
				fileFunctions.CreateDirectory(path)
				fileFunctions.CreateFile(fileName, path)
			} else if !confirmCreateDirectory {
				area.Update(pterm.DefaultBasicText.WithStyle(pterm.FgRed.ToStyle()).Print(confirmCreateDirectoryErr))
			}
		}

	} else if result == "Create a Directory" {
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		fileFunctions.CreateDirectory(path)
	}

	area.Stop()
}
