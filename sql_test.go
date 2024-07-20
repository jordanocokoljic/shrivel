package shrivel_test

import (
	"errors"
	"github.com/jordanocokoljic/shrivel/v2"
	"testing"
)

func TestSql(t *testing.T) {
	tests := []struct {
		name string
		sql  string
		out  string
		err  error
	}{
		{
			name: "CreateTable",
			sql:  "CREATE TABLE tbl\n(\n    pk INTEGER PRIMARY KEY\n);",
			out:  "CREATE TABLE tbl(pk INT PRIMARY KEY)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := []rune(tt.sql)
			n, err := shrivel.Sql(raw, raw)
			if !errors.Is(err, tt.err) {
				t.Fatalf("returned %s, should have been %s", err, tt.err)
			}

			if string(raw[:n]) != tt.out {
				t.Errorf(
					"returned\n%s\nshould have been\n%s",
					string(raw[:n]), tt.out)
			}
		})
	}
}
