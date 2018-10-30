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

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "注册",
	Long:  `该命令用于注册`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		phone, _ := cmd.Flags().GetString("phone")
		email, _ := cmd.Flags().GetString("email")
		if _, exist := entity.GetUserInfo(username); exist {
			fmt.Println("该用户已注册，不可重复注册")
			opfile.WriteLog("Register: Repetition of registration: " + username)
			return
		}
		entity.CreateUser(username, password, phone, email)
		fmt.Println("成功注册!")
		opfile.WriteLog("Register: New successful registration: " + username)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.
	registerCmd.Flags().StringP("user", "u", "", "用户名")
	registerCmd.MarkFlagRequired("user")
	registerCmd.Flags().StringP("password", "p", "", "密码")
	registerCmd.MarkFlagRequired("password")
	registerCmd.Flags().StringP("phone", "n", "", "手机")
	registerCmd.MarkFlagRequired("phone")
	registerCmd.Flags().StringP("email", "e", "", "邮箱")
	registerCmd.MarkFlagRequired("email")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
