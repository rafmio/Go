package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/quii/learn-go-with-tests/reading-files"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, _ := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d ponsts", len(posts), len(fs))
	}
}
