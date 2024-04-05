package UART

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/octavio-luna/EPS_emulator/internal/constants"
	"github.com/octavio-luna/EPS_emulator/internal/dataType"
)

var STIDMap = map[uint64]string{
	0x00: "Bypass",
	//Electrical Power System
	0x11: "PDU",
	0x12: "PBU",
	0x13: "PCU",
	0x1A: "PIU",
	//Attitude and Orbit Control System
	0x21: "IMTQv3",
	//RF Communication
	0x41: "IMDU",
}

var IVIDMap = map[uint64]string{
	0x00: "LATEST",
	0x07: "CURRENT",
}

var STATMap = map[uint64]string{
	0x00: "ACCEPTED",
	0x01: "REJECTED",
	0x02: "REJECTED: INVALID COMMAND CODE",
	0x03: "REJECTED: PARAMETER MISSING",
	0x04: "REJECTED: PARAMETER INVALID",
	0x05: "REJECTED: UNAVAILABLE IN CURRENT MODE/CONFIGURATION",
	0x06: "REJECTED: INVALID SYSTEM TYPE, INTERFACE VERSION OR BOARD ID",
	0x07: "INTERNAL ERROR",
	0x80: "Response read for the first time",
}

type CMDDetails struct {
	CommandName  string
	Len          int
	Fields       []dataType.Fields
	ResponseFunc func(*dataType.EPS, dataType.UARTMessage) (dataType.UARTMessage, error)
}

func (c *CMDDetails) GetFieldValue(payload []uint64) dataType.FieldsValue {
	fieldsValue := dataType.FieldsValue{}
	for _, field := range c.Fields {
		fieldsValue.FieldName = field.FieldName
		if field.Offset+field.Size > len(payload) {
			fieldsValue.Value = payload[field.Offset:]
		} else {
			fieldsValue.Value = payload[field.Offset : field.Offset+field.Size]
		}
	}
	return fieldsValue
}

// ConstructUARTMessage constructs a UART message string from a dataType.UARTMessage struct.
func ConstructUARTMessage(msg dataType.UARTMessage) string {
	return fmt.Sprintf("<%s>%s</%s>", msg.Tag, msg.ToPayload(), msg.Tag)
}

// ParseUARTMessage parses a UART message string into aUARTMessage struct.
func ParseUARTMessage(msg []byte) (dataType.UARTMessage, error) {
	//Msg is a byte array, ithas a form of "<tag>payload</tag>" where the payload is a serie of hex values
	strMsg := string(msg)
	if strings.Contains(strMsg, "<cmd>") && strings.Contains(strMsg, "</cmd>") {
		fmt.Println("Parsing cmd message: ", strMsg)
		init_in := strings.Index(strMsg, "<cmd>")
		fin_id := strings.Index(strMsg, "</cmd>")
		extractedPayload := strMsg[init_in+5 : fin_id]
		pay, err := ExtractPayloadCmd(extractedPayload)
		if err != nil {
			return dataType.UARTMessage{
				Tag: "cmd",
			}, err
		}
		msg := dataType.UARTMessage{
			Tag:        "cmd",
			PayloadCmd: pay,
		}
		return msg, nil
	}
	if strings.HasPrefix(strMsg, "<rsp>") {
		pay, err := ExtractPayloadRsp(strings.TrimSuffix(strings.TrimPrefix(strMsg, "<rsp>"), "</rsp>"))
		if err != nil {
			return dataType.UARTMessage{
				Tag: "rsp",
			}, err
		}
		msg := dataType.UARTMessage{
			Tag:        "rsp",
			PayloadRsp: pay,
		}
		return msg, nil
	}
	return dataType.UARTMessage{}, nil
}

// Payload is a space-separated string of hex values
func hexToUint64(hex string) uint64 {
	value, _ := strconv.ParseUint(hex, 16, 64)
	return value
}

func ExtractPayloadCmd(p string) (payloadCmd dataType.PayloadCmdComponents, err error) {
	// Split the payload string into series of two hex values
	fields := make([]string, 0)
	for i := 0; i < len(p); i += 2 {
		if i+2 <= len(p) {
			fields = append(fields, p[i:i+2])
		}
	}
	fieldsHex := make([]uint64, 0)
	for _, val := range fields {
		fieldsHex = append(fieldsHex, hexToUint64(val))
	}

	// Extract the payload components
	payloadCmd.STID = fieldsHex[0]
	payloadCmd.IVID = fieldsHex[1]
	payloadCmd.Command = fieldsHex[2]
	payloadCmd.BID = fieldsHex[3]

	// Check if the IVID is valid
	if payloadCmd.IVID != constants.IVIDUsed {
		return payloadCmd, constants.ErrInvalidIVID
	}

	cmdDetails, ok := CMDMapDetails[payloadCmd.Command]
	if !ok {
		return payloadCmd, constants.ErrInvalidCmd
	}

	if cmdDetails.Len != -1 && len(fields) != cmdDetails.Len {
		return payloadCmd, constants.ErrInvalidLen
	}

	extraFields := make([]dataType.FieldsValue, 0)
	if cmdDetails.Len == -1 {
		for _, val := range cmdDetails.Fields {
			if val.Size == -1 {
				extraFields = append(extraFields, dataType.FieldsValue{
					FieldName: val.FieldName,
					Value:     fieldsHex[val.Offset:],
				})
			} else {
				extraFields = append(extraFields, dataType.FieldsValue{
					FieldName: val.FieldName,
					Value:     fieldsHex[val.Offset : val.Offset+val.Size],
				})
			}
		}
	} else {
		extraFields = append(extraFields, cmdDetails.GetFieldValue(fieldsHex))
	}

	payloadCmd.Extras = extraFields
	return payloadCmd, nil
}

func ExtractPayloadRsp(p string) (payloadRsp dataType.PayloadRspComponents, err error) {
	// Split the payload string into space-separated hex values
	fields := make([]string, 0)
	for i := 0; i < len(p); i += 2 {
		if i+2 <= len(p) {
			fields = append(fields, p[i:i+2])
		}
	}

	fieldsHex := make([]uint64, 0)
	for _, val := range fields {
		fieldsHex = append(fieldsHex, uint64(hexToUint64(val)))
	}

	// Extract the payload components
	payloadRsp.STID = fieldsHex[0]
	payloadRsp.IVID = fieldsHex[1]
	payloadRsp.RC = fieldsHex[2]
	payloadRsp.BID = fieldsHex[3]
	payloadRsp.STAT = fieldsHex[4]

	//TODO: Implement the logic for this
	return payloadRsp, nil
}
