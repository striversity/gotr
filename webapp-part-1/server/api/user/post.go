package user

import (
	"encoding/json"
	"net/http"
)

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jd := json.NewDecoder(r.Body)

	aUser := &User{}
	err := jd.Decode(aUser)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start of protected code changes
	lock.Lock()
	nextUserID++
	aUser.ID = nextUserID
	db = append(db, aUser)
	// end protected code changes
	lock.Unlock()

	respUser := User{ID: aUser.ID, Username: aUser.Username}
	je := json.NewEncoder(w)
	je.Encode(respUser)
}
