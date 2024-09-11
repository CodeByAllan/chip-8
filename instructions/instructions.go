package instructions

func ClearDisplay(Screen *[64 * 32]byte) {
	for i := range Screen {
		Screen[i] = 0
	}
}
