package ops

import (
	"github.com/octavio-luna/EPS_emulator/internal/dataType"
)

func GetSystemStatus(e *dataType.EPS, msg dataType.UARTMessage) (dataType.UARTMessage, error) {
	rsp := e.BuildResponse(msg)

	//TODO: Check if this is the correct way to add the extra fields or I should handle this case in a different way
	extra := e.SystemStatus()
	extraUint64 := make([]uint64, len(extra))
	for i, val := range extra {
		extraUint64[i] = uint64(val)
	}

	rsp.PayloadRsp.Extras = append(rsp.PayloadRsp.Extras, dataType.FieldsValue{
		FieldName: "SYSTEM_STATUS",
		Value:     extraUint64,
	})

	return rsp, nil
}
