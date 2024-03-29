package main

type Dictionary map[string]string

type DictErr string

const (
	ErrNotFound         = DictErr("could not find the word you were looking for")
	ErrWorldExist       = DictErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictErr("cannot update word because it does not exist")
)

func (e DictErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if _, ok := d[word]; ok {
		return d[word], nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case ErrWorldExist:
		return ErrWorldExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
