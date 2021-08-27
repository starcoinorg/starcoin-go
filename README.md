# starcoin-go

# Starcoin Go SDK

Starcoin Go SDK implementation.

## 用户指南

### 引入依赖

``` 
import "github.com/starcoinorg/starcoin-go"
```

### 节点配置

在 starcoin [github](https://github.com/starcoinorg/starcoin) 下载 starcoin 预编译好的包。通过如下命令启动 dev 测试网，当然你也可以启动其他网络如 barnard/main 等等。

```
starcoin -n dev --http-apis all console
```

注意需要把 starcoin 命令行替换为实际地址，如果不需要使用交互式命令行工具，则不需要加 console。在看到这些信息之后，则表示 starcoin 节点启动成功。

```
2021-08-20T16:39:17.159471+08:00 INFO - Service starcoin_rpc_server::module::pubsub::PubSubService start.
2021-08-20T16:39:17.159801+08:00 INFO - starcoin_rpc_server::module::pubsub::PubSubService service actor started
2021-08-20T16:39:17.160851+08:00 INFO - Service starcoin_stratum::service::StratumService start.
2021-08-20T16:39:17.161045+08:00 INFO - starcoin_stratum::service::StratumService service actor started
2021-08-20T16:39:17.164330+08:00 INFO - Ipc rpc server start at :"/var/folders/wd/skrxtj1n25qf74v3q2l9k39c0000gn/T/0fbe8e453ab9d82068e68053d5fa252d/dev/starcoin.ipc"
2021-08-20T16:39:17.169487+08:00 INFO - Rpc: http server start at :0.0.0.0:9850
2021-08-20T16:39:17.170683+08:00 INFO - Rpc: tcp server start at: 0.0.0.0:9860
2021-08-20T16:39:17.171802+08:00 INFO - Listening for new connections on 0.0.0.0:9870.
2021-08-20T16:39:17.172077+08:00 INFO - Rpc: websocket server start at: 0.0.0.0:9870
2021-08-20T16:39:17.172190+08:00 INFO - Service starcoin_rpc_server::service::RpcService start.
2021-08-20T16:39:17.172279+08:00 INFO - starcoin_rpc_server::service::RpcService service actor started
2021-08-20T16:39:17.176553+08:00 INFO - ChainWater actor started
2021-08-20T16:39:17.185617+08:00 INFO - Start console, disable stderr output.
```

默认的 http 端口是 9850 ，websocket 端口是 9870。

## Examples

### 签名并发送交易

#### 转账
```
	client := NewStarcoinClient("http://localhost:9850")
	privateKeyString := "7ddee640acc92417aee935daccfa34306b7c2b827a1308711d5b1d9711e1bdac"
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey := types.Ed25519PrivateKey(privateKeyBytes)
	addressArray := toAccountAddress("b75994d55eae88219dc57e7e62a11bc0")

	result, err := client.TransferStc(addressArray, privateKey, toAccountAddress("ab4039861ca47ec349b64ddb862293bf"), serde.Uint128{
		High: 0,
		Low:  100000,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)

```

### 根据地址查询链上最新状态或者资源

#### 查询资源
```
	client := NewStarcoinClient("http://localhost:9850")

	result, err := client.GetResource("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetState("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

```

#### 查询最新状态
```
	client := NewStarcoinClient("http://localhost:9850")

	result, err := client.GetState("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
```

#### 查询txn
```
	client := NewStarcoinClient("http://localhost:9850")
    var result interface{}

	result, err = client.GetTransactionByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionInfoByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

```

### 监听链上Event
#### 监听 txn event
```
	client := NewStarcoinClient("ws://localhost:9870")

	c1, err := client.NewTxnSendRecvEventNotifications("0xb75994d55eae88219dc57e7e62a11bc0")

	if err != nil {
		t.Error(err)
	}

	data1 := <-c1

	fmt.Println(data1)
```

#### 监听 pending transaction
```
	client := NewStarcoinClient("ws://localhost:9870")

	c, err := client.NewPendingTransactionsNotifications()

	if err != nil {
		t.Error(err)
	}

	data := <-c

	fmt.Println(data)
```

## License

starcoin-go is licensed as [Apache 2.0](./LICENSE).
