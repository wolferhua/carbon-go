package carbon

import (
	"sync"
	"time"
)

type Carbon struct {
	time   time.Time
	o      sync.Once
	inited bool
}

// Now 获取当前时间
func Now() Carbon {
	return Carbon{
		time: time.Now(),
	}
}

// 初始化
func (c Carbon) init() {
	if c.inited {
		return
	}
	c.o.Do(func() {
		c.inited = true
		if c.time.IsZero() {
			c.time = time.Now()
		}
	})
}

// GetYear 获取年份
func (c Carbon) GetYear() int {
	c.init()
	return c.time.Year()
}

// GetMonth 获取月份[1-12]
func (c Carbon) GetMonth() int {
	c.init()
	return int(c.time.Month())
}

// GetDay 获取日期[1-31]
func (c Carbon) GetDay() int {
	c.init()
	return c.time.Day()
}

// GetHour 获取当前时间[0-23]
func (c Carbon) GetHour() int {
	c.init()
	return c.time.Hour()
}

func (c Carbon) GetTimestamp() int64 {
	c.init()
	return c.time.Unix()
}
