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

// deleteaccountCmd represents the deleteaccount command
var deleteaccountCmd = &cobra.Command{
	Use:   "deleteaccount",
	Short: "删除账号",
	Long: `该指令会删除账号
	
	格式：$ deleteaccount.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, loginned := entity.GetCurrentUser()
		if !loginned {
			fmt.Println("未登录")
			return
		}

		userInfo, _ := entity.GetUserInfo(username)
		for _, title := range userInfo.HostMeetings {
			entity.DeleteMeeting(title)
		}

		for _, title := range userInfo.ParMeetings {
			entity.RemoveParticFromMeeting(title, username)
		}

		entity.DeleteUser(username)
		fmt.Println("操作成功.")
		entity.WriteLog("DeleteAccount: user(" + username + ")")
		entity.SetCurrentUser("")
	},
}

func init() {
	rootCmd.AddCommand(deleteaccountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteaccountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteaccountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
