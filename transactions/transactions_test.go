package transactions

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/node-standalone-pool/go-pool-server/daemonManager"
	"github.com/node-standalone-pool/go-pool-server/utils"
	"log"
	"testing"
)

func TestCreateGeneration(t *testing.T) {
	data := `
{
  "capabilities": [
    "proposal"
  ],
  "version": 536870912,
  "rules": [
    "csv",
    "segwit"
  ],
  "vbavailable": {
  },
  "vbrequired": 0,
  "previousblockhash": "b83b698bed0897ac94819041aec857d1a26a567bf7bf046d60849d5ccf24155e",
  "transactions": [
    {
      "data": "01000000012f8975c900f56662f35c317a0669fecc5fe0e1fb8ee53f4de72f1cb68c07e606010000008a473044022061a9ac17f269f3c69e18b5d67dfa6bf8b6a5a60eb7f9b0c992ffaeb66b5b88fb02202bb6fd7eb539302d97f4b8604bc822c91747e1cb365fecc91b37526b6b8c2c25014104fe67366f857106ee7b4cc48abb4dabd46302e12fe4140f4c933b92bd3ce75b1f4ae45055312f9a6c5ddc1f8d94d4f6d11e2a13372bcd6bfd651e48997b0f767effffffff02e8030000000000001976a914dffec839eba107e556d6c4f25f90765b3d10583288acbb60da04000000001976a914bdd83cf3ab8b7a57ff9b841752c1ae764f2a02ee88ac00000000",
      "txid": "f9b8b0bdd0dc38b2a707faf89acf064f543c3a88d39f54fb126cbd084ffb5ed9",
      "hash": "f9b8b0bdd0dc38b2a707faf89acf064f543c3a88d39f54fb126cbd084ffb5ed9",
      "depends": [
      ],
      "fee": 6450,
      "sigops": 8,
      "weight": 1028
    },
    {
      "data": "0200000001979a795a82096fc375487778939d9193bb284c58525e5df9c3a404c81c9220ef01000000d9004730440220086f0b09ded442c84e602520f5a8b38b41a1bc860fb595bd47834c20fa8db39402200a40cb86c15198302cabfd5c620c24fa6ac9ac5d946394e37dd3f9960b65a0e701473044022049bc0be153a4535196f73455bf82667956f2089019db4eeb57cb35649d8f69b202206ddd411917cb3e54f7b9a89c0693c971a6499eb6e421dce1d5fa9358300525d301475221025ad7eedea4c87b98463b8c7316c139f94c0e75fe4c849f42dab112479e1a1bb7210257591ace4d6a9fc94b8114cffd84df9bd0349c974a792580f7f5afb74f5ba94952ae0000000002102700000000000017a914b75a640760f2caae367c0e0cd6bfb85e8d80755987e17608000000000017a914ef20c4471b54fc47c93d587a318d351e93fbc13b8700000000",
      "txid": "620c724890f76b802714d786d5d3fe13a89106d81e93b74c4eafd6dc04179f37",
      "hash": "620c724890f76b802714d786d5d3fe13a89106d81e93b74c4eafd6dc04179f37",
      "depends": [
      ],
      "fee": 3493,
      "sigops": 8,
      "weight": 1328
    }
  ],
  "coinbaseaux": {
    "flags": ""
  },
  "coinbasevalue": 2500009943,
  "longpollid": "b83b698bed0897ac94819041aec857d1a26a567bf7bf046d60849d5ccf24155e19620",
  "target": "00000000a9490000000000000000000000000000000000000000000000000000",
  "mintime": 1581747579,
  "mutable": [
    "time",
    "transactions",
    "prevblock"
  ],
  "noncerange": "00000000ffffffff",
  "sigoplimit": 80000,
  "sizelimit": 4000000,
  "weightlimit": 4000000,
  "curtime": 1581749398,
  "bits": "1d00a949",
  "height": 1369986
}

`

	var rpcData daemonManager.GetBlockTemplate
	_ = json.Unmarshal([]byte(data), &rpcData)

	pk := utils.P2PKHAddressToScript("QPxrDq3sorCk8DWaYX2GeCkxoePhm1asyY")
	placeholder, _ := hex.DecodeString("f000000ff111111f")

	log.Println(hex.EncodeToString(utils.PackUint32LE(uint32(0))))

	gens := CreateGeneration(&rpcData, pk, placeholder, "POW", true, make(map[string]float64))

	log.Println("0: ", hex.EncodeToString(gens[0]))
	log.Println("1: ", hex.EncodeToString(gens[1]))

	t1, _ := hex.DecodeString("02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff1f0382e71404")
	if !bytes.Contains(gens[0], t1) {
		t.Fail()
	}

	if hex.EncodeToString(gens[1]) != "0c2f627920436f6d6d616e642f0000000001d71f0395000000001976a91424da8749fde8fcdcde60ba1c5afea8d2bd4a4f2688ac000000000a627920436f6d6d616e64" {
		t.Fail()
	}
}

