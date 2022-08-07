package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	x, y := 8, 8
	mf := NewMinefield(x, y, 10)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	render(mf)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<MouseLeft>":
			payload := e.Payload.(ui.Mouse)
			mcy, mcx := (payload.X+1)/2, payload.Y

			if mcy > y || mcy > x {
				continue
			}

			boom := mf.Click(mcx, mcy)
			if boom {
				ui.Close()
				println("boom")
				os.Exit(1)
			}

			render(mf)
		case "<MouseRight>":
			fallthrough
		case "<MouseMiddle>":
			payload := e.Payload.(ui.Mouse)
			mcy, mcx := (payload.X+1)/2, payload.Y

			if mcy > y || mcy > x {
				continue
			}

			mf.ToggleMark(mcx, mcy)
			render(mf)
		}
	}
}

func render(mf minefield) {
	paragraph := widgets.NewParagraph()
	paragraph.Text = mf.String()
	paragraph.SetRect(0, 0, mf.height*2+1, mf.width+2)

	ui.Render(paragraph)
}
