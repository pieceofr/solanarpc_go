package solanarpc

var testResultAccountExsitBase58 string = `{
	"jsonrpc": "2.0",
	"result": {
	  "context": {
		"slot": 1
	  },
	  "value": {
		"data": [
		  "11116bv5nS2h3y12kD1yUKeMZvGcKLSjQgX6BeV7u1FrjeJcKfsHRTPuR3oZ1EioKtYGiYxpxMG5vpbZLsbcBYBEmZZcMKaSoGx9JZeAuWf",
		  "base58"
		],
		"executable": false,
		"lamports": 1000000000,
		"owner": "11111111111111111111111111111111",
		"rentEpoch": 2
	  }
	},
	"id": 1
  }`
var testResultAccountNotExist string = `{
    "jsonrpc": "2.0",
    "result": {
        "context": {
            "slot": 76951452
        },
        "value": null
    },
    "id": 2
}`

var testResultAccountError string = `{
    "jsonrpc": "2.0",
    "error": {
        "code": -32602,
        "message": "Invalid param: WrongSize"
    },
    "id": 2
}`

var testResultBalance01 string = `{
    "jsonrpc": "2.0",
    "result": {
        "context": {
            "slot": 76954704
        },
        "value": 27020041285
    },
    "id": 1
}`

var testResultBlockCommitment01 string = `{
	"jsonrpc":"2.0",
	"result":{
	  "commitment":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,10,32],
	  "totalStake": 42
	},
	"id":1
  }`

var testResultBlockTime01 string = `{
    "jsonrpc": "2.0",
    "result": 1574721591,
    "id": 1
}`

var testResultBlockTime02 string = `{
    "jsonrpc": "2.0",
    "result": null,
    "id": 1
}`

var testResultBlockTime03 string = `{
    "jsonrpc": "2.0",
    "result": -123,
    "id": 1
}`

var testResultClusterNodes01 string = `{
	"jsonrpc": "2.0",
	"result": [
	  {
		"gossip": "10.239.6.48:8001",
		"pubkey": "9QzsJf7LPLj8GkXbYT3LFDKqsj2hHG7TA3xinJHu8epQ",
		"rpc": "10.239.6.48:8899",
		"tpu": "10.239.6.48:8856",
		"version": "1.0.0 c375ce1f"
	  },
	  {
		"featureSet": null,
		"gossip": "3.14.216.138:11000",
		"pubkey": "6xnLs5AnhkTkNcgArVSopx32sheFim1oGQwBWUJJXG1F",
		"rpc": null,
		"tpu": null,
		"version": null
	}
	],
	"id": 1
  }`

var testResultConfirmedBlock01 string = `{
	"jsonrpc": "2.0",
	"result": {
	  "blockTime": null,
	  "blockhash": "3Eq21vXNB5s86c62bVuUfTeaMif1N2kUqRPBmGRJhyTA",
	  "parentSlot": 429,
	  "previousBlockhash": "mfcyqEXB3DnHXki6KjjmZck6YjmZLvpAByy2fj4nh6B",
	  "transactions": [
		{
		  "meta": {
			"err": null,
			"fee": 5000,
			"innerInstructions": [],
			"logMessages": [],
			"postBalances": [
			  499998932500,
			  26858640,
			  1,
			  1,
			  1
			],
			"postTokenBalances": [],
			"preBalances": [
			  499998937500,
			  26858640,
			  1,
			  1,
			  1
			],
			"preTokenBalances": [],
			"status": {
			  "Ok": null
			}
		  },
		  "transaction": {
			"message": {
			  "accountKeys": [
				"3UVYmECPPMZSCqWKfENfuoTv51fTDTWicX9xmBD2euKe",
				"AjozzgE83A3x1sHNUR64hfH7zaEBWeMaFuAN9kQgujrc",
				"SysvarS1otHashes111111111111111111111111111",
				"SysvarC1ock11111111111111111111111111111111",
				"Vote111111111111111111111111111111111111111"
			  ],
			  "header": {
				"numReadonlySignedAccounts": 0,
				"numReadonlyUnsignedAccounts": 3,
				"numRequiredSignatures": 1
			  },
			  "instructions": [
				{
				  "accounts": [
					1,
					2,
					3,
					0
				  ],
				  "data": "37u9WtQpcm6ULa3WRQHmj49EPs4if7o9f1jSRVZpm2dvihR9C8jY4NqEwXUbLwx15HBSNcP1",
				  "programIdIndex": 4
				}
			  ],
			  "recentBlockhash": "mfcyqEXB3DnHXki6KjjmZck6YjmZLvpAByy2fj4nh6B"
			},
			"signatures": [
			  "2nBhEBYYvfaAe16UMNqRHre4YNSskvuYgx3M6E4JP1oDYvZEJHvoPzyUidNgNX5r9sTyN1J9UxtbCXy2rqYcuyuv"
			]
		  }
		}
	  ]
	},
	"id": 1
  }`

var testResultConfirmedBlock02 string = `{
	"jsonrpc": "2.0",
	"result": null,
	"id": 1
  }`

var testResultConfirmedBlock03 string = `{
    "jsonrpc": "2.0",
    "error": {
        "code": -32007,
        "message": "Slot 76884393 was skipped, or missing due to ledger jump to recent snapshot"
    },
    "id": 1
}`

var testResultBlockProduction01 string = `{
	"jsonrpc": "2.0",
	"result": {
	  "context": {
		"slot": 9887
	  },
	  "value": {
		"byIdentity": {
		  "85iYT5RuzRTDgjyRa3cP8SYhM2j21fj7NhfJ3peu1DPr": [
			9888,
			9886
		  ]
		},
		"range": {
		  "firstSlot": 0,
		  "lastSlot": 9887,
		}
	  }
	},
	"id": 1
  }`

var testResultConfirmedBlocks01 string = `{
    "jsonrpc": "2.0",
    "result": [
        77228566,
        77228567,
        77228568,
        77228569,
        77228570
    ],
    "id": 1
}`

var testResultConfirmedBlocks02 string = `{
    "jsonrpc": "2.0",
    "result": [],
    "id": 1
}`

var testResultConfirmedBlocks03 string = `{
    "jsonrpc": "2.0",
    "error": {
        "code": -32602,
        "message": "Slot range too large; max 500000"
    },
    "id": 1
}`
