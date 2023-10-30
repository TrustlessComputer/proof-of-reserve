package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
)

type BitcoinParams struct {
	MasterPubKeys   []string
	NumRequiredSigs int
	TotalSigs       int
	ChainParam      *chaincfg.Params
}

var BitcoinParamsNOS = &BitcoinParams{
	MasterPubKeys: []string{
		"02bf0fd86a31568497c7635d9b48d48194cd12a3083ba04e599c0ccdb1b0ba955b",
		"024237a4f5fe8057ebf5ee890a892be3958ed3691ab7c0af84d72d197ec961bf98",
		"02b0305abe6d6bae7ee95e4a7fede1281b6d3df7f0d841d4edb155b1978b0835f7",
		"029097a2513dff395905fa3d7b4d1dce258608fd936d4552add96c6e7a4d2d5c3a",
		"02d6b28c3e9ca7cc870afc95192c3a6fa6de6a4acc9fb0a7dc8c8ca7fb46c4cfd4",
		"021bf77c146362f5a99f208b3512b52de22b0c5ba5e1eb4d5b7eeb610c013d032e",
		"03479cd4022fa8c4ebc16c7cd4fc9396c4cf6ebc37daed1ac09803390026785e1a",
	},
	NumRequiredSigs: 5,
	TotalSigs:       7,
	ChainParam:      &chaincfg.MainNetParams,
}

// GenerateOTMultisigAddress returns the Bech32 P2WSH multisig address for each NOS address
func GenerateOTMultisigAddress(bitcoinParams *BitcoinParams, nosAddress string) (string, error) {
	if bitcoinParams == nil {
		return "", fmt.Errorf("Invalid config bitcoin params")
	}

	numSigsRequired := bitcoinParams.NumRequiredSigs

	masterPubKeys := [][]byte{}
	for i, pubKey := range bitcoinParams.MasterPubKeys {
		pubKeyBytes, err := hex.DecodeString(pubKey)
		if err != nil {
			return "", fmt.Errorf("Master BTC Public Key (#%v) %v is invalid - Error %v", i, pubKey, err)
		}
		masterPubKeys = append(masterPubKeys, pubKeyBytes)
	}

	if len(masterPubKeys) < numSigsRequired || numSigsRequired < 0 {
		return "", fmt.Errorf("Invalid signature requirement")
	}

	nosAddress = strings.ToLower(strings.TrimPrefix(nosAddress, "0x"))

	pubKeys := [][]byte{}
	if nosAddress == "" {
		pubKeys = masterPubKeys[:]
	} else {
		chainCode := chainhash.HashB([]byte(nosAddress))
		for idx, masterPubKey := range masterPubKeys {
			extendedBTCPublicKey := hdkeychain.NewExtendedKey(bitcoinParams.ChainParam.HDPublicKeyID[:], masterPubKey, chainCode, []byte{}, 0, 0, false)
			extendedBTCChildPubKey, _ := extendedBTCPublicKey.Derive(0)
			childPubKey, err := extendedBTCChildPubKey.ECPubKey()
			if err != nil {
				return "", fmt.Errorf("Master BTC Public Key (#%v) %v is invalid - Error %v", idx, masterPubKey, err)
			}
			pubKeys = append(pubKeys, childPubKey.SerializeCompressed())
		}
	}

	// create redeem script for m of n multi-sig
	builder := txscript.NewScriptBuilder()
	// add the minimum number of needed signatures
	builder.AddOp(byte(txscript.OP_1 - 1 + numSigsRequired))
	// add the public key to redeem script
	for _, pubKey := range pubKeys {
		builder.AddData(pubKey)
	}
	// add the total number of public keys in the multi-sig script
	builder.AddOp(byte(txscript.OP_1 - 1 + len(pubKeys)))
	// add the check-multi-sig op-code
	builder.AddOp(txscript.OP_CHECKMULTISIG)

	redeemScript, err := builder.Script()
	if err != nil {
		return "", fmt.Errorf("Could not build script - Error %v", err)
	}

	// generate P2WSH address
	scriptHash := sha256.Sum256(redeemScript)
	addr, err := btcutil.NewAddressWitnessScriptHash(scriptHash[:], bitcoinParams.ChainParam)
	if err != nil {
		return "", fmt.Errorf("Could not generate address from script - Error %v", err)
	}
	addrStr := addr.EncodeAddress()

	return addrStr, nil
}

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter the nos address in the argument.")
		return
	}

	nosAddress := args[1]
	multisigAddress, err := GenerateOTMultisigAddress(BitcoinParamsNOS, nosAddress)
	if err != nil {
		fmt.Printf("Run test error: %v\n", err)
	} else {
		fmt.Printf("\n================== Generate multisig address successfully ==================\n")
		fmt.Printf("\nNOS address: %v\n", nosAddress)
		fmt.Printf("\nThe corresponding multisig address: %v\n", multisigAddress)
		fmt.Printf("\n============================================================================\n")
	}
}
