package quackgit

import (
	"codeberg.org/lindenii/furgit/object"
	"codeberg.org/lindenii/furgit/object/stored"
)

// Object resolves a revision to its exact object.
func (repo *Repository) Object(rev string) (*stored.Stored[object.Object], error) {
	id, err := repo.Resolve(rev)
	if err != nil {
		return nil, err
	}

	return repo.resolver.ExactObject(id)
}

// Commit resolves a revision to a commit.
func (repo *Repository) Commit(rev string) (*stored.Stored[*object.Commit], error) {
	id, err := repo.Resolve(rev)
	if err != nil {
		return nil, err
	}

	return repo.resolver.PeelToCommit(id)
}

// Tree resolves a revision to a tree.
func (repo *Repository) Tree(rev string) (*stored.Stored[*object.Tree], error) {
	id, err := repo.Resolve(rev)
	if err != nil {
		return nil, err
	}

	return repo.resolver.PeelToTree(id)
}

// Blob resolves a revision to a blob.
func (repo *Repository) Blob(rev string) (*stored.Stored[*object.Blob], error) {
	id, err := repo.Resolve(rev)
	if err != nil {
		return nil, err
	}

	return repo.resolver.PeelToBlob(id)
}

// Tag resolves a revision to an exact tag object.
func (repo *Repository) Tag(rev string) (*stored.Stored[*object.Tag], error) {
	id, err := repo.Resolve(rev)
	if err != nil {
		return nil, err
	}

	return repo.resolver.ExactTag(id)
}
