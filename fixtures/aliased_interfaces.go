package fixtures

import alias "github.com/yuyangjack/counterfeiter/fixtures/another_package"

//go:generate go run github.com/yuyangjack/counterfeiter . AliasedInterface
//go:generate go run github.com/yuyangjack/counterfeiter -generate
//counterfeiter:generate . AliasedInterface

// AliasedInterface is an interface that embeds an interface in an aliased package.
type AliasedInterface interface {
	alias.AnotherInterface
}
