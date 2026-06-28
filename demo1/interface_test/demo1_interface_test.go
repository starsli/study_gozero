/*
提示词
为这个接口生成接口测试用例，用例逻辑如下
1.从配置文件 （文件名叫test_env.yaml） 中获取etcd配置
2.创建etcd客户端，从etcd拉取服务器地址
3.创建rpc客户端，根据从etcd拉取的地址访问服务器接口
*/

package interface_test

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"demo1/demo1_pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	etcdclient "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestConfig struct {
	Etcd struct {
		Hosts []string
		Key   string
	}
}

var configFile = flag.String("f", "test_env.yaml", "the test config file")

func TestDemo1Test(t *testing.T) {
	logx.SetLevel(logx.DebugLevel)

	var c TestConfig
	conf.MustLoad(*configFile, &c)
	t.Logf("Loaded config: etcd hosts=%v, key=%s", c.Etcd.Hosts, c.Etcd.Key)

	etcdCli, err := etcdclient.New(etcdclient.Config{
		Endpoints:   c.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		t.Skipf("Skipping test, etcd not available: %v", err)
	}
	defer etcdCli.Close()

	resp, err := etcdCli.Get(context.Background(), c.Etcd.Key, etcdclient.WithPrefix())
	if err != nil {
		t.Skipf("Skipping test, failed to get server address from etcd: %v", err)
	}
	if len(resp.Kvs) == 0 {
		t.Skipf("Skipping test, no server found for prefix %s", c.Etcd.Key)
	}

	var addresses []string
	for _, kv := range resp.Kvs {
		addresses = append(addresses, string(kv.Value))
	}
	t.Logf("Discovered %d server addresses from etcd: %v", len(addresses), addresses)

	selectedIdx := rand.Intn(len(addresses))
	serverAddr := addresses[selectedIdx]
	t.Logf("Selected server address: %s (index %d)", serverAddr, selectedIdx)

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := demo1_pb.NewDemo1Client(conn)

	tests := []struct {
		name       string
		input      string
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "normal case",
			input:      "hello",
			wantOutput: "welcome:hello",
			wantErr:    false,
		},
		{
			name:       "empty input",
			input:      "",
			wantOutput: "welcome:",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &demo1_pb.Demo1Req{InputParams: tt.input}
			resp, err := client.Demo1Test(context.Background(), req)

			if (err != nil) != tt.wantErr {
				t.Errorf("Demo1Test() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp.OutputParams != tt.wantOutput {
				t.Errorf("Demo1Test() output = %v, want %v", resp.OutputParams, tt.wantOutput)
			}
		})
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	fmt.Println("Starting interface tests...")
	ret := m.Run()
	fmt.Println("Interface tests completed")
	if ret != 0 {
		fmt.Printf("Tests failed with exit code: %d\n", ret)
	}
}
