package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dchest/uniuri"

	"golang.org/x/net/context"

	"github.com/amirrezaask/bank_grpc/api"
	"google.golang.org/grpc"
)

var db = make(map[string]*api.Account)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("could not listen on 8080 :%v", err)
	}
	srv := grpc.NewServer()
	api.RegisterBankServer(srv, &server{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve :%v", err)
	}
}

type server struct{}

func (s *server) NewAccount(ctx context.Context, req *api.AccountReq) (*api.Account, error) {
	id := uniuri.NewLen(8)
	acc := &api.Account{Id: id, Name: req.GetName(), LastName: req.GetLastName(), Credit: req.GetCredit()}
	db[id] = acc
	return acc, nil
}
func (s *server) GetCredit(ctx context.Context, id *api.ID) (*api.Credit, error) {
	log.Println(id.GetId())
	log.Println(db[id.GetId()].GetName())
	if db[id.GetId()] == nil {
		return nil, fmt.Errorf("could not find this id :%v", id.GetId())

	}
	return &api.Credit{Credit: db[id.GetId()].GetCredit()}, nil

}
func (s *server) Transfer(ctx context.Context, t *api.TransactionReq) (*api.Transaction, error) {
	if db[t.GetSenderID()].GetCredit() < t.GetAmount() {
		return nil, fmt.Errorf("credit is not enough")
	}

	return &api.Transaction{
		Id:       uniuri.NewLen(8),
		SenderID: t.SenderID,
		RecvID:   t.RecvID,
		Amount:   t.Amount,
	}, nil

}
