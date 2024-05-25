package taosRestful

import (
	"fmt"
	"testing"
)

func Test_newTaosConn(t *testing.T) {
	resp, err := httpClient.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := httpClient.Get("https://www.baidu.com")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func BenchmarkName2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := client.R().Get("https://www.baidu.com")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
