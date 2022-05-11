package sqlservice

import (
	"database/sql"
	"fmt"
	"mymod/users"

	_ "github.com/go-sql-driver/mysql"
)

func AddSQl(u *users.User, db *sql.DB) (idReturn int) {
	result, err := db.Query("INSERT INTO users VALUES (?,?,?)", u.ID, u.Name, u.Age)
	if err != nil {
		fmt.Println(err)
	} else {
		result.Close()
	}
	results1, err := db.Query("SELECT id,user,age FROM users")
	if err != nil {
		panic(err.Error())

	}
	for results1.Next() {
		var id, age int
		var name string
		err = results1.Scan(&id, &name, &age)
		if err != nil {
			panic(err.Error())
		}
		if name == u.Name && age == u.Age {
			idReturn = id
		}
	}

	return
}

func Make_friends_SQL(id1 int, id2 int, db *sql.DB) (string, string) {
	var friends1, friends2 string
	result, err := db.Query("INSERT INTO friends VALUES (?,?)", id1, id2)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Close()

	results2, err := db.Query("INSERT INTO friends VALUES (?,?)", id2, id1)
	if err != nil {
		fmt.Println(err)
	}
	defer results2.Close()

	results, err := db.Query("SELECT id,user FROM users")
	if err != nil {
		panic(err.Error())

	}
	for results.Next() {
		var id int
		var name string
		err = results.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		if id1 == id {
			friends1 = name
		}
		if id2 == id {
			friends2 = name
		}

	}
	fmt.Println(friends1, "and", friends2, "friends")
	return friends1, friends2

}

func Delete_SQL(idDelete int, db *sql.DB) string {
	var nameDel string

	results1, err := db.Query("SELECT id,user FROM users")
	if err != nil {
		panic(err.Error())

	}
	for results1.Next() {
		var id int
		var name string
		err = results1.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		if id == idDelete {
			nameDel = name
		}
	}

	results2, err := db.Query("DELETE FROM friends WHERE id_user =? OR id_friends =?", idDelete, idDelete)
	if err != nil {
		panic(err.Error())

	}
	defer results2.Close()

	results3, err := db.Query("DELETE FROM users WHERE id=?", idDelete)
	if err != nil {
		panic(err.Error())

	}
	defer results3.Close()

	return nameDel

}

func GetFriends_SQL(idUSER int, db *sql.DB) (string, string) {
	var nameDD string

	var st string

	results1, err := db.Query("SELECT id,user FROM users")
	if err != nil {
		panic(err.Error())

	}
	for results1.Next() {
		var id int
		var name string
		err = results1.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		if idUSER == id {
			nameDD = name
		}
	}

	results, err := db.Query(
		"SELECT users.user FROM friends JOIN users ON users.id = id_friends WHERE id_user = (?)", idUSER)
	if err != nil {
		panic(err.Error())

	}
	for results.Next() {
		var name string
		err = results.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		st += name + " "
	}
	return st, nameDD
}

func ReplacementAgeSQL(nID int, newAge int, db *sql.DB) string {
	var nameDD string

	results1, err := db.Query("SELECT id,user FROM users")
	if err != nil {
		panic(err.Error())

	}
	for results1.Next() {
		var id int
		var name string
		err = results1.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		if nID == id {
			nameDD = name
		}
	}

	results, err := db.Query("UPDATE users SET age = (?) WHERE id = (?)", newAge, nID)
	if err != nil {
		panic(err.Error())

	}
	defer results.Close()
	return nameDD

}
