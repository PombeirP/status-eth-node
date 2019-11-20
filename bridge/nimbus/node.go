// +build nimbus

package nimbusbridge

// https://golang.org/cmd/cgo/

/*
#include <stddef.h>
#include <stdbool.h>
#include <stdlib.h>
#include <libnimbus.h>
*/
import "C"
import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/status-im/status-eth-node/types"
	"go.uber.org/zap"
)

func Init() {
	runtime.LockOSThread()
}

func StartNimbus(privateKey *ecdsa.PrivateKey, listenAddr string, staging bool) error {
	C.NimMain()

	port, err := strconv.Atoi(strings.Split(listenAddr, ":")[1])
	if err != nil {
		return fmt.Errorf("failed to parse port number from %s", listenAddr)
	}

	privateKeyC := C.CBytes(crypto.FromECDSA(privateKey))
	defer C.free(privateKeyC)
	if !C.nimbus_start(C.ushort(port), true, false, 0.002, (*C.uchar)(privateKeyC), C.bool(staging)) {
		return errors.New("failed to start Nimbus node")
	}

	return nil
}

type nimbusNodeWrapper struct {
}

func NewNodeBridge() types.Node {
	return &nimbusNodeWrapper{}
}

func (w *nimbusNodeWrapper) NewENSVerifier(_ *zap.Logger) enstypes.ENSVerifier {
	panic("not implemented")
}

func (w *nimbusNodeWrapper) NewWhisper() (types.Whisper, error) {
	return NewNimbusWhisperWrapper(), nil
}
