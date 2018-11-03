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
	"strings"

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

		hostname, loginned := entity.GetCurrentUser()
		if !loginned {
			fmt.Println("未登录")
			return
		}

		title, _ := cmd.Flags().GetString("title")
		meetingInfo, meetingExist := entity.GetMeetingInfo(title)
		if !meetingExist {
			fmt.Println("该会议不存在.")
			return
		}
		if !strings.EqualFold(meetingInfo.Host, hostname) {
			fmt.Println("您无该会议的操作权.")
			return
		}

		particName, _ := cmd.Flags().GetString("participator")
		_, parcExist := entity.GetUserInfo(particName)
		if !parcExist {
			fmt.Println("该用户不存在.")
			return
		}

		if _, hasParc := entity.UserHasParcMeeting(particName, title); hasParc {
			fmt.Println("该用户已是会议成员.")
			return
		}

		if _, overlap := entity.IsTimeOverlapForUser(particName, meetingInfo.StartTime, meetingInfo.EndTime); overlap {
			fmt.Println("该用户时间冲突.")
			return
		}

		fmt.Println("操作成功.")
		entity.WriteLog("AddParticipator: host(" + hostname + ") adds participator (" + particName + ") to meeting [" + title + "]")
		meetingInfo.Partics = append(meetingInfo.Partics, particName)
		entity.AddUserMeetingParc(particName, title)

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
