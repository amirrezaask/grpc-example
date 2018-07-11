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

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial grpc host:%v", err)
	}
	defer conn.Close()
	c := api.NewBankClient(conn)
	ctx := context.Background()
	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			log.Fatalln("Usage: \n client add name")
		}
		account, err := c.NewAccount(ctx, &api.AccountReq{
			Name:     os.Args[2],
			LastName: "",
			Credit:   10000,
		})
		if err != nil {
			log.Fatalf("could not create new account :%v", err)
		}
		log.Printf("New Account is : %+v", account)
	case "credit":
		if len(os.Args) < 3 {
			log.Fatalln("Usage: \n client credit account_id")
		}
		credit, err := c.GetCredit(ctx, &api.ID{Id: os.Args[2]})
		if err != nil {
			log.Fatalf("could not get credit  :%v", err)
		}
		log.Printf("Credit is :%v", credit.GetCredit())
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
