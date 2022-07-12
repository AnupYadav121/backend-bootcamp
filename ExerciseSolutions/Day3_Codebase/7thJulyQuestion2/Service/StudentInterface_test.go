package Service

import (
	Utils "7thJulyQuestion2/DB_Utils/mock"
	"7thJulyQuestion2/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"testing"
	_ "testing"
)

var (
	student1 = Models.Student{
		ID:        1,
		FirstName: "Anup",
		LastName:  "Yadav",
		DOB:       "5/1/02",
		Address:   "Home",
	}
	student2 = Models.Student{
		ID:        2,
		FirstName: "Test",
		LastName:  "User",
		DOB:       "15/12/01",
		Address:   "Home",
	}
)

type testsForStudent struct {
	arg      int
	expected *Models.Student
}

func TestFindStudent(t *testing.T) {
	tests := []testsForStudent{
		{
			arg:      int(student1.ID),
			expected: &student1,
		},
		{
			arg:      int(student2.ID),
			expected: &student2,
		},
	}

	// mockgen -destination=Utils/mock/mock.go -package=Utils 7thJulyQuestion2/Utils InterfaceDB

	ctrl := gomock.NewController(t)
	iMock := Utils.NewMockInterfaceDB(ctrl)

	for _, _ = range tests {
		iMock.EXPECT().IsPresent(gomock.Any(), gomock.Any()).Return(&student1, nil).Times(2)
		studentService := NewStudent(iMock)

		c := &gin.Context{}
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: "1",
			},
		}
		cc := &gin.Context{}
		cc.Params = gin.Params{
			{
				Key:   "id",
				Value: "1",
			},
		}
		studentService.FindStudent(c)
		studentService.FindStudent(cc)
	}
}
