package dup_packages // import "github.com/yuyangjack/counterfeiter/fixtures/dup_packages"

//counterfeiter:generate . DupAB
type DupAB interface {
	DupA
	DupB
}
