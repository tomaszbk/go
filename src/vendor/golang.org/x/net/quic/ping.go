package quic

import "time"

func (c *Conn) ping(space numberSpace) {
	c.sendMsg(func(now time.Time, c *Conn) {
		c.testSendPing.setUnsent()
		c.testSendPingSpace = space
	})
}
