// https://blog.codecentric.de/gomock-tutorial
// Source:
// https://github.com/sgreben/testing-with-gomock/blob/master/user/user.go
package user

import "github.com/sgreben/testing-with-gomock/tree/master/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello, GoMock")
}
