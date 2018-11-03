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
	"time"

	"github.com/spf13/cobra"
)

// querymCmd represents the querym command
var querymCmd = &cobra.Command{
	Use:   "querym",
	Short: "查询会议",
	Long: `该指令用于查询某个用户某一时间段的全部会议
	
	格式：$querym -s [startTime] -e [endTime]
	示例：$querym -s [2018-10-1 20:00] -e [2018-10-7 20:00]`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取参数
		startTimeStr, _ := cmd.Flags().GetString("startTime")
		startTime, err := time.Parse(service.TimeFormat(), startTimeStr)
		if err != nil {
			panic(err)
		}
		startTimeUnix := startTime.Unix()
		endTimeStr, _ := cmd.Flags().GetString("endTime")
		endTime, err := time.Parse(service.TimeFormat(), endTimeStr)
		if err != nil {
			panic(err)
		}
		endTimeUnix := endTime.Unix()

		//调用服务
		sucess, errMsg, hosts, partics := service.QueryMeeting(startTimeUnix, endTimeUnix)
		if !sucess {
			fmt.Println("操作失败: " + errMsg)
		} else {
			fmt.Println("发起的会议：")
			for _, msg := range hosts {
				fmt.Print(msg)
				fmt.Print("---------------------------------------")
			}
			fmt.Println("参与的会议：")
			for _, msg := range partics {
				fmt.Print(msg)
				fmt.Print("---------------------------------------")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(querymCmd)

	// Here you will define your flags and configuration settings.
	querymCmd.Flags().StringP("startTime", "s", "", "起始时间")
	querymCmd.MarkFlagRequired("startTime")
	querymCmd.Flags().StringP("endTime", "e", "", "结束时间")
	querymCmd.MarkFlagRequired("endTime")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
