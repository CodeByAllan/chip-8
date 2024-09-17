package keyboard

type keyStorage interface {
	SetKey(index byte, value byte)
	ClearKeys()
}
