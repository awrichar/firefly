// Code generated by mockery v2.15.0. DO NOT EDIT.

package networkmapmocks

import (
	context "context"

	ffapi "github.com/hyperledger/firefly-common/pkg/ffapi"
	core "github.com/hyperledger/firefly/pkg/core"

	mock "github.com/stretchr/testify/mock"

	networkmap "github.com/hyperledger/firefly/internal/networkmap"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// GetDIDDocForIndentityByDID provides a mock function with given fields: ctx, did
func (_m *Manager) GetDIDDocForIndentityByDID(ctx context.Context, did string) (*networkmap.DIDDocument, error) {
	ret := _m.Called(ctx, did)

	var r0 *networkmap.DIDDocument
	if rf, ok := ret.Get(0).(func(context.Context, string) *networkmap.DIDDocument); ok {
		r0 = rf(ctx, did)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmap.DIDDocument)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, did)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDIDDocForIndentityByID provides a mock function with given fields: ctx, id
func (_m *Manager) GetDIDDocForIndentityByID(ctx context.Context, id string) (*networkmap.DIDDocument, error) {
	ret := _m.Called(ctx, id)

	var r0 *networkmap.DIDDocument
	if rf, ok := ret.Get(0).(func(context.Context, string) *networkmap.DIDDocument); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*networkmap.DIDDocument)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentities provides a mock function with given fields: ctx, filter
func (_m *Manager) GetIdentities(ctx context.Context, filter ffapi.AndFilter) ([]*core.Identity, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.Identity); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Identity)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetIdentitiesWithVerifiers provides a mock function with given fields: ctx, filter
func (_m *Manager) GetIdentitiesWithVerifiers(ctx context.Context, filter ffapi.AndFilter) ([]*core.IdentityWithVerifiers, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.IdentityWithVerifiers
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.IdentityWithVerifiers); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.IdentityWithVerifiers)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetIdentityByDID provides a mock function with given fields: ctx, did
func (_m *Manager) GetIdentityByDID(ctx context.Context, did string) (*core.Identity, error) {
	ret := _m.Called(ctx, did)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Identity); ok {
		r0 = rf(ctx, did)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, did)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentityByDIDWithVerifiers provides a mock function with given fields: ctx, did
func (_m *Manager) GetIdentityByDIDWithVerifiers(ctx context.Context, did string) (*core.IdentityWithVerifiers, error) {
	ret := _m.Called(ctx, did)

	var r0 *core.IdentityWithVerifiers
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.IdentityWithVerifiers); ok {
		r0 = rf(ctx, did)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.IdentityWithVerifiers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, did)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentityByID provides a mock function with given fields: ctx, id
func (_m *Manager) GetIdentityByID(ctx context.Context, id string) (*core.Identity, error) {
	ret := _m.Called(ctx, id)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Identity); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentityByIDWithVerifiers provides a mock function with given fields: ctx, id
