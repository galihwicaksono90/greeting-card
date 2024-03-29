package main

import "github.com/gdamore/tcell/v2"

var treeLarge = []string{
	".{{}}}}}}.",
	"{{{{{{(`)}}}.",
	"{{{(`)}}}}}}}}}",
	"}}}}}}}}}{{(`){{{",
	"}}}}{{{{(`)}}{{{{",
	"{{{(`)}}}}}}}{}}}}}",
	"{{{{{{{{(`)}}}}}}}}}}",
	"{{{{{{{}{{{{(`)}}}}}}",
	" {{{{{(`)   {{{{(`)}'",
	"  `\"\"'\" |   | \"'\"'`",
	"  (`)  /     \\       ",
}

var shamrock = []string{
	"  _  ",
	" ( ) ",
	"(_X_)",
	"  j  ",
}

var treeSmall = []string{
	"	 	 ####   ",
	"   #o###   ",
	" #####o### ",
	"#o#\\#|#/###",
	" ###\\|/#o# ",
	"  # }|{  # ",
	"    }|{    ",
}

type Trees struct {
	treeLarge []string
	treeSmall []string
	shamrock  []string
}

func (t *Trees) render(a *Animation, x, y int) {
	right, _ := a.screen.Size()
	pos := right - a.frame
	pos2 := right - (a.frame / 2)
	if pos < -150 {
		return
	}
	a.drawObject(t.treeSmall, pos2-x, y-2, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	if a.frame > 50 {
		a.drawObject(t.treeLarge, pos-x+60, y, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	}
}

func NewTrees() *Trees {
	return &Trees{
		treeLarge: treeLarge,
		shamrock:  shamrock,
		treeSmall: treeSmall,
	}
}
