package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/my/main/controller"
	"github.com/my/main/entity"
	"github.com/my/main/services"
)

var (
	s   services.Service      = services.New()
	c   controller.Controller = controller.New(s)
	err error
	db  *pg.DB
)

func main() {

	router := gin.Default()

	adminGroup := router.Group("admin")
	{
		adminGroup.GET("/bank", GetBank)
		adminGroup.GET("/bank/:id", GetBankByid)
		adminGroup.POST("/bank/", BankInsertVal)
		adminGroup.DELETE("/bank/:id", DeleteBank)
		adminGroup.PUT("/bank/:id", UpdateBank)

		adminGroup.GET("/branch", GetBranch)
		adminGroup.GET("/branch/:id", GetBranchByid)
		adminGroup.POST("/branch/", BranchInsertVal)
		adminGroup.DELETE("/branch/:id", DeleteBranch)
		//adminGroup.PUT("/branch/:id", UpdateBranch) */
	}

	//adminGroup := router.Group("admin/brnach/:id")
	employeeGroup := router.Group("employee")
	{
		employeeGroup.GET("/account", GetAccount)
		employeeGroup.GET("/account/:id", GetAccountByid)
		employeeGroup.POST("/account/", AccountInsertVal)
		employeeGroup.DELETE("/account/:id", DeleteAccount)
		//employeeGroup.PUT("/account/:id", UpdateAccount)

		employeeGroup.GET("/customer", GetCustomer)
		employeeGroup.GET("/customer/:id", GetCustomerByid)
		employeeGroup.POST("/customer/", CustomerInsertVal)
		employeeGroup.DELETE("/customer/:id", DeleteCustomer)
		//employeeGroup.PUT("/customer/:id", UpdateCustomer)

		employeeGroup.GET("/loan", GetLoan)
		employeeGroup.GET("/loan/:id", GetLoanByid)
		employeeGroup.POST("/loan/", LoanInsertVal)
		employeeGroup.DELETE("/loan/:id", DeleteLoan)
		//employeeGroup.PUT("/loan/:id", UpdateLoan)

		employeeGroup.GET("/transaction", GetTransaction)
		employeeGroup.GET("/transaction/:id", GetTransactionByid)
		employeeGroup.POST("/transaction", Trans)
	}

	router.Run(":8080")
	db.Close()
	/*cust := &entity.Customer{
		ID:      103,
		Name:    "random",
		Address: "somewhere",
		City:    "newyork",
		PhoneNo: "9584965454",
		Email:   "rand@gmail.com",
		Time:    time.Now(),
	}*/
	//services.InsertTable(db, cust)
}
func GetBank(ctx *gin.Context) {
	var tab []entity.Bank
	var res interface{}
	res, err = c.SelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}

func GetBankByid(ctx *gin.Context) {
	var tab entity.Bank
	var res interface{}
	res, err = c.SelectID(&tab, ctx)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			ctx.JSON(200, gin.H{"error": "not found"})
		} else {
			ctx.JSON(200, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(200, res)
	}
}

