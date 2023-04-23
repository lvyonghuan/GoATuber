package service

import (
	"sync"
	"time"
)

type buffer[T any] struct {
	bufCap        int             //缓冲区最大容量
	buf           []T             //缓冲区
	lastFlushTime time.Time       //上次刷新缓冲区的时间
	frequency     time.Duration   //自动刷新的频率
	flushFunc     func(items []T) //刷新缓冲区的回调方法
	async         bool            // 是否异步刷新，默认false
	free          chan struct{}   //用于通知缓冲区已被释放
	lock          sync.Mutex
}

func newBuffer[T any](bufCap int, freq time.Duration, async bool, flushFunc func(items []T)) *buffer[T] {
	if bufCap <= 0 || flushFunc == nil {
		return nil
	}
	b := &buffer[T]{
		bufCap:        bufCap,
		buf:           make([]T, 0, bufCap),
		lastFlushTime: time.Now(),
		frequency:     freq,
		flushFunc:     flushFunc,
		async:         async,
		free:          make(chan struct{}),
	}
	go b.tick()
	return b
}

// 周期性刷新缓冲区
func (b *buffer[T]) tick() {
	t := time.NewTicker(b.frequency)
	defer t.Stop()
	offset := time.Second
	for {
		select {
		case <-b.free:
			return
		case now := <-t.C:
			interval := now.Sub(b.lastFlushTime)
			//判断两次刷新时间间隔是否满足条件
			if b.buf == nil || len(b.buf) == 0 || interval+offset < b.frequency {
				continue
			}
			b.lock.Lock()
			//再次判断是否满足刷新条件
			interval = now.Sub(b.lastFlushTime)
			if b.buf != nil && len(b.buf) != 0 && interval+offset >= b.frequency {
				b.flush()
			}
			b.lock.Unlock()
		}
	}
}

func (b *buffer[T]) Put(item T) {
	b.buf = append(b.buf, item)
	if len(b.buf) < b.bufCap {
		return
	}
	b.lock.Lock()
	defer b.lock.Unlock()
	//再次判断
	if len(b.buf) >= b.bufCap {
		b.flush()
	}
}

func (b *buffer[T]) flush() {
	if b.async {
		b.flushAsync()
	} else {
		b.flushSync()
	}
}

// 同步刷新
func (b *buffer[T]) flushSync() {
	b.flushFunc(b.buf)
	b.buf = make([]T, 0, b.bufCap)
	b.lastFlushTime = time.Now()
}

// 异步刷新
func (b *buffer[T]) flushAsync() {
	go b.flushFunc(b.buf)
	b.buf = make([]T, 0, b.bufCap)
	b.lastFlushTime = time.Now()
}

// MustFlush 强制刷新缓冲区，同步操作
func (b *buffer[T]) MustFlush() {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.buf != nil && len(b.buf) != 0 {
		b.flushSync()
	}
}

// Free 释放缓冲区，被释放后不应该再调用 Put 方法，否则会出现空指针异常
func (b *buffer[T]) Free() {
	close(b.free)
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buf = nil
	b.flushFunc = nil
}
