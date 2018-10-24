# 笔记
一点小测试
## cobra

- [官方教程](https://github.com/spf13/cobra#getting-started)
- [中文教程](https://www.cnblogs.com/borey/p/5715641.html)
- 安装不成功： [clone from github](https://github.com/golang/text)


## JSON

- [参考教程](https://blog.go-zh.org/json-and-go)

## User信息

- KEY: id int
- username string
- password(md5) string
- email string
- phone string


## Meeting信息

- KEY: title string
- creator int //会议发起人id
- attenders []int //会议参与者id
- start_time Date 
- end_time Date

## Global包（全局包）

> 记录全局数据 - 用户登陆状态等

- log_state 登录状态
- log_user_id 登陆用户id

## 持久化要求

```
持久化要求：
使用 json 存储 User 和 Meeting 实体
当前用户信息存储在 curUser.txt 中
```
- 每次运行相关命令读取整个User/Meeting实体，并在结束时保存
- 学习使用 io/ioutils 进行文件读写
    - [文档](https://go-zh.org/pkg/io/ioutil/)
    - [教程](https://blog.csdn.net/wangshubo1989/article/details/74777112/)

## 目录逻辑 

- cmd：存放命令实现代码
- entity：存放 users 和 meetings 对象读写与处理逻辑
- log.txt：使用```log```包记录命令执行
- curUser.txt: 当前登陆用户的存储（**是否需要加密**）

## 耦合

还是需要具有耦合性的记录，以便于增删的操作
- 在Meeting里记录参与者
- 在User里记录参与的会议


# 命令与参数设计

## help

> 显示帮助信息，系统自带

## login

> 执行登陆功能

参数列表：

- 用户名(-u --user)
- password(-p --password)

功能：
- 若用户名或密码为空，返回提示信息
- 若用户已登陆返回提示信息
- 从```entity/users.txt```文件中读取用户信息，确认登陆并保存状态到```curUser.txt```
- 若用户名或密码错误，返回提示信息

## register

> 执行注册功能

参数列表：
- 用户名(-u --user)
- password(-p --password)

功能：
- 若用户名或密码为空，返回提示信息
- 在```entity/users.txt```中检测用户名是否重复
- 保存用户信息，自动登陆，保存登陆信息到```curUser.txt```

## logout

> 执行登出操作

参数列表：（空）

功能：
- 检测```curUser.txt```的登陆状态，记录保存并返回提示

## createm

> 创建会议

## adda

> 增加会议参与者

## removea

> 删除会议参与者

## querym

> 查询某个时段的会议

## cancelm

> 会议发起者取消某个会议

## exitm

> 会议参与者退出某个会议

## clearm

> 会议发起者清空所有发起的会议
