# Pingq

## 简介
这是个非常简单的ICMP包发送小工具，常用来测网速。
## 实例
```go
package pingq

import (
	"testing"
	"time"
)

func TestPingq(t *testing.T) {

	pingTime, err := Pingq("8.8.8.8", 5*time.Second)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(pingTime)

}
func TestPingqTimeout(t *testing.T) {

	pingTime, err := Pingq("118.8.8.8", 5*time.Second)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(pingTime)

}

```