package slices2

// FromChannel drains the channel ch into a slice.
//
// It collects ALL the values from channel ch until the channel is closed.
func FromChannel[E any](ch <-chan E) []E {
	if ch == nil {
		return nil
	}

	var out []E // capacity is unknown
	for v := range ch {
		out = append(out, v)
	}

	return out
}

// ToChannel converts a slice to read-only channel.
//
// All values from slice s will be sent to new channel.
func ToChannel[S ~[]E, E any](s S, bufferSize int) <-chan E {
	ch := make(chan E, bufferSize)

	go func() {
		defer close(ch) // close later

		for _, v := range s {
			ch <- v // send to channel
		}
	}()

	return ch
}
