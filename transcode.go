package main

var CodeList string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func id2url(id int64) string {
	urlByteList := []byte{}
	for id >= 62 {
		urlByteList = append(urlByteList,CodeList[id%62])
		id /= 62
	}
	urlByteList = append(urlByteList,CodeList[id])
	reverse(urlByteList)
	return string(urlByteList)
}

func url2id(url string) int64 {
	n := len(url)
	var id int64
	for i := 0;i < n;i++ {
		if url[i] >= '0' && url[i] <= '9' {
			id = id * 62 + int64(url[i]) - 48
		}else if url[i] >= 'A' && url[i] <= 'Z' {
			id = id * 62 + int64(url[i]) - 55
		}else if url[i] >= 'a' && url[i] <= 'z' {
			id = id * 62 + int64(url[i]) - 61
		}
	}
	return id
}

func reverse(list []byte){
	n := len(list)
	for i := 0;i < n / 2;i++ {
		list[i],list[n-1-i] = list[n-1-i],list[i]
	}
}
