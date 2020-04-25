package main

import (
	"fmt"
	"github.com/Quantumoffices/ethrpc"
)

func main()  {
	rpc:=ethrpc.NewEthRPC("http://192.168.0.220:8545")
	version,_:=rpc.Web3ClientVersion()
	fmt.Println(version)
	count,_:=rpc.NetPeerCount()
	fmt.Println(count)
	sync,_:=rpc.EthSyncing()
	fmt.Println(sync)
	addr,_:=rpc.EthCoinbase()
	fmt.Println(addr)
}
