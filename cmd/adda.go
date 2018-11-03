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

// addaCmd represents the adda command
var addaCmd = &cobra.Command{
	Use:   "adda",
	Short: "增加会议参与者",
	Long: `该指令用于增加会议参与者
	
	格式:$ adda -t [title] -p [participator].
	示例:$ adda -t testMeeting -p testUser`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取指令
		title, _ := cmd.Flags().GetString("title")
		particName, _ := cmd.Flags().GetString("participator")

		// 执行服务
		success, errorMsg := service.AddParticToMeeting(title, particName)
		if success {
			fmt.Println("操作成功.")
		} else {
			fmt.Println("操作失败: " + errorMsg)
		}

	},
}

func init() {
	rootCmd.AddCommand(addaCmd)

	// Here you will define your flags and configuration settings.
	addaCmd.Flags().StringP("title", "t", "", "标题")
	addaCmd.MarkFlagRequired("title")
	addaCmd.Flags().StringP("participator", "p", "", "参与者")
	addaCmd.MarkFlagRequired("participator")
}
