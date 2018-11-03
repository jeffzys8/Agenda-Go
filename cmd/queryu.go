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
	"fmt"

	"github.com/spf13/cobra"
)

// queryuCmd represents the queryu command
var queryuCmd = &cobra.Command{
	Use:   "queryu",
	Short: "查询用户",
	Long: `该命令用于查询已注册用户的信息
	
	示例 $ queryu -u [username]`,
	Run: func(cmd *cobra.Command, args []string) {
		_, haslogin := entity.GetCurrentUser()
		if !haslogin {
			fmt.Println("未登录.")
			return
		}
		username, _ := cmd.Flags().GetString("user")
		userinfo, exist := entity.GetUserInfo(username)
		if !exist {
			fmt.Println("查无此用户")
			return
		}

		fmt.Println("-------------------------")
		fmt.Println("Name: " + username)
		fmt.Println("Phone: " + userinfo.Phone)
		fmt.Println("Email: " + userinfo.Email)
	},
}

func init() {
	rootCmd.AddCommand(queryuCmd)

	queryuCmd.Flags().StringP("user", "u", "", "用户名")
	queryuCmd.MarkFlagRequired("user")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
