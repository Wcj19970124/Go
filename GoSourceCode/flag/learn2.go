package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type interval []time.Duration

//实现String接口
func (i *interval) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval already exists")
	}

	for _, d := range strings.Split(value, ",") {
		t, err := time.ParseDuration(d)
		if err != nil {
			return err
		}
		*i = append(*i, t)
	}
	return nil
}

var flagTime interval

func init() {
	flag.Var(&flagTime, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {

	flag.Parse()
	fmt.Println(flagTime)

}
