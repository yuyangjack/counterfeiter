package dup_packages // import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages"

import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a/foo"

//counterfeiter:generate . DupA
type DupA interface {
	A() foo.S
}
