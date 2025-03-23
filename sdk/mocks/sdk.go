// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	errors "github.com/hantdev/certs/errors"
	mock "github.com/stretchr/testify/mock"

	sdk "github.com/hantdev/certs/sdk"
)

// MockSDK is an autogenerated mock type for the SDK type
type MockSDK struct {
	mock.Mock
}

type MockSDK_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSDK) EXPECT() *MockSDK_Expecter {
	return &MockSDK_Expecter{mock: &_m.Mock}
}

// DeleteCert provides a mock function with given fields: entityID
func (_m *MockSDK) DeleteCert(entityID string) errors.SDKError {
	ret := _m.Called(entityID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCert")
	}

	var r0 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) errors.SDKError); ok {
		r0 = rf(entityID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.SDKError)
		}
	}

	return r0
}

// MockSDK_DeleteCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCert'
type MockSDK_DeleteCert_Call struct {
	*mock.Call
}

// DeleteCert is a helper method to define mock.On call
//   - entityID string
func (_e *MockSDK_Expecter) DeleteCert(entityID interface{}) *MockSDK_DeleteCert_Call {
	return &MockSDK_DeleteCert_Call{Call: _e.mock.On("DeleteCert", entityID)}
}

func (_c *MockSDK_DeleteCert_Call) Run(run func(entityID string)) *MockSDK_DeleteCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_DeleteCert_Call) Return(_a0 errors.SDKError) *MockSDK_DeleteCert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSDK_DeleteCert_Call) RunAndReturn(run func(string) errors.SDKError) *MockSDK_DeleteCert_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadCA provides a mock function with given fields: token
func (_m *MockSDK) DownloadCA(token string) (sdk.CertificateBundle, errors.SDKError) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for DownloadCA")
	}

	var r0 sdk.CertificateBundle
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) (sdk.CertificateBundle, errors.SDKError)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) sdk.CertificateBundle); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(sdk.CertificateBundle)
	}

	if rf, ok := ret.Get(1).(func(string) errors.SDKError); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_DownloadCA_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadCA'
type MockSDK_DownloadCA_Call struct {
	*mock.Call
}

// DownloadCA is a helper method to define mock.On call
//   - token string
func (_e *MockSDK_Expecter) DownloadCA(token interface{}) *MockSDK_DownloadCA_Call {
	return &MockSDK_DownloadCA_Call{Call: _e.mock.On("DownloadCA", token)}
}

func (_c *MockSDK_DownloadCA_Call) Run(run func(token string)) *MockSDK_DownloadCA_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_DownloadCA_Call) Return(_a0 sdk.CertificateBundle, _a1 errors.SDKError) *MockSDK_DownloadCA_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_DownloadCA_Call) RunAndReturn(run func(string) (sdk.CertificateBundle, errors.SDKError)) *MockSDK_DownloadCA_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadCert provides a mock function with given fields: token, serialNumber
func (_m *MockSDK) DownloadCert(token string, serialNumber string) (sdk.CertificateBundle, errors.SDKError) {
	ret := _m.Called(token, serialNumber)

	if len(ret) == 0 {
		panic("no return value specified for DownloadCert")
	}

	var r0 sdk.CertificateBundle
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string, string) (sdk.CertificateBundle, errors.SDKError)); ok {
		return rf(token, serialNumber)
	}
	if rf, ok := ret.Get(0).(func(string, string) sdk.CertificateBundle); ok {
		r0 = rf(token, serialNumber)
	} else {
		r0 = ret.Get(0).(sdk.CertificateBundle)
	}

	if rf, ok := ret.Get(1).(func(string, string) errors.SDKError); ok {
		r1 = rf(token, serialNumber)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_DownloadCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadCert'
type MockSDK_DownloadCert_Call struct {
	*mock.Call
}

// DownloadCert is a helper method to define mock.On call
//   - token string
//   - serialNumber string
func (_e *MockSDK_Expecter) DownloadCert(token interface{}, serialNumber interface{}) *MockSDK_DownloadCert_Call {
	return &MockSDK_DownloadCert_Call{Call: _e.mock.On("DownloadCert", token, serialNumber)}
}

func (_c *MockSDK_DownloadCert_Call) Run(run func(token string, serialNumber string)) *MockSDK_DownloadCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockSDK_DownloadCert_Call) Return(_a0 sdk.CertificateBundle, _a1 errors.SDKError) *MockSDK_DownloadCert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_DownloadCert_Call) RunAndReturn(run func(string, string) (sdk.CertificateBundle, errors.SDKError)) *MockSDK_DownloadCert_Call {
	_c.Call.Return(run)
	return _c
}

