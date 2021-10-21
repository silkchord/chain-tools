package mocks

import (
	"github.com/songqianba/chain-tools/signer/element"
)

type MockData struct {
	MemeChar     string
	Address      string
	Instructions []element.SigningInstructionTemp
	RawTx        string
}

// MockMap ...
var MockMap = map[string]MockData{
	"2c1826322ec81743e7932fe856ac2c27017de1fa0fbb3b6e1969210a7df3e3bc": {
		Instructions: LingerTestInstructions,
		RawTx:        RawTx,
		MemeChar:     "salon invest obtain tomato submit jelly clump drum grape usual heavy wear",
		Address:      "tn1qd4eqqn4el57asjjps9hewwxxe3fsgv4ckkvjm9",
	},
}

// LingerTestInstructions LingerTestInstructions for test
var LingerTestInstructions = []element.SigningInstructionTemp{
	&element.SigningInstruction{
		DerivationPath: []string{
			"2c000000",
			"99000000",
			"01000000",
			"00000000",
			"01000000",
		},
		SignData: []string{"638c129da58d2c0069ce70e64522a08917f4b153b000658b43c9b07f676fb36f"},
		Pubkey:   "28c4340e6338691e0d9a061510d376a2eec5b87267f5a871dc6ea8b5b0a38ccf",
	},
}

var RawTx = "070100010161015f853fd39851e88fdf026384f0977f340251b75758afe671817451074190b7e9e2ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80c8afa02501011600146d72004eb9fd3dd84a41816f9738c6cc530432b80022012028c4340e6338691e0d9a061510d376a2eec5b87267f5a871dc6ea8b5b0a38ccf02010049ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80c2d72f01220020ff05b61619aff6545d79dba2fd5a856cee730bf7f2729e517d8d336ac3529d9f000001003effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80d9f5eb24011600146d72004eb9fd3dd84a41816f9738c6cc530432b80000"
