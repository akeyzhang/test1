package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	 "strings"
)

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title: "TEST",
		MinSize: Size{600,400},
		Layout: VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo:&inTE},
					TextEdit{AssignTo:&outTE,ReadOnly:true},
				},
			},
			PushButton{
				Text:"確定",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()


}


