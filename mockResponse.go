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

var testResultConfirmedBlock04 string = `{
    "jsonrpc": "2.0",
    "result": {
        "blockTime": 1620374821,
        "blockhash": "8FBzBHVg5eLY6LJ2uhzT1vYxDWE6raafGfePssNKNdYQ",
        "parentSlot": 77228599,
        "previousBlockhash": "67kqu56ayRW91kwT2yDYYhx2QQHsz3YZtcARZwoKXc88",
        "rewards": [
            {
                "lamports": 1817500,
                "postBalance": 912525496459,
                "pubkey": "EvnRmnMrd69kFdbLMxWkTn1icZ7DCceRhvmb2SJXqDo4",
                "rewardType": "Fee"
            }
        ]
    },
    "id": 1
}`
var testResultConfirmedBlock05 string = `{
    "jsonrpc": "2.0",
    "result": {
        "blockTime": 1620374821,
        "blockhash": "8FBzBHVg5eLY6LJ2uhzT1vYxDWE6raafGfePssNKNdYQ",
        "parentSlot": 77228599,
        "previousBlockhash": "67kqu56ayRW91kwT2yDYYhx2QQHsz3YZtcARZwoKXc88",
        "rewards": [
            {
                "lamports": 1817500,
                "postBalance": 912525496459,
                "pubkey": "EvnRmnMrd69kFdbLMxWkTn1icZ7DCceRhvmb2SJXqDo4",
                "rewardType": "Fee"
            }
        ],
        "signatures": [
            "3sZKvqgwCSDraErrqE8WQiXuWf85zhezR2VEQFNBAMyYJMUw5psX61iMALxmxRohAwSkEDCTRtDTTXEGVzjdtfJQ",
            "5sUcdRKENq8cy1XWXXAJnrFHnwzsgo1H2vdPCjunCQcvhkRDtPk7bDtLEJm7KDugSifVT4pRy6XFASe4PL5WGpEK",
            "4pXBJxxUrBBepVFi5LmHbkB6paY89cKbj9XGWq1C61ehkxojEGJ5NcJbVMpzSLURqfbYMEsEDzzubq6SDLeDFvtv",
            "39YeURzS3Eku3CcukdTim4CMru1BY3WDSqqN4BtrH694FNsP1CsDSVJNsc97BLgEcZhpzmhwVQSYhy4HbjsXtBgm",
            "JtXRg9woJWWgZwJTTvdy6YdtdY6akGkb2Jk6nJmQbQ3reF6EDPg4jwzbBNapt9cMuMyWCXhRMpwWhzk7Kosurxj",
            "5BP9ZkWnht9uEHgjjphUuki6V7rQwNzY42vaTEdu8EuJVRDrMNeAFYx34xX5EtGayyqbdFMVok89j4V7EpZ7K5mC",
            "5Mdq1gfis4nzN8CZXNvi3sdyj4vzrQYGfQzwq776KEDP5PnJyJGuqXHKca43U3aeH2DRmUJ7mtMWMuGZ8tT8AtxH"
        ]
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

var testResultConfirmedBlockWithLimit01 string = `{
    "jsonrpc": "2.0",
    "result": [
        77118633,
        77118634,
        77118635
    ],
    "id": 1
}`

var testResultConfirmedSignaturesForAddress201 string = `{
    "jsonrpc": "2.0",
    "result": [
        {
            "blockTime": 1620403540,
            "confirmationStatus": "confirmed",
            "err": null,
            "memo": null,
            "signature": "67k4Uzed5ZBBgbSdPTCSUCx2pLyJTDXTZzzAKfQejviPdUJM5Yz1njXeHdrbg7nQKCsBfWBsVte7b34mfKeypW3B",
            "slot": 77279819
        }
    ],
    "id": 1
}`

var testResultTokenSupply01 string = `{
  "jsonrpc": "2.0",
  "result": {
      "context": {
          "slot": 77292120
      },
      "value": {
          "amount": "555000000000000",
          "decimals": 6,
          "uiAmount": 555000000,
          "uiAmountString": "555000000"
      }
  },
  "id": 1
}`
