package filters

type filter int

type (
	message       filter
	callbackQuery filter
)

var (
	Message       = new(message)
	CallbackQuery = new(callbackQuery)
)
