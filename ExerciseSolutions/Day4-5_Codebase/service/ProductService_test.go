package service

//
//import (
//	repoMock "Day4-5_Codebase/db_utils/mock"
//	"Day4-5_Codebase/models"
//	"github.com/gin-gonic/gin"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	_ "github.com/stretchr/testify/assert"
//	_ "github.com/stretchr/testify/mock"
//	"testing"
//	_ "testing"
//)
//
//func TestSaveProduct(t *testing.T) {
//	assertion := assert.New(t)
//
//	ctrl := gomock.NewController(t)
//	repo := repoMock.NewMockInterfaceDB(ctrl)
//	productService := NewProductService(repo)
//
//	data := models.Product{
//		ID:          1,
//		RetailerID:  2,
//		ProductName: "corn",
//		Price:       30,
//		Quantity:    2,
//	}
//
//	expectedModel := models.Product{
//		ID:          data.ID,
//		RetailerID:  data.RetailerID,
//		ProductName: data.ProductName,
//		Price:       data.Price,
//		Quantity:    data.Quantity,
//	}
//
//	repo.EXPECT().DoCreate(&expectedModel).Return().Times(1)
//	repo.EXPECT().IsPresentR("2", gomock.Any()).Return(nil).Times(1)
//
//	c := &gin.Context{}
//	_, err := productService.SaveProduct(c, &data)
//	assertion.Nil(err)
//}
//
//func TestFindMyProduct(t *testing.T) {
//	assertion := assert.New(t)
//
//	ctrl := gomock.NewController(t)
//	repo := repoMock.NewMockInterfaceDB(ctrl)
//	productService := NewProductService(repo)
//
//	data := models.Product{
//		ID:          1,
//		RetailerID:  2,
//		ProductName: "corn",
//		Price:       30,
//		Quantity:    2,
//	}
//
//	repo.EXPECT().IsPresentR(gomock.Any(), gomock.Any()).Return(nil).Times(1)
//
//	c := &gin.Context{}
//	c.Params = gin.Params{
//		{
//			Key:   "id",
//			Value: "1",
//		},
//	}
//	_, err := productService.FindMyProduct(c, &data)
//	assertion.Nil(err)
//}
