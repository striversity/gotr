package types

// Ref: https://golang.org/pkg/encoding/csv/

import (
	"encoding/xml"
	"strconv"
)

var header = []string{`id`, `fname`, `lname`, `username`, `password`, `email`}

type (
	User struct {
		Id        int    `json:"id" xml:"id,attr"`
		FirstName string `json:"firstname" xml:"name>first"`
		LastName  string `json:"lastname" xml:"name>last"`
		Username  string `json:"username,omitempty" xml:"secret>username"`
		Password  string `json:"password,omitempty" xml:"secret>password"`
		Email     string `json:"email,omitempty" xml:"email"`
	}

	UserDb struct {
		XMLName xml.Name `json:"-" xml:"users"`
		Type    string   `json:"type,omitempty" xml:"type"`
		Users   []User   `json:"users,omitempty" xml:"user"`
	}
)

func GetHeader() []string {
	return header
}

func (user User) EncodeAsStrings() (ss []string) {
	ss = make([]string, 6)
	ss[0] = strconv.Itoa(user.Id)
	ss[1] = user.FirstName
	ss[2] = user.LastName
	ss[3] = user.Username
	ss[4] = user.Password
	ss[5] = user.Email
	return
}

func (user *User) FromCSV(ss []string) {
	if nil == user {
		return
	}

	if nil == ss {
		return
	}

	user.Id, _ = strconv.Atoi(ss[0])
	user.FirstName = ss[1]
	user.LastName = ss[2]
	user.Username = ss[3]
	user.Password = ss[4]
	user.Email = ss[5]
}
