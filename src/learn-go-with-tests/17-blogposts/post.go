package blogposts

// Post represents a post on a blog
type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

const {
	titleSeparator = "Title:"
	descriptionSeparator = "Description:"
	tagSeparator = "Tags:"
}

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)
}