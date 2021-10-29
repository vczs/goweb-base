package user

import (
	"fmt"
	"testing"
)

//开始执行测试函数之前 进行的操作
func TestMain(m *testing.M) {
	fmt.Println("即将开始执行测试函数！")
	m.Run() //执行该类里的测试函数Test开头的函数:TestUser
}

//开始执行测试函数
func TestUser(t *testing.T) {
	t.Run("开始执行 添加用户:先预编译再添加！", testAdd_1)     //执行测试函数testAdd_1
	t.Run("开始执行 添加用户:不用预编译直接添加！", testAdd_2)   //执行测试函数testAdd_2
	t.Run("开始执行 通过id查询用户 测试函数！", testQueryRow) //执行测试函数testQueryRow
	t.Run("开始执行 查询所有用户 测试函数！", testQuery)      //执行测试函数testQuery
}

//测试函数 添加用户:先预编译再添加
func testAdd_1(t *testing.T) {
	user1 := &User{UserName: "admin1", Password: "123456", Email: "admin1@golang.com"}
	user1.Add_1()
}

//测试函数 添加用户:不用预编译直接添加
func testAdd_2(t *testing.T) {
	user2 := &User{UserName: "admin2", Password: "654321", Email: "admin2@golang.com"}
	user2.Add_2()
}

//测试函数 通过id查询用户
func testQueryRow(t *testing.T) {
	user := &User{Id: 1}
	users, err := user.GetUserById()
	if err != nil {
		fmt.Println("通过id查询用户出错：", err)
		return
	}
	fmt.Println("查询成功：", users)
}

//测试函数 查询所有用户
func testQuery(t *testing.T) {
	user := &User{}
	users, err := user.GetUser()
	if err != nil {
		fmt.Println("查询所有用户出错：", err)
		return
	}
	fmt.Println("查询成功：")
	for _, v := range users {
		fmt.Printf("%v\n", *v)
	}
}
