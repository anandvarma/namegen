package namegen

// DictType represents a unique type of word corpus.
type DictType byte

const (
	// Adjective word list.
	Adjectives DictType = iota
	// Color word list.
	Colors DictType = iota
	// Animal word list.
	Animals DictType = iota
	// Max value of DictType. Used internally for bounds checking.
	numDicts DictType = iota
)

var dicts [numDicts][]string

func init() {
	dicts[Adjectives] = adjectives
	dicts[Colors] = colors
	dicts[Animals] = animals
}

// Word corpus of interesting adjectives.
var adjectives = []string{
	"abbreviated",
	"aberrant",
	"abhorrent",
	"abiding",
	"abnormal",
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
	"antsy",
	"anxious",
	"arrogant",
	"articulate",
	"ashamed",
	"attractive",
	"auspicious",
	"awesome",
	"awful",
	"bald",
	"barren",
	"baseless",
	"bashful",
	"basic",
	"beautiful",
	"best",
	"big",
	"bitter",
	"bizarre",
	"blissful",
	"bored",
	"boring",
	"boujee",
	"brainy",
	"bright",
	"broad",
	"broken",
	"busy",
	"calm",
	"capable",
	"careful",
	"careless",
	"caring",
	"cautious",
	"charming",
	"chatty",
	"cheerful",
	"chubby",
	"clean",
	"clever",
	"clumsy",
	"cold",
	"colorful",
	"comfortable",
	"concerned",
	"confused",
	"curious",
	"curly",
	"curteous",
	"curvy",
	"cute",
	"daft",
	"dainty",
	"damaged",
	"damp",
	"dangerous",
	"dark",
	"deadly",
	"deep",
	"defective",
	"delicate",
	"determined",
	"different",
	"dirty",
	"docile",
	"dry",
	"dusty",
	"early",
	"educated",
	"efficient",
	"elderly",
	"elegant",
	"embarrassed",
	"empty",
	"encouraging",
	"enthusiastic",
	"excellent",
	"exciting",
	"expensive",
	"fabulous",
	"fair",
	"faithful",
	"famous",
	"fancy",
	"fantastic",
	"fast",
	"fearful",
	"fearless",
	"fertile",
	"flamboyant",
	"foolish",
	"forgetful",
	"friendly",
	"funny",
	"gentle",
	"glamorous",
	"glorious",
	"goofy",
	"gorgeous",
	"graceful",
	"grateful",
	"great",
	"greedy",
	"hairy",
	"handsome",
	"happy",
	"harsh",
	"hasty",
	"healthy",
	"heavy",
	"helpful",
	"hilarious",
	"historical",
	"hostile",
	"hot",
	"huge",
	"humorous",
	"hungry",
	"hyperactive",
	"hyperbolic",
	"hysterical",
	"ignorant",
	"illegal",
	"imaginary",
	"impolite",
	"important",
	"impossible",
	"innocent",
	"inspiring",
	"intelligent",
	"interesting",
	"jealous",
	"jolly",
	"juicy",
	"jumpy",
	"juvenile",
	"killer",
	"kind",
	"large",
	"lazy",
	"legal",
	"light",
	"literate",
	"little",
	"lively",
	"livid",
	"lonely",
	"loopy",
	"loud",
	"lovely",
	"lucky",
	"macho",
	"magical",
	"magnificent",
	"massive",
	"mature",
	"mean",
	"mega",
	"messy",
	"mighty",
	"miniature",
	"modern",
	"muscular",
	"narrow",
	"nasty",
	"naughty",
	"nervous",
	"new",
	"noisy",
	"nutritious",
	"obedient",
	"obese",
	"old",
	"overconfident",
	"peaceful",
	"pious",
	"polite",
	"poor",
	"powerful",
	"precious",
	"pretty",
	"proud",
	"quick",
	"quiet",
	"rapid",
	"rare",
	"remarkable",
	"responsible",
	"rich",
	"romantic",
	"royal",
	"rude",
	"scintillating",
	"secretive",
	"selfish",
	"serious",
	"sharp",
	"shiny",
	"shocking",
	"short",
	"shy",
	"silly",
	"sincere",
	"skinny",
	"slim",
	"slow",
	"small",
	"snooty",
	"soft",
	"space",
	"spicy",
	"spiritual",
	"splendid",
	"strong",
	"successful",
	"supercharged",
	"sweet",
	"tactful",
	"talented",
	"tall",
	"tangible",
	"tasteful",
	"tasty",
	"temperate",
	"tenacious",
	"tender",
	"tense",
	"terrific",
	"thankful",
	"thick",
	"thin",
	"thorough",
	"thoughtful",
	"tiny",
	"unique",
	"untidy",
	"upbeat",
	"upset",
	"victorious",
	"violent",
	"warm",
	"weak",
	"wealthy",
	"weary",
	"wide",
	"wise",
	"witty",
	"wonderful",
	"worried",
	"young",
	"youthful",
	"zealous",
	"zen",
}

