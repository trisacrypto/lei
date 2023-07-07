package lei

type LEI string

func Parse(s string) (LEI, error) {
	entity := LEI(s)
	if err := entity.Check(); err != nil {
		return LEI(""), err
	}
	return entity, nil
}

func (s LEI) Check() error {
	if len(s) != 20 {
		return InvalidLength(len(s))
	}

	// TODO: validate checksum

	return nil
}

func Random() LEI {
	return LEI("")
}

func Mod97(s string) (int, error) {
	return 0, nil
}
