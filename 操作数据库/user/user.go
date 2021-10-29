package user

import (
	"database/db"
	"fmt"
)

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

//向mySQL数据库添加User  方法一：先预编译再添加 可以防止SQL注入
func (user *User) Add_1() error {
	//编写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//预编译
	inStmt, err := db.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("预编译出错：%v\n", err)
		return err
	}
	//执行
	res, err := inStmt.Exec(user.UserName, user.Password, user.Email)
	if err != nil {
		fmt.Printf("执行添加出错：%v\n", err)
		return err
	}
	fmt.Println("添加成功：", res)
	return nil
}

//向mySQL数据库添加User 方法二：不用预编译直接添加
func (user *User) Add_2() error {
	//编写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)" //删除或修改只需将sql语句改为删除语句或查询语句即可
	//执行
	res, err := db.Db.Exec(sqlStr, user.UserName, user.Password, user.Email)
	if err != nil {
		fmt.Printf("执行添加出错：%v\n", err)
		return err
	}
	fmt.Println("添加成功：", res)
	return nil
}

//通过id查询用户user
func (user *User) GetUserById() (*User, error) {
	//编写sql语句
	sqlStr := "select id,username,password,email from users where id = ?"
	//执行
	row := db.Db.QueryRow(sqlStr, user.Id)
	//声明变量接收查询结果
	var id int
	var username string
	var password string
	var email string
	//扫描查询结果并赋值给 声明的接收查询结果变量
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	//创建*user实例
	users := &User{id, username, password, email}
	return users, nil
}

//查询所有用户
func (user *User) GetUser() ([]*User, error) {
	//编写sql语句
	sqlStr := "select id,username,password,email from users"
	//执行
	rows, err := db.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建*User切片
	var users []*User
	//声明变量接收查询结果
	var id int
	var username string
	var password string
	var email string
	//将查询到的结果逐个添加到*User切片中
	for rows.Next() { //如果rows里还有内容就添加
		//扫描查询结果并赋值给 声明的接收查询结果变量
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		//创建*user实例
		u := &User{id, username, password, email}
		users = append(users, u) //将*user实例添加到*User切片中
	}
	return users, nil
}
