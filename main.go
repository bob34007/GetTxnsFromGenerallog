package main

import (
	"errors"
	"os"
)

func main() {
	var infile string = ""
	var outfile string = ""
	var err error
	var key string = ""
	/*
		if len(os.Args) != 4 {
			fmt.Println(" param error")
			return
		}
	*/
	//get the parameters in a simple way
	//fmt.Println("args[%v]:%v", i, v)
	if os.Args[1] == "split" {
		//infile name
		infile = os.Args[2]
		//outfile path
		outfile = os.Args[3]
		err = splitGenerallogByConn(infile, outfile)
		if err != nil {
			panic(err)
		}

	} else if os.Args[1] == "gettxn" {
		infile = os.Args[2]
		outfile = os.Args[3]
		key = os.Args[4]
		err = findTxns(infile, outfile, key)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("error param"))
	}

}
