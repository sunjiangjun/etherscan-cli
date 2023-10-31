# etherscan-cli

According to the monitoring address set in the configuration file, monitor transactions involving address changes in new Ethereum blocks and save them to the specified file or push them to the client.

tips:

go version>=1.20


## config.json 
``````

{

   "key": [ // token in etherscan
     "37CTHB671QQHEEQFNNV64D8MACMJMHPSG9",
     "E4GM1H5Y7AEXS3S9XIT4U2VUBNVJBPMNFI",
     "8VZRWD329TBQGGATEHUA57Z5E7GHQJ7P9Y"
   ],

  "address": [ // monitor address
    "0xde66f7b6ec7e6d857437aad56e679175fbab9360",
    "0xd7a1edf88871112dd4257885b20f523ab387163d"
  ]

}

``````

## runing
./etherscan-cli [options]

  -config string the path of config file



