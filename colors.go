package main

const (
	colorBlack         = "black"
	colorGray          = "gray"
	colorRed           = "red"
	colorBrightRed     = "brightred"
	colorGreen         = "green"
	colorBrightGreen   = "brightgreen"
	colorBrown         = "brown"
	colorYellow        = "yellow"
	colorBlue          = "blue"
	colorBrightBlue    = "brightblue"
	colorMagenta       = "magenta"
	colorBrightMagenta = "brightmagenta"
	colorCyan          = "cyan"
	colorBrightCyan    = "brightcyan"
	colorLightGray     = "lightgray"
	colorWhite         = "white"
)

type colorsType map[string]bool

func (c colorsType) exists(names ...string) string {
	for _, name := range names {
		if _, ok := c[name]; !ok {
			return name
		}
	}
	return ""
}

var colors = colorsType{
	colorBlack:         true,
	colorGray:          true,
	colorRed:           true,
	colorBrightRed:     true,
	colorGreen:         true,
	colorBrightGreen:   true,
	colorBrown:         true,
	colorYellow:        true,
	colorBlue:          true,
	colorBrightBlue:    true,
	colorMagenta:       true,
	colorBrightMagenta: true,
	colorCyan:          true,
	colorBrightCyan:    true,
	colorLightGray:     true,
	colorWhite:         true,
}
