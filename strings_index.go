package main

import (
	"fmt"
	"strings"
)

func main() {

	detail := `
#### 受众人群包类型

| 值 | 描述 |
| --- | --- |
| RETARGETING\_INCLUDE | 包含人群包 |
| RETARGETING\_EXCLUDE | 排除人群包 |


#### 抖音达人级别

| 值 | 描述 |  |
| --- | --- | --- |
| FIRST\_LEVEL | 一级达人分类 |  |
| SECOND\_LEVEL | 二级达人分类 |  |
| THIRD\_LEVEL | 三级达人分类 |  |
| KEYWORD\_AWEME | 抖音达人 |  |

#### 协议号

| 值 | 描述 |
| --- | --- |
| 5 | 通投智选协议 |`

	tableName := "抖音达人级别"
	tableIndex := fmt.Sprintf("#### %s", tableName)
	pos := strings.Index(detail, tableIndex)
	if pos < 0 {
		fmt.Println("not found", tableName)
		return
	}

	endPos := strings.Index(detail[pos+len(tableIndex):], "####")

	var tableContent string
	if endPos < 0 {
		tableContent = detail[pos+len(tableIndex):]
	} else {
		tableContent = detail[pos+len(tableIndex):pos+len(tableIndex)+endPos]
	}

	fmt.Println(strings.TrimSpace(tableContent))
}
