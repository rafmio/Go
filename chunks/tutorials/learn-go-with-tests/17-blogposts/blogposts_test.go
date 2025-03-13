package main

import (
	"blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Post1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fsTestMapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assertNoError(t, err)

	assertPostsLength(t, posts, fs)

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
	World`,
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertPostsLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
