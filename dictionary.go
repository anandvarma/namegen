package namegen

type Dictionary byte

const (
	Adjectives Dictionary = iota
	Animals    Dictionary = iota
	Colours    Dictionary = iota
)

var adjectives = []string{
	"abundant",
	"adorable",
	"adventurous",
	"aggressive",
	"alien",
	"aloof",
	"ambitious",
	"ancient",
	"angry",
	"animated",
	"annoying",
	"anxious",
	"arrogant",
	"ashamed",
	"attractive",
	"auspicious",
	"awesome",
	"awful",
	"abbreviated",
	"aberrant",
	"abhorrent",
	"abiding",
}

var colours = []string{
	"pink",
	"pale",
	"tan",
}

var animals = []string{
	"hammerhead",
	"coorgi",
	"mako",
	"dogfish",
	"frilled",
}
