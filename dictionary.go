package namegen

type DictType byte

const (
	Adjectives DictType = iota
	Colors     DictType = iota
	Animals    DictType = iota
	numDicts   DictType = iota
)

var dicts [numDicts][]string

func init() {
	dicts[Adjectives] = adjectives
	dicts[Colors] = colors
	dicts[Animals] = animals
}

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

var colors = []string{
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
