package worker

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"context"
	"crontab/common"
	"github.com/coreos/etcd/mvcc/mvccpb"
)


// 任务管理器
type JobMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}


var (
	G_jobMgr *JobMgr
)


// 初始化管理器
func InitJobMgr() (err error) {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		lease clientv3.Lease
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints: G_config.EtcdEndpoints, // 集群地址
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond, // 连接超时
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	// 赋值单例
	G_jobMgr = &JobMgr{
		client: client,
		kv: kv,
		lease: lease,
	}
	return
}



//监听任务变化 
func (jobMgr *JobMgr) watchJobs() (err error) {
	var (
		getResp *clientv3.GetResponse
		kvpair *mvccpb.KeyValue
		job *common.Job
	)
	//获取所有任务
	if getResp,err =jobMgr.kv.Get(context.TODO(),common.JOB_SAVE_DIR,clientv3.WithPrefix());err !=nil{
		return
	}
	//当前有那些任务
	for _,kvpair= range getResp.Kvs{

		if job,err= common.UnpackJob(kvpair.Value);err ==nil{
			//TODO是吧这个任务同步给调度写成

		}
	}

	go func() {

	}()

	return
}