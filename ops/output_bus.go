package ops

import (
	"errors"

	"github.com/octavio-luna/EPS_emulator/internal/constants"
	"github.com/octavio-luna/EPS_emulator/internal/dataType"
)

func OutputBusGroupOn(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)
	if e.OpMode != constants.OpModeNominal {
		rsp.PayloadRsp.STAT = constants.STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG
		return rsp, errors.New("Command not available in current mode")
	}

	//convert any extra fields to a bit array
	bitLevel := make([]bool, 0)
	for _, val := range msg.PayloadCmd.Extras {
		for _, v := range val.Value {
			//v is a 2 byte value, convert it to a bit array
			for i := 0; i < 16; i++ {
				bitLevel = append(bitLevel, v&(1<<i) != 0)
			}
		}
	}

	for i, v := range bitLevel {
		if v {
			e.CEOBC[i] = true
		}
	}
	//TODO: We can add logic to check if all bus ids are in range

	return rsp, nil
}

func OutputBusGroupOff(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)

	if e.OpMode != constants.OpModeNominal {
		rsp.PayloadRsp.STAT = constants.STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG
		return rsp, errors.New("Command not available in current mode")
	}

	//convert any extra fields to a bit array
	bitLevel := make([]bool, 0)
	for _, val := range msg.PayloadCmd.Extras {
		for _, v := range val.Value {
			//v is a 2 byte value, convert it to a bit array
			for i := 0; i < 16; i++ {
				bitLevel = append(bitLevel, v&(1<<i) != 0)
			}
		}
	}

	for i, v := range bitLevel {
		if v {
			e.CEOBC[i] = false
		}
	}

	//TODO: We can add logic to check if all bus ids are in range

	return rsp, nil
}

func OutputBusGroupState(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)

	if e.OpMode != constants.OpModeNominal {
		rsp.PayloadRsp.STAT = constants.STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG
		return rsp, errors.New("Command not available in current mode")
	}

	//convert any extra fields to a bit array
	bitLevel := make([]bool, 0)
	for _, val := range msg.PayloadCmd.Extras {
		for _, v := range val.Value {
			//v is a 2 byte value, convert it to a bit array
			for i := 0; i < 16; i++ {
				bitLevel = append(bitLevel, v&(1<<i) != 0)
			}
		}
	}

	for i, v := range bitLevel {
		e.CEOBC[i] = v
	}

	//TODO: We can add logic to check if all bus ids are in range
	return rsp, nil
}

func OutputBusChannelOnSingle(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)

	if e.OpMode != constants.OpModeNominal {
		rsp.PayloadRsp.STAT = constants.STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG
		return rsp, errors.New("Command not available in current mode")
	}

	e.CEOBC[int(msg.PayloadCmd.Extras[0].Value[0])] = true
	return rsp, nil
}

func OutputBusChannelOffSingle(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)

	if e.OpMode != constants.OpModeNominal {
		rsp.PayloadRsp.STAT = constants.STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG
		return rsp, errors.New("Command not available in current mode")
	}

	e.CEOBC[int(msg.PayloadCmd.Extras[0].Value[0])] = false
	return rsp, nil
}

func SwitchToNominalMode(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)
	e.SetOpMode(constants.OpModeNominal) //TODO: We can simulate a reject here simulating low battery
	return rsp, nil
}

func SwitchToSafetyMode(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)
	e.SetOpMode(constants.OpModeSafety)
	return rsp, nil
}
