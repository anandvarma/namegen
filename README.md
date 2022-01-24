# namegen <div> [![GitHub license](https://badgen.net/github/license/anandvarma/namegen)](https://github.com/anandvarma/namegen/blob/main/LICENSE) [![GoReportCard](https://goreportcard.com/badge/github.com/anandvarma/namegen)](https://goreportcard.com/report/github.com/anandvarma/namegen) [![Go Reference](https://pkg.go.dev/badge/github.com/anandvarma/namegen.svg)](https://pkg.go.dev/github.com/anandvarma/namegen) </div>

![preview](https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GOPHER_AVATARS.jpg)

## What's this?
Go package and associated command line utility to generate random yet human-readable names and identifiers. 
Somewhat inspired by the random names that Docker gives its containers and GCP uses as project-ids. 

Use it to generate email-ids, gamer-tags, component-ids, passwords or startup names.

For instance, it came up with gems like: `boujee-coorgi`, `young-flamingo`, `articulate-otter-92`, `ancient-mako-7ab1g3` and `mighty-pale-pikachu-7133`.


### Features
- 1 Million+ unique possibilities with just names and nearly infinite named ids
- Zero Allocations! No additional memory allocations other than the `string` returned
- Blazing fast!! `(9,000,000+ Op/s)`

<hr/>


### API & Usage

> Generate names using default config.

	ngen := namegen.New()
	fmt.Println(ngen.Get()) // twisted-flamingo

> Configure name generator to generate named ids of the form "adjective-color-animal-xxxx".

	name_schema := []namegen.DictType{
		namegen.Adjectives,
		namegen.Colors,
		namegen.Animals,
	}
	ngen := namegen.NewWithPostfixId(name_schema, namegen.Numeric, 4)
	fmt.Println(ngen.Get()) // energetic-pink-pug-3729
	
[Go Reference](https://pkg.go.dev/github.com/anandvarma/namegen)

### Benchmarks

