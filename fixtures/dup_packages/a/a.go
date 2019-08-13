package a // import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a"

import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a/foo"

//go:generate go run github.com/yuyangjack/counterfeiter . A
type A interface {
	V1() foo.I
}
