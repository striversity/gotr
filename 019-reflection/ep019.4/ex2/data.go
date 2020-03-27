package main

import "fmt"

// people is sample table with the following structures
// ID(int) | Fname(string) | Lname(string) | Age(uint8) | Height(float64)
// --------+---------------+---------------+------------+-----------------
// 1       | Jane          | Doe           | 41         | 5.5
// 2       | Mark          | Smith         | 23         | 5.9
// 3       | Anne          | Marie         | 89         | 5.2
var (
	row1 = []TableField{
		{Name: "ID", Type: "int", Value: "1"},
		{Name: "Fname", Type: "string", Value: "Jane"},
		{Name: "Lname", Type: "string", Value: "Doe"},
		{Name: "Age", Type: "uint8", Value: "41"},
		{Name: "Height", Type: "float64", Value: "5.5"},
	}
	row2 = []TableField{
		{Name: "ID", Type: "int", Value: "2"},
		{Name: "Fname", Type: "string", Value: "Mark"},
		{Name: "Lname", Type: "string", Value: "Smith"},
		{Name: "Age", Type: "uint8", Value: "23"},
		{Name: "Height", Type: "float64", Value: "5.9"},
	}
	row3 = []TableField{
		{Name: "ID", Type: "int", Value: "3"},
		{Name: "Fname", Type: "string", Value: "Anne"},
		{Name: "Lname", Type: "string", Value: "Marie"},
		{Name: "Age", Type: "uint8", Value: "89"},
		{Name: "Height", Type: "float64", Value: "5.2"},
	}
	// PeopleData is our collection of data populated from RDBMS
	PeopleData = [][]TableField{row1, row2, row3}
)
