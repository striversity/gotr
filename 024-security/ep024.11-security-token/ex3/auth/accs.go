package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	// these are our account types
	ACCOUNT_NONE     = iota
	ACCOUNT_SAVING   = iota
	ACCOUNT_CHECKING = iota
	ACCOUNT_MORTGAGE = iota
)

type (
	Account struct {
		OwnerID int
		ID      string
		Type    int // one of account type ACCOUNT_X
		Balance float64
	}

	AccountsDB struct {
		accounts map[ing][]*Account
	}
)

func (adb *AccountsDB) init() {
	adb.accounts = make(map[int][]*Account)
	adb.addAccount(1, ACCOUNT_SAVING, 3.14)
	adb.addAccount(1, ACCOUNT_SAVING, 11.04)
	adb.addAccount(2, ACCOUNT_SAVING, 47.9)
	adb.addAccount(3, ACCOUNT_MORTGAGE, 1972)
}

func (adb *AccountsDB) addAccount(ownerID int, accountType int, startingBalance float64) {
	accID := genAccountID(len(adb.accounts))
	u := &Account{ID: accID, OwnerID: ownerID, Type: accountType, Balance: startingBalance}

	adb.accounts[ownerID] = append(adb.accounts[ownerID], u)
}

func (adb*AccountsDB) getAccounts(ownerID int)[]*Account{
	return adb.accounts[ownerID]
}

func (adb *AccountsDB) getAccount(ownerID int, accountType int) *Account {
	accs := adb.accounts[ownerID]

	for _, a := range accs {
		if a.Type == accountType {
			return a
		}
	}

	return nil
}

func genAccountID(id int) string {
	str := fmt.Sprintf("%v", id)
	m := md5.Sum([]byte(str))
	h := hex.EncodeToString(m[:])
	return string(h)
}