func BankInsertVal(ctx *gin.Context) {
	var tab entity.Bank
	err = c.InsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteBank(ctx *gin.Context) {
	var tab entity.Bank
	res, id, er := c.DeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateBank(ctx *gin.Context) {

	//updateval := ctx.Query("updval")
	//updateval := ctx.Params.ByName("updval")
	tab := entity.Bank{
		//ID:   id,
		//Name: updateval,
	}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := c.UpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//-----------------------
func GetBranch(ctx *gin.Context) {
	var tab []entity.Branch
	var res interface{}
	res, err = c.SelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}

func GetBranchByid(ctx *gin.Context) {
	var tab entity.Branch
	var res interface{}
	res, err = c.SelectID(&tab, ctx)
	if err.Error() == "pg: no rows in result set" {
		ctx.JSON(200, gin.H{"error": "not found"})
	} else if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

func BranchInsertVal(ctx *gin.Context) {
	var tab entity.Branch
	err = c.InsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteBranch(ctx *gin.Context) {
	var tab entity.Branch
	res, id, er := c.DeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateBranch(ctx *gin.Context) {
	//updateval := ctx.Query("updval")
	//updateval := ctx.Params.ByName("updval")
	tab := entity.Branch{
		//ID:   id,
		//Address: updateval,
	}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := c.UpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//--------------------
func GetAccount(ctx *gin.Context) {
	var tab []entity.Account
	var res interface{}
	res, err = c.SelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}

func GetAccountByid(ctx *gin.Context) {
	var tab entity.Account
	var res interface{}
	res, err = c.SelectID(&tab, ctx)
	if err.Error() == "pg: no rows in result set" {
		ctx.JSON(200, gin.H{"error": "not found"})
	} else if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

func AccountInsertVal(ctx *gin.Context) {
	var tab entity.Account
	err = c.InsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteAccount(ctx *gin.Context) {
	var tab entity.Account
	res, id, er := c.DeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateAccount(ctx *gin.Context) {
	tab := entity.Account{}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := c.UpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//-------------------
func GetCustomer(ctx *gin.Context) {
	var tab []entity.Customer
	var res interface{}
	res, err = c.SelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}

func GetCustomerByid(ctx *gin.Context) {
	var tab entity.Customer
	var res interface{}
	res, err = c.SelectID(&tab, ctx)
	if err.Error() == "pg: no rows in result set" {
		ctx.JSON(200, gin.H{"error": "not found"})
	} else if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

func CustomerInsertVal(ctx *gin.Context) {
	var tab entity.Customer
	err = c.InsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteCustomer(ctx *gin.Context) {
	var tab entity.Customer
	res, id, er := c.DeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateCustomer(ctx *gin.Context) {
	tab := entity.Customer{}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := c.UpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//---------------------
func GetLoan(ctx *gin.Context) {
	var tab []entity.Loan
	var res interface{}
	res, err = c.SelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}

func GetLoanByid(ctx *gin.Context) {
	var tab entity.Loan
	var res interface{}
	res, err = c.SelectID(&tab, ctx)
	if err.Error() == "pg: no rows in result set" {
		ctx.JSON(200, gin.H{"error": "not found"})
	} else if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}

func LoanInsertVal(ctx *gin.Context) {
	var tab entity.Loan
	err = c.InsertTable(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": "inserted successfully"})
	}
}

func DeleteLoan(ctx *gin.Context) {
	var tab entity.Loan
	res, id, er := c.DeleteRow(&tab, ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "deleted"})
	}
}

func UpdateLoan(ctx *gin.Context) {
	tab := entity.Loan{}
	err := ctx.BindJSON(&tab)
	id := strconv.Itoa(tab.ID)
	res, er := c.UpdateRow(&tab, tab.ID)

	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else if res == 0 {
		ctx.JSON(200, gin.H{id: "not found"})
	} else {
		ctx.JSON(200, gin.H{id: "updated"})
	}
}

//-------------------------
func GetTransaction(ctx *gin.Context) {
	var tab []entity.Transaction
	var res interface{}
	res, err = c.SelectAll(&tab)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
		log.Println(res)
		log.Println(tab)
	}
}

func GetTransactionByid(ctx *gin.Context) {
	var tab entity.Transaction
	var res interface{}
	res, err = c.SelectID(&tab, ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, res)
	}
}
func Trans(ctx *gin.Context) {

	er := c.Transactions(ctx)
	if er != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"value": "transacted"})
}

/*
{
    "title" : "hello1",
    "description" : "testing1" ,
    "url" : "https://www.youtube.com/watch?v=qR0WnWL2o1Q&t=32ss" ,
    "author" : {
                    "name" : "Ayushi2",
                    "age" : 23,
                    "email" : "Ayushi2@gmail.com"
                }
}*/
