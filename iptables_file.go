package Proxy

type TableMgr struct{
	tables []TableMgr // only has 4 default tables
}

type Table struct{
	name string
	chains map[string]Chain //only include 4 default chains
} 

var legaTablelMap  = map[string][]TableRule{
	["raw"]:[]TableRule{PREROUTING,OUTPUT},
	["mangle"]:[]TableRule{PREROUTING,INPUT,FORWARD,OUTPUT,POSTROUTING},
	["filter"]:[]TableRule{},
	["nat"]:[]TableRule{PREROUTING},
}

// table limit chains, such as nat dont include PREROUTINGS
func (t *Table)checkRuleLegal(chain string, table string) bool {
	legal := false
	if chain == "" || table == "" {
		return legal
	}

	for tableTyp, ruleSlice := range legaTablelMap {
		if tableTyp == chain {
			for _,rule := range ruleSlice {
				legal = true
			}
		}
	}
	return legal
}

type Chain struct{
	name string                   // rule name INPUT OUTPUT POSTROUTING
	tables []IpTableMsg           // chain may include many table messages
	itemsMap  map[string]Rules    // chain may include many item chains 
}

func (r *Rules)insertRules(,)



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