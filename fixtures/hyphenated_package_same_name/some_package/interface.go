package some_package // import "github.com/yuyangjack/counterfeiter/fixtures/hyphenated_package_same_name/some_package"

import "github.com/yuyangjack/counterfeiter/fixtures/hyphenated_package_same_name/hyphen-ated/some_package"

//go:generate go run github.com/yuyangjack/counterfeiter . SomeInterface
type SomeInterface interface {
	CreateThing() some_package.Thing
}
