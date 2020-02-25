package testpkg

import "sync/atomic"

type mgr struct {
	data int64
}

func NewMgr() *mgr {
	return &mgr{data: 1}
}

func (m *mgr) Add() {
	atomic.AddInt64(&m.data, 1)
}

func (m *mgr) Get() int64 {
	return m.data
}
