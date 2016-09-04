package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Llongfile)
	// tree("")
	//recursiveRename()
	removeDuplicates()
}

func removeDuplicates() {
	index := make(map[int64][]string)
	findDuplicates(".", index)
	for _, names := range index {
		if len(names) > 1 {
			for i := 0; i < len(names); i++ {
				if names[i] == "" {
					continue
				}
				for j := i + 1; j < len(names); j++ {
					if names[j] == "" {
						continue
					}

					fa := names[i]
					a, err := ioutil.ReadFile(names[i])
					if err != nil {
						log.Fatal(err)
					}

					fb := names[j]
					b, err := ioutil.ReadFile(names[j])
					if err != nil {
						log.Fatal(err)
					}

					fn := fa
					idx := i
					if len(fn) < len(fb) {
						fn = fb
						idx = j
					}

					if bytes.Compare(a, b) == 0 {
						fmt.Printf("Removing %s...\n", fn)
						if err := os.Remove(fn); err != nil {
							log.Fatal(err)
						}
						names[idx] = ""
					} else {
						fmt.Printf("\nCandidates:\n")
						fmt.Printf("\n>> %q\n>> %q\n\n", fa, fb)

						fmt.Printf("Should remove %q?", fn)
						var op int
						count, err := fmt.Scanf("%d", &op)
						if err == nil && count == 1 && op == 1 {
							fmt.Printf("> Removing %s...\n", fb)
							if err := os.Remove(fn); err != nil {
								log.Fatal(err)
							}
							names[idx] = ""
						}
					}
				}
			}
		}
	}
}

func tree(prefix string) {
	file, err := os.Open("./")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	infos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		fmt.Printf("%s-> %s\n", prefix, info.Name())
		if info.IsDir() {
			if err := os.Chdir(info.Name()); err != nil {
				log.Fatal(err)
			}
			tree(prefix + "  ")
			if err := os.Chdir("../"); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func recursiveRename() {
	file, err := os.Open("./")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	infos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		oldpath := info.Name()
		newpath := info.Name()
		newpath = strings.Replace(newpath, " ", "", -1)
		newpath = strings.Replace(newpath, "-", "_", -1)
		if oldpath != newpath {
			fmt.Printf("Moving %q -> %q\n", oldpath, newpath)
			if err := os.Rename(oldpath, newpath); err != nil {
				log.Printf("ERR: Failed to move %q -> %q\n", oldpath, newpath)
			}
		}
		if info.IsDir() {
			if err := os.Chdir(newpath); err != nil {
				log.Fatal(err)
			}
			recursiveRename()
			if err := os.Chdir("../"); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func findDuplicates(path string, index map[int64][]string) {
	file, err := os.Open("./")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	infos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			if err := os.Chdir(info.Name()); err != nil {
				log.Fatal(err)
			}
			findDuplicates(path+"/"+info.Name(), index)
			if err := os.Chdir("../"); err != nil {
				log.Fatal(err)
			}
		} else {
			name := path + "/" + info.Name()
			index[info.Size()] = append(index[info.Size()], name)
		}
	}
}
