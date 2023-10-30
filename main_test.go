package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTMultisigAddress(t *testing.T) {
	nosAddress := "0xA73795E3caaED8F37c92530Fb939175054927175"
	multisigAddress, err := GenerateOTMultisigAddress(BitcoinParamsNOS, nosAddress)

	assert.Equal(t, nil, err)

	fmt.Printf("NOS address: %v\n", nosAddress)
	fmt.Printf("The corresponding multisig address: %v\n", multisigAddress)
}
