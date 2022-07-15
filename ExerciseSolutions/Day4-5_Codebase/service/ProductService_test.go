package service

import (
	repoMock "Day4-5_Codebase/db_utils/mock"
	"Day4-5_Codebase/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"testing"
	_ "testing"
)

func TestSaveProduct(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repoMock.NewMockInterfaceDB(ctrl)
	productService := NewProductService(repo)

	data := models.Product{
		ID:          1,
		RetailerID:  2,
		ProductName: "corn",
		Price:       30,
		Quantity:    2,
	}

	expectedModel := models.Product{
		ID:          data.ID,
		RetailerID:  data.RetailerID,
		ProductName: data.ProductName,
		Price:       data.Price,
		Quantity:    data.Quantity,
	}

	repo.EXPECT().DoCreate(&expectedModel).Return(nil).Times(1)
	repo.EXPECT().IsPresentR("2", gomock.Any()).Return(nil).Times(1)

	res, err := productService.SaveProduct(&data)
	assertion.Nil(err)
	assertion.Equal(res.ProductName, data.ProductName)
	assertion.Equal(res.RetailerID, data.RetailerID)
	assertion.Equal(res.Quantity, data.Quantity)
	assertion.Equal(res.Price, data.Price)
	assertion.Equal(res.ID, data.ID)
}

func TestFindMyProduct(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repoMock.NewMockInterfaceDB(ctrl)
	productService := NewProductService(repo)

	data := models.Product{
		ID:          1,
		RetailerID:  2,
		ProductName: "corn",
		Price:       30,
		Quantity:    2,
	}

	repo.EXPECT().IsPresent(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	c := &gin.Context{}
	c.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	res, err := productService.FindMyProduct(c, &data)
	assertion.Nil(err)
	assertion.Equal(res.ProductName, data.ProductName)
	assertion.Equal(res.RetailerID, data.RetailerID)
	assertion.Equal(res.Quantity, data.Quantity)
	assertion.Equal(res.Price, data.Price)
	assertion.Equal(res.ID, data.ID)
}

func TestUpdateProduct(t *testing.T) {
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repoMock.NewMockInterfaceDB(ctrl)
	productService := NewProductService(repo)

	data2 := models.Product{
		ID:          1,
		RetailerID:  2,
		ProductName: "corn",
		Price:       60,
		Quantity:    1,
	}

	repo.EXPECT().IsPresent(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	repo.EXPECT().DoUpdate(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	c := &gin.Context{}
	c.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	res, err := productService.UpdateProduct(c, &data2)
	assertion.Nil(err)
	assertion.Equal(res.ProductName, data2.ProductName)
	assertion.Equal(res.RetailerID, data2.RetailerID)
	assertion.Equal(res.Quantity, data2.Quantity)
	assertion.Equal(res.Price, data2.Price)
	assertion.Equal(res.ID, data2.ID)
}
