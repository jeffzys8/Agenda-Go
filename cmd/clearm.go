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

// clearmCmd represents the clearm command
var clearmCmd = &cobra.Command{
	Use:   "clearm",
	Short: "清空会议",
	Long: `该指令用于清空用户发起的所有会议 - 用户一定是作为发起人
	
	格式: $clearm.`,
	Run: func(cmd *cobra.Command, args []string) {

		// 无参数

		// 执行服务
		success, errorMsg := service.ClearMeetings()
		if success {
			fmt.Println("操作成功.")
		} else {
			fmt.Println("操作失败: " + errorMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
