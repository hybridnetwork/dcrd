// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

// BlockOneLedgerMainNet is the block one output ledger for the main
// network.
var BlockOneLedgerMainNet = []*TokenPayout{}

// BlockOneLedgerTestNet is the block one output ledger for the test
// network.
var BlockOneLedgerTestNet = []*TokenPayout{
	{"TsmWaPM77WSyA3aiQ2Q1KnwGDVWvEkhipBc", 100000 * 1e8},
}

// BlockOneLedgerTestNet2 is the block one output ledger for the 2nd test
// network.
var BlockOneLedgerTestNet2 = []*TokenPayout{
	{"Tsn9NzUG7fVgr9en42RDJ8BhhkoLvuZfZAZ", 20000000 * 1e8},
	{"Tsi1DJUyenj5gc4ZLR8faBzneRgRnDKYUMZ", 20000000 * 1e8},
	{"TsnxFr2Hi1mmgHt4EP6zi4wP5vbLgFFScWs", 20000000 * 1e8},
	{"TsgFKaxAzpxnzxpBNRMiBAd5inANSmqvtLS", 20000000 * 1e8},
	{"TsgmR7mZ7suZqa7gbood9vj8m6ASd61Qez7", 20000000 * 1e8},
	{"TsVaH1Xf2hCJk2nvzLgXsesidNQfuYpDwFs", 20000000 * 1e8},
	{"TsRKGkJBhztSThMwPCiCjRsirL7Rv9RneSL", 20000000 * 1e8},
	{"TsW1FrJKLMLh6hqAFfj9KAHKL5sHFqqkRkS", 20000000 * 1e8},
	{"TsUyg1Bj8u9iMzZqGP8PD72mTmMzpzXLKhr", 20000000 * 1e8},
	{"TsRvKhZAmNj6g4954h2UjWhYxSZEdbZwRbe", 20000000 * 1e8},
}

// BlockOneLedgerSimNet is the block one output ledger for the simulation
// network. See under "Decred organization related parameters" in params.go
// for information on how to spend these outputs.
var BlockOneLedgerSimNet = []*TokenPayout{
	{"Sshw6S86G2bV6W32cbc7EhtFy8f93rU6pae", 100000 * 1e8},
	{"SsjXRK6Xz6CFuBt6PugBvrkdAa4xGbcZ18w", 100000 * 1e8},
	{"SsfXiYkYkCoo31CuVQw428N6wWKus2ZEw5X", 100000 * 1e8},
}
