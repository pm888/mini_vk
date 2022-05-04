package users

import "fmt"

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Friends   []string `json:"friends"`
	FriendsID []int
}

type FriendsStr struct {
	SourceFriends int `json:"source_id"`
	TargetFriends int `json:"target_id"`
}

type DeleteFriends struct {
	UserIDToDelete int `json:"target_id"`
}

type ReplacementAge struct {
	NewAge int `json:"new age"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("ID - %d,user name - %s,age - %d, friends - %s\n", u.ID, u.Name, u.Age, u.Friends)
}
