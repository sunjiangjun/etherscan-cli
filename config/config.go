package config

/**
{

   "key": [
     "37CTHB671QQHEEQFNNV64D8MACMJMHPSG9",
     "E4GM1H5Y7AEXS3S9XIT4U2VUBNVJBPMNFI",
     "8VZRWD329TBQGGATEHUA57Z5E7GHQJ7P9Y"
   ],

  "address": [
    "0xdac17f958d2ee523a2206206994597c13d831ec7"
  ]

}

*/

type Config struct {
	Address []string `json:"address" gorm:"column:address"`
	Key     []string `json:"key" gorm:"column:key"`
}
