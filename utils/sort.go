package utils

import "main.go/model/common/response"

// 用户聊天列表排序
func UpSort(ucl []*response.ResUserChartList) {
	for k, v := range ucl {
		if v.AdminUnRead != 0 {
			for i := 0; i < len(ucl); i++ {
				if v.AdminUnRead > ucl[i].AdminUnRead {
					if k < i {
						ucl = append(ucl[:k], append(ucl[k+1:i], append([]*response.ResUserChartList{v}, ucl[i:]...)...)...)
					} else {
						ucl = append(ucl[:i], append([]*response.ResUserChartList{v}, append(ucl[i:k], ucl[k+1:]...)...)...)
					}
					break
				}
			}
		}
	}
}
