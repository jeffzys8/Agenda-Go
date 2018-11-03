package cmd

import (
	"Agenda/service"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// createmCmd represents the createm command
var createmCmd = &cobra.Command{
	Use:   "createm",
	Short: "创建会议",
	Long: `该命令用于创建会议
	
	格式: $ createm -t [title] -s [startTime] -e [endTime] -p [participator]
	其中participator为第一个会议参与者；
	
	示例: $ createm -t exampleMeeting -s '2018-10-31 17:00' -e '2018-10-31 18:00 -p 'testUser'`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取参数
		title, _ := cmd.Flags().GetString("title")
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
		participatorStr, _ := cmd.Flags().GetString("participator")

		// 调用服务
		success, errorMsg := service.CreateMeating(title, startTimeUnix, endTimeUnix, participatorStr)
		if success {
			fmt.Println("操作成功.")
		} else {
			fmt.Println("操作失败: " + errorMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(createmCmd)

	// Here you will define your flags and configuration settings.
	createmCmd.Flags().StringP("title", "t", "", "标题")
	createmCmd.MarkFlagRequired("title")
	createmCmd.Flags().StringP("startTime", "s", "", "起始时间")
	createmCmd.MarkFlagRequired("startTime")
	createmCmd.Flags().StringP("endTime", "e", "", "结束时间")
	createmCmd.MarkFlagRequired("endTime")
	createmCmd.Flags().StringP("participator", "p", "", "首个会议参与者")
	createmCmd.MarkFlagRequired("participator")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
