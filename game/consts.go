package game

const (
	ScreenWidth  = 256
	ScreenHeight = 240
)

const (
	BlockWidth  = 10
	BlockHeight = 10
)

const (
	FieldBlockCountX = 10 // Number of blocks horizontally
	FieldBlockCountY = 20 // Number of blocks vertically

	FieldWidth  = BlockWidth * FieldBlockCountX
	FieldHeight = BlockHeight * FieldBlockCountY

	TopMargin = (FieldHeight-BlockHeight*6)/3 - BlockHeight*3 // Vertical gap between label sections
)

const (
	InputInitialDelay = 10 // Delay before auto-repeat starts
	InputRepeatRate   = 3  // Repeat every few frames
)

const (
	BaseDropInterval = 60 // Base drop interval in frames at level 0
	MinDropInterval  = 3  // Fastest possible drop interval
	LinesPerLevel    = 10 // Number of cleared lines to increase one level
)
