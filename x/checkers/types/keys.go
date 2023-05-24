package types

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey             = "SystemInfo/value/"
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventBlack     = "black"            // Is it relevant to me?
	GameCreatedEventRed       = "red"              // Is it relevant to me?

	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
)
