package the_aliased_package // import "github.com/yuyangjack/counterfeiter/fixtures/aliased_package"

//go:generate go run github.com/yuyangjack/counterfeiter . InAliasedPackage
type InAliasedPackage interface {
	Stuff(int) string
}
