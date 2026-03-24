package quackgit

import (
	"errors"
	"fmt"

	"codeberg.org/lindenii/furgit/refstore"
)

// ErrEmptyRevision indicates that a revision input was empty.
var ErrEmptyRevision = errors.New("quackgit: empty revision")

// ErrEmptyReferenceName indicates that a reference name was empty.
var ErrEmptyReferenceName = errors.New("quackgit: empty reference name")

// RevisionNotFoundError indicates that a revision could not be resolved.
type RevisionNotFoundError struct {
	Revision string
}

// Error implements error.
func (err *RevisionNotFoundError) Error() string {
	return fmt.Sprintf("quackgit: revision %q not found", err.Revision)
}

// ReferenceNotFoundError indicates that a reference could not be resolved.
type ReferenceNotFoundError struct {
	Name string
}

// Error implements error.
func (err *ReferenceNotFoundError) Error() string {
	return fmt.Sprintf("quackgit: reference %q not found", err.Name)
}

func isRefNotFound(err error) bool {
	return errors.Is(err, refstore.ErrReferenceNotFound)
}
