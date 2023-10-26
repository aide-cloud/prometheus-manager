package conf

import "time"

func (c *Bootstrap) GetFlag() string {
	if c == nil {
		return ""
	}
	return c.flag
}

func (c *Bootstrap) GetData() *Data {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *Bootstrap) GetServer() *Server {
	if c == nil {
		return nil
	}
	return c.Server
}

func (c *Bootstrap) GetKafka() *Kafka {
	if c == nil {
		return nil
	}

	return c.Kafka
}

func (c *Bootstrap) GetAlert() *Alert {
	if c == nil {
		return nil
	}
	return c.Alert
}

func (c *Bootstrap) GetPushStrategy() *PushStrategy {
	if c == nil {
		return nil
	}
	return c.PushStrategy
}

func (c *Server) GetHttp() *Http {
	if c.Http == nil {
		return nil
	}
	return c.Http
}

func (c *Server) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *Server) GetMode() string {
	if c == nil {
		return ""
	}
	return c.Mode
}

func (c *Server) GetMetadata() map[string]string {
	if c == nil {
		return nil
	}
	return c.Metadata
}

func (c *Data) GetMysql() *Mysql {
	if c.Mysql == nil {
		c.Mysql = nil
	}
	return c.Mysql
}

func (c *Mysql) GetDSN() string {
	if c == nil {
		return ""
	}
	return c.DSN
}

func (c *Mysql) GetDebug() bool {
	if c == nil {
		return false
	}
	return c.Debug
}

func (c *Http) GetHost() string {
	if c == nil {
		return ""
	}
	return c.Host
}

func (c *Http) GetPort() int {
	if c == nil {
		return 0
	}
	return c.Port
}

func (c *Kafka) GetEnable() bool {
	if c == nil {
		return false
	}
	return c.Enable
}

func (c *Kafka) GetEndpoints() []string {
	if c == nil {
		return nil
	}
	return c.Endpoints
}

func (c *Alert) GetAlertTopic() string {
	if c == nil {
		return ""
	}
	return c.Topic
}

func (c *Alert) GetEnable() bool {
	if c == nil {
		return false
	}
	return c.Enable
}

func (c *PushStrategy) GetEnable() bool {
	if c == nil {
		return false
	}
	return c.Enable
}

func (c *PushStrategy) GetInterval() time.Duration {
	if c == nil || c.Interval < 10 {
		return 10
	}

	return c.Interval
}

func (c *PushStrategy) GetTopic() string {
	if c == nil {
		return ""
	}
	return c.Topic
}
