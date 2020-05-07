package ethrpc

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

const PWD = "luce!1989@9922sdf" //密钥
//const  TAddress  = "0xa376Bd485433699D3D2c1A5322931879503330da"// 归集地址
//const TADDR  = "0x031CdC6C680Dd7889605f1d9EAE15D22797d0b3E" //归集地址二
const TADDR = "0x59f8b414805E0Bb0246A0b0CaF72a889cfb92a16"     //归集地址三
const GASADDRES = "0x7260c1661793170694344bc813be6857ed16e58c" //手续费地址
const Fromadd = "0x7f5CeDBb4CAC1e31dE6Aa8F02b6Bf882a04Fca35"   //转账地址

//收款地址：0x59f8b414805E0Bb0246A0b0CaF72a889cfb92a16
//USDTCollect()//归集
//version, err := client.Web3ClientVersion()
//if err != nil {
//	log.Println("ERR", err.Error())
//}
//list:=listAccount()  //账户列表
//var sumUSDT=0.00
//var count=0
//
//for i:=0;i< len(newlist); i++ {
//
//	sumUSDT+=newlist[i].USDT
//	if newlist[i].ETH>0 && newlist[i].USDT>0{
//		fmt.Println(i," 地址：",newlist[i].Address,"	ETH 余额：",newlist[i].ETH, "  USDT 余额：",newlist[i].USDT)
//		count++
//	}
//
//}

//USDT 归集
func USDTCollect() {
	//list := ListAccount() //账户列表
	/*	for i := 1; i < len(list); i++ {
			if list[i].Address != GASADDRES {
				//fmt.Println(i," 地址：",list[i].Address,"	ETH 余额：",list[i].ETH, "  USDT 余额：",list[i].USDT)
				//手续费不足发送0.001手续到有USDT的账户里面
				if list[i].ETH < 0.0001 && list[i].USDT >= 1 {
					txhash, _ := sendGas(list[i].Address, 0.0005)
					log.Println("GAS交易Hash：", txhash)
					//等待1分钟
					//time.Sleep(1*60*time.Second)
					////判断手续费是否确认
					//tr,err:=client.EthGetTransactionReceipt(txhash)
					//if err !=nil {
					//	log.Println("ERR",err.Error())
					//	return
					//}
					//
					//s,_ :=strconv.Atoi(strings.Split(tr.Status,"0x")[1])
					//if s==1{
					//	//解锁
					//	l,err:=	UnLock(list[i].Address,PWD)
					//	if err !=nil || l ==false{
					//		log.Println("解锁失败，错误提示：",err.Error())
					//		return
					//	}
					//	//开始收集USDT
					//	thash,_:=tokenTraction(list[i].Address,list[i].USDT)
					//	fmt.Println("归集交易hash:",thash)
					//}
				}
			}
		}
	*/
	//等待1分钟
	time.Sleep(60 * time.Second)
	//开始归集
	//判断手续费是否确认
	//tr,err:=client.EthGetTransactionReceipt(txhash)
	//if err !=nil {
	//	log.Println("ERR",err.Error())
	//	return
	//}

	//s,_ :=strconv.Atoi(strings.Split(tr.Status,"0x")[1])
	//if s==1{
	//}

	/*	for i := 1; i < len(list); i++ {
		if list[i].Address != GASADDRES {
			if list[i].USDT >= 1 {
				//解锁
				l, err := UnLock(list[i].Address, PWD)
				fmt.Println("归集账户：", list[i].Address)
				if err != nil || l == false {
					log.Println("解锁失败，错误提示：", err.Error())
					return
				}
				//开始收集USDT
				thash, _ := Traction(list[i].Address, list[i].USDT)
				fmt.Println("归集交易hash:", thash)
			}
		}
	}*/

	//time.Sleep(1*60*time.Second)
	//newlist:=listAccount()  //账户列表
	//for i:=1;i< len(newlist); i++ {
	//	fmt.Println(i," 地址：",newlist[i].Address,"	ETH 余额：",newlist[i].ETH, "  USDT 余额：",newlist[i].USDT)
	//
	//}
}

