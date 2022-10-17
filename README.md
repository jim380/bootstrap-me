# Bootstrap Me

A simple little tool to help you launch a node on any Cosmos chain faster than it would've been. **This is not intended to be a Terraform-based "node launcher", but rather a tool to give you better visibility of a chain while launching a node**. Lots of interesting ideas to build out. Stay tuned anon.

## Sample

```bash
$ go run main.go --chain cosmoshub --scan true
```

<details close>

```
45.79.249.253:26656 is unreachable
34.122.34.67:26656 is unreachable
34.125.169.233:26656 is unreachable
34.83.209.166:26656 is unreachable
94.130.98.157:26656 is unreachable
206.189.4.227:26656 is unreachable
99.79.60.15:26656 is unreachable
51.15.94.196:26656 is unreachable
54.180.225.240:26656 is unreachable
44.239.140.195:26656 is unreachable
173.249.50.25:26656 is unreachable
52.9.212.125:26656 is unreachable
142.93.157.186:26656 is unreachable
84.201.128.115:26656 is unreachable
167.172.59.196:26656 is unreachable
95.179.136.131:26656 is unreachable
sentries.us-east1.iqext.net:26656 is unreachable
public-seed.cosmos.vitwit.com:26656 is unreachable
173.249.50.25:26656 is unreachable
84.201.128.115:26656 is unreachable
51.15.94.196:26656 is unreachable
cosmos-seed.sunshinevalidation.io:31038 is unreachable
********************************************************************
**     Unreachable hosts are removed in the final lists below     **
********************************************************************
Persistent Peers: 585794737e6b318957088e645e17c0669f3b11fc@54.160.123.34:26656,d6318b3bd51a5e2b8ed08f2e520d50289ed32bf1@52.79.43.100:26656,ec779a2741da6dd2ccdaa6dfc0bebb10e595dfa4@50.18.113.67:26656,cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656,b0e746acb6fbed7a0311fe21cfb2ee94581ca3bc@51.79.21.187:26656,64bd8eaf08b05f17ccd88425f80b59ab48934004@157.90.18.35:26656

Seeds: cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656,ba3bacc714817218562f743178228f23678b2873@public-seed-node.cosmoshub.certus.one:26656,3b67739570f921cc5e0db4b3efe488ce184155a9@seeds.pupmos.network:2000,ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0@seeds.polkachu.com:14956,20e1000e88125698264454a884812746c2eb4807@seeds.lavenderfive.com:14956,57a5297537b9b6ef8b105c08a8ad3f6ac452c423@seeds.goldenratiostaking.net:1618

********************************************************************
** Domain names in this section are redacted for security reasons **
********************************************************************
Port Scanning for Persistent Peers
Port 22 is open on 54.160.123.34
Port 9090 is open on 52.79.43.100

Port Scanning for Seeds
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
