package main

import "fmt"

func main() {
	/*数组的遍历1：for循环根据索引值直接遍历*/
	// skills := [5]string{"计算机组成原理", "操作系统基础", "算法与数据结构", "计算机网络", "内网渗透"}
	// for i := 0; i < len(skills); i++ {
	// 	fmt.Println(skills[i])	//skills[0] skills[1]
	// }

	/*数组的遍历2：for range*/
	// 	skills := [...]string{"操作系统基础", "计算机网络", "代码审计", "WEB安全", "信息收集", "中间件漏洞", "python工具", "内网渗透"}
	// 	for i, v := range skills {
	// 		fmt.Printf("%d %v\n", i, v)
	// 	}
	// }
	/*数组的遍历*/
	skills := [...]string{"ipconfig /all", "systeminfo | findstr \"OS\"", "%PROCESSOR_ARCHITECTURE%", "wmic product get name,version", "wmic service list brief"}
	for i, v := range skills {
		fmt.Printf("%d %v\n", i, v)
	}

}

/*多维数组*/
//[[1,2] [3,4] [5,6]]
// var a11 = [3][2]int
