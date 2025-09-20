package dblib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getDBConnection() *sql.DB {
	connStr := "admin:MyPassword@tcp(mysqldb1.cora8c66s2x6.us-east-1.rds.amazonaws.com:3306)/dbname"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic("Problem in getting database connection")
	}
	return db
}

func  Create(emp Emp) error {
	dbConnection := getDBConnection()
	defer dbConnection.Close()
	query := "INSERT INTO EmpTable (empno, ename, salary) VALUES (?, ?, ?)"
	_, err := dbConnection.Exec(query, emp.EmpNo, emp.EName, emp.Salary)
	return err
}


func  Update(emp Emp) error {
	dbConnection := getDBConnection()
	defer dbConnection.Close()
	query := "UPDATE EmpTable set ename=?, salary=? where empno = ?"
	_, err := dbConnection.Exec(query, emp.EName, emp.Salary,  emp.EmpNo)
	return err
}

func  Delete(empno int) error {
	dbConnection := getDBConnection()
	defer dbConnection.Close()
	query := "delete from  EmpTable where empno = ?"
	_, err := dbConnection.Exec(query, empno)
	return err
}


func  List() ([]Emp, error) {
	dbConnection := getDBConnection()
	defer dbConnection.Close()
	query := "SELECT * FROM EmpTable"
	rows, err := dbConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []Emp
	for rows.Next() {
		var emp Emp
		rows.Scan(&emp.EmpNo, &emp.EName, &emp.Salary)
		employees = append(employees, emp)
	}
	return employees, nil
}