func TestGenerateOutputTransactions(t *testing.T) {
	publicKey := utils.P2PKHAddressToScript("QPxrDq3sorCk8DWaYX2GeCkxoePhm1asyY")
	recipients := make(map[string]float64)
	data := `
{
  "capabilities": [
    "proposal"
  ],
  "version": 536870912,
  "rules": [
    "csv",
    "segwit"
  ],
  "vbavailable": {
  },
  "vbrequired": 0,
  "previousblockhash": "b83b698bed0897ac94819041aec857d1a26a567bf7bf046d60849d5ccf24155e",
  "transactions": [
    {
      "data": "01000000012f8975c900f56662f35c317a0669fecc5fe0e1fb8ee53f4de72f1cb68c07e606010000008a473044022061a9ac17f269f3c69e18b5d67dfa6bf8b6a5a60eb7f9b0c992ffaeb66b5b88fb02202bb6fd7eb539302d97f4b8604bc822c91747e1cb365fecc91b37526b6b8c2c25014104fe67366f857106ee7b4cc48abb4dabd46302e12fe4140f4c933b92bd3ce75b1f4ae45055312f9a6c5ddc1f8d94d4f6d11e2a13372bcd6bfd651e48997b0f767effffffff02e8030000000000001976a914dffec839eba107e556d6c4f25f90765b3d10583288acbb60da04000000001976a914bdd83cf3ab8b7a57ff9b841752c1ae764f2a02ee88ac00000000",
      "txid": "f9b8b0bdd0dc38b2a707faf89acf064f543c3a88d39f54fb126cbd084ffb5ed9",
      "hash": "f9b8b0bdd0dc38b2a707faf89acf064f543c3a88d39f54fb126cbd084ffb5ed9",
      "depends": [
      ],
      "fee": 6450,
      "sigops": 8,
      "weight": 1028
    },
    {
      "data": "0200000001979a795a82096fc375487778939d9193bb284c58525e5df9c3a404c81c9220ef01000000d9004730440220086f0b09ded442c84e602520f5a8b38b41a1bc860fb595bd47834c20fa8db39402200a40cb86c15198302cabfd5c620c24fa6ac9ac5d946394e37dd3f9960b65a0e701473044022049bc0be153a4535196f73455bf82667956f2089019db4eeb57cb35649d8f69b202206ddd411917cb3e54f7b9a89c0693c971a6499eb6e421dce1d5fa9358300525d301475221025ad7eedea4c87b98463b8c7316c139f94c0e75fe4c849f42dab112479e1a1bb7210257591ace4d6a9fc94b8114cffd84df9bd0349c974a792580f7f5afb74f5ba94952ae0000000002102700000000000017a914b75a640760f2caae367c0e0cd6bfb85e8d80755987e17608000000000017a914ef20c4471b54fc47c93d587a318d351e93fbc13b8700000000",
      "txid": "620c724890f76b802714d786d5d3fe13a89106d81e93b74c4eafd6dc04179f37",
      "hash": "620c724890f76b802714d786d5d3fe13a89106d81e93b74c4eafd6dc04179f37",
      "depends": [
      ],
      "fee": 3493,
      "sigops": 8,
      "weight": 1328
    }
  ],
  "coinbaseaux": {
    "flags": ""
  },
  "coinbasevalue": 2500009943,
  "longpollid": "b83b698bed0897ac94819041aec857d1a26a567bf7bf046d60849d5ccf24155e19620",
  "target": "00000000a9490000000000000000000000000000000000000000000000000000",
  "mintime": 1581747579,
  "mutable": [
    "time",
    "transactions",
    "prevblock"
  ],
  "noncerange": "00000000ffffffff",
  "sigoplimit": 80000,
  "sizelimit": 4000000,
  "weightlimit": 4000000,
  "curtime": 1581749398,
  "bits": "1d00a949",
  "height": 1369986
}

`
	var rpcData daemonManager.GetBlockTemplate
	json.Unmarshal([]byte(data), &rpcData)
	log.Println(hex.EncodeToString(GenerateOutputTransactions(publicKey, recipients, &rpcData)))
}

