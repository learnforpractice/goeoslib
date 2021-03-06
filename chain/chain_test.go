package chain

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestUnpack(t *testing.T) {
	rawTx, _ := hex.DecodeString("11faa561a8198cb93a22000000000120428a97721aa36a000000204fb5d5a401305d67594ed7599600000000a8ed32326f020000000000000018a620338845455fbcde0e205ef4da44965e5c6e434c3fa9b780c50f43cd955c01f32bb6ab1a4d44bb892ff0129239df890100000000350c000000000000000000000000001830303030313961386161333166333363386362393361323258836db6fb4abc160000")
	tx := &Transaction{}
	tx.Unpack(rawTx)
	s, _ := json.MarshalIndent(tx, "", "  ")
	t.Logf("%s", s)
}

func TestChainApi(t *testing.T) {
	chain := NewChainApi("http://127.0.0.1:8888")
	block, err := chain.GetBlockTrace(100)
	if err != nil {
		t.Fatal(err)
	}

	value, err := block.GetString("status")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("++++++%v", value)
}
