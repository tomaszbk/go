package testdata

import "log"

func m() {
	maps := make(map[string]string)
	for k, _ := range maps { // want "simplify range expression"
		log.Println(k)
	}
	for _ = range maps { // want "simplify range expression"
	}
	for _, _ = range maps { // want "simplify range expression"
	}
	for _, v := range maps { // nope
		println(v)
	}
	for range maps { // nope
	}
}
