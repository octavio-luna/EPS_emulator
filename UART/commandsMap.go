package UART

import (
	"github.com/octavio-luna/EPS_emulator/internal/dataType"
	"github.com/octavio-luna/EPS_emulator/ops"
)

var CMDMapDetails = map[uint64]CMDDetails{
	//Operational commands
	0xAA: {
		CommandName: "System reset",
		Len:         5,
		Fields: []dataType.Fields{
			{"RESET_KEY", 4, 1},
		},
		ResponseFunc: ops.SystemReset,
	},
	0x02: {
		CommandName:  "No-operation",
		Len:          4,
		ResponseFunc: ops.NoOperation,
	},
	0x04: {
		CommandName:  "Cancel operation",
		Len:          4,
		ResponseFunc: ops.CancelOperation,
	},
	0x06: {
		CommandName:  "Watchdog",
		Len:          4,
		ResponseFunc: ops.WatchDogCommand,
	},
	0x10: { //TODO: Table says itcan be 6 or 8 bytes
		CommandName: "Output Bus Group On",
		Len:         8,
		Fields: []dataType.Fields{
			{"CH_BF", 4, 2},
			{"CH_EXT_BF", 6, 2},
		},
		ResponseFunc: ops.OutputBusGroupOn,
	},
	0x12: { //TODO: Table says itcan be 6 or 8 bytes
		CommandName: "Output Bus Group Off",
		Len:         8,
		Fields: []dataType.Fields{
			{"CH_BF", 4, 2},
			{"CH_EXT_BF", 6, 2},
		},
		ResponseFunc: ops.OutputBusGroupOff,
	},
	0x14: { //TODO: Table says itcan be 6 or 8 bytes
		CommandName: "Output Bus Group State",
		Len:         8,
		Fields: []dataType.Fields{
			{"CH_BF", 4, 2},
			{"CH_EXT_BF", 6, 2},
		},
		ResponseFunc: ops.OutputBusGroupState,
	},
	0x16: {
		CommandName: "Output Bus Channel On",
		Len:         5,
		Fields: []dataType.Fields{
			{"CH_IDX", 4, 1},
		},
		ResponseFunc: ops.OutputBusChannelOnSingle,
	},
	0x18: {
		CommandName: "Output Bus Channel Off",
		Len:         5,
		Fields: []dataType.Fields{
			{"CH_IDX", 4, 1},
		},
		ResponseFunc: ops.OutputBusChannelOffSingle,
	},
	//Mode switching commands
	0x30: {
		CommandName:  "Switch to Nominal Mode",
		Len:          4,
		ResponseFunc: ops.SwitchToNominalMode,
	},
	0x32: {
		CommandName:  "Switch to Safety Mode",
		Len:          4,
		ResponseFunc: ops.SwitchToSafetyMode,
	},
	//Data request commands
	0x40: {
		CommandName:  "Get System Status",
		Len:          4,
		ResponseFunc: ops.GetSystemStatus,
	},
	0x42: {
		CommandName: "Get PDU/PIU Overcurrent Fault Status",
		Len:         4,
	},
	0x44: {
		CommandName: "Get PBU ABF Placed State",
		Len:         4,
	},
	0x50: {
		CommandName: "Get PDU Housekeeping Data(Raw)",
		Len:         4,
	},
	0x52: {
		CommandName: "Get PDU Housekeeping Data(Engineering)",
		Len:         4,
	},
	0x54: {
		CommandName: "Get PDU Housekeeping Data(Running Average)",
		Len:         4,
	},
	0x60: {
		CommandName: "Get PBU Housekeeping Data(Raw)",
		Len:         4,
	},
	0x62: {
		CommandName: "Get PBU Housekeeping Data(Engineering)",
		Len:         4,
	},
	0x64: {
		CommandName: "Get PBU Housekeeping Data(Running Average)",
		Len:         4,
	},
	0x70: {
		CommandName: "Get PCU Housekeeping Data(Raw)",
		Len:         4,
	},
	0x72: {
		CommandName: "Get PCU Housekeeping Data(Engineering)",
		Len:         4,
	},
	0x74: {
		CommandName: "Get PCU Housekeeping Data(Running Average)",
		Len:         4,
	},
	//Configuration commands
	0x82: {
		CommandName: "Get Configuration Parameter",
		Len:         6,
		Fields: []dataType.Fields{
			{"PAR_ID", 4, 2},
		},
	},
	//TODO: Add magic because 0x84 is 6 bytes + 1-8 parameter bytes
	0x84: {
		CommandName: "Set Configuration Parameter",
		Len:         -1,
		Fields: []dataType.Fields{
			{"PAR_ID", 4, 2},
			{"PAR_VAL", 6, -1}, // This is a variable length field
		},
	},
	0x86: {
		CommandName: "Reset Configuration Parameter",
		Len:         6,
		Fields: []dataType.Fields{
			{"PAR_ID", 4, 2},
		},
	},
	0x90: {
		CommandName: "Reset Configuration",
		Len:         5,
		Fields: []dataType.Fields{
			{"CONF_KEY", 4, 1},
		},
	},
	0x92: {
		CommandName: "Load Configuration",
		Len:         5,
		Fields: []dataType.Fields{
			{"CONF_KEY", 4, 1},
		},
	},
	0x94: {
		CommandName: "Save Configuration",
		Len:         7,
		Fields: []dataType.Fields{
			{"CONF_KEY", 4, 1},
			{"CHECKSUM", 5, 2},
		},
	},
	//Data Request Commands
	0xA0: {
		CommandName: "Get PIU Housekeeping Data(Raw)",
		Len:         4,
	},
	0xA2: {
		CommandName: "Get PIU Housekeeping Data(Engineering)",
		Len:         4,
	},
	0xA4: {
		CommandName: "Get PIU Housekeeping Data(Running Average)",
		Len:         4,
	},
	//Other commands
	0xC4: {
		CommandName: "Correct Time",
		Len:         8,
		Fields: []dataType.Fields{
			{"CORRECTION", 4, 4},
		},
	},
	0xC6: {
		CommandName: "Zero Reset Cause Counters",
		Len:         5,
		Fields: []dataType.Fields{
			{"ZERO_KEY", 4, 1},
		},
	},
}
