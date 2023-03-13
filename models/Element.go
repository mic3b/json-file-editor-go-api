package models

type Element struct {
	// Dont mind about this "Id", our primary key is "Key"!!!
	Id    int
	Name  string
	Value interface{}
	Key   string
}
