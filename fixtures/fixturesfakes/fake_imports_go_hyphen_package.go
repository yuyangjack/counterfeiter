// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/yuyangjack/counterfeiter/fixtures"
	hyphenpackage "github.com/yuyangjack/counterfeiter/fixtures/go-hyphenpackage"
)

type FakeImportsGoHyphenPackage struct {
	UseHyphenTypeStub        func(hyphenpackage.HyphenType)
	useHyphenTypeMutex       sync.RWMutex
	useHyphenTypeArgsForCall []struct {
		arg1 hyphenpackage.HyphenType
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenType(arg1 hyphenpackage.HyphenType) {
	fake.useHyphenTypeMutex.Lock()
	fake.useHyphenTypeArgsForCall = append(fake.useHyphenTypeArgsForCall, struct {
		arg1 hyphenpackage.HyphenType
	}{arg1})
	fake.recordInvocation("UseHyphenType", []interface{}{arg1})
	fake.useHyphenTypeMutex.Unlock()
	if fake.UseHyphenTypeStub != nil {
		fake.UseHyphenTypeStub(arg1)
	}
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenTypeCallCount() int {
	fake.useHyphenTypeMutex.RLock()
	defer fake.useHyphenTypeMutex.RUnlock()
	return len(fake.useHyphenTypeArgsForCall)
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenTypeCalls(stub func(hyphenpackage.HyphenType)) {
	fake.useHyphenTypeMutex.Lock()
	defer fake.useHyphenTypeMutex.Unlock()
	fake.UseHyphenTypeStub = stub
}

func (fake *FakeImportsGoHyphenPackage) UseHyphenTypeArgsForCall(i int) hyphenpackage.HyphenType {
	fake.useHyphenTypeMutex.RLock()
	defer fake.useHyphenTypeMutex.RUnlock()
	argsForCall := fake.useHyphenTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImportsGoHyphenPackage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.useHyphenTypeMutex.RLock()
	defer fake.useHyphenTypeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImportsGoHyphenPackage) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.ImportsGoHyphenPackage = new(FakeImportsGoHyphenPackage)
