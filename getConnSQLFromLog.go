package main

import (
	"fmt"
	"strings"
)

func getConnID(log string) string {
	end := 0
	begin := strings.LastIndex(log, "[conn=")
	if begin == -1 {
		return ""
	} else {
		for k := begin; k <= len(log)-1; k++ {
			if log[k] == ']' {
				end = k
				break
			}
		}
	}
	if end > 0 {
		return log[begin+6 : end-1]
	} else {
		return ""
	}
}

func writeSQLByConnID(m map[string]string, outfile string) error {
	for k, v := range m {
		filename := outfile + "/" + k
		err := writeFile(filename, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func splitGenerallogByConn(infile, outfile string) error {
	m := make(map[string]string)
	var connID string
	data, err := readAllDataFromFile(infile)
	if err != nil {
		return err
	}
	sqls := strings.Split(data, "\n")
	for _, v := range sqls {
		if strings.LastIndex(v, "GENERAL_LOG") == -1 {
			continue
		} else {
			connID = getConnID(v)
			//fmt.Println(connID)
			if connID == "" {
				continue
			} else {
				txns := m[connID]
				txns += v + "\n"
				m[connID] = txns
			}
		}
	}
	fmt.Println("conn num is ", len(m))
	return writeSQLByConnID(m, outfile)
}
