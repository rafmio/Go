// https://blog.codecentric.de/gomock-tutorial
// Source file:
// https://github.com/sgreben/testing-with-gomock/blob/master/doer/doer.go
package doer

type Doer interface {
	DoSomething(int, string) error
}
