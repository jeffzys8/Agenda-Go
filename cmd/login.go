// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Agenda/entity"
	"Agenda/opfile"
	"fmt"
	"strings"
	"DES"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登陆",
	Long:  `该命令用于登陆`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		_password, _ := cmd.Flags().GetString("password")

		pass, error := DES.TripleDesDecrypt([]byte(_password), []byte("sfe023f_sefiel#fi32lf3e!"))
		if error != nil {
			panic(error)
		}
		des_password := string(pass[:])
		// read the current file
		password, err := DES.TripleDesDecrypt([]byte(des_password), []byte("sfe023f_sefiel#fi32lf3e!"))
		if err != nil{
			panic(err)
		}
		_, exist := opfile.GetCurrentUser()
		if exist {
			fmt.Println("已经登陆，无需重复登陆")
			return
		}
		user, exist := entity.GetUserInfo(username)
		if !exist {
			fmt.Println("账户不存在，请核对")
			opfile.WriteLog("Login: Invalid username: " + username)
			return
		}
		if strings.EqualFold(user.Password, string(password)) == false {
			fmt.Println("密码错误，请核对")
			opfile.WriteLog("Login: Wrong password: " + username)
			return
		}
		fmt.Println("登陆成功!")
		opfile.SetCurrentUser(username)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	loginCmd.Flags().StringP("user", "u", "", "用户名")
	loginCmd.MarkFlagRequired("user")
	loginCmd.Flags().StringP("password", "p", "", "密码")
	loginCmd.MarkFlagRequired("password")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
