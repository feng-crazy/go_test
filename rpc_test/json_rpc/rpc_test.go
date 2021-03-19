package utils

import (
	"errors"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith struct {
}

func (t *Arith) Multiply(args []Args, reply *int) error {
	for _, arg := range args {
		*reply = *reply + (arg.A * arg.B)
	}

	return nil
}
func (t *Arith) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero!")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func TestRpcServer(t *testing.T) {
	rpcserver := RpcServer{}
	defer rpcserver.DeInit()
	arith := new(Arith)
	if err := rpcserver.Init(":9999", arith); err != nil {
		log.Error(err)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func TestRpcClient(t *testing.T) {
	rpcclient := RpcClient{}
	if err := rpcclient.Connect(":9999"); err != nil {
		log.Error(err)
	}

	for {
		// args := Args{7, 8}

		arg1 := map[string]interface{}{
			"a": 7,
			"b": 8,
		}

		arg2 := map[string]interface{}{
			"a": 1,
			"b": 2,
		}

		args := []map[string]interface{}{
			arg1,
			arg2,
		}

		var reply int
		if err := rpcclient.Call("Arith.Multiply", args, &reply, 2); err != nil {
			log.Error(err)
		}
		log.Info("====>", reply)
		time.Sleep(1 * time.Second)
	}
}
