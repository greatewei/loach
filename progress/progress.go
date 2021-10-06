package progress

import (
	"fmt"
	"strconv"
)

type Progress struct {
	totalProgress   uint
	currentProgress uint
	percent         string
	with            int
	curWith         int
	style           StyleChars
}

func NewProgress(totalProgress uint, style int) *Progress {
	stype := BarStyles[style]
	return &Progress{
		totalProgress:   totalProgress,
		currentProgress: 0,
		percent:         "0.0%",
		style:           stype,
		with:            30,
		curWith:         0,
	}
}

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

func (prog *Progress) Render() {
	//fmt.Print(string(prog.style.Completed))
	//fmt.Print(string(prog.style.Processing))
	//fmt.Print(string(prog.style.Remaining))
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
	fmt.Print(text)
}
