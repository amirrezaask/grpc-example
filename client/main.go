package main

import (
	"os"
	"strconv"

	"github.com/amirrezaask/grpc-example/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatalln("Usage : \n client command")
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not dial grpc host:%v", err)
	}
	defer conn.Close()
	c := api.NewBankClient(conn)
	ctx := context.Background()
	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			logrus.Fatalln("Usage: \n client add name")
		}
		account, err := c.NewAccount(ctx, &api.AccountReq{
			Name:     os.Args[2],
			LastName: "",
			Credit:   10000,
		})
		if err != nil {
			logrus.Fatalf("could not create new account :%v", err)
		}
		logrus.Printf("New Account is : %+v", account)
	case "credit":
		if len(os.Args) < 3 {
			logrus.Fatalln("Usage: \n client credit account_id")
		}
		credit, err := c.GetCredit(ctx, &api.ID{Id: os.Args[2]})
		if err != nil {
			logrus.Fatalf("could not get credit  :%v", err)
		}
		logrus.Printf("Credit is :%v", credit.GetCredit())
	case "transfer":
		if len(os.Args) < 5 {
			logrus.Fatalln("Usage: \n client transfer sender_id recv_id amount")
		}
		amount, err := strconv.Atoi(os.Args[4])
		if err != nil {
			logrus.Fatalf("amount should be a number :%v", err)
		}
		trans, err := c.Transfer(ctx, &api.TransactionReq{
			SenderID: os.Args[2], RecvID: os.Args[3], Amount: int64(amount),
		})
		if err != nil {
			logrus.Fatalf("could not transfer money:%v", err)
		}
		logrus.Printf("Transfer Completed :%v", trans)
	default:
		logrus.Fatalln("INVALID COMMAND")

	}

}
