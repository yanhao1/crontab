package worker

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	EtcdEndpoints []string `json:"etcdEndpoints"`
	EtcdDialTimeout int `json:"etcdDialTimeout"`
	MongodbUri string `json:"mongodbUri"`
	MongodbConnectTimeout int `json:"mongodbConnectTimeout"`
	JobLogBatchSize int `json:"jobLogBatchSize"`
	JobLogCommitTimeout int `json"jobLogCommitTimeout"`
}

//设置单利模式
var (
	G_config *Config
)


//加载配置
func InitConfig(filename string)(err error) {
	var (
		content []byte
		conf    Config
	)

	//读取配置文件
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	//做JSON反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	//赋值给单利模式
	G_config= &conf
	return
}