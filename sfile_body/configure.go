package sfile

import (
	"fmt"
	"strings"
)

func ConfigureSfile(args []string) {
	switch len(args) {
	case 1:
		// configure the remote information
		if strings.ContainsRune(args[0], '@') && strings.Count(args[0], ":") >= 2 && len(args[0]) > 10 {
			//example:jake:jake123@localhost:2999
			addarr := strings.Split(args[0], "@")
			if strings.ContainsRune(addarr[0], ':') && len(addarr[0]) > 3 && strings.ContainsRune(addarr[1], ':') && len(addarr[1]) > 6 {
				originlist := ParseList(siteconf)
				originlist["cloud"] = args[0]
				if !FormatList(originlist, siteconf) {
					fmt.Println("write to site.cnf failed")
				}
			} else {
				fmt.Println("your remote information format is incorrect :-(")
			}
		} else {
			fmt.Println("your remote information format is incorrect :-(")
		}
	case 2:
		switch args[0] {
		case "show":
			switch args[1] {
			case "remoteinfo":
				ShowRemoteInfo()
			}
		}
	}
}

// show remote info
func ShowRemoteInfo() {
	originlist := ParseList(siteconf)
	if _, ok := originlist["cloud"]; !ok {
		fmt.Println("you were not configure remote information")
		return
	}
	if CheckRemoteInfo(originlist["cloud"]) {
		addarr := strings.Split(originlist["cloud"], "@")
		hostaddarr := strings.Split(addarr[len(addarr)-1], ":")
		frontinfo := strings.Join(addarr[:len(addarr)-1], "@")
		userinfo := strings.Split(frontinfo, ":")
		user := userinfo[0]
		authkey := strings.Join(userinfo[1:], ":")
		fmt.Printf("%-20v %v\n%-20v %v\n%-20v %v\n%-20v %v\n", "username", user, "authkey", authkey, "host address", hostaddarr[0], "port", hostaddarr[1])
	} else {
		fmt.Println("your remote information is incorrect,sfile remote format is user:authkey@hostadd .here is an example=> jake:jake123@localhost:2999")
	}
}
