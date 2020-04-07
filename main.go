package main

import (
	"fmt"
	"log"
)

type raizError struct {
	lat, long string
	err error
}

func (re raizError) Error() string {
	return fmt.Sprintf("Error matemático: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := raiz(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("Error sin formatear")
		e := fmt.Errorf("(Error con formato :::: %v)", f)
		return 0, raizError{
			lat:  "50.548 N",
			long: "89.374 W",
			err:  e,
		}
	}
	return 42, nil
}
