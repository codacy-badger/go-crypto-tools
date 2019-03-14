package masternodes_online

import (
	"fmt"
	"testing"
)

func TestMNList(t *testing.T) {
	list, page, err := parseIPTable([]byte(getTestMNListData()))
	if err != nil {
		fmt.Printf("err: %v \n", err)
		t.Fail()
		return
	}
	if page != 36 {
		fmt.Printf("page: %v \n", page)
		t.Fail()
		return
	}
	if len(list) != 50 {
		fmt.Printf("lenList: %v \n", len(list))
		t.Fail()
		return
	}
}

func getTestMNListData() string {
	return `
<br />
<h3>Masternode list</h3>
<div class="row">
<div class="col-lg-12">
<div style="overflow-x:auto;" id="pubkey_list">
<table class="table table-striped table-hover">
<tr>
<td><strong>Masternode IP</strong></td><td><strong>Port</strong></td><td><strong>Pubkey</strong></td><td><strong>Status</strong></td><td><strong>Protocol</strong></td><td><strong>Last paid</strong></td></tr>
<tr><td>89.33.6.243</td><td>51472</td><td><a href="/pubkey/DEHCmfp5DLnrUxJmQ5BTwWZhqSE9qRXEGQ/">DEHCmfp5DLnrUxJmQ5BTwWZhqSE9qRXEGQ</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>3m 6s ago</td></tr>
<tr><td>2001:19f0:5801:19e9::2</td><td>51472</td><td><a href="/pubkey/DGMojuqhftjtbftFiMTMbT52PAyRJC83ak/">DGMojuqhftjtbftFiMTMbT52PAyRJC83ak</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>3m 18s ago</td></tr>
<tr><td>45.76.121.149</td><td>51472</td><td><a href="/pubkey/D7HiWFRa7q6mzF5i2ySBD2v5a5xbYkzfW3/">D7HiWFRa7q6mzF5i2ySBD2v5a5xbYkzfW3</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>4m 22s ago</td></tr>
<tr><td>2001:470:1f11:1d4:cc0:bfff:fe84:b280</td><td>51472</td><td><a href="/pubkey/D8ocdwDYcUhwzYJo2d7Y5wwGk1G1kDmLD2/">D8ocdwDYcUhwzYJo2d7Y5wwGk1G1kDmLD2</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>6m 52s ago</td></tr>
<tr><td>188.40.177.100</td><td>51472</td><td><a href="/pubkey/DQY88dgzn1jxEkM7ALgiei1mjU7BpcUiep/">DQY88dgzn1jxEkM7ALgiei1mjU7BpcUiep</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>7m 41s ago</td></tr>
<tr><td>2001:19f0:6801:1e04:2003::1</td><td>51472</td><td><a href="/pubkey/DPPLs3TdoMYXW9JDufLuqWrdmTk88pEMZJ/">DPPLs3TdoMYXW9JDufLuqWrdmTk88pEMZJ</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>10m 1s ago</td></tr>
<tr><td>185.162.250.158</td><td>51472</td><td><a href="/pubkey/DLkQVRRWLQY3f5UoAYQBrLMAnmwLN4D566/">DLkQVRRWLQY3f5UoAYQBrLMAnmwLN4D566</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>10m 14s ago</td></tr>
<tr><td>94.156.189.143</td><td>51472</td><td><a href="/pubkey/D59NnzEwQ8LPgFkwBR9vFrxjdSgD2ZJX17/">D59NnzEwQ8LPgFkwBR9vFrxjdSgD2ZJX17</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>10m 37s ago</td></tr>
<tr><td>2a0d:3001:2100:a001:9::3</td><td>51472</td><td><a href="/pubkey/DLUGqpZJYdKwexEcno8zrL2A8zn66nsXoL/">DLUGqpZJYdKwexEcno8zrL2A8zn66nsXoL</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>11m 54s ago</td></tr>
<tr><td>oisqrbedquvewswu.onion</td><td>51472</td><td><a href="/pubkey/DJQE1kHLcVn9VaFta981sfkoYUKDQsNcHJ/">DJQE1kHLcVn9VaFta981sfkoYUKDQsNcHJ</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>12m 21s ago</td></tr>
<tr><td>45.32.164.20</td><td>51472</td><td><a href="/pubkey/D8WL6oAx2RbzcEeXtXXQfyM5buZAXgAVHk/">D8WL6oAx2RbzcEeXtXXQfyM5buZAXgAVHk</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>12m 35s ago</td></tr>
<tr><td>107.191.49.249</td><td>51472</td><td><a href="/pubkey/D7rQATA2ueFh3bgoMBDsfmZFQy6SkePM6c/">D7rQATA2ueFh3bgoMBDsfmZFQy6SkePM6c</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>14m 52s ago</td></tr>
<tr><td>51.15.60.21</td><td>51472</td><td><a href="/pubkey/DCbFpPtvLfYcknYEBizmPvReFTTEVKg7j5/">DCbFpPtvLfYcknYEBizmPvReFTTEVKg7j5</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>15m 27s ago</td></tr>
<tr><td>108.61.202.131</td><td>51472</td><td><a href="/pubkey/DKtfBWyptyJekpbXBQESaS9vvJnx7M8buv/">DKtfBWyptyJekpbXBQESaS9vvJnx7M8buv</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>17m 35s ago</td></tr>
<tr><td>45.32.84.126</td><td>51472</td><td><a href="/pubkey/DB56o8ZyP3ER64vGyczA2tPmfndu8iuxxU/">DB56o8ZyP3ER64vGyczA2tPmfndu8iuxxU</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>18m 28s ago</td></tr>
<tr><td>45.63.68.26</td><td>51472</td><td><a href="/pubkey/DRiSjeFsxyt33savARWBpkUDno4bTQ2ZfX/">DRiSjeFsxyt33savARWBpkUDno4bTQ2ZfX</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>19m 21s ago</td></tr>
<tr><td>fypaivlhjxmcvhkk.onion</td><td>51472</td><td><a href="/pubkey/DK5V4MdweY5Fx8Sa3a8THNK7NjCu7XFCyp/">DK5V4MdweY5Fx8Sa3a8THNK7NjCu7XFCyp</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>19m 38s ago</td></tr>
<tr><td>45.77.121.13</td><td>51472</td><td><a href="/pubkey/DPf7RmUvb26z6j7HEiU4hLrJgWWVziTFUo/">DPf7RmUvb26z6j7HEiU4hLrJgWWVziTFUo</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>20m 21s ago</td></tr>
<tr><td>6txjwbxxdkaji3tf.onion</td><td>51472</td><td><a href="/pubkey/DGUbi89PXy6yKc4U7Wjzm3Bgd9cToMT8xK/">DGUbi89PXy6yKc4U7Wjzm3Bgd9cToMT8xK</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>21m 38s ago</td></tr>
<tr><td>163.172.216.135</td><td>51472</td><td><a href="/pubkey/DAfzCf8QeFPnj7fb6Uj99M43c6Zi7A85Ci/">DAfzCf8QeFPnj7fb6Uj99M43c6Zi7A85Ci</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>21m 43s ago</td></tr>
<tr><td>213.136.80.93</td><td>51472</td><td><a href="/pubkey/D5HK2d9pJsDGpNF2VqnKxSCJrhxLa5o8Fm/">D5HK2d9pJsDGpNF2VqnKxSCJrhxLa5o8Fm</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>21m 50s ago</td></tr>
<tr><td>fixn7p2ofruyuowi.onion</td><td>51472</td><td><a href="/pubkey/D5AvUFeukCVYAjSYCQCxYZyoqakWmKYHDG/">D5AvUFeukCVYAjSYCQCxYZyoqakWmKYHDG</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>22m 29s ago</td></tr>
<tr><td>rq4tbaryj6jpxgrs.onion</td><td>51472</td><td><a href="/pubkey/D5gJz2nLxKSBCjpTwHZXEeQV146iQRwmFf/">D5gJz2nLxKSBCjpTwHZXEeQV146iQRwmFf</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>23m 35s ago</td></tr>
<tr><td>2001:470:1f11:1d4:b8e3:58ff:fe33:41f5</td><td>51472</td><td><a href="/pubkey/D6RfsspXdUBDQ2AraY9xAHaCW4KKihM1Xi/">D6RfsspXdUBDQ2AraY9xAHaCW4KKihM1Xi</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>24m 22s ago</td></tr>
<tr><td>51.15.98.10</td><td>51472</td><td><a href="/pubkey/DATdhPaSByL5dsChrKh8fyexbeZF3CFsjL/">DATdhPaSByL5dsChrKh8fyexbeZF3CFsjL</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>25m 34s ago</td></tr>
<tr><td>217.69.9.154</td><td>51472</td><td><a href="/pubkey/DN3t9pdYiENpBpGpATTYPveJWjfBnNQdud/">DN3t9pdYiENpBpGpATTYPveJWjfBnNQdud</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>27m 14s ago</td></tr>
<tr><td>149.28.63.125</td><td>51472</td><td><a href="/pubkey/DHkQ7R6dhTJ878kSEbjK8MgoiHjViPGxR9/">DHkQ7R6dhTJ878kSEbjK8MgoiHjViPGxR9</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>28m 13s ago</td></tr>
<tr><td>qzha6vrwu4scqbfk.onion</td><td>51472</td><td><a href="/pubkey/DKGunFwuB5Q9c4npK9x8MW2xc8kiKAJWFc/">DKGunFwuB5Q9c4npK9x8MW2xc8kiKAJWFc</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>29m 6s ago</td></tr>
<tr><td>f64uqjlnyonoeutw.onion</td><td>51472</td><td><a href="/pubkey/DBT8CjQXcwqnczbxJFqrujDqejvqWsveSF/">DBT8CjQXcwqnczbxJFqrujDqejvqWsveSF</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>29m 29s ago</td></tr>
<tr><td>45.77.116.214</td><td>51472</td><td><a href="/pubkey/DEAMHCgm2t7CSTGfxGRzmMcBfkmaDbjAES/">DEAMHCgm2t7CSTGfxGRzmMcBfkmaDbjAES</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>30m 44s ago</td></tr>
<tr><td>45.32.220.251</td><td>51472</td><td><a href="/pubkey/D7CWTBAqJ2RPUaNx3HmHkzSDbMGH5XnS2o/">D7CWTBAqJ2RPUaNx3HmHkzSDbMGH5XnS2o</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>31m 31s ago</td></tr>
<tr><td>207.148.76.169</td><td>51472</td><td><a href="/pubkey/DBZEoVdraE7mwzivJAv5RFKfE7eSC4uUXF/">DBZEoVdraE7mwzivJAv5RFKfE7eSC4uUXF</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>31m 53s ago</td></tr>
<tr><td>149.28.212.235</td><td>51472</td><td><a href="/pubkey/DCEAqYtCN23nfziwQUAbzNHUVZzgYZLfW6/">DCEAqYtCN23nfziwQUAbzNHUVZzgYZLfW6</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>32m 10s ago</td></tr>
<tr><td>hg64itvvjfon4tha.onion</td><td>51472</td><td><a href="/pubkey/DNKiK5zwBpBZXtGT5Ldee2UeHqqENTUhcs/">DNKiK5zwBpBZXtGT5Ldee2UeHqqENTUhcs</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>33m 9s ago</td></tr>
<tr><td>e6qanlm4jlhxe2ba.onion</td><td>51472</td><td><a href="/pubkey/DEELFYbQRDEkvWEjCYsg1N92pyAmpKcJxM/">DEELFYbQRDEkvWEjCYsg1N92pyAmpKcJxM</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>33m 53s ago</td></tr>
<tr><td>5.135.187.212</td><td>51472</td><td><a href="/pubkey/DNZoTppKhgpovSHrJjUPKiLbic8LBfx5ZJ/">DNZoTppKhgpovSHrJjUPKiLbic8LBfx5ZJ</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>34m 57s ago</td></tr>
<tr><td>176.31.116.11</td><td>51472</td><td><a href="/pubkey/DALkypPG3HcwprUZfethJDy6y7SXXaZSN4/">DALkypPG3HcwprUZfethJDy6y7SXXaZSN4</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>37m 19s ago</td></tr>
<tr><td>104.131.78.235</td><td>51472</td><td><a href="/pubkey/DKpS2ZD68G6dKfUb9qUAdenizwjy9cUPdT/">DKpS2ZD68G6dKfUb9qUAdenizwjy9cUPdT</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>37m 38s ago</td></tr>
<tr><td>45.77.174.8</td><td>51472</td><td><a href="/pubkey/DHq7UTEUk2xgXvnsZXcFjSJjt7CJnWoM91/">DHq7UTEUk2xgXvnsZXcFjSJjt7CJnWoM91</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>38m 51s ago</td></tr>
<tr><td>54.38.198.155</td><td>51472</td><td><a href="/pubkey/DTZX1UaBqogH92DfAXcspg8unCVX49yscY/">DTZX1UaBqogH92DfAXcspg8unCVX49yscY</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>39m 59s ago</td></tr>
<tr><td>185.92.221.32</td><td>51472</td><td><a href="/pubkey/DGksLufPv6tv2zir4HXMysschMaoY9aD12/">DGksLufPv6tv2zir4HXMysschMaoY9aD12</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>40m 4s ago</td></tr>
<tr><td>ho56qg4capbi6fll.onion</td><td>51472</td><td><a href="/pubkey/DP7V14T47JHDanwZCcpV2FGTZnSadzuUQV/">DP7V14T47JHDanwZCcpV2FGTZnSadzuUQV</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>41m 3s ago</td></tr>
<tr><td>2001:19f0:4401:992::12</td><td>51472</td><td><a href="/pubkey/DF267WH4bD4h4F8NCKXDgmSDdn1KTYPncm/">DF267WH4bD4h4F8NCKXDgmSDdn1KTYPncm</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>41m 10s ago</td></tr>
<tr><td>51.38.109.225</td><td>51472</td><td><a href="/pubkey/DGYG1btiT1RZKUXtAHc3woC7WwU91PqfHm/">DGYG1btiT1RZKUXtAHc3woC7WwU91PqfHm</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>41m 58s ago</td></tr>
<tr><td>2001:470:1f11:1d4:5455:65ff:fe17:963</td><td>51472</td><td><a href="/pubkey/DEALojR1mccJHYnQdbmeJYdFzZM1teySmd/">DEALojR1mccJHYnQdbmeJYdFzZM1teySmd</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>42m 58s ago</td></tr>
<tr><td>2001:470:1f11:1d4:8c44:ebff:fec9:3039</td><td>51472</td><td><a href="/pubkey/DEJLwaPUcwJh4QJfYtoQhHoJkydSmem7bE/">DEJLwaPUcwJh4QJfYtoQhHoJkydSmem7bE</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>43m 42s ago</td></tr>
<tr><td>66.42.48.22</td><td>51472</td><td><a href="/pubkey/D6gcTZXSkJdxLTmd32wNgPAEnPx4tLaWqp/">D6gcTZXSkJdxLTmd32wNgPAEnPx4tLaWqp</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>44m 52s ago</td></tr>
<tr><td>77.170.32.79</td><td>51472</td><td><a href="/pubkey/DDY7hbdkyFjGRqCGFJf2necXqWYWXBPcMY/">DDY7hbdkyFjGRqCGFJf2necXqWYWXBPcMY</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>46m 10s ago</td></tr>
<tr><td>95.179.162.129</td><td>51472</td><td><a href="/pubkey/D8WKbWL7Pa5zhehmYYNVWe4jQwLYzNsRqX/">D8WKbWL7Pa5zhehmYYNVWe4jQwLYzNsRqX</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>47m 50s ago</td></tr>
<tr><td>2001:470:1f11:1d4:3c5c:9eff:fe92:aca4</td><td>51472</td><td><a href="/pubkey/DCNn3t5oq6XwivL8Z3KG3zstqrzKDmGoKG/">DCNn3t5oq6XwivL8Z3KG3zstqrzKDmGoKG</a></td><td><span class="text-success">ENABLED</span></td><td><span class="text-success" title="wallet is up-to-date">70914</span></td><td>48m 1s ago</td></tr>
</table>
</div>
<div align="right">
 <ul class="pagination"><li class="active"><a href="#">1</a></li><li><a href="/currencies/PIVX/?page=2">2</a></li><li><a href="/currencies/PIVX/?page=3">3</a></li><li><a href="/currencies/PIVX/?page=4">4</a></li><li><a href="/currencies/PIVX/?page=2">Next</a></li><li><a href="/currencies/PIVX/?page=36">Last</a></li></ul></div>
<hr />
<br />
<p align="center"><strong>MNO is a masternode coin monitoring and stats service. MNO does not research or recommend any coin. Do your own research and invest at your own risk.</strong></p>
<p align="center">ROI changes often and is not the most important factor.  Please consider <strong>Dev Team - Community - PURPOSE/Platform - Liquidity - Wallet</strong> when making masternode purchases.</p>
</div>
</div>
<footer>
      <div class="row">
          <div class="col-lg-12">
<hr />
`
}
