package progress

import (
	"fmt"
	"strconv"
)

// Progress is progress bar
type Progress struct {
	totalProgress   uint
	currentProgress uint
	percent         string
	with            int
	curWith         int
	style           StyleChars
}

// NewProgress is instantiate the progress bar
func NewProgress(totalProgress uint, style int) *Progress {
	if style >= len(BarStyles) {
		style = 0
	}
	cs := BarStyles[style]
	return &Progress{
		totalProgress:   totalProgress,
		currentProgress: 0,
		percent:         "0.0%",
		style:           cs,
		with:            50,
		curWith:         0,
	}
}

// AddProgress is add progress
func (prog *Progress) AddProgress(progress uint) {
	if progress == 0 {
		return
	}
	prog.currentProgress += progress
	if prog.currentProgress > prog.totalProgress {
		prog.currentProgress = prog.totalProgress
	}
	prog.curWith = int(((float64)(prog.currentProgress) / float64(prog.totalProgress)) * float64(prog.with))
	prog.percent = strconv.FormatFloat(float64(prog.currentProgress)/float64(prog.totalProgress)*100, 'f', 1, 64) + "%"
	prog.Render()
}

// Render is apply colours to a drawing the picture
func (prog *Progress) Render() {
	text := ""
	for i := 1; i <= prog.with; i++ {
		if i < prog.curWith || prog.curWith == prog.with {
			text += string(prog.style.Completed)
		} else if i == prog.curWith {
			text += string(prog.style.Processing)
		} else {
			text += string(prog.style.Remaining)
		}
	}
	text += " " + prog.percent
	fmt.Print("\x0D\x1B[2K")
	if prog.curWith == prog.with {
		fmt.Println(text)
	} else {
		fmt.Print(text)
	}
}
