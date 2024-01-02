package main

//https://www.educative.io/answers/what-is-the-factory-design-pattern-in-go
import "fmt"

type Post struct {
	author, website, content string
}

// type PostFactory struct {
// 	website, author string
// }

func newPostFactory(website string, author string) func(content string) *Post {
	return func(content string) *Post {
		return &Post{content, website, author}
	}
}

func main() {
	devPostFact := newPostFactory("Dev.to", "Ravi")
	mediumPostFact := newPostFactory("medium.com", "Kiran")

	devPost := devPostFact("We the Best Engineers For Better World")
	mediumPost := mediumPostFact("We the Best Medium of Idea's")

	fmt.Println(devPost)
	fmt.Println(mediumPost)

}
