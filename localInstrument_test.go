package wire

// mockLocalInstrument creates a LocalInstrument
func mockLocalInstrument() *LocalInstrument {
	li := NewLocalInstrument()
	li.LocalInstrumentCode = "ANSI"
	li.ProprietaryCode = ""
	return li
}
