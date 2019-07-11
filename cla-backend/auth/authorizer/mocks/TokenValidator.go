// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT
package mocks

import http "net/http"
import jwt "gopkg.in/square/go-jose.v2/jwt"
import mock "github.com/stretchr/testify/mock"

// TokenValidator is an autogenerated mock type for the TokenValidator type
type TokenValidator struct {
	mock.Mock
}

// Claims provides a mock function with given fields: _a0, _a1, _a2
func (_m *TokenValidator) Claims(_a0 *http.Request, _a1 *jwt.JSONWebToken, _a2 ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Request, *jwt.JSONWebToken, ...interface{}) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateRequest provides a mock function with given fields: _a0
func (_m *TokenValidator) ValidateRequest(_a0 *http.Request) (*jwt.JSONWebToken, error) {
	ret := _m.Called(_a0)

	var r0 *jwt.JSONWebToken
	if rf, ok := ret.Get(0).(func(*http.Request) *jwt.JSONWebToken); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.JSONWebToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
