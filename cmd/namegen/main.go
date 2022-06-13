package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/anandvarma/namegen"
)

func main() {
	// Command line flags.
	lPtr := flag.Int("postfix", 4, "The length of the postfix id in number of characters.")
	aPtr := flag.Bool("alphanum", false, "If set, the postfix is represented as an alpha-numeric id instead of a decimal number.")
	flag.Parse()

	// Parse command line arguments.
	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("Unexpected command line arguments. Run 'namegen --help' to see usage.")
		os.Exit(1)
	}
	numIters := 1
	if len(args) == 1 {
		if num, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("Received bad value of <num-iterations>. Run 'namegen --help' to see usage.")
			os.Exit(1)
		} else {
			numIters = num
		}
	}

	// Initialize a nameGen instance with the received options.
	nameSchema := []namegen.DictType{
		namegen.Adjectives,
		namegen.Animals,
	}
	var pIdType namegen.PostfixIdType = namegen.Numeric
	if *aPtr {
		pIdType = namegen.AlphaNumeric
	}
	ngen := namegen.NewWithPostfixId(nameSchema, pIdType, *lPtr /* pIdLen */)

	for i := 0; i < numIters; i++ {
		fmt.Println(ngen.Get())
	}
}
