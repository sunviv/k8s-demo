//go:build !k8s

package config

var (
	DB = DBConfig{
		DSN: "root:root@tcp(localhost:30006)/k8s-demo",
	}
	Redis = RedisConfig{
		Addr: "localhost:6379",
	}
)
