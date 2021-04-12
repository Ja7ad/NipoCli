package module

import "errors"

// Nipo Promoter
const nipoPromoter = "nipo >>> "

// Global Errors
var (
	errInput     = errors.New("Incorrect your input, Please try 'help'")
	errInterrupt = errors.New("No interrupt handler")
)

// Shell is interactive struct manage
type Shell struct {
}
