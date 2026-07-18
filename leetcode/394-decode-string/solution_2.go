package main

import (
	"strconv"
	"strings"
)

// 有个变量记录当前层 [] 之间的所有内容。
// 遇到 ] 弹一层栈。
// 一层栈存什么？存 [] 之间的内容，包括 k times + k 之前的内容。
// 3[a]：times = 3, content = a
// 3[a2[c]]：times = 3，content = a2[c] = a + times 2 + content 3
// 3[a2[cd7[ff]]]：times = 3，content = a2[cd7[ff]] = a + times 2 + content cd7[ff] = cd + times 7 + content ff
// 一层有：前缀 + 次数 + 内容
// 一路往右遍历：
// 1. 遇到数字记录次数。把 prefix = content。
// 2. 遇到 [ 说明当前层结束，进入下一层。记录 prefix + times
// 3. 遇到其它记录为 content。
// 4. 遇到 ] 就弹出上一个栈（其中保存了上层的数据）：prefix + times * content
func decodeString2(s string) string {
	var timesBuilder strings.Builder
	var contentBuilder strings.Builder
	var frames []frame
	for i := 0; i < len(s); i++ {
		currChar := s[i]
		switch {
		case isNumber2(currChar):
			timesBuilder.WriteByte(currChar)
		case currChar == '[':
			times, _ := strconv.Atoi(timesBuilder.String())
			prefix := contentBuilder.String()
			frames = append(frames, frame{prefix: prefix, times: times})
			contentBuilder.Reset()
			timesBuilder.Reset()
		case currChar == ']':
			frame := frames[len(frames)-1]
			tmp := contentBuilder.String()
			contentBuilder.Reset()
			contentBuilder.WriteString(frame.prefix)
			for j := 0; j < frame.times; j++ {
				contentBuilder.WriteString(tmp)
			}
			frames = frames[:len(frames)-1]
		default:
			contentBuilder.WriteByte(currChar)
		}
	}

	return contentBuilder.String()
}

type frame struct {
	prefix string
	times  int
}

func isNumber2(b byte) bool {
	return b >= '0' && b <= '9'
}
