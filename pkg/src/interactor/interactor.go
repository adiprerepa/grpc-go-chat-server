package interactor

import (
	"context"
	"go_generated"
	"github.com/jroimartin/gocui"
	"google.golang.org/grpc"
	"log"
)

var conn* grpc.ClientConn

func Connect(g* gocui.Gui, addr string) {
	var err error
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// register client
	clientStub := chat.NewChatServiceClient(conn)
	regRes, _ := clientStub.RegisterClient(context.Background(),
		&chat.RegisterClient{
			Username:  nil,
			RoomId:    nil,
		}, )

	_, _ = g.SetViewOnTop("messages")
	_, _ = g.SetViewOnTop("users")
	_, _ = g.SetViewOnTop("input")
	_, _ = g.SetCurrentView("input")

	messagesView, _ := g.View("messages")
	usersView, _ := g.View("users")

	go func() {
		for {
			// get data

		}
	}()
}

func RegisterCli() {

}