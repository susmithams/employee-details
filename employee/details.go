package employee

import (
    "net/http"
	"strconv"
	"time"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"


)
// var employees []EmployeeDetails
type employeeValidator struct {
	validator *validator.Validate
}
func (ev *employeeValidator)Validate(i interface{})error{
	return ev.validator.Struct(i)
}
type EmployeeDetails struct{
	ID int `json:"id"`
	Name string `json:"name" validate:"required"`
	Age int `json:"age" validate:"required"`
	Salary int `json:"salary" validate:"required"`
	Experience int `json:"experience"`
	Created_at time.Time `json:"created_at"`

}
func createEmployees(c echo.Context)error{
	var reqEmployee EmployeeDetails
	e.Validator= (&employeeValidator{validator: v})
	
	if err:=c.Bind(&reqEmployee);err!=nil{
		return c.JSON(http.StatusBadRequest, "Invalid request body")

	}
	if err:=c.Validate(reqEmployee);err!=nil{
		return err
	}
	query:=`INSERT INTO Employees (name,age,salary,experience,created_at)VALUES($1,$2,$3,$4,$5)RETURNING id`
	err:=db.QueryRow(query,reqEmployee.Name,reqEmployee.Age,reqEmployee.Salary,reqEmployee.Experience,time.Now()).Scan(&reqEmployee.ID)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "Failed to insert employee")

	}
	return c.JSON(http.StatusOK, reqEmployee)

// 	reqEmployee.ID=len(employees)+1
// 	reqEmployee.Created_at=time.Now()
// 	employees=append(employees, reqEmployee)
// 	return c.JSON(http.StatusOK,reqEmployee)
 }



func getEmployees(c echo.Context)error{
	var employee EmployeeDetails
	rows, err:=db.Query("SELECT id,name,age,salary,experience,created_at FROM Employees")
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "Failed to fetch employees")
    }
	defer rows.Close()
      employees:=[]EmployeeDetails{}
	for rows.Next(){
		err:=rows.Scan(&employee.ID,&employee.Name,&employee.Age,&employee.Salary,&employee.Experience,&employee.Created_at)
		if err!=nil{
			return c.JSON(http.StatusInternalServerError, "Failed to parse employee data")
        }
		employees=append(employees, employee)

	}
	return c.JSON(http.StatusOK,employees)
}



func getEmployee(c echo.Context)error{
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		return c.JSON(http.StatusBadRequest, "Invalid employee ID")

	}
	var emp EmployeeDetails
	err =db.QueryRow("SELECT id,name,age,salary,experience,created_at FROM Employees WHERE id=$1",id).Scan(&emp.ID,&emp.Name,&emp.Age,&emp.Salary,&emp.Experience,&emp.Created_at)
	
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve employee")

	}
	return c.JSON(http.StatusOK, emp)

}
	// if eid >0 && eid<=len(employees){
	// 	employee:=employees[eid-1]
	// 	return c.JSON(http.StatusOK,employee)
	// }
	


func updateEmployees(c echo.Context)error{
	e.Validator = &employeeValidator{validator: v}

	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		return c.JSON(http.StatusBadRequest, "Invalid employee ID")

	}
	var updateEmploye EmployeeDetails
	if err:=c.Bind(&updateEmploye);err!=nil{
		return c.JSON(http.StatusBadRequest, "Invalid request body")

	}
	if err:=c.Validate(updateEmploye);err!=nil{
		return c.JSON(http.StatusBadRequest, "Validation error")
	}
	query:=`UPDATE Employees SET name=$1,age=$2,salary=$3,experience=$4,created_at=$5 WHERE id=$6 RETURNING id,name,age,salary,experience,created_at;`
	err=db.QueryRow(query,updateEmploye.Name,updateEmploye.Age,updateEmploye.Salary,updateEmploye.Experience,time.Now(),id).Scan(&updateEmploye.ID,&updateEmploye.Name,&updateEmploye.Age,&updateEmploye.Salary,&updateEmploye.Experience,&updateEmploye.Created_at)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, "Failed to update employee")

	}
	return c.JSON(http.StatusOK, updateEmploye)
}
// for i,emp:=range employees{
	// 	if emp.ID == eid{
	// 		updateEmploye.ID=eid
	// 		updateEmploye.Created_at=time.Now()
	// 		employees[i]=updateEmploye
	// 		return c.JSON(http.StatusOK,updateEmploye)
	// 	}

// 	}
// 	return c.JSON(http.StatusNotFound,"Employee not found")
// }



func updateEmployee(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid employee ID")
	}

	var partialUpdate map[string]interface{}

	if err := c.Bind(&partialUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Retrieve the current employee details from the database
	var emp EmployeeDetails
	err = db.QueryRow("SELECT id, name, age, salary, experience, created_at FROM Employees WHERE id = $1", id).Scan(&emp.ID, &emp.Name, &emp.Age, &emp.Salary, &emp.Experience, &emp.Created_at)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve employee")
	}

	if name, ok := partialUpdate["name"].(string); ok && name != "" {
		emp.Name = name
	}

	if age, ok := partialUpdate["age"].(float64); ok {
		emp.Age = int(age)
	}

	if salary, ok := partialUpdate["salary"].(float64); ok {
		emp.Salary = int(salary)
	}

	if experience, ok := partialUpdate["experience"].(float64); ok {
		emp.Experience = int(experience)
	}

	query := `UPDATE Employees SET name=$1, age=$2, salary=$3, experience=$4, created_at=$5 WHERE id=$6`
	_, err = db.Exec(query, emp.Name, emp.Age, emp.Salary, emp.Experience, time.Now(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update employee")
	}

	return c.JSON(http.StatusOK, emp)
}

func deleteEmployee(c echo.Context)error{
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		return c.JSON(http.StatusBadRequest, "Invalid employee ID")
 }
 _,err= db.Query("DELETE FROM Employees WHERE id=$1",id)
 if err != nil {
	return c.JSON(http.StatusInternalServerError, "Failed to delete employee")
}
return c.JSON(http.StatusOK, "Deleted")
}
	//  for i,emp :=range employees{
	// 	if emp.ID==eid{
	// 		employees=append(employees[:i],employees[i+1:]... )
	// 		return c.JSON(http.StatusNoContent,"deleted")
	// 	}

	// }
// 	return c.JSON(http.StatusNotFound,"employee not found")
// }
