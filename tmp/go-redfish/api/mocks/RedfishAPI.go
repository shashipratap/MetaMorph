// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "opendev.org/airship/go-redfish/client"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RedfishAPI is an autogenerated mock type for the RedfishAPI type
type RedfishAPI struct {
	mock.Mock
}

// CreateVirtualDisk provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RedfishAPI) CreateVirtualDisk(_a0 context.Context, _a1 string, _a2 string, _a3 client.CreateVirtualDiskRequestBody) (client.RedfishError, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 client.RedfishError
	if rf, ok := ret.Get(0).(func(context.Context, string, string, client.CreateVirtualDiskRequestBody) client.RedfishError); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(client.RedfishError)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, client.CreateVirtualDiskRequestBody) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, client.CreateVirtualDiskRequestBody) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteVirtualdisk provides a mock function with given fields: _a0, _a1, _a2
func (_m *RedfishAPI) DeleteVirtualdisk(_a0 context.Context, _a1 string, _a2 string) (*http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *http.Response); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EjectVirtualMedia provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RedfishAPI) EjectVirtualMedia(_a0 context.Context, _a1 string, _a2 string, _a3 map[string]interface{}) (client.RedfishError, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 client.RedfishError
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]interface{}) client.RedfishError); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(client.RedfishError)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, map[string]interface{}) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FirmwareInventory provides a mock function with given fields: _a0
func (_m *RedfishAPI) FirmwareInventory(_a0 context.Context) (client.Collection, *http.Response, error) {
	ret := _m.Called(_a0)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(context.Context) client.Collection); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(client.Collection)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FirmwareInventoryDownloadImage provides a mock function with given fields: _a0, _a1
func (_m *RedfishAPI) FirmwareInventoryDownloadImage(_a0 context.Context, _a1 client.InlineObject) (client.RedfishError, *http.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.RedfishError
	if rf, ok := ret.Get(0).(func(context.Context, client.InlineObject) client.RedfishError); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.RedfishError)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, client.InlineObject) *http.Response); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, client.InlineObject) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetManager provides a mock function with given fields: _a0, _a1
func (_m *RedfishAPI) GetManager(_a0 context.Context, _a1 string) (client.Manager, *http.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.Manager
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Manager); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.Manager)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetManagerVirtualMedia provides a mock function with given fields: _a0, _a1, _a2
func (_m *RedfishAPI) GetManagerVirtualMedia(_a0 context.Context, _a1 string, _a2 string) (client.VirtualMedia, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 client.VirtualMedia
	if rf, ok := ret.Get(0).(func(context.Context, string, string) client.VirtualMedia); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(client.VirtualMedia)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRoot provides a mock function with given fields: _a0
func (_m *RedfishAPI) GetRoot(_a0 context.Context) (client.Root, *http.Response, error) {
	ret := _m.Called(_a0)

	var r0 client.Root
	if rf, ok := ret.Get(0).(func(context.Context) client.Root); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(client.Root)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetSystem provides a mock function with given fields: _a0, _a1
func (_m *RedfishAPI) GetSystem(_a0 context.Context, _a1 string) (client.ComputerSystem, *http.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.ComputerSystem
	if rf, ok := ret.Get(0).(func(context.Context, string) client.ComputerSystem); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.ComputerSystem)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTask provides a mock function with given fields: _a0, _a1
func (_m *RedfishAPI) GetTask(_a0 context.Context, _a1 string) (client.Task, *http.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Task); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.Task)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTaskList provides a mock function with given fields: _a0
func (_m *RedfishAPI) GetTaskList(_a0 context.Context) (client.Collection, *http.Response, error) {
	ret := _m.Called(_a0)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(context.Context) client.Collection); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(client.Collection)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetVolumes provides a mock function with given fields: _a0, _a1, _a2
func (_m *RedfishAPI) GetVolumes(_a0 context.Context, _a1 string, _a2 string) (client.Collection, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string, string) client.Collection); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(client.Collection)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InsertVirtualMedia provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RedfishAPI) InsertVirtualMedia(_a0 context.Context, _a1 string, _a2 string, _a3 client.InsertMediaRequestBody) (client.RedfishError, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 client.RedfishError
	if rf, ok := ret.Get(0).(func(context.Context, string, string, client.InsertMediaRequestBody) client.RedfishError); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(client.RedfishError)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string, client.InsertMediaRequestBody) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, client.InsertMediaRequestBody) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListManagerVirtualMedia provides a mock function with given fields: _a0, _a1
func (_m *RedfishAPI) ListManagerVirtualMedia(_a0 context.Context, _a1 string) (client.Collection, *http.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string) client.Collection); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.Collection)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListManagers provides a mock function with given fields: _a0
func (_m *RedfishAPI) ListManagers(_a0 context.Context) (client.Collection, *http.Response, error) {
	ret := _m.Called(_a0)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(context.Context) client.Collection); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(client.Collection)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListSystems provides a mock function with given fields: _a0
func (_m *RedfishAPI) ListSystems(_a0 context.Context) (client.Collection, *http.Response, error) {
	ret := _m.Called(_a0)

	var r0 client.Collection
	if rf, ok := ret.Get(0).(func(context.Context) client.Collection); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(client.Collection)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResetSystem provides a mock function with given fields: _a0, _a1, _a2
func (_m *RedfishAPI) ResetSystem(_a0 context.Context, _a1 string, _a2 client.ResetRequestBody) (client.RedfishError, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 client.RedfishError
	if rf, ok := ret.Get(0).(func(context.Context, string, client.ResetRequestBody) client.RedfishError); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(client.RedfishError)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, client.ResetRequestBody) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, client.ResetRequestBody) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetSystem provides a mock function with given fields: _a0, _a1, _a2
func (_m *RedfishAPI) SetSystem(_a0 context.Context, _a1 string, _a2 client.ComputerSystem) (client.ComputerSystem, *http.Response, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 client.ComputerSystem
	if rf, ok := ret.Get(0).(func(context.Context, string, client.ComputerSystem) client.ComputerSystem); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(client.ComputerSystem)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, client.ComputerSystem) *http.Response); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, client.ComputerSystem) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateService provides a mock function with given fields: _a0
func (_m *RedfishAPI) UpdateService(_a0 context.Context) (client.UpdateService, *http.Response, error) {
	ret := _m.Called(_a0)

	var r0 client.UpdateService
	if rf, ok := ret.Get(0).(func(context.Context) client.UpdateService); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(client.UpdateService)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateServiceSimpleUpdate provides a mock function with given fields: _a0, _a1
func (_m *RedfishAPI) UpdateServiceSimpleUpdate(_a0 context.Context, _a1 client.SimpleUpdateRequestBody) (client.RedfishError, *http.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 client.RedfishError
	if rf, ok := ret.Get(0).(func(context.Context, client.SimpleUpdateRequestBody) client.RedfishError); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(client.RedfishError)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, client.SimpleUpdateRequestBody) *http.Response); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, client.SimpleUpdateRequestBody) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}