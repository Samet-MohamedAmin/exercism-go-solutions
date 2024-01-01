package erratum

// Use performs some operation Frob(input) on the resource returned by calling its
// ResourceOpener o. The function will keep trying to open the resource in the case
// that calling o returns a TransientError. If calling Frob(input) panics with a
// FrobError, the resource will be Defrob'd with the FrobError's defrobTag. The
// resource is guaranteed to be closed after Use returns.
// description from: https://exercism.io/tracks/go/exercises/error-handling/solutions/a228d124c605489d8466f55fb1b7af88
func Use(o ResourceOpener, input string) error {
	r, err := open(o)
	if err != nil {
		return err
	}
	defer r.Close()
	return frob(r, input)
}

func open(o ResourceOpener) (r Resource, err error) {
	for ok := true; ok; {
		r, err = o()
		_, ok = err.(TransientError)
	}

	return
}

func frob(r Resource, input string) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch e := rec.(type) {
			case FrobError:
				r.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()
	r.Frob(input)

	return
}
