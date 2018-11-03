package cmd

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

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// cancelmCmd represents the cancelm command
var cancelmCmd = &cobra.Command{
	Use:   "cancelm",
	Short: "取消会议",
	Long: `该指令用于取消某个会议 - 仅发起人可以使用
	
	格式: $cancelm -t [title]`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取参数
		title, _ := cmd.Flags().GetString("title")

		// 执行服务
		success, errorMsg := service.CancelMeeting(title)
		if success {
			fmt.Println("操作成功.")
		} else {
			fmt.Println("操作失败: " + errorMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(cancelmCmd)

	// Here you will define your flags and configuration settings.
	cancelmCmd.Flags().StringP("title", "t", "", "标题")
	cancelmCmd.MarkFlagRequired("title")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
