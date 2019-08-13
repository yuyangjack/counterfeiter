package foo // import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/foo"

import (
	"github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a/foo"
	bfoo "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/b/foo"
)

type S struct{}

//go:generate go run github.com/yuyangjack/counterfeiter . MultiAB
type MultiAB interface {
	Mine() S
	foo.I
	bfoo.I
}
