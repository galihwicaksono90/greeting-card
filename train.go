package main

var train = []string{
	"   ____||_=======_____ ______________________________   . . . o o o o o o o    ",
	"  |                 | |                            ||           _____      o   ",
	"  |                 | |                            |  _____====  ]DD|_n_n__][. ",
	"  |_________________|_|____________________________|_[_________]_|__|________)>",
	"    `o!o         o!o' `o!o!o                  o!o!o'  ooo   ooo  'oo OOOOO oo\\_",
	"                                                                               ",
}

func renderTracks(tick int, length int) string {
	tracks := ""
	for i := 0; i < length; i++ {
		if (tick+i)%4 == 0 {
			tracks += "+"
		} else {
			tracks += "-"
		}
	}
	return tracks
}

func RenderTrain(tick int) []string {
	train[len(train)-1] = renderTracks(tick, len(train[0]))
	return train
}
