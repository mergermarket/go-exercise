package main

import (
	"testing"
	"reflect"
	"./friends"
)

func TestMakeFriendsMutable(t *testing.T) {

	bob := friends.New("bob")
	alice := friends.New("alice")
	friends.MakeFriendsMutable(&bob, &alice)

	if !reflect.DeepEqual(bob.ListFriends(), []string{"alice"}) {
		t.Errorf("expected alice, got %s", bob.ListFriends())
	}

	if !reflect.DeepEqual(alice.ListFriends(), []string{"bob"}) {
		t.Errorf("expected bob, got %s", alice.ListFriends())
	}
}

func TestMakeFriendsImmutable(t *testing.T) {
	bob := friends.New("bob")
	alice := friends.New("alice")

	bob, alice = friends.MakeFriendsImmutable(bob, alice)

	if len(alice.ListFriends()) != 1 {
		t.Fail()
	}

	if len(bob.ListFriends()) != 1 {
		t.Fail()
	}
}
