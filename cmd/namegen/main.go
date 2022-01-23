package main

import (
	"fmt"

	"github.com/anandvarma/namegen"
)

func main() {
	name_schema := []namegen.DictType{
		namegen.Adjectives,
		namegen.Animals,
	}
	ngen := namegen.NewWithPostfixId(name_schema, namegen.Numeric, 4)
	fmt.Println(ngen.Get())
}
