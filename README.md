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
- Blazing fast!! `(10,000,000+ Op/s)`

## API & Usage

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

---

### Benchmarks
	goos: linux
	goarch: amd64
	pkg: github.com/anandvarma/namegen
	cpu: AMD Ryzen 5 3600 6-Core Processor              
	BenchmarkNew-12                 1000000000               0.2561 ns/op          0 B/op          0 allocs/op
	BenchmarkGet-12                 14617033                82.80 ns/op           18 B/op          1 allocs/op
	BenchmarkGetForId-12            21623218                51.25 ns/op           16 B/op          1 allocs/op
	BenchmarkGetWithPostfix-12      10882941               111.7 ns/op            26 B/op          1 allocs/op
	BenchmarkGetPostfixId-12        31611556                37.07 ns/op           20 B/op          0 allocs/op
	PASS
	ok      github.com/anandvarma/namegen   5.293s
