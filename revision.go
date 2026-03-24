package quackgit

import (
	"codeberg.org/lindenii/furgit/objectid"
	furgitref "codeberg.org/lindenii/furgit/ref"
)

// Resolve resolves a revision to an object ID.
func (repo *Repository) Resolve(rev string) (ObjectID, error) {
	if rev == "" {
		return ObjectID{}, ErrEmptyRevision
	}

	detached, err := repo.repo.Refs().ResolveToDetached(rev)
	if err == nil {
		return detached.ID, nil
	}

	if !isRefNotFound(err) {
		return ObjectID{}, err
	}

	id, err := objectid.ParseHex(repo.repo.Algorithm(), rev)
	if err == nil {
		return id, nil
	}

	return ObjectID{}, &RevisionNotFoundError{Revision: rev}
}

// Reference resolves a reference name.
func (repo *Repository) Reference(name string) (furgitref.Ref, error) {
	if name == "" {
		return nil, ErrEmptyReferenceName
	}

	ref, err := repo.repo.Refs().Resolve(name)
	if err != nil {
		if isRefNotFound(err) {
			return nil, &ReferenceNotFoundError{Name: name}
		}

		return nil, err
	}

	return ref, nil
}
