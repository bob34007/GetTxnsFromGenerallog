package main

import (
	"strings"
)

/*
else if strings.LastIndex(v, "autocommit") != -1 && begin == true {
			begin = true
			end = true
		}
*/
func findTxnsSQL(data string, businessName string) []string {
	var txns []string
	var tmp string
	var begin bool = false
	var end bool = false
	var findkey bool = false
	sqls := strings.Split(data, "\n")
	for _, v := range sqls {
		if strings.LastIndex(v, "autocommit=0") != -1 && begin == false {
			tmp = ""
			begin = true
		} else if strings.LastIndex(v, "commit") != -1 && begin == true {
			begin = false
			end = true
		} else if strings.LastIndex(v, businessName) != -1 && begin == true {
			findkey = true
		}
		tmp += v + "\n"
		if end == true && findkey == true {
			txns = append(txns, tmp)
			tmp = ""
			end = false
			findkey = false
		}
	}
	return txns
}

func getTxnHash(txn string) (uint64, error) {
	var hstr string
	sqls := strings.Split(txn, "\n")
	for _, v := range sqls {
		pos := strings.LastIndex(v, "[arguments:")
		if pos == -1 {
			continue
		} else {
			hstr += v[0:pos]
		}
	}
	return HashString(txn)
}

func findTxns(infile, outfile, key string) error {
	var data string
	var m = make(map[uint64]string)
	var err error
	fs, err := getFileListFromDir(infile)
	if err != nil {
		return err
	}
	for _, v := range fs {
		data, err = readAllDataFromFile(infile + "/" + v)
		if err != nil {
			return err
		}
		txns := findTxnsSQL(data, key)
		for _, vv := range txns {
			hid, err := getTxnHash(vv)
			if err != nil {
				return err
			}
			m[hid] = vv
		}
	}

	data = ""
	for _, v := range m {
		data += v
	}
	return writeFile(outfile, data)
}

/*
func findTxns(infile, outfile, key string) error {
	var data string = ""
	var txns string = ""
	var err error
	data, err = readAllDataFromFile(infile)
	if err != nil {
		//panic("read file fail" + infile + err.Error())
		return err
	}
	txns = findTxnsSQL(data, key)
	err = writeFile(outfile, txns)
	if err != nil {
		//panic("write file fail " + outfile + err.Error())
		return err
	}
	return err
}
*/
