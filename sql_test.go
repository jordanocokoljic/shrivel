package shrivel_test

import (
	"github.com/jordanocokoljic/shrivel"
	"testing"
)

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

const complexSqlSample = `
-- Lasciate ogne speranza, voi ch'intrate

CREATE TABLE "THIS
is          supported
by
minify"
(
    id             BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "يونيكود صحيح" VARCHAR NOT NULL,
    "🤣😍  "         VARCHAR NOT NULL
);

insert into "THIS
is          supported
by
minify" ("يونيكود صحيح", "🤣😍  ")
VALUES ('涵盖
多种
语言  ', 'ඞ ℻');

/*
With support for block comments.

-- And line comments within it.

/*
and nested blocks too!
 */

Throw in some unterminated strings for good measure: "

ועוד שפה או שתיים?
конечно!
 */

select id, "يونيكود صحيح", "🤣😍  "
from "THIS
is          supported
by
minify";

DROP TABLE "THIS
is          supported
by
minify";


`

const minifiedComplexSqlSample = `CREATE TABLE"THIS
is          supported
by
minify"(id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,"يونيكود صحيح" VARCHAR NOT NULL,"🤣😍  " VARCHAR NOT NULL);insert into"THIS
is          supported
by
minify"("يونيكود صحيح","🤣😍  ")VALUES('涵盖
多种
语言  ','ඞ ℻');select id,"يونيكود صحيح","🤣😍  "from"THIS
is          supported
by
minify";DROP TABLE"THIS
is          supported
by
minify";`
