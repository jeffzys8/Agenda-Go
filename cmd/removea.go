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

	"github.com/spf13/cobra"
)

// removeaCmd represents the removea command
var removeaCmd = &cobra.Command{
	Use:   "removea",
	Short: "删除会议参与者",
	Long: `该指令用于删除会议参与者
	
	格式: $removea -t [title] -p [participator].`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname, loginned := opfile.GetCurrentUser()
		if !loginned {
			fmt.Println("未登录")
			return
		}

		title, _ := cmd.Flags().GetString("title")
		meetingInfo, meetingExist := entity.GetMeetingInfo(title)
		if !meetingExist {
			fmt.Println("该会议不存在.")
			opfile.WriteLog("RemoveParticipator: Non-exist meeting. user:" + hostname)
			return
		}
		if !strings.EqualFold(meetingInfo.Host, hostname) {
			fmt.Println("您无该会议的操作权.")
			opfile.WriteLog("RemoveParticipator: No right to add participators. user:" + hostname)
			return
		}

		particName, _ := cmd.Flags().GetString("participator")
		_, parcExist := entity.GetUserInfo(particName)
		if !parcExist {
			fmt.Println("该用户不存在.")
			opfile.WriteLog("RemoveParticipator: Invalid participator. user:" + hostname)
			return
		}

		parIndex, hasPar := entity.UserHasParcMeeting(particName, title)
		if !hasPar {
			fmt.Println("该用户不是会议参与者.")
			opfile.WriteLog("RemoveParticipator: Participator repetition. user:" + hostname)
			return
		}

		entity.RemoveParticFromMeeting(title, parIndex)
		entity.RemovePartMeetingFromUser(particName, title)
	},
}

func init() {
	rootCmd.AddCommand(removeaCmd)

	// Here you will define your flags and configuration settings.
	removeaCmd.Flags().StringP("title", "t", "", "标题")
	removeaCmd.MarkFlagRequired("title")
	removeaCmd.Flags().StringP("participator", "p", "", "参与者")
	removeaCmd.MarkFlagRequired("participator")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
