//go:build k8s

package config

var (
	DB = DBConfig{
		DSN: "root:root@tcp(k8s-demo-mysql:13306)/k8s-demo",
	}
	Redis = RedisConfig{
		Addr: "k8s-demo-redis:6379",
	}
)
