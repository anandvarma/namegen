package main

import (
	"fmt"

	"github.com/anandvarma/namegen"
)

func main() {
	name_schema := []namegen.DictType{
		namegen.Adjectives,
		namegen.Nouns,
	}
	//ngen := namegen.New(name_schema)
	ngen := namegen.NewWithPostfixId(name_schema, namegen.Numeric, 6)
	ngen.SetPostfixIdLen(2)
	ngen.SetDelimiter("_")
	fmt.Println(ngen.Get())
}
