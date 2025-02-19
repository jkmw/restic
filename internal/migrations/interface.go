package migrations

import (
	"context"

	"github.com/restic/restic/internal/restic"
)

type RepositoryCheckOptions struct {
}

// Migration implements a data migration.
type Migration interface {
	// Check returns true if the migration can be applied to a repo.
	Check(context.Context, restic.Repository) (bool, error)

	RepoCheckOptions() *RepositoryCheckOptions

	// Apply runs the migration.
	Apply(context.Context, restic.Repository) error

	// Name returns a short name.
	Name() string

	// Descr returns a description what the migration does.
	Desc() string
}
