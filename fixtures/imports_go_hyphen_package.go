package fixtures

import (
	"github.com/yuyangjack/counterfeiter/fixtures/go-hyphenpackage"
)

//counterfeiter:generate . ImportsGoHyphenPackage
type ImportsGoHyphenPackage interface {
	UseHyphenType(hyphenpackage.HyphenType)
}
