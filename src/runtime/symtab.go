package runtime

type Frames struct {
	//
}

type Frame struct {
	Function string

	// Func is the Func value of this call frame. This may be nil
	// for non-Go code or fully inlined functions.
	Func *Func

	File string
	Line int
	PC   uintptr
}

func CallersFrames(callers []uintptr) *Frames {
	return nil
}

func (ci *Frames) Next() (frame Frame, more bool) {
	return Frame{}, false
}
