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
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登陆",
	Long:  `该命令用于登陆`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取参数
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		// 调用服务
		success, errorMsg := service.Login(username, password)

		if success {
			fmt.Println("登陆成功.")
		} else {
			fmt.Println("登陆失败: " + errorMsg)
		}
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
