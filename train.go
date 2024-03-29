package main

import "github.com/gdamore/tcell/v2"

type Train struct {
	body []string
}

var train = []string{
	"   ____||_=======_____ ___________________________ .                           ",
	"  |                 | |                            ||           _____      O   ",
	"  |                 | |                            |  _____====  ]DD|_n_n__][. ",
	"  |_________________|_|____________________________|_[_________]_|__|________)>",
	"    `o!o         o!o' `o!o!o                  o!o!o'  ooo   ooo  'oo OOOOO oo\\_",
}

func (t *Train) renderTracks(frame, length int) []string {
	tracks := ""
	for i := 0; i < length; i++ {
		if (frame+i)%4 == 0 {
			tracks += "+"
		} else {
			tracks += "-"
		}
	}
	return []string{tracks}
}

func (t *Train) renderTrain() []string {
	return t.body
}

func (t *Train) renderSmoke(frame int) []string {
	smoke := ". . ."
	length := 13

	for i := length; i > 0; i-- {
		pos := frame % length
		if i%2 == 0 {
			smoke += " "
		} else if pos == i || pos == i-2 {
			smoke += "O"
		} else {
			smoke += "o"
		}
	}
	return []string{smoke}
}

func (t *Train) render(a *Animation, x, y int) {
	width, _ := a.screen.Size()

	train := t.renderTrain()
	tracks := t.renderTracks(a.frame, width)
	smoke := t.renderSmoke(a.frame)

	centerX := x - len(train[0])/2
	centerY := y + 10 - len(train)/2

	a.drawObject(train, centerX, centerY, tcell.StyleDefault)
	a.drawObject(tracks, 0, centerY+(len(train)), tcell.StyleDefault.Foreground(tcell.ColorYellow))
	a.drawObject(smoke, centerX+57, centerY, tcell.StyleDefault.Foreground(tcell.Color248))
}

func NewTrain() *Train {
	return &Train{
		body: train,
	}
}
