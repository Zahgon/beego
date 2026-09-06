package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/beego/beego/v2/core/config"
)

type EtcdConfiger struct {
	prefix string
	client *clientv3.Client
	config.BaseConfiger
}

func newEtcdConfiger(client *clientv3.Client, prefix string) *EtcdConfiger {
	_ = "STUB: not implemented"
	return nil
}

func (e *EtcdConfiger) reader(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *EtcdConfiger) Set(key, val string) error { _ = "STUB: not implemented"; return nil }

func (e *EtcdConfiger) DIY(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EtcdConfiger) GetSection(section string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EtcdConfiger) SaveConfigFile(filename string) error { _ = "STUB: not implemented"; return nil }

func (e *EtcdConfiger) Unmarshaler(prefix string, obj interface{}, opt ...config.DecodeOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *EtcdConfiger) Sub(key string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (e *EtcdConfiger) OnChange(key string, fn func(value string)) {
	_ = "STUB: not implemented"
	return
}

type EtcdConfigerProvider struct{}

func (provider *EtcdConfigerProvider) Parse(key string) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func (provider *EtcdConfigerProvider) ParseData(data []byte) (config.Configer, error) {
	_ = "STUB: not implemented"
	return *new(config.Configer), nil
}

func get(client *clientv3.Client, key string) (*clientv3.GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() {
	config.Register("etcd", &EtcdConfigerProvider{})
}
