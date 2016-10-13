package friends

type Person struct {
	name string
	friends []Person
}

func New(name string) Person {
	return Person{name,nil}
}

func MakeFriendsMutable(first, second *Person) {
	first.friends = append(first.friends, *second)
	second.friends = append(second.friends, *first)
}

func MakeFriendsImmutable(first, second Person) (Person, Person) {
	first.friends = append(first.friends, second)
	second.friends = append(second.friends, first)

	return first, second
}

func (p Person) ListFriends() (result []string) {
	for _, person := range p.friends {
		result = append(result, person.name)
	}
	return
}
