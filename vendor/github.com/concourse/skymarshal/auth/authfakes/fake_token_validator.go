// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"net/http"
	"sync"

	"github.com/concourse/skymarshal/auth"
)

type FakeTokenValidator struct {
	IsAuthenticatedStub        func(r *http.Request) bool
	isAuthenticatedMutex       sync.RWMutex
	isAuthenticatedArgsForCall []struct {
		r *http.Request
	}
	isAuthenticatedReturns struct {
		result1 bool
	}
	isAuthenticatedReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenValidator) IsAuthenticated(r *http.Request) bool {
	fake.isAuthenticatedMutex.Lock()
	ret, specificReturn := fake.isAuthenticatedReturnsOnCall[len(fake.isAuthenticatedArgsForCall)]
	fake.isAuthenticatedArgsForCall = append(fake.isAuthenticatedArgsForCall, struct {
		r *http.Request
	}{r})
	fake.recordInvocation("IsAuthenticated", []interface{}{r})
	fake.isAuthenticatedMutex.Unlock()
	if fake.IsAuthenticatedStub != nil {
		return fake.IsAuthenticatedStub(r)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isAuthenticatedReturns.result1
}

func (fake *FakeTokenValidator) IsAuthenticatedCallCount() int {
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	return len(fake.isAuthenticatedArgsForCall)
}

func (fake *FakeTokenValidator) IsAuthenticatedArgsForCall(i int) *http.Request {
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	return fake.isAuthenticatedArgsForCall[i].r
}

func (fake *FakeTokenValidator) IsAuthenticatedReturns(result1 bool) {
	fake.IsAuthenticatedStub = nil
	fake.isAuthenticatedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeTokenValidator) IsAuthenticatedReturnsOnCall(i int, result1 bool) {
	fake.IsAuthenticatedStub = nil
	if fake.isAuthenticatedReturnsOnCall == nil {
		fake.isAuthenticatedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isAuthenticatedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeTokenValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTokenValidator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ auth.TokenValidator = new(FakeTokenValidator)