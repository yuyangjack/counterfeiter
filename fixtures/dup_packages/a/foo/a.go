package foo // import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages/a/foo"

type S struct{}

//go:generate go run github.com/yuyangjack/counterfeiter . I
type I interface {
	FromA() S
}
