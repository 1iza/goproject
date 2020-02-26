package main

import (
	"context"
	"fmt"
	"github.com/1iza/syncAdapter"
	"time"
)

var (
	ADAPTER = syncAdapter.NewAdapter()
	CHAN    = make(chan request)
)

type request struct {
	srvname string
	data    string
	seqid   uint64
}

type response struct {
	srvname string
	data    string
	seqid   uint64
}

//模拟异步发送
func AsyncCall(req *request) {
	CHAN <- *req
}

//模拟异步响应
func AsyncRespHandler() {
	for req := range CHAN {
		if c, ok := ADAPTER.Get(req.srvname, req.seqid); !ok {
			continue
		} else {
			//模拟处理耗时
			time.Sleep(5 * time.Second)
			c <- []byte("ok")
		}
	}
}

//把异步调用转换成同步调用
func AsyncToSync(req *request) (resp *response, err error) {
	resp = &response{}
	userctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	promise := ADAPTER.New(req.srvname, req.seqid)
	//异步调用的方法，结果从别处返回
	go AsyncCall(req)

	defer ADAPTER.Delete(req.srvname, req.seqid)
	for {
		select {
		case <-userctx.Done():
			err = fmt.Errorf("time out")
			resp = nil
			return
		case d := <-promise:
			err = nil
			resp.data = string(d)
			return
		}
	}
}

func main() {
	go AsyncRespHandler()
	req := &request{
		srvname: "myapp",
		data:    "test",
		seqid:   1,
	}
	resp, err := AsyncToSync(req)
	if err != nil {
		fmt.Printf("err:%v \n", err)
		return
	}
	fmt.Printf("call success resp:%v \n", resp)
}
