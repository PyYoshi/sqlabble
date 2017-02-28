package statement

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Table struct {
	name string
}

func NewTable(name string) Table {
	return Table{
		name: name,
	}
}

func (t Table) As(alias string) TableAs {
	return TableAs{
		table: t,
		alias: alias,
	}
}

func (t Table) Join(table TableOrAlias) Join {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t Table) InnerJoin(table TableOrAlias) Join {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t Table) LeftJoin(table TableOrAlias) Join {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t Table) RightJoin(table TableOrAlias) Join {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}

func (t Table) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(t)
}

func (t Table) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	if t.name == "" {
		return nil, nil
	}
	return tokenizer.NewLine(token.Word(t.name)), nil
}

func (t Table) previous() Prever {
	return nil
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (t Table) isTableOrAlias() bool {
	return true
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (t Table) isTableOrAliasOrJoiner() bool {
	return true
}
