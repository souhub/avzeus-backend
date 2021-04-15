package db

import "testing"

func TestParseSqlFile(t *testing.T) {
	given := "hoge"
	got := parseSqlFile(given)
	want := "./pkg/db/sql/hoge.sql"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
