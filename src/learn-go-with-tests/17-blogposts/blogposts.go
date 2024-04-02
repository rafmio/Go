package blogposts

import (
	"io/fs"
)

// NewPostsFromFS return a collection of blog posts from a file system.
// If it does not conform to the format then it'll return an error
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
}
