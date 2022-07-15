package service

import (
	repoMock "Day4-5_Codebase/db_utils/mock"
	"Day4-5_Codebase/dto"
	"Day4-5_Codebase/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthCustomer(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repoMock.NewMockInterfaceDB(ctrl)
	orderService := NewCustomerService(repo)

	data := models.Customer{
		ID:       1,
		Name:     "Anup",
		Password: "1234",
	}

	repo.EXPECT().DoCreateC(gomock.Any()).Return(nil).Times(1)

	res, err := orderService.AuthCustomer(&data)
	assertion.Nil(err)
	assertion.Equal(res.Password, data.Password)
	assertion.Equal(res.Name, data.Name)
}

func TestOrderCreation(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repoMock.NewMockInterfaceDB(ctrl)
	orderService := NewCustomerService(repo)

	data := models.Order{
		ID:         1,
		CustomerID: 2,
		ProductID:  3,
		Quantity:   4,
		Status:     "failed",
	}

	repo.EXPECT().IsPresentC(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	repo.EXPECT().IsPresent(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	repo.EXPECT().DoCreateO(gomock.Any()).Return(nil).Times(1)

	res, err := orderService.OrderCreation(&data)
	assertion.Nil(err)
	assertion.Equal(res.ID, data.ID)
	assertion.Equal(res.ProductID, data.ProductID)
	assertion.Equal(res.Quantity, data.Quantity)
	assertion.Equal(res.CustomerID, data.CustomerID)
	assertion.Equal(res.Status, data.Status)
}

func TestMultipleOrderCreation(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repoMock.NewMockInterfaceDB(ctrl)
	orderService := NewCustomerService(repo)

	data := dto.Order{ID: 1, CustomerID: 1}
	data.ProductID = append(data.ProductID, 1)
	data.ProductID = append(data.ProductID, 2)
	data.Quantity = append(data.Quantity, 3)
	data.Quantity = append(data.Quantity, 5)

	repo.EXPECT().IsPresentC(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	repo.EXPECT().IsPresent(gomock.Any(), gomock.Any()).Return(nil).Times(2)
	repo.EXPECT().DoCreateO(gomock.Any()).Return(nil).Times(2)

	err := orderService.MultipleOrderCreation(&data)
	assertion.Nil(err)
}
