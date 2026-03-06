package problems

import (
	"fmt"
	"testing"
)

func TestAllocateInterconnectSubnets(t *testing.T) {
	tests := []struct {
		name      string
		pool      string
		links     []Link
		wantCount int
		wantErr   bool
		errMsg    string
		validate  func(t *testing.T, result []SubnetAllocation)
	}{
		// === Happy Path ===
		{
			name: "单条连接-/24池",
			pool: "172.168.1.0/24",
			links: []Link{
				{SourceDevice: "SW-A", SourcePort: "Eth1", TargetDevice: "SW-B", TargetPort: "Eth1"},
			},
			wantCount: 1,
			validate: func(t *testing.T, result []SubnetAllocation) {
				if result[0].Subnet != "172.168.1.0/30" {
					t.Errorf("期望子网 172.168.1.0/30, 实际 %s", result[0].Subnet)
				}
				if result[0].SourceIP != "172.168.1.1/30" {
					t.Errorf("期望源IP 172.168.1.1/30, 实际 %s", result[0].SourceIP)
				}
				if result[0].TargetIP != "172.168.1.2/30" {
					t.Errorf("期望目标IP 172.168.1.2/30, 实际 %s", result[0].TargetIP)
				}
			},
		},
		{
			name: "两条连接-连续分配",
			pool: "172.168.1.0/24",
			links: []Link{
				{SourceDevice: "SW-A", SourcePort: "Eth1", TargetDevice: "SW-B", TargetPort: "Eth1"},
				{SourceDevice: "SW-A", SourcePort: "Eth2", TargetDevice: "SW-B", TargetPort: "Eth2"},
			},
			wantCount: 2,
			validate: func(t *testing.T, result []SubnetAllocation) {
				if result[0].Subnet != "172.168.1.0/30" {
					t.Errorf("第1个子网错误: %s", result[0].Subnet)
				}
				if result[1].Subnet != "172.168.1.4/30" {
					t.Errorf("第2个子网错误: %s", result[1].Subnet)
				}
				if result[1].SourceIP != "172.168.1.5/30" {
					t.Errorf("第2个源IP错误: %s", result[1].SourceIP)
				}
			},
		},
		{
			name:      "10条连接-/24池",
			pool:      "10.0.0.0/24",
			links:     generateLinks(10),
			wantCount: 10,
		},
		{
			name:      "64条连接-/24池满载",
			pool:      "192.168.100.0/24",
			links:     generateLinks(64),
			wantCount: 64,
		},
		{
			name:      "16条连接-/26池满载",
			pool:      "10.10.10.0/26",
			links:     generateLinks(16),
			wantCount: 16,
		},
		{
			name:      "32条连接-/25池满载",
			pool:      "172.16.0.0/25",
			links:     generateLinks(32),
			wantCount: 32,
		},
		{
			name:      "8条连接-/27池满载",
			pool:      "192.168.50.0/27",
			links:     generateLinks(8),
			wantCount: 8,
		},
		{
			name:      "4条连接-/28池满载",
			pool:      "10.1.1.0/28",
			links:     generateLinks(4),
			wantCount: 4,
		},
		{
			name:      "128条连接-/23池",
			pool:      "172.20.0.0/23",
			links:     generateLinks(128),
			wantCount: 128,
		},
		{
			name:      "256条连接-/22池",
			pool:      "10.100.0.0/22",
			links:     generateLinks(256),
			wantCount: 256,
		},

		// === Boundary Cases ===
		{
			name:      "1条连接-最小分配",
			pool:      "192.168.1.0/24",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "63条连接-接近满载",
			pool:      "172.168.1.0/24",
			links:     generateLinks(63),
			wantCount: 63,
		},
		{
			name:      "/28池-3条连接",
			pool:      "10.0.0.0/28",
			links:     generateLinks(3),
			wantCount: 3,
		},
		{
			name:      "/27池-7条连接",
			pool:      "192.168.1.0/27",
			links:     generateLinks(7),
			wantCount: 7,
		},
		{
			name:      "/26池-15条连接",
			pool:      "10.10.0.0/26",
			links:     generateLinks(15),
			wantCount: 15,
		},
		{
			name:      "/25池-31条连接",
			pool:      "172.30.0.0/25",
			links:     generateLinks(31),
			wantCount: 31,
		},
		{
			name:      "/23池-127条连接",
			pool:      "192.168.0.0/23",
			links:     generateLinks(127),
			wantCount: 127,
		},
		{
			name:      "/22池-255条连接",
			pool:      "10.200.0.0/22",
			links:     generateLinks(255),
			wantCount: 255,
		},
		{
			name:      "/21池-512条连接",
			pool:      "172.16.0.0/21",
			links:     generateLinks(512),
			wantCount: 512,
		},
		{
			name:      "/20池-1024条连接",
			pool:      "10.0.0.0/20",
			links:     generateLinks(1024),
			wantCount: 1024,
		},

		// === Edge Cases - 资源耗尽 ===
		{
			name:    "/24池-65条连接超载",
			pool:    "172.168.1.0/24",
			links:   generateLinks(65),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/26池-17条连接超载",
			pool:    "10.0.0.0/26",
			links:   generateLinks(17),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/28池-5条连接超载",
			pool:    "192.168.1.0/28",
			links:   generateLinks(5),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/27池-9条连接超载",
			pool:    "10.10.10.0/27",
			links:   generateLinks(9),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/25池-33条连接超载",
			pool:    "172.20.0.0/25",
			links:   generateLinks(33),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/23池-129条连接超载",
			pool:    "192.168.0.0/23",
			links:   generateLinks(129),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/22池-257条连接超载",
			pool:    "10.100.0.0/22",
			links:   generateLinks(257),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/24池-100条连接大幅超载",
			pool:    "172.168.1.0/24",
			links:   generateLinks(100),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/28池-10条连接严重超载",
			pool:    "10.0.0.0/28",
			links:   generateLinks(10),
			wantErr: true,
			errMsg:  "已耗尽",
		},
		{
			name:    "/26池-20条连接超载",
			pool:    "192.168.50.0/26",
			links:   generateLinks(20),
			wantErr: true,
			errMsg:  "已耗尽",
		},

		// === Edge Cases - 无效输入 ===
		{
			name:    "无效CIDR格式",
			pool:    "192.168.1.0",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "无效IP地址",
			pool:    "999.999.999.999/24",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "前缀长度超出范围-/32",
			pool:    "192.168.1.0/32",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "前缀长度超出范围-/31",
			pool:    "192.168.1.0/31",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "前缀长度过小-/8",
			pool:    "10.0.0.0/8",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "空字符串池",
			pool:    "",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "格式错误-缺少斜杠",
			pool:    "192.168.1.024",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "格式错误-多余字符",
			pool:    "192.168.1.0/24/extra",
			links:   generateLinks(1),
			wantErr: true,
		},

		// === Edge Cases - 空连接列表 ===
		{
			name:      "空连接列表",
			pool:      "192.168.1.0/24",
			links:     []Link{},
			wantCount: 0,
		},

		// === 特殊IP段 ===
		{
			name:      "私有IP-10段",
			pool:      "10.255.255.0/24",
			links:     generateLinks(5),
			wantCount: 5,
		},
		{
			name:      "私有IP-172段",
			pool:      "172.31.255.0/24",
			links:     generateLinks(5),
			wantCount: 5,
		},
		{
			name:      "私有IP-192段",
			pool:      "192.168.255.0/24",
			links:     generateLinks(5),
			wantCount: 5,
		},
		{
			name:      "边界IP-0.0.0.0段",
			pool:      "0.0.0.0/24",
			links:     generateLinks(3),
			wantCount: 3,
		},
		{
			name:      "高位IP段",
			pool:      "223.255.255.0/24",
			links:     generateLinks(3),
			wantCount: 3,
		},

		// === 连接信息完整性 ===
		{
			name: "验证连接信息透传",
			pool: "172.168.1.0/24",
			links: []Link{
				{SourceDevice: "Core-A", SourcePort: "100GE1/1/1", TargetDevice: "Core-B", TargetPort: "100GE1/1/1"},
				{SourceDevice: "Core-A", SourcePort: "100GE1/1/2", TargetDevice: "Core-B", TargetPort: "100GE1/1/2"},
			},
			wantCount: 2,
			validate: func(t *testing.T, result []SubnetAllocation) {
				if result[0].Link.SourceDevice != "Core-A" {
					t.Errorf("源设备信息丢失")
				}
				if result[0].Link.SourcePort != "100GE1/1/1" {
					t.Errorf("源端口信息丢失")
				}
				if result[1].Link.TargetDevice != "Core-B" {
					t.Errorf("目标设备信息丢失")
				}
			},
		},

		// === 子网连续性验证 ===
		{
			name:      "验证子网连续性-5条",
			pool:      "10.0.0.0/24",
			links:     generateLinks(5),
			wantCount: 5,
			validate: func(t *testing.T, result []SubnetAllocation) {
				expected := []string{
					"10.0.0.0/30", "10.0.0.4/30", "10.0.0.8/30",
					"10.0.0.12/30", "10.0.0.16/30",
				}
				for i, exp := range expected {
					if result[i].Subnet != exp {
						t.Errorf("第%d个子网期望 %s, 实际 %s", i+1, exp, result[i].Subnet)
					}
				}
			},
		},
		{
			name:      "验证IP配对-3条",
			pool:      "192.168.1.0/24",
			links:     generateLinks(3),
			wantCount: 3,
			validate: func(t *testing.T, result []SubnetAllocation) {
				pairs := [][2]string{
					{"192.168.1.1/30", "192.168.1.2/30"},
					{"192.168.1.5/30", "192.168.1.6/30"},
					{"192.168.1.9/30", "192.168.1.10/30"},
				}
				for i, pair := range pairs {
					if result[i].SourceIP != pair[0] {
						t.Errorf("第%d对源IP错误: 期望 %s, 实际 %s", i+1, pair[0], result[i].SourceIP)
					}
					if result[i].TargetIP != pair[1] {
						t.Errorf("第%d对目标IP错误: 期望 %s, 实际 %s", i+1, pair[1], result[i].TargetIP)
					}
				}
			},
		},

		// === 大规模测试 ===
		{
			name:      "大规模-500条连接",
			pool:      "10.0.0.0/21",
			links:     generateLinks(500),
			wantCount: 500,
		},
		{
			name:      "大规模-1000条连接",
			pool:      "172.16.0.0/20",
			links:     generateLinks(1000),
			wantCount: 1000,
		},

		// === 边界前缀长度 ===
		{
			name:      "/16池-大容量",
			pool:      "10.0.0.0/16",
			links:     generateLinks(100),
			wantCount: 100,
		},
		{
			name:      "/17池",
			pool:      "172.16.0.0/17",
			links:     generateLinks(50),
			wantCount: 50,
		},
		{
			name:      "/18池",
			pool:      "192.168.0.0/18",
			links:     generateLinks(50),
			wantCount: 50,
		},
		{
			name:      "/19池",
			pool:      "10.10.0.0/19",
			links:     generateLinks(50),
			wantCount: 50,
		},

		// === 特殊设备名和端口名 ===
		{
			name: "特殊字符-设备名",
			pool: "172.168.1.0/24",
			links: []Link{
				{SourceDevice: "SW-A_01", SourcePort: "Eth1", TargetDevice: "SW-B-02", TargetPort: "Eth1"},
			},
			wantCount: 1,
		},
		{
			name: "长设备名",
			pool: "172.168.1.0/24",
			links: []Link{
				{SourceDevice: "CoreSwitch-DataCenter-A-Floor3-Rack12", SourcePort: "Eth1", TargetDevice: "CoreSwitch-DataCenter-B-Floor3-Rack13", TargetPort: "Eth1"},
			},
			wantCount: 1,
		},
		{
			name: "空设备名",
			pool: "172.168.1.0/24",
			links: []Link{
				{SourceDevice: "", SourcePort: "Eth1", TargetDevice: "", TargetPort: "Eth1"},
			},
			wantCount: 1,
		},

		// === 结果顺序验证 ===
		{
			name: "验证结果顺序与输入一致",
			pool: "10.0.0.0/24",
			links: []Link{
				{SourceDevice: "SW-1", SourcePort: "P1", TargetDevice: "SW-2", TargetPort: "P1"},
				{SourceDevice: "SW-3", SourcePort: "P2", TargetDevice: "SW-4", TargetPort: "P2"},
				{SourceDevice: "SW-5", SourcePort: "P3", TargetDevice: "SW-6", TargetPort: "P3"},
			},
			wantCount: 3,
			validate: func(t *testing.T, result []SubnetAllocation) {
				if result[0].Link.SourceDevice != "SW-1" {
					t.Errorf("顺序错误: 第1个应该是 SW-1")
				}
				if result[1].Link.SourceDevice != "SW-3" {
					t.Errorf("顺序错误: 第2个应该是 SW-3")
				}
				if result[2].Link.SourceDevice != "SW-5" {
					t.Errorf("顺序错误: 第3个应该是 SW-5")
				}
			},
		},

		// === /30 边界验证 ===
		{
			name:      "验证/30对齐-起始地址",
			pool:      "172.168.1.0/24",
			links:     generateLinks(1),
			wantCount: 1,
			validate: func(t *testing.T, result []SubnetAllocation) {
				// 第一个子网必须是 .0/30
				if result[0].Subnet != "172.168.1.0/30" {
					t.Errorf("第一个子网应该从 .0 开始")
				}
			},
		},
		{
			name:      "验证/30对齐-第10个",
			pool:      "10.0.0.0/24",
			links:     generateLinks(10),
			wantCount: 10,
			validate: func(t *testing.T, result []SubnetAllocation) {
				// 第10个子网应该是 .36/30 (9*4=36)
				if result[9].Subnet != "10.0.0.36/30" {
					t.Errorf("第10个子网错误: %s", result[9].Subnet)
				}
			},
		},
		{
			name:      "验证/30对齐-最后一个",
			pool:      "192.168.1.0/24",
			links:     generateLinks(64),
			wantCount: 64,
			validate: func(t *testing.T, result []SubnetAllocation) {
				// 第64个子网应该是 .252/30 (63*4=252)
				if result[63].Subnet != "192.168.1.252/30" {
					t.Errorf("最后一个子网错误: %s", result[63].Subnet)
				}
			},
		},

		// === 前缀长度边界测试 ===
		{
			name:    "前缀长度过小-/15",
			pool:    "10.0.0.0/15",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "前缀长度过大-/29",
			pool:    "192.168.1.0/29",
			links:   generateLinks(1),
			wantErr: true,
		},
		{
			name:    "前缀长度过大-/30",
			pool:    "192.168.1.0/30",
			links:   generateLinks(1),
			wantErr: true,
		},

		// === nil 输入测试 ===
		{
			name:      "nil连接列表",
			pool:      "192.168.1.0/24",
			links:     nil,
			wantCount: 0,
		},

		// === 跨越256边界测试 ===
		{
			name:      "跨越256边界-从.252开始",
			pool:      "10.0.0.0/22",
			links:     generateLinks(65),
			wantCount: 65,
			validate: func(t *testing.T, result []SubnetAllocation) {
				// 第64个子网: 10.0.0.252/30
				if result[63].Subnet != "10.0.0.252/30" {
					t.Errorf("第64个子网错误: %s", result[63].Subnet)
				}
				// 第65个子网应该跨越到下一个256段: 10.0.1.0/30
				if result[64].Subnet != "10.0.1.0/30" {
					t.Errorf("第65个子网应该跨越边界: 期望 10.0.1.0/30, 实际 %s", result[64].Subnet)
				}
			},
		},
		{
			name:      "跨越多个256边界",
			pool:      "172.16.0.0/22",
			links:     generateLinks(200),
			wantCount: 200,
			validate: func(t *testing.T, result []SubnetAllocation) {
				// 第200个子网: (199*4=796) -> 172.16.3.28/30
				if result[199].Subnet != "172.16.3.28/30" {
					t.Errorf("第200个子网错误: %s", result[199].Subnet)
				}
			},
		},

		// === 地址池末尾边界测试 ===
		{
			name:      "地址池末尾-刚好用完",
			pool:      "192.168.1.0/26",
			links:     generateLinks(16),
			wantCount: 16,
			validate: func(t *testing.T, result []SubnetAllocation) {
				// 最后一个子网: 192.168.1.60/30 (15*4=60)
				if result[15].Subnet != "192.168.1.60/30" {
					t.Errorf("最后一个子网错误: %s", result[15].Subnet)
				}
			},
		},
		{
			name:    "地址池末尾-超出1个",
			pool:    "192.168.1.0/26",
			links:   generateLinks(17),
			wantErr: true,
			errMsg:  "已耗尽",
		},

		// === 更多前缀长度组合 ===
		{
			name:      "/24池-1条连接",
			pool:      "10.1.1.0/24",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/24池-32条连接",
			pool:      "10.2.2.0/24",
			links:     generateLinks(32),
			wantCount: 32,
		},
		{
			name:      "/24池-48条连接",
			pool:      "10.3.3.0/24",
			links:     generateLinks(48),
			wantCount: 48,
		},
		{
			name:      "/27池-1条连接",
			pool:      "192.168.10.0/27",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/27池-4条连接",
			pool:      "192.168.11.0/27",
			links:     generateLinks(4),
			wantCount: 4,
		},
		{
			name:      "/26池-1条连接",
			pool:      "172.20.20.0/26",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/26池-8条连接",
			pool:      "172.20.21.0/26",
			links:     generateLinks(8),
			wantCount: 8,
		},
		{
			name:      "/25池-1条连接",
			pool:      "10.50.50.0/25",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/25池-16条连接",
			pool:      "10.50.51.0/25",
			links:     generateLinks(16),
			wantCount: 16,
		},
		{
			name:      "/23池-1条连接",
			pool:      "192.168.100.0/23",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/23池-64条连接",
			pool:      "192.168.102.0/23",
			links:     generateLinks(64),
			wantCount: 64,
		},
		{
			name:      "/22池-1条连接",
			pool:      "10.200.0.0/22",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/22池-128条连接",
			pool:      "10.201.0.0/22",
			links:     generateLinks(128),
			wantCount: 128,
		},
		{
			name:      "/21池-1条连接",
			pool:      "172.30.0.0/21",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/21池-256条连接",
			pool:      "172.31.0.0/21",
			links:     generateLinks(256),
			wantCount: 256,
		},
		{
			name:      "/20池-1条连接",
			pool:      "10.100.0.0/20",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/20池-512条连接",
			pool:      "10.101.0.0/20",
			links:     generateLinks(512),
			wantCount: 512,
		},
		{
			name:      "/19池-1条连接",
			pool:      "192.168.0.0/19",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/19池-100条连接",
			pool:      "192.168.32.0/19",
			links:     generateLinks(100),
			wantCount: 100,
		},
		{
			name:      "/18池-1条连接",
			pool:      "10.64.0.0/18",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/18池-200条连接",
			pool:      "10.128.0.0/18",
			links:     generateLinks(200),
			wantCount: 200,
		},
		{
			name:      "/17池-1条连接",
			pool:      "172.128.0.0/17",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/17池-300条连接",
			pool:      "172.64.0.0/17",
			links:     generateLinks(300),
			wantCount: 300,
		},
		{
			name:      "/16池-1条连接",
			pool:      "10.0.0.0/16",
			links:     generateLinks(1),
			wantCount: 1,
		},
		{
			name:      "/16池-500条连接",
			pool:      "192.168.0.0/16",
			links:     generateLinks(500),
			wantCount: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AllocateInterconnectSubnets(tt.pool, tt.links)

			if tt.wantErr {
				if err == nil {
					t.Errorf("期望返回错误，但成功了")
				} else if tt.errMsg != "" && !contains(err.Error(), tt.errMsg) {
					t.Errorf("错误信息不匹配，期望包含 %q, 实际 %q", tt.errMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("不期望错误，但返回了: %v", err)
			}

			if len(result) != tt.wantCount {
				t.Errorf("结果数量错误: 期望 %d, 实际 %d", tt.wantCount, len(result))
			}

			if tt.validate != nil {
				tt.validate(t, result)
			}
		})
	}
}

// generateLinks 生成指定数量的测试连接
func generateLinks(n int) []Link {
	links := make([]Link, n)
	for i := 0; i < n; i++ {
		links[i] = Link{
			SourceDevice: fmt.Sprintf("Device-A-%d", i+1),
			SourcePort:   fmt.Sprintf("Eth%d", i+1),
			TargetDevice: fmt.Sprintf("Device-B-%d", i+1),
			TargetPort:   fmt.Sprintf("Eth%d", i+1),
		}
	}
	return links
}

// contains 检查字符串是否包含子串
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
