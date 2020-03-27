package main

type (

	// TableField describe a RDBMS table field's value
	TableField struct {
		Name  string // ex: Lname, Fname, Age, etc.
		Type  string // ex: int, string, float64, etc
		Value string // all values stored as string
	}
)
