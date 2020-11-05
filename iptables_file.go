package Proxy

type IpTableMsg struct {
	pkts int32 //对应规则匹配到的报文的个数。

	bytes int32 //对应匹配到的报文包的大小总和。

	target string //规则对应的target，往往表示规则对应的"动作"，即规则匹配成功后需要采取的措施。

	prot string //表示规则对应的协议，是否只针对某些协议应用此规则。

	opt string //表示规则对应的选项。

	in string //表示数据包由哪个接口(网卡)流入，我们可以设置通过哪块网卡流入的报文需要匹配当前规则。

	out string //表示数据包由哪个接口(网卡)流出，我们可以设置通过哪块网卡流出的报文需要匹配当前规则。

	source string //表示规则对应的源头地址，可以是一个IP，也可以是一个网段。

	destination string //表示规则对应的目标地址。可以是一个IP，也可以是一个网段。

	extends []string
}