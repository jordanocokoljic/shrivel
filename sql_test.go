package shrivel_test

import (
	_ "embed"
	"github.com/jordanocokoljic/shrivel"
	"testing"
)

//go:embed testdata/complex.raw.sql
var complexSqlSample string

//go:embed testdata/complex.min.sql
var minifiedComplexSqlSample string

func TestSql_CanOperateWithTwoSlices(t *testing.T) {
	src := []rune(complexSqlSample)
	dst := make([]rune, len(src))

	n := shrivel.Sql(dst, src)

	if string(dst[:n]) != minifiedComplexSqlSample {
		t.Errorf("minification was mangled")
	}
}

func TestSql_CanOperateInPlace(t *testing.T) {
	sql := []rune(complexSqlSample)
	sql = sql[:shrivel.Sql(sql, sql)]

	if string(sql) != minifiedComplexSqlSample {
		t.Errorf("minification was mangled")
	}
}
