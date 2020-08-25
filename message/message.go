package message

type Message struct {
	Id      int
	name    string
	address string
	phone   int
}

func New(id, phone int, name, addr string) Message {
	return Message{
		Id:      id,
		name:    name,
		address: addr,
		phone:   phone,
	}
}

type Option func(msg *Message)

var DEFAULT_MESSAGE = Message{Id: 1, name: "-1", address: "-1", phone: 1}

func WithID(id int) Option {
	return func(m *Message) {
		m.Id = id
	}
}

func WithName(name string) Option {
	return func(m *Message) {
		m.name = name
	}
}

func WithAddress(addr string) Option {
	return func(m *Message) {
		m.address = addr
	}
}

func WithPhone(phone int) Option {
	return func(m *Message) {
		m.phone = phone
	}
}

func NewByOption(opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	for _, o := range opts {
		o(&msg)
	}
	return msg
}
