package retry

import "time"

func Do(times int, waitFor time.Duration, action func() error) error {
	var lastError error
	for i := 0; i < times; i++ {
		err := action()
		if err == nil {
			return nil
		}

		lastError = err

		time.Sleep(waitFor)
	}

	return lastError
}
