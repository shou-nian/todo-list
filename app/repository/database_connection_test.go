package repository

import "testing"

func TestNewDb(t *testing.T) {
	db, err := NewDb()
	if err != nil {
		t.Error(err.Error())
	}

	_, err = db.Exec("show databases;")
	if err != nil {
		t.Error(err.Error())
	}
}
