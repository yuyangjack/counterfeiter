package another_package // import "github.com/yuyangjack/counterfeiter/fixtures/another_package"

type SomeType int

//go:generate go run github.com/yuyangjack/counterfeiter . AnotherInterface
type AnotherInterface interface {
	AnotherMethod([]SomeType, map[SomeType]SomeType, *SomeType, SomeType, chan SomeType)
}
