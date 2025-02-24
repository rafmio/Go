package blogposts

import (
	"io/fs"
)

// NewPostsFromFS return a collection of blog posts from a file system.
// If it does not conform to the format then it'll return an error
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {

	// ReadDir reads the named directory and returns a list of
	// directory entries sorted by filename.
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err // TODO: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

// Where can fs.FS can come from?
// For e.g.:
// dir := os.DirFS("/path/to/directory")