// GetCAToken provides a mock function with given fields:
func (_m *MockSDK) GetCAToken() (sdk.Token, errors.SDKError) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCAToken")
	}

	var r0 sdk.Token
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func() (sdk.Token, errors.SDKError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() sdk.Token); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sdk.Token)
	}

	if rf, ok := ret.Get(1).(func() errors.SDKError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_GetCAToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCAToken'
type MockSDK_GetCAToken_Call struct {
	*mock.Call
}

// GetCAToken is a helper method to define mock.On call
func (_e *MockSDK_Expecter) GetCAToken() *MockSDK_GetCAToken_Call {
	return &MockSDK_GetCAToken_Call{Call: _e.mock.On("GetCAToken")}
}

func (_c *MockSDK_GetCAToken_Call) Run(run func()) *MockSDK_GetCAToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSDK_GetCAToken_Call) Return(_a0 sdk.Token, _a1 errors.SDKError) *MockSDK_GetCAToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_GetCAToken_Call) RunAndReturn(run func() (sdk.Token, errors.SDKError)) *MockSDK_GetCAToken_Call {
	_c.Call.Return(run)
	return _c
}

// IssueCert provides a mock function with given fields: entityID, ttl, ipAddrs, opts
func (_m *MockSDK) IssueCert(entityID string, ttl string, ipAddrs []string, opts sdk.Options) (sdk.Certificate, errors.SDKError) {
	ret := _m.Called(entityID, ttl, ipAddrs, opts)

	if len(ret) == 0 {
		panic("no return value specified for IssueCert")
	}

	var r0 sdk.Certificate
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string, string, []string, sdk.Options) (sdk.Certificate, errors.SDKError)); ok {
		return rf(entityID, ttl, ipAddrs, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, []string, sdk.Options) sdk.Certificate); ok {
		r0 = rf(entityID, ttl, ipAddrs, opts)
	} else {
		r0 = ret.Get(0).(sdk.Certificate)
	}

	if rf, ok := ret.Get(1).(func(string, string, []string, sdk.Options) errors.SDKError); ok {
		r1 = rf(entityID, ttl, ipAddrs, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_IssueCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IssueCert'
type MockSDK_IssueCert_Call struct {
	*mock.Call
}

// IssueCert is a helper method to define mock.On call
//   - entityID string
//   - ttl string
//   - ipAddrs []string
//   - opts sdk.Options
func (_e *MockSDK_Expecter) IssueCert(entityID interface{}, ttl interface{}, ipAddrs interface{}, opts interface{}) *MockSDK_IssueCert_Call {
	return &MockSDK_IssueCert_Call{Call: _e.mock.On("IssueCert", entityID, ttl, ipAddrs, opts)}
}

func (_c *MockSDK_IssueCert_Call) Run(run func(entityID string, ttl string, ipAddrs []string, opts sdk.Options)) *MockSDK_IssueCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].([]string), args[3].(sdk.Options))
	})
	return _c
}

func (_c *MockSDK_IssueCert_Call) Return(_a0 sdk.Certificate, _a1 errors.SDKError) *MockSDK_IssueCert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_IssueCert_Call) RunAndReturn(run func(string, string, []string, sdk.Options) (sdk.Certificate, errors.SDKError)) *MockSDK_IssueCert_Call {
	_c.Call.Return(run)
	return _c
}

