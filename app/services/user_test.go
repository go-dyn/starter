package services

import (
	"testing"

	"github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	"github.com/go-dyn/app/models"
)

type UserTestSuite struct {
	suite.Suite
	user IUser
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, &UserTestSuite{
		user: NewUserImpl(),
	})
}

func (s *UserTestSuite) SetupTest() {

}

func (s *UserTestSuite) TestCreate() {
	mockOrm, mockDb, _, _ := mock.Orm()
	mockOrm.On("Query").Return(mockDb).Once()
	mockDb.On("Create", &models.User{
		Name: "name",
	}).Return(nil).Once()
	user, err := s.user.Create()
	s.Nil(err)
	s.Equal("name", user.Name)
	mockOrm.AssertExpectations(s.T())
	mockDb.AssertExpectations(s.T())
}
