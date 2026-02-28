package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func InitNacos() string {

	// 配置nacos服务信息
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "172.17.0.2",
			Port:   8848,
		},
	}

	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",   //当namespace是public时，此处填空字符串
		TimeoutMs:           5000, // 请求 Nacos 的超时时间（毫秒）
		NotLoadCacheAtStart: true, // 启动时是否不读取本地缓存
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug", // 日志级别
	}

	// 创建动态配置客户端的另一种方式
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err) // 如果失败,直接退出
	}

	// 从nacos拉取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "goweb-config.yaml",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		panic(err) // 如果失败,直接退出
	}

	return content
}
