package main

//colors enum
const (
	def = iota
	purple
	yellow
	gray
	fire
	blue
	raimbow
)

func (m *model) nextPalette() {
	m.color++
	switch m.color {
	case def:
		m.palette = []string{"#421E0F", "#19071A", "#09012F", "#040449", "#000764", "#0C2C8A", "#1852B1", "#397DD1", "#86B5E5", "#D3ECF8", "#F1E9BF", "#F8C95F", "#FFAA00", "#CC8000", "#995700", "#6A3403"}
	case purple:
		m.palette = []string{"#83a1f3", "#5c61c5", "#443d93", "#5c54ae", "#465c9c"}
	case yellow:
		m.palette = []string{"#f7f9bd", "#141409", "#947c3d", "#a4b164", "#54642c"}
	case gray:
		m.palette = []string{"#3D3D3D", "#E0E0E0", "#C8C8C8", "#888888", "#707070", "#505050", "#383838", "#282828", "#101010"}
	case fire:
		m.palette = []string{"#ffe808", "#B62203", "#D73502", "#FC6400", "#FF7500", "#FAC000"}
	case blue:
		m.palette = []string{"#328ca9", "#0e6ea9", "#2c4ea3", "#193882", "#102446"}
	case raimbow:
		m.palette = []string{"#ff0000", "#ff4103", "#ff6703", "#ffab03", "#ffc903", "#ffe703", "#eaff03", "#ccff03", "#adff03", "#71ff03", "#03ff95", "#03ffc2", "#03fff6", "#03d3ff", "#03b5ff", "#0380ff", "#0344ff", "#5803ff", "#8603ff", "#ba03ff", "#f603ff", "#ff03cb", "#ff0379"}
	default:
		m.color = -1
		m.nextPalette()
	}
}
