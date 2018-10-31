package cmd

import (
	"Agenda/entity"
	"Agenda/opfile"
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

		hostname, loginned := opfile.GetCurrentUser()
		if !loginned {
			fmt.Println("未登录")
			return
		}

		title, _ := cmd.Flags().GetString("title")
		if _, exist := entity.GetMeetingInfo(title); exist {
			fmt.Println("该会议已存在.")
			opfile.WriteLog("CreateMeeting: Same title error.")
			return
		}

		startTimeStr, _ := cmd.Flags().GetString("startTime")
		startTime, err := time.Parse(entity.TimeFormat, startTimeStr)
		if err != nil {
			panic(err)
		}
		startTimeUnix := startTime.Unix()

		endTimeStr, _ := cmd.Flags().GetString("endTime")
		endTime, err := time.Parse(entity.TimeFormat, endTimeStr)
		if err != nil {
			panic(err)
		}
		endTimeUnix := endTime.Unix()

		if endTimeUnix <= startTimeUnix || startTimeUnix < time.Now().Unix() {
			fmt.Println("不合法的时间")
			opfile.WriteLog("CreateMeeting: Illegal time. user:" + hostname)
			return
		}

		// 检查是否和host时间重合
		if meetingName, overlap := entity.IsTimeOverlapForUser(hostname, startTimeUnix, endTimeUnix); overlap {
			fmt.Println("该时间与您的会议[" + meetingName + "]时间冲突")
			opfile.WriteLog("CreateMeeting: time overlap. user:" + hostname)
			return
		}

		// 检查part是否存在
		participatorStr, _ := cmd.Flags().GetString("participator")
		_, parExist := entity.GetUserInfo(participatorStr)
		if !parExist {
			fmt.Println("输入的用户不存在")
			opfile.WriteLog("CreateMeeting: participator not exist. user:" + hostname)
			return
		}
		// 检查是否和part时间重合
		if meetingName, overlap := entity.IsTimeOverlapForUser(participatorStr, startTimeUnix, endTimeUnix); overlap {
			fmt.Println("该时间与参与者(" + participatorStr + ")的会议[" + meetingName + "]时间冲突")
			opfile.WriteLog("CreateMeeting: time overlap for participator. user:" + hostname)
			return
		}

		// 创建会议
		entity.CreateMeeting(title, startTimeUnix, endTimeUnix, hostname, participatorStr)
		entity.AddUserMeetingHost(hostname, title)
		entity.AddUserMeetingParc(participatorStr, title)
		fmt.Println("创建成功!")
		opfile.WriteLog("CreateMeeting: Meeting created [" + title + "] by (" + hostname + ")")
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