// IssueFromCSR provides a mock function with given fields: entityID, ttl, csr
func (_m *MockSDK) IssueFromCSR(entityID string, ttl string, csr string) (sdk.Certificate, errors.SDKError) {
	ret := _m.Called(entityID, ttl, csr)

	if len(ret) == 0 {
		panic("no return value specified for IssueFromCSR")
	}

	var r0 sdk.Certificate
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string, string, string) (sdk.Certificate, errors.SDKError)); ok {
		return rf(entityID, ttl, csr)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) sdk.Certificate); ok {
		r0 = rf(entityID, ttl, csr)
	} else {
		r0 = ret.Get(0).(sdk.Certificate)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) errors.SDKError); ok {
		r1 = rf(entityID, ttl, csr)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_IssueFromCSR_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IssueFromCSR'
type MockSDK_IssueFromCSR_Call struct {
	*mock.Call
}

// IssueFromCSR is a helper method to define mock.On call
//   - entityID string
//   - ttl string
//   - csr string
func (_e *MockSDK_Expecter) IssueFromCSR(entityID interface{}, ttl interface{}, csr interface{}) *MockSDK_IssueFromCSR_Call {
	return &MockSDK_IssueFromCSR_Call{Call: _e.mock.On("IssueFromCSR", entityID, ttl, csr)}
}

func (_c *MockSDK_IssueFromCSR_Call) Run(run func(entityID string, ttl string, csr string)) *MockSDK_IssueFromCSR_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockSDK_IssueFromCSR_Call) Return(_a0 sdk.Certificate, _a1 errors.SDKError) *MockSDK_IssueFromCSR_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_IssueFromCSR_Call) RunAndReturn(run func(string, string, string) (sdk.Certificate, errors.SDKError)) *MockSDK_IssueFromCSR_Call {
	_c.Call.Return(run)
	return _c
}

// ListCerts provides a mock function with given fields: pm
func (_m *MockSDK) ListCerts(pm sdk.PageMetadata) (sdk.CertificatePage, errors.SDKError) {
	ret := _m.Called(pm)

	if len(ret) == 0 {
		panic("no return value specified for ListCerts")
	}

	var r0 sdk.CertificatePage
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(sdk.PageMetadata) (sdk.CertificatePage, errors.SDKError)); ok {
		return rf(pm)
	}
	if rf, ok := ret.Get(0).(func(sdk.PageMetadata) sdk.CertificatePage); ok {
		r0 = rf(pm)
	} else {
		r0 = ret.Get(0).(sdk.CertificatePage)
	}

	if rf, ok := ret.Get(1).(func(sdk.PageMetadata) errors.SDKError); ok {
		r1 = rf(pm)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_ListCerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCerts'
type MockSDK_ListCerts_Call struct {
	*mock.Call
}

// ListCerts is a helper method to define mock.On call
//   - pm sdk.PageMetadata
func (_e *MockSDK_Expecter) ListCerts(pm interface{}) *MockSDK_ListCerts_Call {
	return &MockSDK_ListCerts_Call{Call: _e.mock.On("ListCerts", pm)}
}

func (_c *MockSDK_ListCerts_Call) Run(run func(pm sdk.PageMetadata)) *MockSDK_ListCerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sdk.PageMetadata))
	})
	return _c
}

func (_c *MockSDK_ListCerts_Call) Return(_a0 sdk.CertificatePage, _a1 errors.SDKError) *MockSDK_ListCerts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_ListCerts_Call) RunAndReturn(run func(sdk.PageMetadata) (sdk.CertificatePage, errors.SDKError)) *MockSDK_ListCerts_Call {
	_c.Call.Return(run)
	return _c
}

