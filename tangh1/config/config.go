package config

import "time"

var Config = &Configure{}

type Configure struct {
	last int64
	dura time.Duration
	conf *structa
}

func (c *Configure) Get() *structa {
	t := time.Now().UnixMilli()
	if (t - c.last) >= int64(c.dura) {
		c.last = t
	}
	return c.conf
}
