// Automatically generated by MockGen. DO NOT EDIT!
// Source: manager.go

package virtwrap

import (
	gomock "github.com/golang/mock/gomock"
	libvirt_go "github.com/libvirt/libvirt-go"

	v1 "kubevirt.io/kubevirt/pkg/api/v1"
)

// Mock of DomainManager interface
type MockDomainManager struct {
	ctrl     *gomock.Controller
	recorder *_MockDomainManagerRecorder
}

// Recorder for MockDomainManager (not exported)
type _MockDomainManagerRecorder struct {
	mock *MockDomainManager
}

func NewMockDomainManager(ctrl *gomock.Controller) *MockDomainManager {
	mock := &MockDomainManager{ctrl: ctrl}
	mock.recorder = &_MockDomainManagerRecorder{mock}
	return mock
}

func (_m *MockDomainManager) EXPECT() *_MockDomainManagerRecorder {
	return _m.recorder
}

func (_m *MockDomainManager) SyncVM(_param0 *v1.VM) error {
	ret := _m.ctrl.Call(_m, "SyncVM", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) SyncVM(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SyncVM", arg0)
}

func (_m *MockDomainManager) KillVM(_param0 *v1.VM) error {
	ret := _m.ctrl.Call(_m, "KillVM", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) KillVM(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillVM", arg0)
}

// Mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockConnectionRecorder
}

// Recorder for MockConnection (not exported)
type _MockConnectionRecorder struct {
	mock *MockConnection
}

func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &_MockConnectionRecorder{mock}
	return mock
}

func (_m *MockConnection) EXPECT() *_MockConnectionRecorder {
	return _m.recorder
}

func (_m *MockConnection) LookupDomainByName(name string) (VirDomain, error) {
	ret := _m.ctrl.Call(_m, "LookupDomainByName", name)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) LookupDomainByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LookupDomainByName", arg0)
}

func (_m *MockConnection) DomainDefineXML(xml string) (VirDomain, error) {
	ret := _m.ctrl.Call(_m, "DomainDefineXML", xml)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) DomainDefineXML(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainDefineXML", arg0)
}

func (_m *MockConnection) Close() (int, error) {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockConnection) DomainEventLifecycleRegister(callback libvirt_go.DomainEventLifecycleCallback) error {
	ret := _m.ctrl.Call(_m, "DomainEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventLifecycleRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventLifecycleRegister", arg0)
}

func (_m *MockConnection) ListAllDomains(flags libvirt_go.ConnectListAllDomainsFlags) ([]VirDomain, error) {
	ret := _m.ctrl.Call(_m, "ListAllDomains", flags)
	ret0, _ := ret[0].([]VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) ListAllDomains(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllDomains", arg0)
}

func (_m *MockConnection) NewStream(flags libvirt_go.StreamFlags) (Stream, error) {
	ret := _m.ctrl.Call(_m, "NewStream", flags)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) NewStream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewStream", arg0)
}

// Mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *_MockStreamRecorder
}

// Recorder for MockStream (not exported)
type _MockStreamRecorder struct {
	mock *MockStream
}

func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &_MockStreamRecorder{mock}
	return mock
}

func (_m *MockStream) EXPECT() *_MockStreamRecorder {
	return _m.recorder
}

func (_m *MockStream) Read(p []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStreamRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0)
}

func (_m *MockStream) Write(p []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStreamRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}

func (_m *MockStream) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStreamRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockStream) UnderlyingStream() *libvirt_go.Stream {
	ret := _m.ctrl.Call(_m, "UnderlyingStream")
	ret0, _ := ret[0].(*libvirt_go.Stream)
	return ret0
}

func (_mr *_MockStreamRecorder) UnderlyingStream() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnderlyingStream")
}

// Mock of VirDomain interface
type MockVirDomain struct {
	ctrl     *gomock.Controller
	recorder *_MockVirDomainRecorder
}

// Recorder for MockVirDomain (not exported)
type _MockVirDomainRecorder struct {
	mock *MockVirDomain
}

func NewMockVirDomain(ctrl *gomock.Controller) *MockVirDomain {
	mock := &MockVirDomain{ctrl: ctrl}
	mock.recorder = &_MockVirDomainRecorder{mock}
	return mock
}

func (_m *MockVirDomain) EXPECT() *_MockVirDomainRecorder {
	return _m.recorder
}

func (_m *MockVirDomain) GetState() (libvirt_go.DomainState, int, error) {
	ret := _m.ctrl.Call(_m, "GetState")
	ret0, _ := ret[0].(libvirt_go.DomainState)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockVirDomainRecorder) GetState() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetState")
}

func (_m *MockVirDomain) Create() error {
	ret := _m.ctrl.Call(_m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Create() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create")
}

func (_m *MockVirDomain) Resume() error {
	ret := _m.ctrl.Call(_m, "Resume")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Resume() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Resume")
}

func (_m *MockVirDomain) Destroy() error {
	ret := _m.ctrl.Call(_m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Destroy() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Destroy")
}

func (_m *MockVirDomain) GetName() (string, error) {
	ret := _m.ctrl.Call(_m, "GetName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetName")
}

func (_m *MockVirDomain) GetUUIDString() (string, error) {
	ret := _m.ctrl.Call(_m, "GetUUIDString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetUUIDString() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUUIDString")
}

func (_m *MockVirDomain) GetXMLDesc(flags libvirt_go.DomainXMLFlags) (string, error) {
	ret := _m.ctrl.Call(_m, "GetXMLDesc", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetXMLDesc(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetXMLDesc", arg0)
}

func (_m *MockVirDomain) Undefine() error {
	ret := _m.ctrl.Call(_m, "Undefine")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Undefine() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Undefine")
}

func (_m *MockVirDomain) OpenConsole(devname string, stream *libvirt_go.Stream, flags libvirt_go.DomainConsoleFlags) error {
	ret := _m.ctrl.Call(_m, "OpenConsole", devname, stream, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) OpenConsole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenConsole", arg0, arg1, arg2)
}

func (_m *MockVirDomain) Free() error {
	ret := _m.ctrl.Call(_m, "Free")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Free() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Free")
}