// OCSP provides a mock function with given fields: serialNumber, cert
func (_m *MockSDK) OCSP(serialNumber string, cert string) (sdk.OCSPResponse, errors.SDKError) {
	ret := _m.Called(serialNumber, cert)

	if len(ret) == 0 {
		panic("no return value specified for OCSP")
	}

	var r0 sdk.OCSPResponse
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string, string) (sdk.OCSPResponse, errors.SDKError)); ok {
		return rf(serialNumber, cert)
	}
	if rf, ok := ret.Get(0).(func(string, string) sdk.OCSPResponse); ok {
		r0 = rf(serialNumber, cert)
	} else {
		r0 = ret.Get(0).(sdk.OCSPResponse)
	}

	if rf, ok := ret.Get(1).(func(string, string) errors.SDKError); ok {
		r1 = rf(serialNumber, cert)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_OCSP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OCSP'
type MockSDK_OCSP_Call struct {
	*mock.Call
}

// OCSP is a helper method to define mock.On call
//   - serialNumber string
//   - cert string
func (_e *MockSDK_Expecter) OCSP(serialNumber interface{}, cert interface{}) *MockSDK_OCSP_Call {
	return &MockSDK_OCSP_Call{Call: _e.mock.On("OCSP", serialNumber, cert)}
}

func (_c *MockSDK_OCSP_Call) Run(run func(serialNumber string, cert string)) *MockSDK_OCSP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockSDK_OCSP_Call) Return(_a0 sdk.OCSPResponse, _a1 errors.SDKError) *MockSDK_OCSP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_OCSP_Call) RunAndReturn(run func(string, string) (sdk.OCSPResponse, errors.SDKError)) *MockSDK_OCSP_Call {
	_c.Call.Return(run)
	return _c
}

// RenewCert provides a mock function with given fields: serialNumber
func (_m *MockSDK) RenewCert(serialNumber string) errors.SDKError {
	ret := _m.Called(serialNumber)

	if len(ret) == 0 {
		panic("no return value specified for RenewCert")
	}

	var r0 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) errors.SDKError); ok {
		r0 = rf(serialNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.SDKError)
		}
	}

	return r0
}

// MockSDK_RenewCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RenewCert'
type MockSDK_RenewCert_Call struct {
	*mock.Call
}

// RenewCert is a helper method to define mock.On call
//   - serialNumber string
func (_e *MockSDK_Expecter) RenewCert(serialNumber interface{}) *MockSDK_RenewCert_Call {
	return &MockSDK_RenewCert_Call{Call: _e.mock.On("RenewCert", serialNumber)}
}

func (_c *MockSDK_RenewCert_Call) Run(run func(serialNumber string)) *MockSDK_RenewCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_RenewCert_Call) Return(_a0 errors.SDKError) *MockSDK_RenewCert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSDK_RenewCert_Call) RunAndReturn(run func(string) errors.SDKError) *MockSDK_RenewCert_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveCertDownloadToken provides a mock function with given fields: serialNumber
func (_m *MockSDK) RetrieveCertDownloadToken(serialNumber string) (sdk.Token, errors.SDKError) {
	ret := _m.Called(serialNumber)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveCertDownloadToken")
	}

	var r0 sdk.Token
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) (sdk.Token, errors.SDKError)); ok {
		return rf(serialNumber)
	}
	if rf, ok := ret.Get(0).(func(string) sdk.Token); ok {
		r0 = rf(serialNumber)
	} else {
		r0 = ret.Get(0).(sdk.Token)
	}

	if rf, ok := ret.Get(1).(func(string) errors.SDKError); ok {
		r1 = rf(serialNumber)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_RetrieveCertDownloadToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveCertDownloadToken'
type MockSDK_RetrieveCertDownloadToken_Call struct {
	*mock.Call
}

// RetrieveCertDownloadToken is a helper method to define mock.On call
//   - serialNumber string
func (_e *MockSDK_Expecter) RetrieveCertDownloadToken(serialNumber interface{}) *MockSDK_RetrieveCertDownloadToken_Call {
	return &MockSDK_RetrieveCertDownloadToken_Call{Call: _e.mock.On("RetrieveCertDownloadToken", serialNumber)}
}

func (_c *MockSDK_RetrieveCertDownloadToken_Call) Run(run func(serialNumber string)) *MockSDK_RetrieveCertDownloadToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_RetrieveCertDownloadToken_Call) Return(_a0 sdk.Token, _a1 errors.SDKError) *MockSDK_RetrieveCertDownloadToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_RetrieveCertDownloadToken_Call) RunAndReturn(run func(string) (sdk.Token, errors.SDKError)) *MockSDK_RetrieveCertDownloadToken_Call {
	_c.Call.Return(run)
	return _c
}

