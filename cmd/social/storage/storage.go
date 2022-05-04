package storage

import "mymod/users"

var baza = make(map[int]*users.User)

func RemoveIndexString(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndexInt(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func Put(u *users.User) {
	baza[u.ID] = u
}

func GetAll() string {
	respons := ""
	for _, user := range baza {
		respons += user.ToString()
	}
	return respons
}

func Make_friends(id_one int, id_two int) (string, string) {
	var friends1, friends2 string
	var teFr1, teFr2 int

	for i := range baza {
		if i == id_one {
			friends1 = baza[i].Name
			teFr1 = baza[i].ID
		}
	}
	for j := range baza {
		if j == id_two {
			friends2 = baza[j].Name
			teFr2 = baza[j].ID
		}

	}
	baza[id_one].Friends = append(baza[id_one].Friends, friends2)
	baza[id_two].Friends = append(baza[id_two].Friends, friends1)
	baza[id_one].FriendsID = append(baza[id_one].FriendsID, teFr2)
	baza[id_two].FriendsID = append(baza[id_two].FriendsID, teFr1)

	return friends1, friends2

}

func Delete(id_delet int) string {
	namedel := baza[id_delet].Name
	for _, user := range baza[id_delet].FriendsID {
		for i, friendID := range baza[user].FriendsID {

			if friendID == id_delet {

				baza[user].Friends = RemoveIndexString(baza[user].Friends, i)
				baza[user].FriendsID = RemoveIndexInt(baza[user].FriendsID, i)
			}
		}
	}
	delete(baza, id_delet)
	return namedel
}

func GetFriends(newStrInt int) (string, string) {
	var st, name string
	name = baza[newStrInt].Name
	for j := range baza {
		if j == newStrInt {
			for _, i := range baza[j].Friends {
				st += i + " "

			}
		}
	}
	return name, st
}

func ReplacementAge(nID int, newAge int) string {
	for id := range baza {
		if id == nID {
			baza[id].Age = newAge
		}
	}
	return baza[nID].Name
}
