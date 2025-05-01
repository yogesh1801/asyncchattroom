package asyncchatroom

import (
	"fmt"
	"strconv"
)

const MAX_MESSAGE_LENGTH int = 512
const MAX_HEADER_LENGTH int = 4
const TOTAL_DATA_LENGTH int = MAX_HEADER_LENGTH + MAX_MESSAGE_LENGTH

type Message struct {
	bodyLength int
	data       [TOTAL_DATA_LENGTH]byte
}

func NewMessage(body string) *Message {
	m := &Message{}
	if len(body) > MAX_MESSAGE_LENGTH {
		body = body[:MAX_MESSAGE_LENGTH]
	}

	m.bodyLength = len(body)
	copy(m.data[MAX_HEADER_LENGTH:], []byte(body))
	m.EncodeHeader()
	return m
}

func (m *Message) EncodeHeader() {
	header := fmt.Sprintf("%04d", m.bodyLength)
	copy(m.data[:MAX_HEADER_LENGTH], []byte(header))
}

func (m *Message) DecodeHeader() bool {
	headerStr := string(m.data[:MAX_HEADER_LENGTH])
	val, err := strconv.Atoi(headerStr)
	if err != nil || val > MAX_MESSAGE_LENGTH {
		m.bodyLength = 0
		return false
	}
	m.bodyLength = val
	return true
}

func (m *Message) GetData() []byte {
	return m.data[:MAX_HEADER_LENGTH+m.bodyLength]
}

func (m *Message) GetBody() string {
	return string(m.data[MAX_HEADER_LENGTH : MAX_HEADER_LENGTH+m.bodyLength])
}

func (m *Message) PrintMessage() {
	fmt.Println("Message received: ", m.GetBody())
}
