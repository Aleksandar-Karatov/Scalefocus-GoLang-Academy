package main

import (
	"errors"
	"fmt"
	"log"
)

type Action func() error

func main() {

	err := SafeExec(doNoErrors)
	if err != nil {
		log.Fatal(err())
	} else {
		fmt.Println("No errors :D")
	}
	err = SafeExec(doAnError) // to test the handling of panics the section for doAnError has to be put in a coment and
	if err != nil {           // the section for doAPanic has to be uncommented
		log.Fatal(err())
	}
	// err := SafeExec(doAPanic)
	// if err != nil {
	// 	log.Fatal(err())
	// }

}

func SafeExec(a Action) Action {
	var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Errorf("Recovered : %s", r))
		}
	}()
	if r := recover(); r == nil {
		if a() != nil {
			err = fmt.Errorf("safe exec: %w", a())
		} else {
			return nil
		}

	}
	if err != nil {

		return func() error { return err }
	}
	return nil

}

func doAPanic() error {
	panic("bOOOOOm")
}

func doAnError() error {
	return errors.New("This is a test error")
}
func doNoErrors() error {
	return nil
}
