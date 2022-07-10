package main

import "fmt"


type salary interface {
  getSalary() int
}

type Employee struct {
  department string
  number_time_units int
}



func (emp Employee) getSalary() int {
   BasicPays := make(map[string]int)
   BasicPays["Full time"]=500
   BasicPays["Contractor"]=100
   BasicPays["Freelancer"]=10
   if emp.department == "Full time"{
      return BasicPays["Full time"]*emp.number_time_units
   }else if emp.department == "Contractor"{
      return BasicPays["Contractor"]*emp.number_time_units
   }else{
      return BasicPays["Freelancer"]*emp.number_time_units
   }
}


func main() { 
  var s salary
  s = Employee{"Full time",30}
  fmt.Println("The total salary of particular employee is :", s.getSalary())
  s = Employee{"Contractor",30}
  fmt.Println("The total salary of particular employee is :", s.getSalary())
  s = Employee{"Freelancer",20}
  fmt.Println("The total salary of particular employee is :", s.getSalary())
}
