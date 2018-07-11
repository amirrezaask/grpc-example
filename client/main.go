package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/amirrezaask/bank_grpc/api"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: \n client command")
	}
	ctx := context.Background()
	c := &client{}
	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			log.Fatalln("Usage: \n client add name")
		}
		account, err := c.NewAccount(ctx, &api.AccountReq{})
		if err != nil {
			log.Fatalf("could not create new account :%v", err)
		}
		log.Printf("New Account is : %+v", account)
	case "credit":
		if len(os.Args) < 3 {
			log.Fatalln("Usage: \n client credit account_id")
		}
		credit, err := c.GetCredit(ctx, &api.ID{})
		if err != nil {
			log.Fatalf("could not get credit  :%v", err)
		}
		log.Printf("Credit is :%v", credit)
	case "transfer":
		if len(os.Args) < 5 {
			log.Fatalln("Usage: \n client sender_id recv_id amount")
		}
		trans, err := c.Transfer(ctx, &api.TransactionReq{})
		if err != nil {
			log.Fatalf("could not transfer money:%v", err)
		}
		log.Printf("Transfer Completed :%v", trans)
	default:
		log.Fatalln("INVALID COMMAND")

	}

}

type client struct{}

func (c *client) NewAccount(ctx context.Context, in *api.AccountReq, opts ...grpc.CallOption) (*api.Account, error) {

}
func (c *client) GetCredit(ctx context.Context, in *api.ID, opts ...grpc.CallOption) (*api.Credit, error) {
}
func (c *client) Transfer(ctx context.Context, in *api.TransactionReq, opts ...grpc.CallOption) (*api.Transaction, error) {
}
