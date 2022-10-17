# Bootstrap Me

## Overview

A simple little tool to help you launch a node on any Cosmos chain a bit easier and faster than it would've been. **This is not intended to be a Terraform-based "node launcher", but rather a tool to give you better visibility of a chain while launching a node**. Lots of interesting ideas to build out. Stay tuned anon.

## Features

- Grab `chain.json` for any chain on the [Cosmos Chain Registry](https://github.com/cosmos/chain-registry)
- Fetch the persistent peers and seeds listed and check if they are still reachable
- (Optional) Perform a port scanning on the reachable hosts. If a host has ports open that shouldn't be exposed, you can then decide whether to take it out from the final list
- (Optional) Fetch geo data for the reachable hosts
- Remove unreachable hosts and generate a final, concatenated list of reachable persistent peer/seeds

## Sample

```bash
$ go run main.go --chain <chain> --scan true > log
$ go run main.go --help # see all supported chains
```

<details close>

```
********************************************************************
              Checking Reachability of Persistent Peers
********************************************************************
82772547c4575c18dfe6e75aafe521cf7d4dc8de@142.93.157.186:26656 is unreachable
241b17dba97a2ed3c3747d12781fb86c9706e2d4@95.179.136.131:26656 is unreachable
f122129f53b7c584df6cee77716dcc636d5c5e18@167.172.59.196:26656 is unreachable
f1b16c603f3a0e59f0ce5179dc80f549a7ecd0e2@sentries.us-east1.iqext.net:26656 is unreachable
ee27245d88c632a556cf72cc7f3587380c09b469@45.79.249.253:26656 is unreachable
538ebe0086f0f5e9ca922dae0462cc87e22f0a50@34.122.34.67:26656 is unreachable
1bfda3d59e70290a3dada9bb809dd954371850d3@54.180.225.240:26656 is unreachable
d3209b9f88eec64f10555a11ecbf797bb0fa29f4@34.125.169.233:26656 is unreachable
6ee94c2093505e8790442c054e6e1e0211d36583@44.239.140.195:26656 is unreachable
654f47a762c8f9257aef4a44c1fb5014916d8b20@99.79.60.15:26656 is unreachable
366ac852255c3ac8de17e11ae9ec814b8c68bddb@51.15.94.196:26656 is unreachable
5b4ed476e01c49b23851258d867cc0cfc0c10e58@206.189.4.227:26656 is unreachable
d72b3011ed46d783e369fdf8ae2055b99a1e5074@173.249.50.25:26656 is unreachable
047f723806ee702b211e7227f89eacd829aabd86@52.9.212.125:26656 is unreachable
3c7cad4154967a294b3ba1cc752e40e8779640ad@84.201.128.115:26656 is unreachable
585794737e6b318957088e645e17c0669f3b11fc@54.160.123.34:26656 is unreachable
11dfe200894f38e411beca77928e9dd118e66813@94.130.98.157:26656 is unreachable
bdc2c3d410ca7731411b7e46a252012323fbbf37@34.83.209.166:26656 is unreachable
********************************************************************
                   Checking Reachability of Seeds
********************************************************************
bf8328b66dceb4987e5cd94430af66045e59899f@public-seed.cosmos.vitwit.com:26656 is unreachable
bcef90de8a83673c336bf3b3a352445b3a3a1f08@cosmos-seed.sunshinevalidation.io:31038 is unreachable
366ac852255c3ac8de17e11ae9ec814b8c68bddb@51.15.94.196:26656 is unreachable
3c7cad4154967a294b3ba1cc752e40e8779640ad@84.201.128.115:26656 is unreachable
d72b3011ed46d783e369fdf8ae2055b99a1e5074@173.249.50.25:26656 is unreachable
********************************************************************
                     Reachable Persistent Peers
********************************************************************
ec779a2741da6dd2ccdaa6dfc0bebb10e595dfa4@50.18.113.67:26656,b0e746acb6fbed7a0311fe21cfb2ee94581ca3bc@51.79.21.187:26656,d6318b3bd51a5e2b8ed08f2e520d50289ed32bf1@52.79.43.100:26656,64bd8eaf08b05f17ccd88425f80b59ab48934004@157.90.18.35:26656,cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656
********************************************************************
                          Reachable Seeds
********************************************************************
57a5297537b9b6ef8b105c08a8ad3f6ac452c423@seeds.goldenratiostaking.net:1618,cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656,ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0@seeds.polkachu.com:14956,ba3bacc714817218562f743178228f23678b2873@public-seed-node.cosmoshub.certus.one:26656,20e1000e88125698264454a884812746c2eb4807@seeds.lavenderfive.com:14956,3b67739570f921cc5e0db4b3efe488ce184155a9@seeds.pupmos.network:2000
********************************************************************
                          Port Scanning
********************************************************************
!! Domain names in this section are redacted for security reasons !!

>> Scanning Reachable Persistent Peers
Port 22 is open on 54.160.123.34
Port 9090 is open on 52.79.43.100

>> Scanning Reachable Seeds
Port 2200 is open on seeds.xxxxxx.com
Port 5333 is open on seeds.xxxxxx.com
Port 6333 is open on seeds.xxxxxx.com
Port 9184 is open on seeds.xxxxxx.com
Port 11856 is open on seeds.xxxxxx.com
Port 12256 is open on seeds.xxxxxx.com
Port 12856 is open on seeds.xxxxxx.com
Port 13056 is open on seeds.xxxxxx.com
Port 13156 is open on seeds.xxxxxx.com
Port 13356 is open on seeds.yyyyyy.com
Port 14156 is open on seeds.yyyyyy.com
Port 15556 is open on seeds.xxxxxx.com
Port 16756 is open on seeds.yyyyyy.com

----------------------------------------------------------
```

</details>
