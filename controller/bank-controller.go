package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/my/main/entity"
	"github.com/my/main/services"
)

type Controller interface {
	InsertTable(tab interface{}, ctx *gin.Context) error
	SelectID(tab interface{}, ctx *gin.Context) (interface{}, error)
	SelectAll(tab interface{}) (interface{}, error)
	DeleteRow(tab interface{}, ctx *gin.Context) (int, string, error)
	UpdateRow(tab interface{}, id int) (int, error)
	Transactions(cts *gin.Context) error
}

type controller struct {
	s services.Service
}

func New(ser services.Service) Controller {
	return &controller{
		s: ser,
	}
}

func (c *controller) InsertTable(tab interface{}, ctx *gin.Context) error {

	err := ctx.ShouldBindJSON(&tab)
	log.Printf("insert")
	err = c.s.InsertTable(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (c *controller) SelectID(tab interface{}, ctx *gin.Context) (interface{}, error) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	tab, err := c.s.SelectID(tab, id)
	if err != nil {
		log.Printf("error while inserting :%v", err)
	}
	return tab, err
}
func (c *controller) SelectAll(tab interface{}) (interface{}, error) {
	tab, err := c.s.SelectAll(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}
func (c *controller) DeleteRow(tab interface{}, ctx *gin.Context) (int, string, error) {
	iid := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(iid)
	res, delerr := c.s.DeleteRow(tab, id)

	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	return res, iid, delerr
}

func (c *controller) UpdateRow(tab interface{}, id int) (int, error) {
	//iid := ctx.Params.ByName("id")
	//id, _ := strconv.Atoi(iid)
	res, upderr := c.s.UpdateRow(tab, id)
	if upderr != nil {
		log.Printf("error %s", upderr)
		return 0, upderr
	}
	return res, nil
}

func (c *controller) Transactions(ctx *gin.Context) error {

	tab := &entity.Transaction{}
	err := ctx.ShouldBindJSON(&tab)
	if err != nil {
		log.Printf("error in json:%v", err)
	}
	log.Println(tab)
	errr := c.s.Transactions(tab)
	if errr != nil {
		log.Printf("error in transactions:%v", errr)
	}
	return errr

}
