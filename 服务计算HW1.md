# 笔记

- [ ] cmd-entity实现接口模式
- [ ] 三层模型 cmd-service-data
- [ ] 归纳 指针 meetings[title] = &MeetingInfo{...}  --> map的value要用指针的原因
- [ ] append ...操作

## Go项目库

- [链接](https://github.com/golang/go/wiki/Projects)

## cobra

- [官方教程](https://github.com/spf13/cobra#getting-started)
- [中文教程](https://www.cnblogs.com/borey/p/5715641.html)
- 安装不成功： [clone from github](https://github.com/golang/text)
- [强制要求参数](https://github.com/spf13/cobra#required-flags)


## JSON

- [参考教程](https://blog.go-zh.org/json-and-go)

## User信息

- KEY: Name string
- Password string
- Email string
- Phone string

数据结构: map, key 为 Name


## Meeting信息

- KEY: title string //会议标题
- creator string    //会议发起人
- partics []string  //会议参与者
- start_time int    //使用UNIX时间戳进行记录
- end_time int      //使用UNIX时间戳进行记录

数据结构：map, key 为 title

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

- username(-u --user)
- password(-p --password)

功能：
- 若用户已登陆返回提示信息
- 从```entity/users.txt```文件中读取用户信息，确认登陆并保存状态到```curUser.txt```
- 若用户名或密码错误，返回提示信息

## register

> 执行注册功能

参数列表：
- username(-u --user)
- password(-p --password)
- phone(-ph --phone)
- email(-e --email)

功能：
- 在```entity/users.txt```中检测用户名是否重复
- 保存用户信息，自动登陆，保存登陆信息到```curUser.txt```

## logout

> 执行登出操作

参数列表：（空）

功能：
- 检测```curUser.txt```的登陆状态，记录保存并返回提示

## queryu

> 查询已注册用户

参数列表

- username (-u --user)

功能：

- 查询已注册的用户
- 未登陆无法使用该功能

## deleteaccount

> 删除用户

参数列表：空

功能：
- 删除现在登陆的账号
- 取消所有host的会议（先删参与者，后删会议）
- 删除所有参与的会议（从会议中删除，并检测会议是否要被删除）

## createm

> 创建会议

参数列表：
- 会议标题(-t --title)
- 开始时间(-s --start)
- 结束时间(-e --end)
- 首个参与者(-p --participator)

功能：
- 以当前用户为发起人创建会议
- 检测当前用户是否登陆
- 检测会议是否存在
- 检测给定时间是否合法
- 检测开始时间和结束时间是否和发起人现有会议重叠
- 检测参与者：同adda第4、5条

## adda

> 增加会议参与者

参数列表：
- 会议标题(-t --title)
- 参与用户名(-u --user)

功能：
- 增加某个会议的参与者
- 检测当前用户是否登陆
- 检测会议是否存在
- 检测操作者是否有权限（是该会议host）
- 检测输入用户是否存在且合法（不是会议内人员 & 不是用户本身）
- 检测输入用户是否会议冲突

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
