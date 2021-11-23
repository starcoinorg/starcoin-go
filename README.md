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

#### 部署合约
```
	client := NewStarcoinClient("http://localhost:9850")
	privateKeyString := "7ddee640acc92417aee935daccfa34306b7c2b827a1308711d5b1d9711e1bdac"
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey := types.Ed25519PrivateKey(privateKeyBytes)

	//code,_ := ioutil.ReadFile("/Users/fanngyuan/Documents/workspace/starcoin_java/src/test/resources/contract/MyCounter.mv") 
	//fmt.Println(code)
	code := []byte{161, 28, 235, 11, 2, 0, 0, 0, 9, 1, 0, 4, 2, 4, 4, 3, 8, 25, 5, 33, 12, 7, 45, 78, 8, 123, 32, 10, 155, 1, 5, 12, 160, 1, 81, 13, 241, 1, 2, 0, 0, 1, 1, 0, 2, 12, 0, 0, 3, 0, 1, 0, 0, 4, 2, 1, 0, 0, 5, 0, 1, 0, 0, 6, 2, 1, 0, 1, 8, 0, 4, 0, 1, 6, 12, 0, 1, 12, 1, 7, 8, 0, 1, 5, 9, 77, 121, 67, 111, 117, 110, 116, 101, 114, 6, 83, 105, 103, 110, 101, 114, 7, 67, 111, 117, 110, 116, 101, 114, 4, 105, 110, 99, 114, 12, 105, 110, 99, 114, 95, 99, 111, 117, 110, 116, 101, 114, 4, 105, 110, 105, 116, 12, 105, 110, 105, 116, 95, 99, 111, 117, 110, 116, 101, 114, 5, 118, 97, 108, 117, 101, 10, 97, 100, 100, 114, 101, 115, 115, 95, 111, 102, 248, 175, 3, 221, 8, 222, 73, 216, 30, 78, 253, 158, 36, 192, 57, 204, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 7, 3, 0, 1, 0, 1, 0, 3, 13, 11, 0, 17, 4, 42, 0, 12, 1, 10, 1, 16, 0, 20, 6, 1, 0, 0, 0, 0, 0, 0, 0, 22, 11, 1, 15, 0, 21, 2, 1, 2, 0, 1, 0, 1, 3, 14, 0, 17, 0, 2, 2, 1, 0, 0, 1, 5, 11, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 18, 0, 45, 0, 2, 3, 2, 0, 0, 1, 3, 14, 0, 17, 2, 2, 0, 0, 0}

	moduleId := types.ModuleId{
		toAccountAddress("b75994d55eae88219dc57e7e62a11bc0"),
		"xxxx",
	}

	scriptFunction := types.ScriptFunction{
		moduleId,
		"init",
		[]types.TypeTag{},
		[][]byte{},
	}
	client.DeployContract(toAccountAddress("b75994d55eae88219dc57e7e62a11bc0"), privateKey, scriptFunction, code)
```

### 根据地址查询链上最新状态或者资源

#### 查询资源
```
	client := NewStarcoinClient("http://localhost:9850")

	result, err := client.ListResource("0xa76b896725a088beafb470fe93251c4d")
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
