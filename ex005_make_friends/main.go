package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	friends []*person
}

func makeFriends(a *person, b *person) {
	addAsFriendIfNotAlready(a, b)
	addAsFriendIfNotAlready(b, a)
}

func addAsFriendIfNotAlready(a *person, b *person) {
	if !isFriendsWith(a, b) {
		a.friends = append(a.friends, b)
	}
}

func main() {
	ross := newPerson("Ross")
	rachel := newPerson("Rachel")
	joey := newPerson("Joey")
	phoebe := newPerson("Phoebe")
	chandler := newPerson("Chandler")
	monica := newPerson("Monica")

	fmt.Println("----------- Monica and Ross are siblings")
	makeFriends(&monica, &ross)
	listFriends(&monica)
	listFriends(&ross)

	fmt.Println("----------- Ross and Chandler are friends from high school")
	makeFriends(&ross, &chandler)
	listFriends(&ross)
	listFriends(&chandler)

	fmt.Println("----------- Monica and Rachel are friends from high school")
	makeFriends(&monica, &rachel)
	listFriends(&monica)
	listFriends(&rachel)

	fmt.Println("----------- Monica and Chandler become friends (?) when they grow up")
	makeFriends(&monica, &chandler)
	listFriends(&monica)
	listFriends(&chandler)

	fmt.Println("----------- Monica and Phoebe meet via a roommate ad")
	makeFriends(&monica, &phoebe)
	listFriends(&monica)
	listFriends(&phoebe)

	fmt.Println("----------- Monica introduces Phoebe to Ross and Chandler")
	makeFriends(&phoebe, &ross)
	makeFriends(&phoebe, &chandler)
	listFriends(&phoebe)
	listFriends(&ross)
	listFriends(&chandler)

	fmt.Println("----------- Joey and Chandler meet via a roommate ad")
	makeFriends(&joey, &chandler)
	listFriends(&joey)
	listFriends(&chandler)

	fmt.Println("----------- Chandler introduces Joey to everyone except Rachel")
	makeFriends(&joey, &monica)
	makeFriends(&joey, &phoebe)
	makeFriends(&joey, &ross)
	listFriends(&joey)
	listFriends(&ross)
	listFriends(&monica)
	listFriends(&phoebe)

	fmt.Println("----------- Rachel runs away from her wedding and meets everyone at Central Perk")
	makeFriends(&rachel, &ross)
	makeFriends(&rachel, &chandler)
	makeFriends(&rachel, &joey)
	makeFriends(&rachel, &phoebe)

	fmt.Println("----------- Final friends lists:")
	listFriends(&monica)
	listFriends(&rachel)
	listFriends(&joey)
	listFriends(&ross)
	listFriends(&chandler)
	listFriends(&phoebe)
}

func newPerson(name string) person {
	return person{
		name:    name,
		friends: make([]*person, 0), // !!!!!
	}
}

func listFriends(p *person) {
	friendsNames := make([]string, 0)
	for _, friend := range p.friends {
		friendsNames = append(friendsNames, friend.name)
	}
	fmt.Printf("Friends of %s:  %s\n", p.name, strings.Join(friendsNames, ", "))
}

// isFriendsWith(a, b) returns true if a has b in their friends list.
func isFriendsWith(a *person, b *person) bool {
	for _, friend := range a.friends {
		if friend == b {
			return true
		}
	}
	return false
}
