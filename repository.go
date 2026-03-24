package quackgit

import (
	"os"

	"codeberg.org/lindenii/furgit/object/resolve"
	furgitrepository "codeberg.org/lindenii/furgit/repository"
)

// Repository is a higher-level wrapper around a furgit repository.
type Repository struct {
	repo     *furgitrepository.Repository
	resolver *resolve.Resolver
}

// Open opens a repository from disk.
func Open(path string) (*Repository, error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}

	repo, err := furgitrepository.Open(root)
	if err != nil {
		_ = root.Close()

		return nil, err
	}

	_ = root.Close()

	return &Repository{
		repo:     repo,
		resolver: repo.Resolver(),
	}, nil
}

// Wrap wraps an existing furgit repository.
func Wrap(repo *furgitrepository.Repository) *Repository {
	return &Repository{
		repo:     repo,
		resolver: repo.Resolver(),
	}
}

// Close closes the repository.
func (repo *Repository) Close() error {
	return repo.repo.Close()
}

// Algorithm returns the repository object ID algorithm.
func (repo *Repository) Algorithm() Algorithm {
	return repo.repo.Algorithm()
}

// Furgit returns the underlying furgit repository.
func (repo *Repository) Furgit() *furgitrepository.Repository {
	return repo.repo
}
