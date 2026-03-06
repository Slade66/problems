package problems

import "errors"

// Link 表示一条设备间的点对点连接
type Link struct {
	SourceDevice string
	SourcePort   string
	TargetDevice string
	TargetPort   string
}

// SubnetAllocation 表示为一条连接分配的 /30 子网及其两个可用IP
type SubnetAllocation struct {
	Link     Link
	Subnet   string // "172.168.1.0/30"
	SourceIP string // "172.168.1.1/30"
	TargetIP string // "172.168.1.2/30"
}

// AllocateInterconnectSubnets 从指定地址池中，为每条连接分配一个 /30 子网。
//
// 参数：
//   - pool:  CIDR 格式的地址池，如 "172.168.1.0/24"
//   - links: 需要分配IP的连接列表
//
// 返回：
//   - []SubnetAllocation: 每条连接对应的子网分配结果
//   - error:              若地址池容量不足，返回具体原因
func AllocateInterconnectSubnets(pool string, links []Link) ([]SubnetAllocation, error) {
	panic("not implemented")
	_ = errors.New // 占位，避免 import 报错
	return nil, nil
}
