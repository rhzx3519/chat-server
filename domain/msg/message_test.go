package msg

import (
	"chat-server/persistence"
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	persistence.InitMongoDB()
	defer persistence.PostMongoDB()

	m := NewMessage([]byte("123"))
	m.SerialNo = 2
	Save(m)
}

func TestList(t *testing.T) {
	persistence.InitMongoDB()
	defer persistence.PostMongoDB()
	fmt.Println(List())
}
