package cell

import (
	"./language/tags"
	"./language/modifiers"
	"./language/operators"
	"./language/statements"
)

type Cell struct {
	file string
	parsed string
	tags map[string]tags.Tag
	modifiers map[string]modifiers.Modifier
	operators map[string]operators.Operator
	statements map[string]statements.Statement
}

func (c Cell) cache() {
	//TODO: Create file_name.pcell with correct data header
}

func (c Cell) display() {
}

func (c Cell) parseTemplate() {
	//TODO: Call parser and parse template file.
}