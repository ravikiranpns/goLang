package main

import "fmt"

func main() {
	user1 := &user{
		name: "ravi",
		age:  30,
	}

	user2 := &user{
		name: "kiran",
		age:  23,
	}
	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
