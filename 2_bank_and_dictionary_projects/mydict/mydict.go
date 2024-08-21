package mydict

import (
	"errors"
)

// 2-4. dictionary 1
// Dictionary type
type Dictionary map[string]string // 별명(alias), map[string]string에 대한 가명
// type Money interface
// Money(1) -> 1
// type에 method를 추가할 수 있다.

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists.")
	errCantUpdate = errors.New("Can't Update non-existing word. ")
)

// search for a word
func (d Dictionary) Search(word string) (string, error) { // 탐색한 단어와 탐색 결과가 없을 경우의 에러
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// 2-5. add method
// add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word) // 탐색한 결과 에러가 난다면(= 단어가 없다면) 추가할 수 있다.
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	/* 	if err == errNotFound{
	   		d[word] = def
	   	} else if err == nil {
	   		return errWordExists
	   	}
	   	return nil */
}

// 2-6. update delete
// update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word) // d(dictionary)에서 word를 삭제한다.
	case errNotFound:
		return errNotFound
	}
	return nil
}
