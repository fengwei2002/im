package im

/*
用户组 每一个组拥有一个 name 字段代表群组的名字
*/

type Group struct {
	*mutex

	Gid  int64
	Name string

	online []chan *Message
}

func (g *Group) SendMessage(message *Message) {
	defer g.LockUtilReturn()

	for i := range g.online {
		g.online[i] <- message
	}
}

func (g *Group) Subscribe(client *Client) {

}

func (g *Group) Unsubscribe(client *Client) {

}
