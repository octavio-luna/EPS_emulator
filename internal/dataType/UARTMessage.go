package dataType

import "fmt"

type PayloadCmdComponents struct {
	STID    uint64
	IVID    uint64
	Command uint64
	RC      uint64
	BID     uint64
	Extras  []FieldsValue
}

type PayloadRspComponents struct {
	STID   uint64
	IVID   uint64
	RC     uint64
	BID    uint64
	STAT   uint64
	Extras []FieldsValue
}

// UARTMessage represents the structure for UART messages.
type UARTMessage struct {
	Tag        string // 'cmd' for command, 'rsp' for response
	PayloadCmd PayloadCmdComponents
	PayloadRsp PayloadRspComponents
}

type FieldsValue struct {
	FieldName string
	Value     []uint64
}

type Fields struct {
	FieldName string
	Offset    int
	Size      int
}

//TODO: Complex Datatype: VIPD, BPD, CCD, CCSD need to be defined and implemented

func (m *UARTMessage) ToPayload() string {
	if m.Tag == "cmd" {
		return fmt.Sprintf("%02X%02X%02X%02X%02X", m.PayloadCmd.STID, m.PayloadCmd.IVID, m.PayloadCmd.Command, m.PayloadCmd.RC, m.PayloadCmd.BID)
	}
	if m.Tag == "rsp" {
		p := fmt.Sprintf("%02X%02X%02X%02X%02X", m.PayloadRsp.STID, m.PayloadRsp.IVID, m.PayloadRsp.RC, m.PayloadRsp.BID, m.PayloadRsp.STAT)
		for _, val := range m.PayloadRsp.Extras {
			for _, v := range val.Value {
				p += fmt.Sprintf("%02X", v)
			}
		}
		return p
	}
	return ""
}
