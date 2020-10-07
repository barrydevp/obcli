/*
Copyright © 2020 barrydevp <barrydevp@gmail.com>
Foobla CLI application foo.
*/
package main

import (
	"fmt"
	"github.com/barrydevp/obcli/cmd"
)

func main() {
	printLogo()

	cmd.Execute()
}

func printLogo() {
	fmt.Println(` ____                            _ 				 `)
	fmt.Println(`|  _ \                          | |   				 `)
	fmt.Println(`| |_) | __ _ _ __ _ __ _   _  __| | _____   ___ __  `)
	fmt.Println(`|  _ < / _' | '__| '__| | | |/ _/ |/ _ \ \ / / '_ \ `)
	fmt.Println(`| |_) | (_| | |  | |  | |_| | (_| |  __/\ V /| |_) |`)
	fmt.Println(`|____/ \__,_|_|  |_|   \__, |\__,_|\___| \_/ | .__/ `)
	fmt.Println(`                        __/ |                | |    `)
	fmt.Println(`                       |___/                 |_|    `)
	fmt.Println(``)
	fmt.Println(`https://github.com/barrydevp`)
	fmt.Println(``)
}
