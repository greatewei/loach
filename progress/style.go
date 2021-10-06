package progress

// StyleChars setting for a progress bar. default {'#', '>', ' '}
type StyleChars struct {
	Completed, Processing, Remaining rune
}

// BarStyles some built in StyleChars style collection
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
