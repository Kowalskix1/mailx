package cmd

import (
	"fmt"
	"mailtrap/conf"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "MailTrap <sub-command> <params>",
	Short: "MailTrap is Anti-Spam, Anti-Virus and Anti-Phising mail server.",
	Long:  "MailTrap is tiny and flexible mail serber ,contain Anti-Spam, Anti-Virus and Anti-Phising.",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start MailTrap server",
	Long:  "Start API server and SMTPD server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", conf.BANNER)
		logrus.Infoln("Start to run MailTrap Server.")
		// TODO: 编写启动逻辑
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version info",
	Long:  "Display MailTrap version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", conf.Conf.App.Version)
		fmt.Println("Name   :", "Boxer")
		fmt.Println("Build  :", "2025/05/03")
		fmt.Println("Commit :", "d93o809")
		fmt.Println("Author :", "Kowalskix1")
	},
}
   
func ServerRun() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalln("Failed to start server, error:", err)
	}
}
