package userrepository

import "github.com/stretchr/testify/mock"

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) Create() {
	panic("unimplemented")
}

func (u *UserRepositoryMock) GetByEmail() {
	panic("unimplemented")
}

func (u *UserRepositoryMock) UpdatePassword() {
	panic("unimplemented")
}
