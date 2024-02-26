package state

import (
	"github.com/octavio-luna/EPS_emulator/UART"
	"github.com/octavio-luna/EPS_emulator/internal/constants"
	"github.com/octavio-luna/EPS_emulator/internal/dataType"
)

// Representation of the Electrical Power System and its components
type EPSAPI struct {
	E *dataType.EPS
}

func NewEps() EPSAPI {
	return EPSAPI{
		E: &dataType.EPS{
			OpMode: "off",
			STID:   0x00, //TODO: Define the STID for the EPS
			CEOBC:  make(map[int]bool),
		},
	}
}

// ProcessCommand processes a received UART command and generates a response.
func (E *EPSAPI) ProcessCommand(msg dataType.UARTMessage) (response dataType.UARTMessage, err error) {
	resp := E.E.BuildResponse(msg)
	if msg.Tag != "cmd" {
		resp.PayloadRsp.STAT = constants.STID_REJECTED
		return resp, constants.ErrInvalidTag
	}

	cmdDetails, ok := UART.CMDMapDetails[msg.PayloadCmd.Command]
	if !ok {
		resp.PayloadRsp.STAT = constants.STID_REJECTED_INVALID_COMMAND_CODE
		return resp, constants.ErrInvalidCmd
	}

	if cmdDetails.ResponseFunc == nil {
		resp.PayloadRsp.STAT = constants.STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG
		return resp, constants.ErrCmdHandler
	}

	return cmdDetails.ResponseFunc(E.E, msg)
}
