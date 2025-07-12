package configs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Addr     string
		Password string
		Db       int
	}
	Elasticsearch struct {
		Host string
		Port int
	}
	Alipay struct {
		AppId string
		Key   string
	}
}
type Nacos struct {
	NamespaceId string
	Group       string
	DataId      string
	Port        int
	IpAddr      string
}

var WiseConfig Config
var NacosConfig Nacos

// Init 初始化配置
// 优先从Nacos获取配置，如果失败或不完整则使用本地配置
func Init() {
	// 先设置本地默认配置作为备选

	// 尝试从Nacos获取配置
	if err := initFromNacos(); err != nil {
		log.Printf("从Nacos获取配置失败: %v，继续使用本地配置", err)
		return
	}

	// 检查Nacos配置是否完整
	if WiseConfig.Mysql.Host == "" || WiseConfig.Mysql.Database == "" {
		log.Printf("Nacos配置不完整，Host=%s, Database=%s，使用本地配置",
			WiseConfig.Mysql.Host, WiseConfig.Mysql.Database)
		initLocalConfig()
	} else {
		log.Println("成功从Nacos获取完整配置")
	}
}

// initFromNacos 从Nacos获取配置
func initFromNacos() error {
	v := viper.New()
	v.SetConfigFile("configs/dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	nacosConfig := Nacos{}
	if err := v.Unmarshal(&nacosConfig); err != nil {
		return fmt.Errorf("解析Nacos配置失败: %v", err)
	}
	NacosConfig = nacosConfig

	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConfig.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      NacosConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(NacosConfig.Port),
			Scheme:      "http",
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return fmt.Errorf("创建Nacos客户端失败: %v", err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConfig.DataId,
		Group:  NacosConfig.Group})
	if err != nil {
		return fmt.Errorf("获取Nacos配置失败: %v", err)
	}

	configAdd := Config{}
	if err := json.Unmarshal([]byte(content), &configAdd); err != nil {
		return fmt.Errorf("解析Nacos配置内容失败: %v", err)
	}
	WiseConfig = configAdd
	return nil
}

// initLocalConfig 初始化本地配置
func initLocalConfig() {
	// 设置默认配置
	WiseConfig.Mysql.User = "root"
	WiseConfig.Mysql.Password = "wzydsb"
	WiseConfig.Mysql.Host = "14.103.195.100"
	WiseConfig.Mysql.Port = 3306
	WiseConfig.Mysql.Database = "sxt_his"

	WiseConfig.Redis.Addr = "14.103.149.204:6379"
	WiseConfig.Redis.Password = "f453387f13015d7894fda021ae5fef33"
	WiseConfig.Redis.Db = 0

	log.Println("使用本地默认配置")
	log.Printf("数据库配置: Host=%s, Port=%d, Database=%s, User=%s",
		WiseConfig.Mysql.Host,
		WiseConfig.Mysql.Port,
		WiseConfig.Mysql.Database,
		WiseConfig.Mysql.User)
}
