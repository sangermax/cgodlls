package device

import "fmt"

type Rsu struct {
}

func (c *Rsu) Init(Ip string, Power int, nChannel int, waittime int) int {
	fmt.Println("天线初始化 :IP", Ip, "功率:", Power)
	return 0
}
