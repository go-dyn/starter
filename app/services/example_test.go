package services

import (
	"testing"

	"github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	"github.com/go-dyn/app/models"
)

type ServiceTestSuite struct {
	suite.Suite
	item IExample
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, &ServiceTestSuite{
		item: NewExampleImpl(),
	})
}

func (s *ServiceTestSuite) SetupTest() {

}

func (s *ServiceTestSuite) TestCreate() {
	mockOrm, mockDb, _, _ := mock.Orm()
	mockOrm.On("Query").Return(mockDb).Once()
	model := models.Example{
		Title:   "name",
		Content: "Content",
	}
	mockDb.On("Create", &model).Return(nil).Once()
	user, err := s.item.Create(model)
	s.Nil(err)
	s.Equal("name", user.Title)
	mockOrm.AssertExpectations(s.T())
	mockDb.AssertExpectations(s.T())
}
