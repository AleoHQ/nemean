package main

import "github.com/urfave/cli"

var newAccountCommand = cli.Command{
	Name:     "create",
	Category: "wallet",
	Usage:    "Create a new Aleo account.",
	Description: `
	The create command is used to create a new Aleo account.
	`,
	Action: newAccount,
}

var newTransactionCommand = cli.Command{
	Name:     "send",
	Category: "wallet",
	Usage:    "Create a basic transfer transaction.",
	Description: `
	The send command creates a single transfer transaction that consumes
	a single record and returns a serialized transaction in hex.
	`,
	Action: newTransaction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "to",
			Usage:    "recipient address",
			Required: true,
		},
		cli.StringSliceFlag{
			Name:     "ledger_proof",
			Usage:    "list of ledger proofs for input record",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "amount",
			Usage:    "amount to send",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "fee",
			Usage:    "network fee",
			Required: true,
		},
		cli.StringFlag{
			Name:     "private_key",
			Usage:    "private key to sign transaction",
			Required: true,
		},
		cli.StringFlag{
			Name:     "record",
			Usage:    "JSON input record to consume",
			Required: true,
		},
	},
}

var getBlockCommand = cli.Command{
	Name:     "get_block",
	Category: "rpc",
	Usage:    "Get a block.",
	Description: `
	Gets a block from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlock,
}

var getBlockHashCommand = cli.Command{
	Name:     "get_block_hash",
	Category: "rpc",
	Usage:    "Get a block hash.",
	Description: `
	Gets a block hash from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:     "height",
			Usage:    "block height",
			Required: true,
		},
	},
	Action: getBlockHash,
}

var getBlockHeightCommand = cli.Command{
	Name:     "get_block_height",
	Category: "rpc",
	Usage:    "Get a block height.",
	Description: `
	Gets a block height from SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "hash",
			Usage:    "block hash",
			Required: true,
		},
	},
	Action: getBlockHeight,
}

var sendTransactionCommand = cli.Command{
	Name:     "send_transaction",
	Category: "rpc",
	Usage:    "Broadcasts a transaction.",
	Description: `
	Broadcasts a transaction using SnarkOS.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "txn",
			Usage:    "transaction hex",
			Required: true,
		},
	},
	Action: sendTransaction,
}

var latestLedgerRootCommand = cli.Command{
	Name:     "latest_ledger_root",
	Category: "rpc",
	Usage:    "Gets the latest ledger root.",
	Description: `
	Returns the latest ledger root from the canonical chain.
	`,
	Action: latestLedgerRoot,
}

var getLedgerProofCommand = cli.Command{
	Name:     "ledger_proof",
	Category: "rpc",
	Usage:    "Gets the ledger proof.",
	Description: `
	Returns the ledger proof for the given commitment with the current ledger root.
	`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "commitment",
			Usage:    "The record commitment to generate a ledger proof of inclusion for.",
			Required: true,
		},
	},
	Action: getLedgerProof,
}