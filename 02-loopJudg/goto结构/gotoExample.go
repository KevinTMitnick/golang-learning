package main

import "fmt"

func main()  {
	fmt.Println("123")
	fmt.Println("456")

	goto GAMEOVER

	LASTWORD:
	fmt.Println("789")
	return

	GAMEOVER:
	fmt.Println("Game Over!!!")

	goto LASTWORD
}
