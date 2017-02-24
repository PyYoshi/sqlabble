package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type CreateTable struct {
	Parent
	ifNotExists bool
	table       Table
}

func NewCreateTable(table Table) *CreateTable {
	return &CreateTable{
		ifNotExists: false,
		table:       table,
	}
}

func NewCreateTableIfNotExists(table Table) *CreateTable {
	return &CreateTable{
		ifNotExists: true,
		table:       table,
	}
}

func (c *CreateTable) Definitions(defs ...*Definition) *Definitions {
	ds := NewDefinitions(defs...)
	Contract(c, ds)
	return ds
}

func (c *CreateTable) separator() token.Tokens {
	return nil
}

func (c *CreateTable) nodeize() (tokenizer.Tokenizer, []interface{}) {
	tokens := token.Tokens{
		token.Word(keyword.CreateTable),
		token.Space,
	}
	if c.ifNotExists {
		tokens = tokens.Append(
			token.Word(keyword.IfNotExists),
			token.Space,
		)
	}

	t, values := c.table.nodeize()
	return tokenizer.NewContainer(t.Prepend(tokens...)), values
}
