package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Animation struct {
	screen tcell.Screen
	mu     sync.Mutex
	frame  int
	train  *Train
	logo   *Logo
	trees  *Trees
}

func (a *Animation) DrawText(x1, y1, x2, y2 int, text string, style tcell.Style) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		a.screen.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (a *Animation) centerPosition() (x, y int) {
	x, y = a.screen.Size()
	return x / 2, y / 2
}

type draw func()

func (a *Animation) updateScreen(d draw) {
	a.screen.Clear()
	d()
	a.screen.Show()
}

func NewAnimation() *Animation {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Dim(true)
	screen.SetStyle(defStyle)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	a := &Animation{
		screen: screen,
		frame:  1,
		train:  NewTrain(),
		logo:   NewLogo(),
		trees:  NewTrees(),
	}

	return a
}

func (g *Animation) exit() {
	g.screen.Fini()
	os.Exit(9)
}

func (a *Animation) resizeScreen() {
	a.mu.Lock()
	a.screen.Sync()
	a.mu.Unlock()
}

func (a *Animation) handleKeyBoardEvents() {
	for {
		switch event := a.screen.PollEvent().(type) {
		case *tcell.EventResize:
			a.resizeScreen()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				a.exit()
			}
		}
	}
}

func (a *Animation) drawObject(s []string, x, y int, style tcell.Style) {
	for i, v := range s {
		sLen := len(v)
		a.DrawText(x, y+i, x+sLen, y+i, v, style)
	}
}

func (a *Animation) run() {
	a.updateScreen(func() {})
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	row := 0
	for {
		cx, cy := a.centerPosition()
		select {
		case <-ticker.C:
			a.updateScreen(func() {
				a.trees.render(a, 10, cy-15)
				a.train.render(a, cx, cy)
				a.logo.render(a, 5, 5)
			})
			a.frame++
			row++
			if row > 10 {
				row = 0
			}
		}
	}
}

func StartAnimation() {
	a := NewAnimation()
	go a.run()
	a.handleKeyBoardEvents()
}