func (_m *Manager) GetIdentityByIDWithVerifiers(ctx context.Context, id string) (*core.IdentityWithVerifiers, error) {
	ret := _m.Called(ctx, id)

	var r0 *core.IdentityWithVerifiers
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.IdentityWithVerifiers); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.IdentityWithVerifiers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentityVerifiers provides a mock function with given fields: ctx, id, filter
func (_m *Manager) GetIdentityVerifiers(ctx context.Context, id string, filter ffapi.AndFilter) ([]*core.Verifier, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, id, filter)

	var r0 []*core.Verifier
	if rf, ok := ret.Get(0).(func(context.Context, string, ffapi.AndFilter) []*core.Verifier); ok {
		r0 = rf(ctx, id, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Verifier)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, string, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, id, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, id, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNodeByNameOrID provides a mock function with given fields: ctx, nameOrID
func (_m *Manager) GetNodeByNameOrID(ctx context.Context, nameOrID string) (*core.Identity, error) {
	ret := _m.Called(ctx, nameOrID)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Identity); ok {
		r0 = rf(ctx, nameOrID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nameOrID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodes provides a mock function with given fields: ctx, filter
func (_m *Manager) GetNodes(ctx context.Context, filter ffapi.AndFilter) ([]*core.Identity, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.Identity); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Identity)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrganizationByNameOrID provides a mock function with given fields: ctx, nameOrID
func (_m *Manager) GetOrganizationByNameOrID(ctx context.Context, nameOrID string) (*core.Identity, error) {
	ret := _m.Called(ctx, nameOrID)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Identity); ok {
		r0 = rf(ctx, nameOrID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nameOrID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizations provides a mock function with given fields: ctx, filter
func (_m *Manager) GetOrganizations(ctx context.Context, filter ffapi.AndFilter) ([]*core.Identity, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.Identity); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Identity)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrganizationsWithVerifiers provides a mock function with given fields: ctx, filter
func (_m *Manager) GetOrganizationsWithVerifiers(ctx context.Context, filter ffapi.AndFilter) ([]*core.IdentityWithVerifiers, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.IdentityWithVerifiers
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.IdentityWithVerifiers); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.IdentityWithVerifiers)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetVerifierByHash provides a mock function with given fields: ctx, hash
func (_m *Manager) GetVerifierByHash(ctx context.Context, hash string) (*core.Verifier, error) {
	ret := _m.Called(ctx, hash)

	var r0 *core.Verifier
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Verifier); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Verifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVerifiers provides a mock function with given fields: ctx, filter
func (_m *Manager) GetVerifiers(ctx context.Context, filter ffapi.AndFilter) ([]*core.Verifier, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.Verifier
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.Verifier); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Verifier)
		}
	}

	var r1 *ffapi.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterIdentity provides a mock function with given fields: ctx, dto, waitConfirm
func (_m *Manager) RegisterIdentity(ctx context.Context, dto *core.IdentityCreateDTO, waitConfirm bool) (*core.Identity, error) {
	ret := _m.Called(ctx, dto, waitConfirm)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, *core.IdentityCreateDTO, bool) *core.Identity); ok {
		r0 = rf(ctx, dto, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.IdentityCreateDTO, bool) error); ok {
		r1 = rf(ctx, dto, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterNode provides a mock function with given fields: ctx, waitConfirm
func (_m *Manager) RegisterNode(ctx context.Context, waitConfirm bool) (*core.Identity, error) {
	ret := _m.Called(ctx, waitConfirm)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, bool) *core.Identity); ok {
		r0 = rf(ctx, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterNodeOrganization provides a mock function with given fields: ctx, waitConfirm
func (_m *Manager) RegisterNodeOrganization(ctx context.Context, waitConfirm bool) (*core.Identity, error) {
	ret := _m.Called(ctx, waitConfirm)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, bool) *core.Identity); ok {
		r0 = rf(ctx, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterOrganization provides a mock function with given fields: ctx, org, waitConfirm
func (_m *Manager) RegisterOrganization(ctx context.Context, org *core.IdentityCreateDTO, waitConfirm bool) (*core.Identity, error) {
	ret := _m.Called(ctx, org, waitConfirm)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, *core.IdentityCreateDTO, bool) *core.Identity); ok {
		r0 = rf(ctx, org, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.IdentityCreateDTO, bool) error); ok {
		r1 = rf(ctx, org, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdentity provides a mock function with given fields: ctx, id, dto, waitConfirm
func (_m *Manager) UpdateIdentity(ctx context.Context, id string, dto *core.IdentityUpdateDTO, waitConfirm bool) (*core.Identity, error) {
	ret := _m.Called(ctx, id, dto, waitConfirm)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.IdentityUpdateDTO, bool) *core.Identity); ok {
		r0 = rf(ctx, id, dto, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *core.IdentityUpdateDTO, bool) error); ok {
		r1 = rf(ctx, id, dto, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
