package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error reading data from Xenia")
	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

type Pillar struct {
	Host    string
	Timeout time.Duration
}

func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

type System struct {
	Xenia
	Pillar
}

func Pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func Store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := Pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := Store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}

}

func main() {

}
