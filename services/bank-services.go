package services

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/my/main/entity"
)

type Service interface {
	InsertTable(tab interface{}) error
	SelectID(tab interface{}, idt int) (interface{}, error)
	SelectAll(tab interface{}) (interface{}, error)
	DeleteRow(tab interface{}, id int) (int, error)
	UpdateRow(tab interface{}, id int) (int, error)
	Transactions(tab *entity.Transaction) error
}

type service struct {
	db *pg.DB
}

func New() Service {
	opts := &pg.Options{
		User:     "postgres",
		Password: "Pcmwithsc",
		Addr:     "localhost:5432",
		Database: "project",
	}
	db := pg.Connect(opts)

	if db == nil {
		log.Printf("failed  connection")
		os.Exit(100)
	}
	log.Printf("connected")
	entity.CreateTable(db)

	return &service{
		db: db,
	}
}

func (s *service) InsertTable(tab interface{}) error {

	log.Printf("insert")
	err := s.db.Insert(tab)
	if err != nil {
		log.Printf("error while insert in table:%v", err)
	}
	return err
}

func (s *service) SelectID(tab interface{}, idt int) (interface{}, error) {
	err := s.db.Model(tab).Where("id =?", idt).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}

	return tab, err
}
func (s *service) SelectAll(tab interface{}) (interface{}, error) {
	err := s.db.Model(tab).Select(tab)
	if err != nil {
		log.Printf("error while selecting :%v", err)
	}
	return tab, err
}
func (s *service) DeleteRow(tab interface{}, id int) (int, error) {
	res, delerr := s.db.Model(tab).Where("id=?", id).Delete()
	if delerr != nil {
		log.Printf("error %s", delerr)
	}
	r := res.RowsAffected()
	fmt.Println(r)
	return r, delerr
}

func (s *service) UpdateRow(tab interface{}, id int) (int, error) {
	res, upderr := s.db.Model(tab).Where("id=?", id).Update()
	if upderr != nil {
		log.Printf("error, reached service %s", upderr)
	}
	r := res.RowsAffected()
	fmt.Println(r)
	return r, upderr
}

func (s *service) Transactions(tab *entity.Transaction) error {

	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		var balance float64
		err := tx.Model((*entity.Account)(nil)).
			Column("balance").
			Where("id = ?", tab.FromAccountNo).
			Select(&balance)
		log.Println(balance)
		if err != nil {
			log.Printf("couldnt fetch and error %s", err)
			return err
		}
		if balance < tab.Amount {
			err = fmt.Errorf("amount is greater than the balance:%v of account:%v", balance, tab.FromAccountNo)
			return err
		}
		//log.Println(balance)
		balance = balance - tab.Amount
		fmt.Println(balance, tab.Amount)
		_, err = tx.Model((*entity.Account)(nil)).Set("balance = ?", balance).Where("id= ?", tab.FromAccountNo).Update()
		if err != nil {
			return err
		}
		err = tx.Model((*entity.Account)(nil)).Column("balance").Where("id = ?", tab.ToAccountNo).Select(&balance)
		balance = balance + tab.Amount
		_, err = tx.Model((*entity.Account)(nil)).Set("balance = ?", balance).Where("id= ?", tab.ToAccountNo).Update()
		return err
	})
	//fmt.Println(tab)
	if err == nil {
		s.InsertTable(tab)
	}
	return err

}