// Word corpus of colors.
var colors = []string{
	"pale",
	"pink",
	"tan",
}

// Word corpus of animals.
var animals = []string{
	"albatross",
	"alligator",
	"alpaca",
	"anchovy",
	"ant",
	"anteater",
	"antelope",
	"ape",
	"armadillo",
	"baboon",
	"badger",
	"bandicoot",
	"barnacle",
	"barracuda",
	"bat",
	"beagle",
	"bear",
	"beaver",
	"bee",
	"beetle",
	"bichon",
	"bird",
	"bison",
	"bloodhound",
	"bluejay",
	"bobcat",
	"bonobo",
	"buffalo",
	"bulldog",
	"bullfrog",
	"bumblebee",
	"butterfly",
	"camel",
	"carp",
	"cat",
	"caterpillar",
	"catfish",
	"chameleon",
	"cheetah",
	"chicken",
	"chihuahua",
	"chimpanzee",
	"chipmunk",
	"chowchow",
	"cicadia",
	"clownfish",
	"coorgi",
	"coral",
	"cougar",
	"cow",
	"coyote",
	"crab",
	"crane",
	"croc",
	"crocodile",
	"crow",
	"dalmatian",
	"deer",
	"dodo",
	"doggo",
	"doggy",
	"dolphin",
	"donkey",
	"dragonfly",
	"duck",
	"eagle",
	"eel",
	"elephant",
	"elk",
	"emu",
	"falcon",
	"ferret",
	"firefly",
	"fish",
	"flamingo",
	"flea",
	"flounder",
	"fly",
	"flyingfish",
	"fox",
	"frog",
	"gazelle",
	"gecko",
	"gerbil",
	"gibbon",
	"giraffe",
	"glowworm",
	"goat",
	"goldfish",
	"goose",
	"gopher",
	"gorilla",
	"grasshopper",
	"greyhound",
	"grizzly",
	"groundhog",
	"guineapig",
	"guppy",
	"hammerhead",
	"hamster",
	"hare",
	"harpy",
	"hedgehog",
	"heron",
	"herring",
	"hippo",
	"honeybee",
	"horse",
	"human",
	"hummingbird",
	"husky",
	"hyena",
	"ibex",
	"ibis",
	"iguana",
	"impala",
	"jackal",
	"jaguar",
	"jellyfish",
	"kangaroo",
	"kingfisher",
	"kiwi",
	"koala",
	"koi",
	"komodo",
	"kookaburra",
	"krill",
	"lab",
	"labrador",
	"ladybug",
	"leech",
	"lemur",
	"leopard",
	"liger",
	"lion",
	"lizard",
	"llama",
	"lobster",
	"lynx",
	"macaw",
	"magpie",
	"mako",
	"mantaray",
	"mantis",
	"mastiff",
	"megalodon",
	"mink",
	"mole",
	"mongoose",
	"mongrel",
	"monkey",
	"moose",
	"moth",
	"mouse",
	"mule",
	"neanderthal",
	"octopus",
	"orca",
	"ostrich",
	"otter",
	"oyster",
	"panda",
	"pangolin",
	"panther",
	"parakeet",
	"parrot",
	"peacock",
	"pelican",
	"penguin",
	"pig",
	"pigeon",
	"pika",
	"pikachu",
	"piranha",
	"pitbull",
	"platypus",
	"polarbear",
	"poodle",
	"porcupine",
	"possum",
	"prawn",
	"pufferfish",
	"pug",
	"puma",
	"pup",
	"puppy",
	"quail",
	"rabbit",
	"raccoon",
	"rat",
	"rattlesnake",
	"rhino",
	"rooster",
	"salmon",
	"sardine",
	"sawfish",
	"scorpion",
	"seagull",
	"seahorse",
	"seal",
	"sealion",
	"shark",
	"sheep",
	"shibainu",
	"shihtzu",
	"skunk",
	"sloth",
	"slug",
	"snail",
	"snake",
	"sparrow",
	"sponge",
	"squid",
	"squirrel",
	"starfish",
	"stingray",
	"stork",
	"swan",
	"terrier",
	"tiger",
	"tortoise",
	"treefrog",
	"tuna",
	"turkey",
	"turtle",
	"unicorn",
	"vulture",
	"walrus",
	"wasp",
	"weasel",
	"whale",
	"wildboar",
	"wolf",
	"wolverine",
	"wombat",
	"woodpecker",
	"worm",
	"yak",
	"yeti",
	"zebra",
}
