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
	}

	result, _ := pterm.DefaultInteractiveSelect.WithOptions(selectOptions).Show()

	area.Update(
		result,
	)

	if result == "Create a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		area.Update(fileName, path)
		fileFunctions.CreateFile(fileName, path)
	} else if result == "Find a file" {
		fileName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("File Name").Show()
		path, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Path").Show()

		area.Update(fileName, path)
		fileFunctions.FindFile(fileName, path)
	}

	area.Stop()
}
