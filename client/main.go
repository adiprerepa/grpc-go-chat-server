package main

import (
	"github.com/jroimartin/gocui"
	"interactor"

	"layout"
	"log"
)

const (
	addr = "localhost:3000"
)
func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	//conn, gerr := grpc.Dial(addr, grpc.WithInsecure())
	//if gerr != nil {
	//	log.Fatal(gerr)
	//}
	if err != nil {
		log.Fatal(err)
	}
    defer g.Close()
    g.SetManagerFunc(layout.Layout)
	g.SetKeybinding("username", gocui.KeyEnter, gocui.ModNone, interactor.SetUser)
	g.SetKeybinding("roomId", gocui.KeyEnter, gocui.ModNone, interactor.SetRoom)
	g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, interactor.SendMessage)
	_ = g.MainLoop()
}


/*
conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
 */