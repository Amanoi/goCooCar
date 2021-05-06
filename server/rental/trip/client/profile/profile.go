package profile

import (
	"context"
	"coolcar/shared/id"
)

//Manager defines a profile manager
type Manager struct{}

// Verify verifes accounts status.
func (m *Manager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return id.IdentityID("identity1"), nil
}