//发送手续费
//func sendGas(address string, gas float64) (string, error) {
//	GAS := int64(gas * 1e18)
//	v := T{
//		From:  GASADDRES,
//		To:    address,
//		Value: big.NewInt(GAS),
//	}
//	fmt.Println("手续费转账：", v.Value)
//
//	l, err := UnLock(GASADDRES, "HJBJK5810929")
//
//	if err != nil || l == false {
//		log.Println("解锁失败，错误提示：", err.Error())
//		return "", err
//	}
//
//	hash, err := client.EthSendTransaction(v)
//	if err != nil {
//		log.Println("ERR:", err.Error())
//	}
//
//	return hash, err
//}

//Token 转帐
func (rpc *EthRPC) Traction(toaddress string, value float64) (string, error) {

	taddr := strings.Split(toaddress, "0x")[1]

	//代币数量 十六进制的value值去掉0x并由0补够64位数
	//v:=strconv.FormatFloat(value,0.00000,6,64)
	vl := IntToHex(int(value * 1e6)) //"000000000000000000000000"
	vs := strings.Split(vl, "0x")[1]

	//fmt.Println("value :",vstr ," len:",len(vstr))
	//data拼接： “0x”+"23b872dd"+"from地址去掉0x并由0补够64位数"+"to地址去掉0x并由0补够64位数"+"十六进制的value值去掉0x并由0补够64位数"
	//data:="0x70a08231"+faddr+taddr+vstr //data拼接
	data := "0xa9059cbb" + addPreZero(taddr) + addPreZero(vs) //data拼接 addPreZero(faddr)

	fmt.Println("Data拼装：", data)
	t := T{
		From: Fromadd,
		To:   "NBH", //合约地址
		//Value:
		Data: data,
	}
	//hash,err:= client.EthPerSendTransaction(t,PWD)

	hash, err := rpc.EthSendTransaction(t)

	if err != nil {
		log.Println("错误信息:", err.Error())
	}
	//fmt.Println("交易Hash",hash)

	return hash, err
}

// 补齐64位，不够前面用0补齐
func addPreZero(num string) string {
	t := len(num)
	s := ""
	for i := 0; i < 64-t; i++ {
		s += "0"
	}
	return s + num
}

//创建钱包地址
func (rpc *EthRPC) newAccount(nu string) {

	for i := 0; i <= 1000; i++ {
		account, _ := rpc.EthPerNewAccount("luce!1989@9922sdf")
		fmt.Println(nu, " 数量：", i, "-", account)
		time.Sleep(300)
	}
}

//列出本地钱包
func (rpc *EthRPC) ListAccount() []string {

	list, err := rpc.EthPerListAccounts()
	if err != nil {
		log.Println("列出地址失败：", err.Error())
	}

	/*account := Account{}
	listAccounts := []Account{}
	//fmt.Println("本地钱包数量：",len(list))
	for i := 0; i < len(list); i++ {
		usdt := getUSDTBalance(list[i]) //USDT 余额
		blance := getBalance(list[i])   //获取余额
		if blance > 0 || usdt > 0 {
			account.Address = list[i]
			account.USDT = usdt
			account.ETH = blance
			listAccounts = append(listAccounts, account)
		}
	}*/
	return list
}

//获取ETH余额
//func GetBalance(address string) float64 {
//
//	blance, _ := client.EthGetBalance(address, "latest")
//
//	//wei TO  ETH
//	ethc, _ := ParseBigInt(blance.String())
//	intwei, _ := strconv.ParseInt(ethc.String(), 0, 64)
//	inteth := float64(intwei) / 1e18
//
//	return inteth
//	//fmt.Println("会员钱包数量：",len(acounts))
//	//for i:=0;i<len(acounts);i++  {
//	//
//	//	//client.EthPerUnLockAccount(acounts[i],"123456")
//	//
//	//	blance,_:=client.EthGetBalance(acounts[i],"latest")
//	//
//	//	//wei TO  ETH
//	//	ethc,_ :=eth.ParseBigInt(blance.String())
//	//	intwei,_:=eth.ParseInt(ethc.String())
//	//
//	//	inteth:=float64(intwei)/100000000000000000
//	//
//	//	if inteth>0 {
//	//		fmt.Println("地址：",acounts[i],"   余额：",inteth,"ETH")
//	//	}
//	//
//	//}
//}

