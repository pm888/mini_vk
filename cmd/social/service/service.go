package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	sqlservice "mymod/cmd/social/service/servicesql"
	"mymod/users"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	Store map[int]*users.User
}
type Server struct {
	S *sql.DB
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		defer r.Body.Close()
		var u users.User
		if err := json.Unmarshal(content, &u); err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		id := sqlservice.AddSQl(&u, s.S)

		strID := strconv.Itoa(id)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ID:" + strID))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Server) Make_friends(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()

		var f users.FriendsStr
		if err := json.Unmarshal(content, &f); err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		friends1, friends2 := sqlservice.Make_friends_SQL(f.SourceFriends, f.TargetFriends, s.S)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(friends1 + " " + "and" + " " + friends2 + " " + "friends now"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Server) Delet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		defer r.Body.Close()

		var d users.DeleteFriends

		if err := json.Unmarshal(content, &d); err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		namedel := sqlservice.Delete_SQL(d.UserIDToDelete, s.S)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(namedel + " " + "Delet"))

	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Server) GetFriends(w http.ResponseWriter, r *http.Request) {
	name := ""
	friends := ""
	if r.Method == http.MethodGet {
		pathUrlString := r.URL.Path
		proverka := strings.LastIndex(pathUrlString, "/")
		if proverka != -1 {
			newStr := pathUrlString[strings.LastIndex(pathUrlString, "/")+1:]
			newStrInt, err := strconv.Atoi(newStr)
			if err != nil {
				fmt.Println(err)

			}
			friends, name = sqlservice.GetFriends_SQL(newStrInt, s.S)

			w.Write([]byte(name + " " + "Friends - " + " " + friends))
		}
	}

}

func (s *Server) ReplacementAge1(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		pathUrlString := r.URL.Path
		proverka := strings.LastIndex(pathUrlString, "/")
		if proverka != -1 {
			newStr := pathUrlString[strings.LastIndex(pathUrlString, "/")+1:]
			newStrInt, err := strconv.Atoi(newStr)
			if err != nil {
				fmt.Println(err)
			} else {
				content, err := ioutil.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))

					return
				}
				defer r.Body.Close()
				var newAge users.ReplacementAge
				if err := json.Unmarshal(content, &newAge); err != nil {

					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}

				name := sqlservice.ReplacementAgeSQL(newStrInt, newAge.NewAge, s.S)

				w.Write([]byte("Age" + " " + name + " " + "changed"))
			}

		}
	}
}

func (s *Server) RegisterHandlers() {
	http.HandleFunc("/create", s.Create)
	http.HandleFunc("/make_friends", s.Make_friends)
	http.HandleFunc("/delete", s.Delet)
	http.HandleFunc("/", s.ReplacementAge1)
	http.HandleFunc("/friends/", s.GetFriends)
}
