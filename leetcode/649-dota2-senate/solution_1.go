package main

// 推理：
// 当前轮到谁，谁就禁掉一名对方阵营的议员。
// R 可以禁 D，D 也可以禁 R。
// 被禁的议员以后都会被跳过。
// 扫描到字符串末尾后，从开头继续下一轮。
// 最后哪个阵营还有议员拥有权利，哪个阵营获胜。
//
// 用一个同样长的数组记录每个议员是否被禁。1 被禁，0 还有权。必须记录被票死的，否则下一轮不知道你已经死了，就会重新票你，导致真正要被票的人被挡枪。
// R 永远禁掉最近刚的那个 D。
// 要有个 RD 欠债变量，遍历到 RD 的时候查欠债没有，有的话就 ban，没有的话就记录欠债。
// 要有个变量记录当前字符串中 RD 的个数，当扣到 0，对面就获胜。首先预处理一轮？记录 RD 个数。
// 只要场上还有异营情况存在，就继续轮。每轮每个议员都可以投票。
func predictPartyVictory1(senate string) string {
	radiantCount, direCount := 0, 0
	for _, party := range senate {
		if party == 'R' {
			radiantCount++
		} else {
			direCount++
		}
	}

	bansAgainstRadiant, bansAgainstDire := 0, 0
	bannedSenators := make([]bool, len(senate))

	for radiantCount > 0 && direCount > 0 {
		for index, party := range senate {
			if bannedSenators[index] {
				continue
			}

			if party == 'R' {
				if bansAgainstRadiant > 0 {
					bansAgainstRadiant--
					bannedSenators[index] = true
					radiantCount--
				} else {
					bansAgainstDire++
				}
			} else {
				if bansAgainstDire > 0 {
					bansAgainstDire--
					bannedSenators[index] = true
					direCount--
				} else {
					bansAgainstRadiant++
				}
			}
		}
	}

	if direCount == 0 {
		return "Radiant"
	}

	return "Dire"
}