//获取USDT余额
func (rpc *EthRPC) getUSDTBalance(address string) float64 {

	//0x7260c1661793170694344bC813BE6857ED16e58c
	addr2 := strings.Split(address, "0x")[1]
	data := "0x70a08231000000000000000000000000" + addr2 //data拼接

	t := T{
		From: address,
		To:   "USDT", //USDT合约地址
		Data: data,
	}

	balance, err := rpc.EthCall(t, "latest")
	if err != nil {
		log.Println("错误信息:", err.Error())
	}

	//wei TO  ETH
	ethc, _ := ParseBigInt(balance)
	intwei, _ := strconv.ParseFloat(ethc.String(), 6)
	inteth := float64(intwei) / 1000000

	return inteth
	//fmt.Println("USDT余额：",ethc.String(),"    ",inteth, "USDT")

	/**
	acounts,_:=client.EthAccounts()
	for i:=0;i<len(acounts);i++  {
		a:=strings.Split(acounts[i],"0x")[1]
		data:="0x70a08231000000000000000000000000"+a
		fmt.Println(  "NU:",i," DATA：",data)

		t:=eth.T{
			//From:acounts[i],
			To:"0xdac17f958d2ee523a2206206994597c13d831ec7", //USDT合约地址
			Data:data,
		}

		blance,_:=client.EthCall(t,"latest")
		//wei TO  ETH
		ethc,_ :=eth.ParseBigInt(blance)
		intwei,_:=eth.ParseInt(ethc.String())
		inteth:=float64(intwei)/1e6

		if inteth>0 {
			fmt.Println("地址：",acounts[i],"   USDT余额：",inteth,"ETH")
		}

	}*/
}

//获取合约余额
func (rpc *EthRPC) TokeBalance(address string, token string) float64 {

	//0x7260c1661793170694344bC813BE6857ED16e58c
	addr2 := strings.Split(address, "0x")[1]
	data := "0x70a08231000000000000000000000000" + addr2 //data拼接

	t := T{
		From: address,
		To:   token, //合约地址
		Data: data,
	}

	balance, err := rpc.EthCall(t, "latest")
	if err != nil {
		log.Println("错误信息:", err.Error())
	}

	//wei TO  ETH
	ethc, _ := ParseBigInt(balance)
	intwei, _ := strconv.ParseFloat(ethc.String(), 6)
	inteth := float64(intwei) / 1000000

	return inteth
	//fmt.Println("USDT余额：",ethc.String(),"    ",inteth, "USDT")

	/**
	acounts,_:=client.EthAccounts()
	for i:=0;i<len(acounts);i++  {
		a:=strings.Split(acounts[i],"0x")[1]
		data:="0x70a08231000000000000000000000000"+a
		fmt.Println(  "NU:",i," DATA：",data)

		t:=eth.T{
			//From:acounts[i],
			To:"0xdac17f958d2ee523a2206206994597c13d831ec7", //USDT合约地址
			Data:data,
		}

		blance,_:=client.EthCall(t,"latest")
		//wei TO  ETH
		ethc,_ :=eth.ParseBigInt(blance)
		intwei,_:=eth.ParseInt(ethc.String())
		inteth:=float64(intwei)/1e6

		if inteth>0 {
			fmt.Println("地址：",acounts[i],"   USDT余额：",inteth,"ETH")
		}

	}*/
}

////获取GasPrice
//func GasPrice() (big.Int, error) {
//	return client.EthGasPrice()
//}
//
////客户端的coinbase地址
//func CoinsBase() (string, error) {
//	return client.EthCoinbase()
//}

