package progress

// StyleChars is setting for a progress bar. default {'#', '>', ' '}
type StyleChars struct {
	Completed, Processing, Remaining rune
}

// BarStyles is some built in StyleChars style collection
var BarStyles = []StyleChars{
	{'=', '>', ' '},
	{'=', '>', '-'},
	{'#', '>', ' '},
	{'#', '>', '-'},
	{'*', '>', '-'},
	{'▉', '▉', '░'},
	{'■', '■', ' '},
	{'■', '■', '▢'},
	{'■', '▶', ' '},
}
