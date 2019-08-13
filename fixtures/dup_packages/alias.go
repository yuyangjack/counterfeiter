package dup_packages // import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages"

import (
	"github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a"
	afoo "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a/foo"
	"github.com/yuyangjack/counterfeiter/fixtures/dup_packages/b/foo"
)

//go:generate go run github.com/yuyangjack/counterfeiter -generate
//counterfeiter:generate . AliasV1
type AliasV1 interface {
	a.A
	afoo.I
	foo.I
}
