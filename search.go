package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func search(fn string, key string, fold bool) {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	var line, cline string
	var lo, mid, hi int64
	var n int
	st, _ := os.Stat(fn)
	hi = st.Size()
	const bsz int64 = 1 << 12
	var ba [bsz]byte
	var buf []byte = ba[0:]

	if fold {
		key = strings.ToLower(key)
	}
	var found int64
	var match bool
	for {
		mid = lo + (hi-lo)/2
		//fmt.Println("Looping", lo, mid, hi, hi-lo)
		n, err = f.ReadAt(buf, mid)
		if err != nil && err != io.EOF && n == 0 {
			fmt.Println("ReadAt", mid, n, err)
			log.Fatal(err)
		}
		br := bytes.NewBuffer(buf[:n])

		if (hi - lo) < bsz {
			//fmt.Println("linear")

			n, err = f.ReadAt(buf, lo)
			if err != nil && err != io.EOF && n == 0 {
				fmt.Println("ReadAt", lo, n, err)
				log.Fatal(err)
			}
			br := bytes.NewBuffer(buf[:n])

			curo := lo // current offset
			for {
				//fmt.Println("lo close to hi")
				line, err = br.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						return
					}
					log.Fatal(err)
				}
				if fold {
					cline = strings.ToLower(line)
				} else {
					cline = line
				}
				//fmt.Println("ReadString false", cline)
				curo += int64(len(line))

				if strings.HasPrefix(cline, key) {
					fmt.Print(line)
					match = true
					found = curo
					break
				}
				continue
			}

			if match == true {
				//fmt.Println("true")

				n, err = f.ReadAt(buf, found)
				if err != nil && err != io.EOF && n == 0 {
					fmt.Println("ReadAt", lo, n, err)
					log.Fatal(err)
				}
				br := bytes.NewBuffer(buf[:n])

				for {
					line, err = br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							return
						}
						log.Fatal(err)
					}
					if fold {
						cline = strings.ToLower(line)
					} else {
						cline = line
					}
					//fmt.Println("ReadString true", cline)
					if strings.HasPrefix(cline, key) {
						fmt.Print(line)
					} else {
						return
					}
				}
			}
			//return
		}

		line, err = br.ReadString('\n') // partial line?
		mid += int64(len(line))
		line, err = br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		if fold {
			cline = strings.ToLower(line)
		} else {
			cline = line
		}
		//fmt.Println("ReadString binary", cline)

		if strings.HasPrefix(cline, key) {
			//fmt.Println("HasPrefix", cline, key)
			found = mid
			match = true
			hi = mid
			continue
		}
		if key < cline {
			//fmt.Println(key, "<", cline)
			hi = mid
			continue
		}
		if key > cline {
			//fmt.Println(key, ">", line)
			lo = mid
			continue
		}
	}
}
