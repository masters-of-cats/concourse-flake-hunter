// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeNotifier struct {
	NotifyStub        func() <-chan struct{}
	notifyMutex       sync.RWMutex
	notifyArgsForCall []struct{}
	notifyReturns     struct {
		result1 <-chan struct{}
	}
	notifyReturnsOnCall map[int]struct {
		result1 <-chan struct{}
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNotifier) Notify() <-chan struct{} {
	fake.notifyMutex.Lock()
	ret, specificReturn := fake.notifyReturnsOnCall[len(fake.notifyArgsForCall)]
	fake.notifyArgsForCall = append(fake.notifyArgsForCall, struct{}{})
	fake.recordInvocation("Notify", []interface{}{})
	fake.notifyMutex.Unlock()
	if fake.NotifyStub != nil {
		return fake.NotifyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.notifyReturns.result1
}

func (fake *FakeNotifier) NotifyCallCount() int {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	return len(fake.notifyArgsForCall)
}

func (fake *FakeNotifier) NotifyReturns(result1 <-chan struct{}) {
	fake.NotifyStub = nil
	fake.notifyReturns = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeNotifier) NotifyReturnsOnCall(i int, result1 <-chan struct{}) {
	fake.NotifyStub = nil
	if fake.notifyReturnsOnCall == nil {
		fake.notifyReturnsOnCall = make(map[int]struct {
			result1 <-chan struct{}
		})
	}
	fake.notifyReturnsOnCall[i] = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *FakeNotifier) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeNotifier) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeNotifier) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifier) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNotifier) recordInvocation(key string, args []interface{}) {
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

var _ db.Notifier = new(FakeNotifier)
