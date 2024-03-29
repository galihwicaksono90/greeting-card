package main

import "github.com/gdamore/tcell/v2"

type Logo struct {
	left   []string
	center []string
	right  []string
	text   []string
}

var logoLeft = []string{
	"  ......",
	"........",
	".....   ",
	".....   ",
}

var logoCenter = []string{
	"         ...:.    ",
	"       ::::::.    ",
	"      :::::::.    ",
	"     .::::.       ",
	"    .:::::        ",
	"    :::::         ",
	"  .:::::.         ",
	".:::::::          ",
	".:::::.           ",
	".:...             ",
}

var logoRight = []string{
	"   .....",
	"   .....",
	"........",
	"......  ",
}

var text = []string{
	"  _____        __ _                                          _ ",
	" / ____|      / _| |                                        (_)",
	"| (___   ___ | |_| |___      ____ _ _ __ ___  ___  ___ _ __  _ ",
	" \\___ \\ / _ \\|  _| __\\ \\ /\\ / / _` | '__/ _ \\/ __|/ _ \\ '_ \\| |",
	" ____) | (_) | | | |_ \\ V  V / (_| | | |  __/\\__ \\  __/ | | | |",
	"|_____/ \\___/|_|  \\__| \\_/\\_/ \\__,_|_|  \\___||___/\\___|_| |_|_|",
}

func (l *Logo) render(a *Animation, x int, y int) {
	// right, _ := a.screen.Size()
	// pos := right - (a.frame * 2)
	pos := x
	// if pos < 0-len(l.center) {
	// 	return
	// }
	a.drawObject(l.center, pos+4, y, tcell.StyleDefault.Foreground(tcell.Color35).Bold(true))
	a.drawObject(l.left, pos, y, tcell.StyleDefault.Foreground(tcell.Color28).Bold(true))
	a.drawObject(l.right, pos+14, y+6, tcell.StyleDefault.Foreground(tcell.Color28).Bold(true))
	a.drawObject(l.text, pos+24, y+2, tcell.StyleDefault.Bold(true))
}

func NewLogo() *Logo {
	return &Logo{
		left:   logoLeft,
		center: logoCenter,
		right:  logoRight,
		text:   text,
	}
}
