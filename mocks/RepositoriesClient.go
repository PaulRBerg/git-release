// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	github "github.com/google/go-github/github"
	mock "github.com/stretchr/testify/mock"

	os "os"
)

// RepositoriesClient is an autogenerated mock type for the RepositoriesClient type
type RepositoriesClient struct {
	mock.Mock
}

// CreateRelease provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RepositoriesClient) CreateRelease(_a0 context.Context, _a1 string, _a2 string, _a3 *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.RepositoryRelease) *github.RepositoryRelease); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryRelease)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.RepositoryRelease) *github.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.RepositoryRelease) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteRelease provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RepositoriesClient) DeleteRelease(_a0 context.Context, _a1 string, _a2 string, _a3 int64) (*github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *github.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.Response); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReleaseAsset provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RepositoriesClient) DeleteReleaseAsset(_a0 context.Context, _a1 string, _a2 string, _a3 int64) (*github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *github.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *github.Response); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReleaseByTag provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RepositoriesClient) GetReleaseByTag(_a0 context.Context, _a1 string, _a2 string, _a3 string) (*github.RepositoryRelease, *github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *github.RepositoryRelease); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryRelease)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) *github.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UploadReleaseAsset provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *RepositoriesClient) UploadReleaseAsset(_a0 context.Context, _a1 string, _a2 string, _a3 int64, _a4 *github.UploadOptions, _a5 *os.File) (*github.ReleaseAsset, *github.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 *github.ReleaseAsset
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, *github.UploadOptions, *os.File) *github.ReleaseAsset); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.ReleaseAsset)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64, *github.UploadOptions, *os.File) *github.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, int64, *github.UploadOptions, *os.File) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
