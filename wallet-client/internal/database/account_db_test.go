package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/celsopires1999/fc-ms-wallet-client/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	account   *entity.Account
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table accounts (id varchar(255), balance int, created_at date, updated_at date)")
	s.accountDB = NewAccountDB(db)
	s.account, _ = entity.NewAccount("f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d", 20)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	err := s.accountDB.Save(s.account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByIDFound() {
	err := s.accountDB.Save(s.account)
	s.Nil(err)
	accountDB, err := s.accountDB.FindByID(s.account.ID)
	s.Nil(err)
	s.Equal(s.account.ID, accountDB.ID)
	s.Equal(s.account.Balance, accountDB.Balance)
}

func (s *AccountDBTestSuite) TestFindByIDNotFound() {
	accountDB, err := s.accountDB.FindByID(s.account.ID)
	s.Nil(accountDB)
	s.NotNil(err)
	fmt.Println(err.Error())
}
