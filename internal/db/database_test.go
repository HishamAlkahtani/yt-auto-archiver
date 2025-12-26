package db

import "testing"

func Test_NewDb(t *testing.T) {
	db, err := NewDB("mydb.sqlite")

	if err != nil {
		t.Error(err)
	}

	t.Log(db.Stats())
}
