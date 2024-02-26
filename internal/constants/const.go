package constants

// Define constants for command types
const (
	WatchDogTimeout = 60 //seconds
	IVIDUsed        = 0x07
	ResetKey        = 0x5A

	//STID
	STID_ACCEPTED                                                   uint64 = 0x00
	STID_REJECTED                                                   uint64 = 0x01
	STID_REJECTED_INVALID_COMMAND_CODE                              uint64 = 0x02
	STID_REJECTED_PARAMETER_MISSING                                 uint64 = 0x03
	STID_REJECTED_PARAMETER_INVALID                                 uint64 = 0x04
	STID_REJECTED_UNAVAILABLE_IN_CURRENT_MODE_OR_CONFIG             uint64 = 0x05
	STID_REJECTED_INVALID_SYSTEM_TYPE_INTERFACE_VERSION_OR_BOARD_ID uint64 = 0x06
	STID_INTERNAL_ERROR                                             uint64 = 0x07
	STID_RESPONSE_READ_FOR_THE_FIRST_TIME                           uint64 = 0x80

	OpModeOff                 = "off"
	OpModeStartup             = "startup"
	OpModeNominal             = "nominal"
	OpModeSafety              = "safety"
	OpModeEmergencyLowBattery = "emergency_low_battery"
)

// Custom errors
type CustomError struct {
	Err string
}

func (e *CustomError) Error() string {
	return e.Err
}

var (
	ErrInvalidCmd                   = &CustomError{"Invalid command"}
	ErrInvalidLen                   = &CustomError{"Invalid length"}
	ErrInvalidTag                   = &CustomError{"Invalid tag"}
	ErrInvalidIVID                  = &CustomError{"Invalid IVID"}
	ErrCmdHandler                   = &CustomError{"Command handler not found"}
	ErrCmdNotAvailableInCurrentMode = &CustomError{"Command not available in current mode"}
)