//合约地址转帐
func (rpc *EthRPC) TokenTraction(fromAddress, toAddress, token string, decimal int64, value float64) (string, error) {

	//收款地址截取去掉0x
	toaddr := strings.Split(toAddress, "0x")[1]

	//转账代币数量 十六进制的value值去掉0x并由0补够64位数
	//v:=strconv.FormatFloat(value,0.00000,18,64)
	//vl := IntToHex(int(value * 1e6)) //"000000000000000000000000"  USDT
	//log.Println("十六进制：",IntToHex(int(value * 1e6)))
	//IntToHex(int(value*1e18)) //"000000000000000000000000"

	//这是处理位数的代码段
	amount := big.NewFloat(float64(value))
	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	//log.Println("转账数量：", convertAmount)

	has := fmt.Sprintf("%x", convertAmount) //格式化数据
	//log.Println("Int十六进制：",has)

	//vs := strings.Split(has, "0x")[1]
	//log.Println("BIG十六进制：",vs)

	//data拼接： “0x”+"23b872dd"+"from地址去掉0x并由0补够64位数"+"to地址去掉0x并由0补够64位数"+"十六进制的value值去掉0x并由0补够64位数"
	//data:="0x70a08231"+faddr+taddr+vstr //data拼接 "0xa9059cbb"
	//0xa9059cbb

	data := TTtransferCode + addPreZero(toaddr) + addPreZero(has) //data拼接
	//gaspric,_:= GasPrice()
	//fmt.Println("Data拼装：", data)
	t := T{
		From:     fromAddress,
		Gas:      600000,                 //600000
		GasPrice: big.NewInt(4500000000), //big.NewInt(4500000000) 最快到账 60000000000  2500000000  普通：20000000000
		To:       token,                  //合约地址
		//Value:
		Data: data,
	}
	//hash,err:= client.EthPerSendTransaction(t,PWD)
	//log.Println("GasPrice",gaspric)
	hash, err := rpc.EthSendTransaction(t)

	if err != nil {
		log.Println("转账错误:", err.Error())
	}

	return hash, err

}

//获取合約余额
//func GetContractBalance(address string, token string, decimals int) float64 {
//
//	addrSplit := strings.Split(address, "0x")[1] //地址去掉0x
//	//data数据格式：最前边的“0x70a08231000000000000000000000000”是固定的，后边的是钱包地址（不带“0x”前缀）
//	data := "0x70a08231000000000000000000000000" + addrSplit //data拼接
//
//	t := T{
//		From: address, //查詢地址
//		To:   token,   //合约地址
//		Data: data,    //data
//	}
//
//	//获取代币的余额，要通过rpc接口得到接口为：eth_call
//	balance, err := client.EthCall(t, "latest")
//	if err != nil {
//		log.Println("错误信息:", err.Error())
//	}
//	//fmt.Println("余额：", balance)
//	//單位計算
//	ethc, _ := ParseBigInt(balance) //
//	intwei, _ := strconv.ParseFloat(ethc.String(), decimals)
//	inteth := intwei / math.Pow10(decimals)
//
//	return inteth
//
//}

const (
	ContractInfoCode_Name         = "0x06fdde03" //获取合约名称
	ContractInfoCode_Abbreviation = "0x95d89b41" //合约简称
	ContractInfoCode_Balance      = "0x70a08231" //查询余额
	ContractInfoCode_Accuracy     = "0x313ce567" //合约精度
	ContractInfoCode_Total        = "0x18160ddd" //发行总量

)

//获取合約信息
func (rpc *EthRPC) ContractInfo(from string, token string, code string) (interface{}, error) {
	//地址去掉0x
	addr := strings.TrimPrefix(from, "0x")
	//data数据格式：最前边的“0x70a08231000000000000000000000000”是固定的，后边的是钱包地址（不带“0x”前缀）
	data := code + "000000000000000000000000" + addr
	t := T{
		From: from,  //查詢地址
		To:   token, //合约地址
		Data: data,  //data
	}
	result, err := rpc.EthCall(t, "latest")
	if err != nil {
		return "", err
	}
	//單位計算
	resultBigInt, err := ParseBigInt(result)
	if err != nil {
		return "", err
	}
	switch code {
	case ContractInfoCode_Name, ContractInfoCode_Abbreviation: //获取合约名称
		n, _ := resultBigInt.GobEncode()
		//intwei, _ := strconv.ParseFloat(ethc.String(), 18)
		n, err := resultBigInt.GobEncode()
		if err != nil {
			return "", err
		}
		return string(n), nil
	case ContractInfoCode_Balance: //查询余额 	//
		intwei, err := strconv.ParseFloat(resultBigInt.String(), 18) //18
		if err != nil {
			return "", err
		}
		return intwei, nil
	case ContractInfoCode_Accuracy: //合约精度
		pdecimals, err := strconv.Atoi(resultBigInt.String())
		if err != nil {
			return "", err
		}
		return pdecimals, nil
	case ContractInfoCode_Total: //发行总量
		intwei, _ := strconv.ParseFloat(resultBigInt.String(), 18)
		inteth := intwei / math.Pow10(18)
		return inteth, nil
	}
	return nil, errors.New("not found contract code")
}
