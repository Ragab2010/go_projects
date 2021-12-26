package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	fristName  string
	secondName string
	birthday   string
}

func main() {
	fName := inputData("please Enter you First Name  : ")
	sName := inputData("please Enter you Second Name : ")
	bDay := inputData("please Enter you  Birth day  : ")

	var userReturn *User
	userReturn = InitUserReturn(fName, sName, bDay)
	fmt.Println(*userReturn)

	var userAssign *User
	userAssign.InitUserAssign(fName, sName, bDay)
	fmt.Println(*userReturn)

	var userReturnPointer User
	userReturnPointer = *InitUserByReturnPointer(fName, sName, bDay)
	fmt.Println(userReturnPointer)
	/*اما لغه فيها العبر  صحيح */

}

func inputData(prompt string) (output string) {
	fmt.Print(prompt)
	output, _ = reader.ReadString('\n')
	output = strings.Replace(output, "\n", "", -1)
	return
}

func (user *User) InitUserAssign(fName string, sName string, bDay string) {
	user = &User{fristName: fName, secondName: sName, birthday: bDay}
}

func InitUserReturn(fName string, sName string, bDay string) *User {
	user := User{fristName: fName, secondName: sName, birthday: bDay}
	return &user
}

func InitUserByReturnPointer(fName string, sName string, bDay string) *User {
	user := User{fristName: fName, secondName: sName, birthday: bDay}
	return &user
}
