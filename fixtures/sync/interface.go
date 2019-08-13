package sync // import "github.com/yuyangjack/counterfeiter/fixtures/sync"

//go:generate go run github.com/yuyangjack/counterfeiter . SyncSomething
type SyncSomething interface {
	DoThings(string, uint64) (int, error)
	DoNothing()
	DoASlice([]byte)
	DoAnArray([4]byte)
}
