package layout

import "github.com/jroimartin/gocui"

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	g.Cursor = true

	if messages, err := g.SetView("messages", 0, 0, maxX-20, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		messages.Title = " messages: "
		messages.Autoscroll = true
		messages.Wrap = true
	}

	if input, err := g.SetView("input", 0, maxY-5, maxX-20, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		input.Title = " send: "
		input.Autoscroll = false
		input.Wrap = true
		input.Editable = true
	}

	if users, err := g.SetView("users", maxX-20, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		users.Title = " users: "
		users.Autoscroll = false
		users.Wrap = true
	}

	if name, err := g.SetView("username", maxX/2-10, maxY/2-5, maxX/2+10, maxY/2-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView("username")
		name.Title = " name: "
		name.Autoscroll = false
		name.Wrap = true
		name.Editable = true
	}

	if roomId, err := g.SetView("roomId", maxX/2-10, maxY/2-1, maxX/2+10, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//g.SetCurrentView("roomId")
		roomId.Title = "Room Id: "
		roomId.Autoscroll = false
		roomId.Wrap = true
		roomId.Editable = true
	}
	return nil
}