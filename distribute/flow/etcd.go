package flow

import (
	"context"
	"fmt"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"strconv"
)

var GlobalEtcd *Etcd

type Etcd struct {
	client *clientv3.Client
}

func (e *Etcd) GetClient() *clientv3.Client {
	return e.client
}

func NewEtcd(cluster []string) (*Etcd, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cluster,
		DialTimeout: ETCD_RESPONSE_TIMEOUT,
	})
	etcd := Etcd{
		cli,
	}
	return &etcd, err
}

func (e *Etcd) Get(key string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_RESPONSE_TIMEOUT)
	resp, err := e.client.Get(ctx, key)
	cancel()
	return resp, err
}

func (e *Etcd) GetPrefix(topic string, addNode func(key string, value string)) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_RESPONSE_TIMEOUT)
	resp, _ := e.client.Get(ctx, topic, clientv3.WithPrefix())
	cancel()

	for _, ev := range resp.Kvs {
		// 监听到Put事件，添加节点到本地哈希环中 todo
		addNode(string(ev.Key), string(ev.Value))
	}
}

func (e *Etcd) Put(key string, value string) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_RESPONSE_TIMEOUT)
	resp, err := e.client.Put(ctx, key, value)
	cancel()
	return resp, err
}

func (e *Etcd) Delete(key string) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ETCD_RESPONSE_TIMEOUT)
	delResp, err := e.client.Delete(ctx, key)
	cancel()
	return delResp, err
}

func (e *Etcd) Watch(key string, function func(event *clientv3.Event)) {
	rch := e.client.Watch(context.Background(), key)
	for watchResp := range rch {
		for _, event := range watchResp.Events {
			go function(event)
		}
	}
}

func (e *Etcd) RegisterNode(key string, value string) {
	_, err := e.Put(key, value)
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
	}
	return
}

func (e *Etcd) WatchNode(topic string, addNode func(key string, value string)) {
	watcher := clientv3.NewWatcher(e.client)
	watchChan := watcher.Watch(context.TODO(), topic, clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				// 监听到Put事件，添加节点到本地哈希环中 todo
				addNode(string(event.Kv.Key), string(event.Kv.Value))
				fmt.Println("同步节点信息:", string(event.Kv.Key), string(event.Kv.Value))
				InitNewNodeStep(string(event.Kv.Value))
				SyncDataToNode(string(event.Kv.Value))
			}
		}
	}
}

// SyncDataToNode 动态增加节点时同步数据到新节点
func SyncDataToNode(hashString string) {
	// 根据hash值找到原节点
	hash, _ := strconv.Atoi(hashString)
	localIp := GetLocalIp()[0]
	oldNode := GlobalHashRing.GetNextNodeByHashcode(uint32(hash))
	fmt.Println("get oldNode", oldNode, "localIp", localIp)
	//将etcd的获得的string解析为节点的具体数据
	if oldNode.GetIp() != localIp {
		return
	}
	fmt.Println("oldNode", oldNode.GetVid())
	//Distributor.SyncData(oldNode.GetVid(), uint32(hash))
}

func InitNewNodeStep(hashString string) {
	hash, _ := strconv.Atoi(hashString)
	localIp := GetLocalIp()[0]
	newNode := GlobalHashRing.GetNodeByHashcode(uint32(hash))
	fmt.Println("get oldNode", newNode, "localIp", localIp)
	//将etcd的获得的string解析为节点的具体数据
	if newNode.GetIp() != localIp {
		return
	}
	fmt.Println("newNode", newNode.GetVid())
	//DisManager.NewVirtualNode(newNode.GetVid())
}
