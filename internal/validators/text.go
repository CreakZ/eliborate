package validators

import "fmt"

func ValidateTextQuery(text string) error {
	if text == "" {
		return fmt.Errorf("empty query string provided")
	}

	return nil
}
