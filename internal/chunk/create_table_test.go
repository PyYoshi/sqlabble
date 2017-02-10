package chunk_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/minodisk/sqlabble"
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/diff"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func TestCreateTableType(t *testing.T) {
	// if _, ok := interface{}(chunk.From{}).(grammar.Clause); !ok {
	// 	t.Errorf("chunk.FromClause doesn't implement grammar.Clause")
	// }
}

func TestCreateTableSQL(t *testing.T) {
	for i, c := range []struct {
		statement grammar.Statement
		sql       string
		sqlIndent string
		values    []interface{}
	}{
		{
			chunk.NewCreateTable(
				chunk.NewTable("foo"),
			),
			"CREATE TABLE foo",
			`> CREATE TABLE
>   foo
`,
			[]interface{}{},
		},
		{
			chunk.NewCreateTable(
				chunk.NewTable("foo"),
			).Definitions(),
			"CREATE TABLE foo ()",
			`> CREATE TABLE
>   foo
>   (
>   )
`,
			[]interface{}{},
		},
		{
			chunk.NewCreateTable(
				chunk.NewTable("foo"),
			).Definitions(
				chunk.NewColumn("name").Define("VARCHAR(255)"),
			),
			"CREATE TABLE foo (name VARCHAR(255))",
			`> CREATE TABLE
>   foo
>   (
>     name VARCHAR(255)
>   )
`,
			[]interface{}{},
		},
		{
			chunk.NewCreateTable(
				chunk.NewTable("foo"),
			).Definitions(
				chunk.NewColumn("name").Define("VARCHAR(255)"),
				chunk.NewColumn("gender").Define("ENUM('M', 'F')"),
			),
			"CREATE TABLE foo (name VARCHAR(255), gender ENUM('M', 'F'))",
			`> CREATE TABLE
>   foo
>   (
>     name VARCHAR(255)
>     , gender ENUM('M', 'F')
>   )
`,
			[]interface{}{},
		},
		{
			chunk.NewCreateTable(
				chunk.NewTable("foo"),
			).Definitions(
				chunk.NewColumn("name").Define("VARCHAR(255)"),
				chunk.NewColumn("gender").Define("ENUM('M', 'F')"),
				chunk.NewColumn("birth_date").Define("DATE"),
			),
			"CREATE TABLE foo (name VARCHAR(255), gender ENUM('M', 'F'), birth_date DATE)",
			`> CREATE TABLE
>   foo
>   (
>     name VARCHAR(255)
>     , gender ENUM('M', 'F')
>     , birth_date DATE
>   )
`,
			[]interface{}{},
		},
	} {
		t.Run(fmt.Sprintf("%d Build", i), func(t *testing.T) {
			sql, values := sqlabble.Build(c.statement)
			if sql != c.sql {
				t.Error(diff.SQL(sql, c.sql))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
		t.Run(fmt.Sprintf("%d BuildIndent", i), func(t *testing.T) {
			sql, values := sqlabble.BuildIndent(c.statement, "> ", "  ")
			if sql != c.sqlIndent {
				t.Error(diff.SQL(sql, c.sqlIndent))
			}
			if !reflect.DeepEqual(values, c.values) {
				t.Error(diff.Values(values, c.values))
			}
		})
	}
}
