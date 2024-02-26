package dataType

import (
	"fmt"

	"github.com/octavio-luna/EPS_emulator/internal/constants"
)

type EPS struct {
	STID uint64 //System Type Identifier

	OpMode string //off, startup, nominal, safety, emergency_low_battery

	CEOBC map[int]bool //command-enable output bus channel

	CONF bool //this flag indicates if the parameters have been changed by eps after the last read/write PAGE: 40

	Reset_cause uint64 //this flag indicates the cause of last reset PAGE: 41

	Uptime uint32 //uptime since system start PAGE 41

	FirstError uint16 //First internal error during system control cycle PAGE 41

	RC_CNT_PWRON  uint16 //total restarts since beginning of liefe, should be non-volatile Power On PAGE 41
	RC_CNT_WDG    uint16 //Reset Cause Counter: Watchdog PAGE 41
	RC_CNT_CMD    uint16 //Reset Cause Counter: Commanded resets PAGE 41
	RC_CNT_MCU    uint16 //Reset Cause Counter: EPS controller resets PAGE 41
	RC_CNT_EMLOPO uint16 //Reset Cause Counter: Emergency Low Power PAGE 41

	PREVCMD_ELAPSED uint16 //Elapsed time since last command PAGE 41

	UNIX_TIME   uint32 //seconds since 1970-01-01 00:00:00 PAGE 41
	UNIX_YEAR   uint8  //years since 2000
	UNIX_MONTH  uint8  //calendar month of unix time PAGE 42
	UNIX_DAY    uint8  //calendar day of unix time PAGE 42
	UNIX_HOUR   uint8  //calendar hour of unix time PAGE 42
	UNIX_MINUTE uint8  //calendar minute of unix time PAGE 42
	UNIX_SECOND uint8  //calendar second of unix time PAGE 42
}

func (e *EPS) BuildResponse(m UARTMessage) UARTMessage {
	return UARTMessage{
		Tag: "rsp",
		PayloadRsp: PayloadRspComponents{
			STID: e.STID, //TODO: Check if this is the correct STID to use here
			IVID: constants.IVIDUsed,
			RC:   m.PayloadCmd.Command + 1,
			BID:  m.PayloadCmd.BID, //TODO: Check if this is the correct BID to use here
			STAT: constants.STID_ACCEPTED,
		},
	}
}

func (eps *EPS) SetOpMode(opMode string) {
	switch opMode {
	case constants.OpModeOff:
		eps.OpMode = constants.OpModeOff
	case constants.OpModeStartup:
		eps.OpMode = constants.OpModeStartup
	case constants.OpModeNominal:
		eps.OpMode = constants.OpModeNominal
	case constants.OpModeSafety:
		eps.OpMode = constants.OpModeSafety
	case constants.OpModeEmergencyLowBattery:
		eps.OpMode = constants.OpModeEmergencyLowBattery
	default:
		fmt.Println("Invalid mode")
		eps.OpMode = "nominal"
	}
}

func (e *EPS) SystemStatus() []byte {
	fields := make([]byte, 0)
	fields = append(fields, 0x00) // MODE

	switch e.OpMode {
	case "startup":
		fields[0] = 0
	case "nominal":
		fields[0] = 1
	case "safety":
		fields[0] = 2
	case "emergency_low_battery":
		fields[0] = 3
	default:
		fields[0] = 1
	}

	fields = append(fields, 0x00) // CONF
	if e.CONF {
		fields[1] = 1
	} else {
		fields[1] = 0
	}

	fields = append(fields, byte(e.Reset_cause)) // RESET_CAUSE

	fields = append(fields, byte(e.Uptime)) // UPTIME

	fields = append(fields, byte(e.FirstError)) // FIRST_ERROR

	fields = append(fields, byte(e.RC_CNT_PWRON)) // RC_CNT_PWRON

	fields = append(fields, byte(e.RC_CNT_WDG)) // RC_CNT_WDG

	fields = append(fields, byte(e.RC_CNT_CMD)) // RC_CNT_CMD

	fields = append(fields, byte(e.RC_CNT_MCU)) // RC_CNT_MCU

	fields = append(fields, byte(e.RC_CNT_EMLOPO)) // RC_CNT_EMLOPO

	fields = append(fields, byte(e.PREVCMD_ELAPSED)) // PREVCMD_ELAPSED

	fields = append(fields, byte(e.UNIX_TIME)) // UNIX_TIME

	fields = append(fields, byte(e.UNIX_YEAR)) // UNIX_YEAR

	fields = append(fields, byte(e.UNIX_MONTH)) // UNIX_MONTH

	fields = append(fields, byte(e.UNIX_DAY)) // UNIX_DAY

	fields = append(fields, byte(e.UNIX_HOUR)) // UNIX_HOUR

	fields = append(fields, byte(e.UNIX_MINUTE)) // UNIX_MINUTE

	fields = append(fields, byte(e.UNIX_SECOND)) // UNIX_SECOND

	//TODO: Double check if this way is right
	return fields
}
