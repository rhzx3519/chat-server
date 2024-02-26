package domain

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInitUsers(t *testing.T) {
	//persistence.InitMongoDB()
	//defer persistence.PostMongoDB()
	//
	//result, err := persistence.Database().Collection(CollUser).InsertOne(context.TODO(), &User{
	//	No:       "534956de-78af-44b5-bdf1-45aa50948618",
	//	Email:    "lou@gmail.com",
	//	Nickname: "lou",
	//})
	//assert.NilError(t, err)
	//fmt.Println(result)
	fmt.Println(string(bytes.Trim([]byte("\"a\"bc\""), "\"")))
}
