package client

import (
	"fmt"
	"testing"
)

func TestGetMap(t *testing.T) {
	cmcMap, err := parseNestedJSONMap([]byte(getTestCmcInfoData()))
	if err != nil {
		fmt.Printf("err:%v\n", err)
		t.Fail()
		return
	}
	if len(cmcMap) == 0 {
		t.Fail()
		return
	}

}

func getTestCmcInfoData() string {
	return `{
    "status": {
        "timestamp": "2019-02-22T10:50:20.122Z",
        "error_code": 0,
        "error_message": null,
        "elapsed": 3,
        "credit_count": 1
    },
    "data": {
        "131": {
            "urls": {
                "website": [
                    "https://www.dash.org/"
                ],
                "twitter": [
                    "https://twitter.com/Dashpay"
                ],
                "reddit": [
                    "https://reddit.com/r/dashpay"
                ],
                "message_board": [
                    "https://www.dash.org/forum/"
                ],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=421615.0"
                ],
                "chat": [
                    "https://discord.gg/9z8zX5j",
                    "https://discordapp.com/invite/PXbUxJB"
                ],
                "explorer": [
                    "https://explorer.dash.org",
                    "https://insight.dash.org/insight/",
                    "https://chainz.cryptoid.info/dash/"
                ],
                "source_code": [
                    "https://github.com/dashpay/dash"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/131.png",
            "id": 131,
            "name": "Dash",
            "symbol": "DASH",
            "slug": "dash",
            "date_added": "2014-02-14T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "541": {
            "urls": {
                "website": [
                    "http://syscoin.org"
                ],
                "twitter": [
                    "https://twitter.com/syscoin"
                ],
                "reddit": [
                    "https://reddit.com/r/SysCoin"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=1466445.0"
                ],
                "chat": [
                    "https://t.me/joinchat/AAAAAEHzByog3h1qutnjhQ"
                ],
                "explorer": [
                    "https://chainz.cryptoid.info/sys/"
                ],
                "source_code": [
                    "https://github.com/syscoin/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/541.png",
            "id": 541,
            "name": "Syscoin",
            "symbol": "SYS",
            "slug": "syscoin",
            "date_added": "2014-08-20T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "707": {
            "urls": {
                "website": [
                    "http://blocknet.co/"
                ],
                "twitter": [
                    "https://twitter.com/The_Blocknet"
                ],
                "reddit": [
                    "https://reddit.com/r/theblocknet"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=829576.0"
                ],
                "chat": [
                    "https://discord.gg/2e6s7H8",
                    "https://t.me/Blocknet"
                ],
                "explorer": [
                    "https://chainz.cryptoid.info/block/",
                    "https://block.ccore.online/"
                ],
                "source_code": [
                    "https://github.com/BlocknetDX/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/707.png",
            "id": 707,
            "name": "Blocknet",
            "symbol": "BLOCK",
            "slug": "blocknet",
            "date_added": "2014-11-01T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "720": {
            "urls": {
                "website": [
                    "http://crown.tech/"
                ],
                "twitter": [
                    "https://twitter.com/CrownPlatform"
                ],
                "reddit": [
                    "https://reddit.com/r/Crown"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=815487"
                ],
                "chat": [
                    "https://t.me/crownplatform",
                    "https://discordapp.com/invite/rB3Kr86"
                ],
                "explorer": [
                    "https://chainz.cryptoid.info/crw/"
                ],
                "source_code": [
                    "https://gitlab.crown.tech/crown/crown-core"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/720.png",
            "id": 720,
            "name": "Crown",
            "symbol": "CRW",
            "slug": "crown",
            "date_added": "2014-11-08T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "1169": {
            "urls": {
                "website": [
                    "http://www.pivx.org"
                ],
                "twitter": [
                    "https://twitter.com/_pivx"
                ],
                "reddit": [
                    "https://reddit.com/r/pivx"
                ],
                "message_board": [
                    "https://forum.pivx.org/"
                ],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=1262920"
                ],
                "chat": [],
                "explorer": [
                    "https://chainz.cryptoid.info/pivx/",
                    "https://pivx.tokenview.com/en/",
                    "http://pivx.presstab.pw"
                ],
                "source_code": [
                    "https://github.com/PIVX-Project/PIVX/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/1169.png",
            "id": 1169,
            "name": "PIVX",
            "symbol": "PIVX",
            "slug": "pivx",
            "date_added": "2016-02-13T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "1414": {
            "urls": {
                "website": [
                    "https://zcoin.io"
                ],
                "twitter": [
                    "https://twitter.com/zcoinofficial"
                ],
                "reddit": [
                    "https://reddit.com/r/zcoin"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=1638450"
                ],
                "chat": [
                    "https://discordapp.com/invite/4FjnQ2q"
                ],
                "explorer": [
                    "http://explorer.zcoin.io",
                    "https://chainz.cryptoid.info/xzc/",
                    "https://insight.zcoin.io/"
                ],
                "source_code": [
                    "https://github.com/zcoinofficial/zcoin/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/1414.png",
            "id": 1414,
            "name": "Zcoin",
            "symbol": "XZC",
            "slug": "zcoin",
            "date_added": "2016-10-06T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "1606": {
            "urls": {
                "website": [
                    "http://solariscoin.com/"
                ],
                "twitter": [
                    "https://twitter.com/SolarisCoin"
                ],
                "reddit": [
                    "https://reddit.com/r/solarisxlr"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=1831629.0"
                ],
                "chat": [
                    "https://t.me/solariscoin"
                ],
                "explorer": [
                    "https://solaris.blockexplorer.pro/",
                    "https://chainz.cryptoid.info/xlr/",
                    "https://solaris.blockxplorer.info/"
                ],
                "source_code": [
                    "https://github.com/Solaris-Project/Solaris/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/1606.png",
            "id": 1606,
            "name": "Solaris",
            "symbol": "XLR",
            "slug": "solaris",
            "date_added": "2017-04-01T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "1814": {
            "urls": {
                "website": [
                    "http://lindacoin.com/"
                ],
                "twitter": [
                    "https://twitter.com/lindaproject"
                ],
                "reddit": [],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=5096898"
                ],
                "chat": [
                    "https://discord.gg/SQR4uj4"
                ],
                "explorer": [
                    "https://prohashing.com/explorer/Lindacoin/",
                    "http://193.70.109.114:6005/"
                ],
                "source_code": [
                    "https://github.com/TheLindaProjectInc/Linda"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/1814.png",
            "id": 1814,
            "name": "Linda",
            "symbol": "LINDA",
            "slug": "linda",
            "date_added": "2017-07-15T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "1877": {
            "urls": {
                "website": [
                    "http://www.rupayacoin.org/"
                ],
                "twitter": [
                    "https://twitter.com/rupayacoin"
                ],
                "reddit": [
                    "https://reddit.com/r/RupayaCoin"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2956408"
                ],
                "chat": [
                    "https://t.co/hrjiNGFSUW"
                ],
                "explorer": [
                    "http://node2.rupayacoin.org/",
                    "https://find.rupx.io/",
                    "https://hereismy.rupx.io/"
                ],
                "source_code": [
                    "https://github.com/rupaya-project/rupx"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/1877.png",
            "id": 1877,
            "name": "Rupaya",
            "symbol": "RUPX",
            "slug": "rupaya",
            "date_added": "2017-08-11T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "2158": {
            "urls": {
                "website": [
                    "https://phore.io/"
                ],
                "twitter": [
                    "https://twitter.com/phorecrypto"
                ],
                "reddit": [
                    "https://reddit.com/r/PhoreProject"
                ],
                "message_board": [
                    "https://medium.com/@phoreblockchain"
                ],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=3179265.0"
                ],
                "chat": [
                    "https://t.me/PhoreProject",
                    "https://discord.gg/Aucncz5"
                ],
                "explorer": [
                    "https://chainz.cryptoid.info/phr/"
                ],
                "source_code": [
                    "https://github.com/phoreproject/Phore"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2158.png",
            "id": 2158,
            "name": "Phore",
            "symbol": "PHR",
            "slug": "phore",
            "date_added": "2017-11-07T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "2199": {
            "urls": {
                "website": [
                    "https://alqo.org/"
                ],
                "twitter": [
                    "https://twitter.com/ALQOCOIN"
                ],
                "reddit": [
                    "https://reddit.com/r/Alqo"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2343884.0"
                ],
                "chat": [
                    "https://discord.gg/BJYpbfs",
                    "https://t.me/alqochat"
                ],
                "explorer": [
                    "http://explorer.alqo.org/"
                ],
                "source_code": [
                    "https://github.com/alqocrypto/alqo"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2199.png",
            "id": 2199,
            "name": "ALQO",
            "symbol": "XLQ",
            "slug": "alqo",
            "date_added": "2017-11-19T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "2200": {
            "urls": {
                "website": [
                    "https://gobyte.network/"
                ],
                "twitter": [
                    "https://twitter.com/gobytenetwork"
                ],
                "reddit": [
                    "https://reddit.com/r/gobyte"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2442185.0"
                ],
                "chat": [
                    "https://discord.gg/mDAeFzT",
                    "https://t.me/GoByteNetwork"
                ],
                "explorer": [
                    "http://explorer.gobyte.network:5001/",
                    "https://insight.gobyte.network"
                ],
                "source_code": [
                    "https://github.com/gobytecoin/gobyte/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2200.png",
            "id": 2200,
            "name": "GoByte",
            "symbol": "GBX",
            "slug": "gobyte",
            "date_added": "2017-11-20T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "2260": {
            "urls": {
                "website": [
                    "https://bulwarkcrypto.com/"
                ],
                "twitter": [
                    "https://twitter.com/bulwarkcoin"
                ],
                "reddit": [
                    "https://reddit.com/r/bulwarkcoin"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2499481.0"
                ],
                "chat": [
                    "https://t.me/bulwarkcrypto",
                    "https://discord.gg/a7vhegP"
                ],
                "explorer": [
                    "http://explorer.bulwarkcrypto.com"
                ],
                "source_code": [
                    "https://github.com/bulwark-crypto"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2260.png",
            "id": 2260,
            "name": "Bulwark",
            "symbol": "BWK",
            "slug": "bulwark",
            "date_added": "2017-12-08T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "2359": {
            "urls": {
                "website": [
                    "https://polispay.org/"
                ],
                "twitter": [
                    "https://twitter.com/PolisBlockchain"
                ],
                "reddit": [
                    "https://reddit.com/r/Polispay"
                ],
                "message_board": [
                    "http://forum.polispay.org/"
                ],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2627897.0"
                ],
                "chat": [
                    "https://discord.gg/gwhHv8U"
                ],
                "explorer": [
                    "https://explorer.polispay.org/"
                ],
                "source_code": [
                    "https://github.com/polispay/polis"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2359.png",
            "id": 2359,
            "name": "Polis",
            "symbol": "POLIS",
            "slug": "polis",
            "date_added": "2018-01-10T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "2604": {
            "urls": {
                "website": [
                    "https://www.savebitcoin.io/"
                ],
                "twitter": [
                    "https://twitter.com/btc_green"
                ],
                "reddit": [
                    "https://reddit.com/r/btcgreen"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2827989.0"
                ],
                "chat": [
                    "https://t.me/bitcoin_green",
                    "https://discordapp.com/invite/QaPjqff"
                ],
                "explorer": [
                    "http://explorer.savebitcoin.io/",
                    "https://www.coinexplorer.net/BITG"
                ],
                "source_code": [
                    "https://github.com/bitcoingreen/bitcoingreen"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2604.png",
            "id": 2604,
            "name": "Bitcoin Green",
            "symbol": "BITG",
            "slug": "bitcoin-green",
            "date_added": "2018-03-22T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "2633": {
            "urls": {
                "website": [
                    "https://stakenet.io/"
                ],
                "twitter": [
                    "https://twitter.com/XSNofficial"
                ],
                "reddit": [
                    "https://reddit.com/r/Stakenet"
                ],
                "message_board": [
                    "https://medium.com/stakenet"
                ],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=3213013.0"
                ],
                "chat": [
                    "https://t.me/joinchat/BdGxxw-s3b4_DdBdbChI4g",
                    "https://discord.gg/cyF5yCA"
                ],
                "explorer": [
                    "https://xsnexplorer.io/"
                ],
                "source_code": [
                    "https://github.com/X9Developers/XSN"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2633.png",
            "id": 2633,
            "name": "Stakenet",
            "symbol": "XSN",
            "slug": "stakenet",
            "date_added": "2018-04-11T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "2773": {
            "urls": {
                "website": [
                    "https://gincoin.io"
                ],
                "twitter": [
                    "https://twitter.com/gincoin_crypto"
                ],
                "reddit": [],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2998031"
                ],
                "chat": [
                    "https://t.me/gincoin",
                    "https://discord.gg/t9YRSFN"
                ],
                "explorer": [
                    "https://explorer.gincoin.io"
                ],
                "source_code": [
                    "https://github.com/gincoin-dev/gincoin-core"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2773.png",
            "id": 2773,
            "name": "GINcoin",
            "symbol": "GIN",
            "slug": "gincoin",
            "date_added": "2018-05-23T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "2942": {
            "urls": {
                "website": [
                    "http://aegeus.io"
                ],
                "twitter": [
                    "https://twitter.com/Aegeus_Coin"
                ],
                "reddit": [
                    "https://reddit.com/r/Aegeus_Coin_Official"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=3644459.0"
                ],
                "chat": [
                    "https://t.me/joinchat/H6MhmRMCEGThVbXkPNhV4g",
                    "https://discordapp.com/invite/mKRjwb"
                ],
                "explorer": [
                    "https://chainz.cryptoid.info/aeg/"
                ],
                "source_code": [
                    "https://github.com/AegeusCoin/aegeus/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2942.png",
            "id": 2942,
            "name": "Aegeus",
            "symbol": "AEG",
            "slug": "aegeus",
            "date_added": "2018-07-17T00:00:00.000Z",
            "tags": null,
            "platform": null,
            "category": "coin"
        },
        "2991": {
            "urls": {
                "website": [
                    "https://nixplatform.io/"
                ],
                "twitter": [
                    "https://twitter.com/nixplatform"
                ],
                "reddit": [
                    "https://reddit.com/r/NixPlatform"
                ],
                "message_board": [
                    "https://nixplatform.io/blog/"
                ],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=5099680.0"
                ],
                "chat": [
                    "https://t.me/nixplatform",
                    "https://discordapp.com/invite/agAsvQY"
                ],
                "explorer": [
                    "https://blockchain.nixplatform.io/"
                ],
                "source_code": [
                    "https://github.com/NixPlatform/NixCore"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/2991.png",
            "id": 2991,
            "name": "NIX",
            "symbol": "NIX",
            "slug": "nix",
            "date_added": "2018-07-28T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "3197": {
            "urls": {
                "website": [
                    "https://graphcoin.net"
                ],
                "twitter": [
                    "https://twitter.com/GRPHcoin"
                ],
                "reddit": [
                    "https://reddit.com/r/graphcoin"
                ],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=4691938"
                ],
                "chat": [
                    "https://discord.gg/U8gN4BP"
                ],
                "explorer": [
                    "http://explorer.graphcoin.net/"
                ],
                "source_code": [
                    "https://github.com/Graphcoin/source"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/3197.png",
            "id": 3197,
            "name": "Graphcoin",
            "symbol": "GRPH",
            "slug": "graphcoin",
            "date_added": "2018-08-22T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "3201": {
            "urls": {
                "website": [
                    "http://www.printex.tech"
                ],
                "twitter": [
                    "https://twitter.com/Printex_Team"
                ],
                "reddit": [],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=3755807.0"
                ],
                "chat": [
                    "https://discord.gg/qaneFPm"
                ],
                "explorer": [
                    "https://blocks.printex.tech/"
                ],
                "source_code": [
                    "https://github.com/Printex-official/printex-core/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/3201.png",
            "id": 3201,
            "name": "Printex",
            "symbol": "PRTX",
            "slug": "printex",
            "date_added": "2018-08-22T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "3242": {
            "urls": {
                "website": [
                    "https://beetlecoin.io/"
                ],
                "twitter": [
                    "https://twitter.com/beetlecoin"
                ],
                "reddit": [],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=2961802.0"
                ],
                "chat": [
                    "https://discordapp.com/invite/NsJkV6g"
                ],
                "explorer": [
                    "http://explorer.beetlecoin.io/"
                ],
                "source_code": [
                    "https://github.com/beetledev/Wallet"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/3242.png",
            "id": 3242,
            "name": "Beetle Coin",
            "symbol": "BEET",
            "slug": "beetle-coin",
            "date_added": "2018-09-06T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "3272": {
            "urls": {
                "website": [
                    "https://coin2play.io/"
                ],
                "twitter": [
                    "https://twitter.com/coin2play"
                ],
                "reddit": [],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=4849316"
                ],
                "chat": [
                    "https://t.me/coin2play",
                    "https://discord.gg/35KFpHM"
                ],
                "explorer": [
                    "http://explorer.coin2play.io/"
                ],
                "source_code": [
                    "https://github.com/Coin2Play/"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/3272.png",
            "id": 3272,
            "name": "Coin2Play",
            "symbol": "C2P",
            "slug": "coin2play",
            "date_added": "2018-09-12T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "3521": {
            "urls": {
                "website": [
                    "https://paws.fund/"
                ],
                "twitter": [
                    "https://twitter.com/pawsfund"
                ],
                "reddit": [],
                "message_board": [
                    "https://medium.com/paws-animal-charity",
                    "https://steemit.com/@pawsfund"
                ],
                "announcement": [
                    "https://forum.paws.fund/"
                ],
                "chat": [
                    "https://discord.gg/6wk25BH",
                    "https://t.me/pawsfund"
                ],
                "explorer": [
                    "http://chain.paws.fund/"
                ],
                "source_code": []
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/3521.png",
            "id": 3521,
            "name": "PAWS Fund",
            "symbol": "PAWS",
            "slug": "paws-fund",
            "date_added": "2018-10-30T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        },
        "3613": {
            "urls": {
                "website": [
                    "https://dashgreen.net/"
                ],
                "twitter": [
                    "https://twitter.com/dashgreennet"
                ],
                "reddit": [],
                "message_board": [],
                "announcement": [
                    "https://bitcointalk.org/index.php?topic=5068282.0"
                ],
                "chat": [
                    "https://discord.gg/WgVMS4S"
                ],
                "explorer": [
                    "http://explorer.dashgreen.net/"
                ],
                "source_code": [
                    "https://github.com/DashGreenCoin/dashgreen"
                ]
            },
            "logo": "https://s2.coinmarketcap.com/static/img/coins/64x64/3613.png",
            "id": 3613,
            "name": "Dash Green",
            "symbol": "DASHG",
            "slug": "dash-green",
            "date_added": "2019-01-14T00:00:00.000Z",
            "tags": [
                "mineable"
            ],
            "platform": null,
            "category": "coin"
        }
    }
}`
}
