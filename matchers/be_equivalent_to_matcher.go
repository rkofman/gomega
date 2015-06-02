package matchers

import (
	"fmt"
	"github.com/onsi/gomega/format"
  "github.com/onsi/gomega/matchers/support/deep_equivalent"
)

type BeEquivalentToMatcher struct {
	Expected interface{}
}

func (matcher *BeEquivalentToMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil && matcher.Expected == nil {
		return false, fmt.Errorf("Both actual and expected must not be nil.")
	}

	return deep_equivalent.DeepEquivalent(actual, matcher.Expected), nil
}

func (matcher *BeEquivalentToMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be equivalent to", matcher.Expected)
}

func (matcher *BeEquivalentToMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be equivalent to", matcher.Expected)
}