// RevokeCert provides a mock function with given fields: serialNumber
func (_m *MockSDK) RevokeCert(serialNumber string) errors.SDKError {
	ret := _m.Called(serialNumber)

	if len(ret) == 0 {
		panic("no return value specified for RevokeCert")
	}

	var r0 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) errors.SDKError); ok {
		r0 = rf(serialNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.SDKError)
		}
	}

	return r0
}

// MockSDK_RevokeCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeCert'
type MockSDK_RevokeCert_Call struct {
	*mock.Call
}

// RevokeCert is a helper method to define mock.On call
//   - serialNumber string
func (_e *MockSDK_Expecter) RevokeCert(serialNumber interface{}) *MockSDK_RevokeCert_Call {
	return &MockSDK_RevokeCert_Call{Call: _e.mock.On("RevokeCert", serialNumber)}
}

func (_c *MockSDK_RevokeCert_Call) Run(run func(serialNumber string)) *MockSDK_RevokeCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_RevokeCert_Call) Return(_a0 errors.SDKError) *MockSDK_RevokeCert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSDK_RevokeCert_Call) RunAndReturn(run func(string) errors.SDKError) *MockSDK_RevokeCert_Call {
	_c.Call.Return(run)
	return _c
}

// ViewCA provides a mock function with given fields: token
func (_m *MockSDK) ViewCA(token string) (sdk.Certificate, errors.SDKError) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ViewCA")
	}

	var r0 sdk.Certificate
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) (sdk.Certificate, errors.SDKError)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) sdk.Certificate); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(sdk.Certificate)
	}

	if rf, ok := ret.Get(1).(func(string) errors.SDKError); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_ViewCA_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ViewCA'
type MockSDK_ViewCA_Call struct {
	*mock.Call
}

// ViewCA is a helper method to define mock.On call
//   - token string
func (_e *MockSDK_Expecter) ViewCA(token interface{}) *MockSDK_ViewCA_Call {
	return &MockSDK_ViewCA_Call{Call: _e.mock.On("ViewCA", token)}
}

func (_c *MockSDK_ViewCA_Call) Run(run func(token string)) *MockSDK_ViewCA_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_ViewCA_Call) Return(_a0 sdk.Certificate, _a1 errors.SDKError) *MockSDK_ViewCA_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_ViewCA_Call) RunAndReturn(run func(string) (sdk.Certificate, errors.SDKError)) *MockSDK_ViewCA_Call {
	_c.Call.Return(run)
	return _c
}

// ViewCert provides a mock function with given fields: serialNumber
func (_m *MockSDK) ViewCert(serialNumber string) (sdk.Certificate, errors.SDKError) {
	ret := _m.Called(serialNumber)

	if len(ret) == 0 {
		panic("no return value specified for ViewCert")
	}

	var r0 sdk.Certificate
	var r1 errors.SDKError
	if rf, ok := ret.Get(0).(func(string) (sdk.Certificate, errors.SDKError)); ok {
		return rf(serialNumber)
	}
	if rf, ok := ret.Get(0).(func(string) sdk.Certificate); ok {
		r0 = rf(serialNumber)
	} else {
		r0 = ret.Get(0).(sdk.Certificate)
	}

	if rf, ok := ret.Get(1).(func(string) errors.SDKError); ok {
		r1 = rf(serialNumber)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.SDKError)
		}
	}

	return r0, r1
}

// MockSDK_ViewCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ViewCert'
type MockSDK_ViewCert_Call struct {
	*mock.Call
}

// ViewCert is a helper method to define mock.On call
//   - serialNumber string
func (_e *MockSDK_Expecter) ViewCert(serialNumber interface{}) *MockSDK_ViewCert_Call {
	return &MockSDK_ViewCert_Call{Call: _e.mock.On("ViewCert", serialNumber)}
}

func (_c *MockSDK_ViewCert_Call) Run(run func(serialNumber string)) *MockSDK_ViewCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockSDK_ViewCert_Call) Return(_a0 sdk.Certificate, _a1 errors.SDKError) *MockSDK_ViewCert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSDK_ViewCert_Call) RunAndReturn(run func(string) (sdk.Certificate, errors.SDKError)) *MockSDK_ViewCert_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSDK creates a new instance of MockSDK. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSDK(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSDK {
	mock := &MockSDK{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}