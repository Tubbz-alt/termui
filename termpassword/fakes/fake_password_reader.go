// Package fakes is was generated by counterfeiter. Although generated
// this file is edited to silence the linter.
package fakes

import (
	"sync"

	"github.com/SUSE/termui/termpassword"
)

// FakeReader needs a comment to stop lint errors (will break on regeneration)
type FakeReader struct {
	PromptForPasswordStub        func(promptText string, args ...interface{}) string
	promptForPasswordMutex       sync.RWMutex
	promptForPasswordArgsForCall []struct {
		promptText string
		args       []interface{}
	}
	promptForPasswordReturns struct {
		result1 string
	}
}

// PromptForPassword needs a comment to stop lint errors (will break regeneration)
func (fake *FakeReader) PromptForPassword(promptText string, args ...interface{}) string {
	fake.promptForPasswordMutex.Lock()
	fake.promptForPasswordArgsForCall = append(fake.promptForPasswordArgsForCall, struct {
		promptText string
		args       []interface{}
	}{promptText, args})
	fake.promptForPasswordMutex.Unlock()
	if fake.PromptForPasswordStub != nil {
		return fake.PromptForPasswordStub(promptText, args...)
	} // Edited to silence linter
	return fake.promptForPasswordReturns.result1
}

// PromptForPasswordCallCount needs a comment to stop lint errors (will break regeneration)
func (fake *FakeReader) PromptForPasswordCallCount() int {
	fake.promptForPasswordMutex.RLock()
	defer fake.promptForPasswordMutex.RUnlock()
	return len(fake.promptForPasswordArgsForCall)
}

// PromptForPasswordArgsForCall needs a comment to stop lint errors (will break regeneration)
func (fake *FakeReader) PromptForPasswordArgsForCall(i int) (string, []interface{}) {
	fake.promptForPasswordMutex.RLock()
	defer fake.promptForPasswordMutex.RUnlock()
	return fake.promptForPasswordArgsForCall[i].promptText, fake.promptForPasswordArgsForCall[i].args
}

// PromptForPasswordReturns needs a comment to stop lint errors (will break regeneration)needs a comment to stop lint errors (will break regeneration)needs a comment to stop lint errors (will break regeneration)needs a comment to stop lint errors (will break regeneration)
func (fake *FakeReader) PromptForPasswordReturns(result1 string) {
	fake.PromptForPasswordStub = nil
	fake.promptForPasswordReturns = struct {
		result1 string
	}{result1}
}

var _ termpassword.Reader = new(FakeReader)