// 00000020fb08e0b3cb0f759671af79f108dd2dbd1a378ba27968c176c1c6d64f94741d262a7ca761bb4397d2c1a7f6cf457d680f054d43cc4f860de21b776054ab93a3cafbe34b5effff0f1e00452ef00401000000010000000000000000000000000000000000000000000000000000000000000000ffffffff1f0377ee1404fce34b5e086b3c0000000000000c2f627920436f6d6d616e642f00000000020000000000000000266a24aa21a9ed8a44e041a5a86878a1742f66fe7196400e784fee5cdc70a4becaf51c8f4a4f0266140395000000001976a91424da8749fde8fcdcde60ba1c5afea8d2bd4a4f2688ac0000000001000000000101df2565bde1779eaa6aad06a03a5262d324de29aa51735ba26d2301c0af426ee90100000000ffffffff020000000000000000136a0c0701007ac0010000c0000000530345d47106fa2f0b0000000016001407fa56d069e6174b6fa1ca3e27556be765064e150247304402203fb97652eee91717f61a9a9a66c8c233648ac3f5942aeb246c2217ffdc64b5f70220210653c0bc9c74b026e80e77a3221a6c64c9529d37a2051c53aefb19213f381d012102a56c007c837c6323332f03f2d22190f1da0aec10c2338da50ebcec85100e9a96000000000100000000010129d40378ffb37a1b2b751e4469aed63df827636e538e76550f6629dc49978f0b0100000000f0ffffff0340420f00000000001976a914ab83ab1e9284beca76ecdd1460f732acdeb5a45688ac0bd9460000000000160014756b524ee4ec544d7828cb849951b75bf46cf9d30000000000000000196a1768747470733a2f2f746c74632e6269746170732e636f6d02483045022100ab49baf3f2f0ebc910f2d7453a5a50bc10810a17d9a77ddc98638b1fc1a93929022062d798622d7f5a55655e1356b9e0e7fb0d2c40d70df0c10c5ba91311a085559c012102ab861da09e496373d8aee62107d68f8275df04dca403c182b6ab648eebb4aca50000000001000000000101fdaffc6f8c94565763bdf4c0e50c389c5eb817d2dff247a0ffc519dda211dce90100000000f0ffffff0340420f00000000001976a914ab83ab1e9284beca76ecdd1460f732acdeb5a45688ac29d5440000000000160014d3cb800cd29671af47dfd95fcb759a7e76e4b0dd0000000000000000196a1768747470733a2f2f746c74632e6269746170732e636f6d02473044022060f807e10801d10ba51870bbbaee01d80d3727730b43b9c710af80857c14e4d102201002940096d20427246a7738c358b29675108260b18cc612922b1a01b2a588ab0121031506590ee0b0a9cfa13dbc765d9ac9666e5e01d031c5bd5b5e293bcdeb2932af00000000
// 00000020763600ad521ebbb8be835992a5f7e1e315d3978934ed805bfffd2e88b7d65c7c073172cf11eb40f1b663749268e2c63b1c2b1e54fb80d69467c9b5017eb9437cc0e34b5effff0f1e002aaa0e0101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff1f0370ee1404c1e34b5e0840000000000000000c2f627920436f6d6d616e642f00000000020000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf900f90295000000001976a91424da8749fde8fcdcde60ba1c5afea8d2bd4a4f2688ac00000000
