package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	//var progressBarSteps pterm.BasicTextPrinter
	var area pterm.AreaPrinter = pterm.DefaultArea
	var title pterm.BigTextPrinter = pterm.DefaultBigText

	var pbList = strings.Split("Initializing-App "+
		"Loading-Functions "+"Loading-Colors", " ")
	progressBar, pbErr := pterm.DefaultProgressbar.WithTotal(len(pbList)).WithTitle("Downloading").Start()

	for i := 0; i < progressBar.Total; i++ {
		if pbErr != nil {
			pterm.Error.Println("An error has occured")
		}
		if i == 6 {
			time.Sleep(time.Second * 3)
		}
		progressBar.UpdateTitle("Downloading " + pbList[i])
		pterm.Success.Println("Downloading " + pbList[i])
		progressBar.Increment()
		progressBar.WithRemoveWhenDone(true)
		time.Sleep(time.Millisecond * 350)
	}

	area.Center = true

	title.WithLetters(
		putils.LettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightGreen)),
		putils.LettersFromStringWithStyle("_", pterm.NewStyle(pterm.FgLightYellow)),
		putils.LettersFromStringWithStyle("TERMINAL", pterm.NewStyle(pterm.FgLightGreen))).
		Render()

	var selectOptions []string = []string{
		"Create a file",
		"Edit a file",
	}

	result, _ := pterm.DefaultInteractiveSelect.WithOptions(selectOptions).Show()

	area.Update(
		result,
	)

	area.Stop()
}
