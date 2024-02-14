package customerrors

import "fmt"

type RateLimitError struct {
	Limit int
}

func (rle RateLimitError) Error() string {
	return fmt.Sprintf("Limit of %d request exceded", rle.Limit)
}
