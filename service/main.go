package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/amirrezaask/bank_grpc/api"
	"google.golang.org/grpc"
)

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

func (s *server) NewAccount(context.Context, *api.AccountReq) (*api.Account, error) {
	return &api.Account{}, nil
}
func (s *server) GetCredit(context.Context, *api.ID) (*api.Credit, error) {
	return &api.Credit{}, nil

}
func (s *server) Transfer(context.Context, *api.TransactionReq) (*api.Transaction, error) {
	return &api.Transaction{}, nil

}
