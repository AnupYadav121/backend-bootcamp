package main

import ("fmt"
"errors"
)

type MatrixElements [][]int


type Matrix struct {
   rows int
   columns int
   elements MatrixElements
}


func (m Matrix) getRows() int{
   return m.rows
}

func (m Matrix) getColumns() int{
   return m.columns
}

func (m Matrix) getElement(i,j int) int{
   return m.elements[i][j]
}

func (m Matrix) getUpdatedMatrix(newMatrix MatrixElements) (MatrixElements,error){
   if ( len(newMatrix)!= m.rows || len(newMatrix[0]) != m.columns ){
      return [][]int{},errors.New("The lenth of both matrices is not equal so, can not add")
   }
   
   
   for i:=0;i<m.rows;i++{
       for j:=0 ; j< m.columns; j++{
	      newMatrix[i][j]=newMatrix[i][j]+m.elements[i][j]
	   }
   }
   
   return newMatrix, nil
   
}



func main() {
   elements := [][]int{{1, 2}, {3, 4}}
   matrix1 := Matrix{2,2,elements }
   fmt.Println(matrix1.getRows())
   fmt.Println(matrix1.getColumns())
   fmt.Println(matrix1.getElement(1,1))
   addMatrix := [][]int{{1,2}, {3, 4}}
   fmt.Println(matrix1.getUpdatedMatrix(addMatrix))
}
