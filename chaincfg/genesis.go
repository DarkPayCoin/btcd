// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/DarkPayCoin/btcd/chaincfg/chainhash"
	"github.com/DarkPayCoin/btcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).

// 04ffff001d0104454461726b506179436f696e2c206a757374206c696b65204c69676874506179436f696e206f6e20737465726f69647320776974686f7574207468652062756c6c7368697421

var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, 
				0x44,0x61,0x72,0x6b,0x50,0x61,0x79,0x43,0x6f,0x69,
				0x6e,0x2c,0x20,0x6a,0x75,0x73,0x74,0x20,0x6c,0x69,
				0x6b,0x65,0x20,0x4c,0x69,0x67,0x68,0x74,0x50,0x61,
				0x79,0x43,0x6f,0x69,0x6e,0x20,0x6f,0x6e,0x20,0x73,
				0x74,0x65,0x72,0x6f,0x69,0x64,0x73,0x20,0x77,0x69,
				0x74,0x68,0x6f,0x75,0x74,0x20,0x74,0x68,0x65,0x20,
				0x62,0x75,0x6c,0x6c,0x73,0x68,0x69,0x74,0x21,
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value:    0,
			PkScript: []byte{},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x18, 0xaa, 0x09, 0xa6, 0xe4, 0xb9, 0x2d, 0x48, 0x9c,
	0xf6, 0x5f, 0x76, 0xb9, 0x3d, 0x1c, 0x32, 0xa5, 0x14,
	0xc7, 0xf2, 0xbb, 0x8e, 0x1c, 0x20, 0xfb, 0xbd, 0x4f,
	0xd6, 0x5b, 0x0d, 0x00, 0x00,
})

//chaque octet à l'envers :
//         assert(hashGenesisBlock == uint256("0x00,00,0d,5b,d6,4f,bd,fb,20,1c,8e,bb,f2,c7,14,a5,32,1c,3d-b9-76-5f-f6-9c-48-2d-b9-e4a609aa18"));
// assert(genesis.hashMerkleRoot == uint256("0xd8,c2,ab,19,9e,75,91,13,74,47,90,37,41,36,4e,59,cf,4a,5d,e5,5c,68,7a,78,4e,d3,22,c8,94,0a,50,45"));

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x45, 0x50, 0x0a, 0x94, 0xc8, 0x22, 0xd3, 0x4e, 0x78,
	0x7a, 0x68, 0x5c, 0xe5, 0x5d, 0x4a, 0xcf, 0x59, 0x4e,
	0x36, 0x41, 0x37, 0x90, 0x47, 0x74, 0x13, 0x91, 0x75,
	0x9e, 0x19, 0xab, 0xc2, 0xd8,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1536012000, 0), // Monday 3 September 2018 22:00:00 GMT
		Bits:       0x1e0ffff0,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      2189760,                    // 2083236893
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}