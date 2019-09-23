package interactor

import (
	"context"
	"fmt"
	"github.com/jroimartin/gocui"
	chat "go_generated"
	"google.golang.org/grpc"
	"log"
	"time"
)

var conn* grpc.ClientConn
const addr = "localhost:3000"
var username string
var roomId string

func SendMessage(g* gocui.Gui, v* gocui.View) error {
	message := v.Buffer()
	clientStub := chat.NewChatServiceClient(conn)
	curTime := time.Now().String()
	req := chat.ChatMessage{
		Username: &username,
		Payload:  &message,
		Time:     &curTime,
		RoomId:   &roomId,
	}
	// ignore resp for now - it just sends back msg status
	_, err := clientStub.SendMessage(context.Background(), &req)
	v.Clear()
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}
}

// in thread
func MultiUpdate(g* gocui.Gui) {
	messagesView, _ := g.View("messages")
	usersView, _ := g.View("users")
	clientStub := chat.NewChatServiceClient(conn)
	for {
		// req
		req := chat.MessageUpdateReq{
			ChatRoomId: &roomId,
			Username:   &username,
		}
		// call rpc
		cli, err := clientStub.UpdateMessage(context.Background(), &req)
		if err != nil {
			log.Fatal(err)
		}
		// clients for fprintf
		var clients string
		// recieve message - in goroutine
		chatMsg, rerr := cli.Recv()
		if rerr != nil {
			log.Fatal(chatMsg)
		}
		// append client list for fmprintf
		for _, user := range chatMsg.RoomEntities {
			clients += *user.Username + "\n"
		}
		cliMessage := *chatMsg.Message.Username + " -> " + *chatMsg.Message.Payload + "\n"
		// update gui
		g.Update(func(gui *gocui.Gui) error {
			usersView.Title = fmt.Sprintf(" %d users: ", len(chatMsg.RoomEntities))
			usersView.Clear()
			fmt.Fprintln(usersView, clients)
			return nil
		})

		g.Update(func(gui *gocui.Gui) error {
			fmt.Fprintln(messagesView, cliMessage)
			return nil
		})
	}
}

func RegisterCli(g *gocui.Gui) {
	var err error
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	
	clientStub := chat.NewChatServiceClient(conn)
	clientStub.RegisterClient(context.Background(),
		&chat.RegisterClient{
			Username: &username,
			RoomId:   &roomId,
		})
	g.SetViewOnTop("messages")
	g.SetViewOnTop("users")
	g.SetViewOnTop("input")
	g.SetCurrentView("input")
	go MultiUpdate(g)
}

func SetRoom(g *gocui.Gui, v *gocui.View) error {
	roomId = v.Buffer()
	RegisterCli(g)
	return nil
}

func SetUser(g *gocui.Gui, v *gocui.View) error {
	username = v.Buffer()
	g.SetCurrentView("roomId")
	return nil
}
