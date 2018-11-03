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

// exitmCmd represents the exitm command
var exitmCmd = &cobra.Command{
	Use:   "exitm",
	Short: "退出会议",
	Long: `该指令用于会议参与者退出会议 - 会议发起者应使用cancelm取消会议
	
	格式: $exitm -t [title]`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取参数
		title, _ := cmd.Flags().GetString("title")

		// 调用服务
		success, errMsg := service.ExitMeeting(title)
		if success {
			fmt.Println("操作成功.")
		} else {
			fmt.Println("操作失败: " + errMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(exitmCmd)

	// Here you will define your flags and configuration settings.
	exitmCmd.Flags().StringP("title", "t", "", "标题")
	exitmCmd.MarkFlagRequired("title")
	// Cobra supports Persistent Flags which will work for this command

}
