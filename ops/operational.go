package ops

import (
	"errors"

	"github.com/octavio-luna/EPS_emulator/internal/constants"
	"github.com/octavio-luna/EPS_emulator/internal/dataType"
)

func SystemReset(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)

	if msg.PayloadCmd.Extras[0].Value[0] != constants.ResetKey { // Page 29
		rsp.PayloadRsp.STAT = constants.STID_REJECTED
		return rsp, errors.New("invalid reset key")
	}

	//TODO: Implement the logic for simulating the system reset
	return rsp, nil
}

func NoOperation(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)
	return rsp, nil
}

func CancelOperation(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)
	//check if there's any CEOBC=true
	for id, v := range e.CEOBC {
		if v {
			e.CEOBC[id] = false
		}
	}

	return rsp, nil
}

func WatchDogCommand(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)
	return rsp, nil
}
