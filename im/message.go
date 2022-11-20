package im

/*
message 用于根据 id 来判断这个消息的类型，相当于做了一层格式封装
*/
import "encoding/json"

type Action int64

const (
	_ Action = iota
	ActionApi
	ActionUserLogin
	ActionUserRegister
	ActionUserGetInfo
	ActionUserEditInfo
	ActionUserLogout

	ActionMessage
	ActionGroupMessage
	ActionChatMessage

	ActionHeartbeat
)

func (a Action) IsApi() bool {
	return a > ActionApi && a < ActionMessage
}

func (a Action) IsMessage() bool {
	return a > ActionMessage && a < ActionHeartbeat
}

func (a Action) IsHeartbeat() bool {
	return a == ActionHeartbeat
}

func (a Action) ActionName() string {
	return ""
}

type Message struct {
	Req    string
	Action Action
	Data   interface{}
}

// DeserializeMessage 将二进制数据流序列化为 Message 的格式
func DeserializeMessage(data []byte) (*Message, error) {
	m := &Message{}
	err := json.Unmarshal(data, m)
	return m, err
}

// Serialize 将一个 Message 类型的数据反序列化为一个具体的二进制数据流
func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}
