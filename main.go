package main

func main() {
	if err := scheduler(); err != nil {
		panic(err)
	}
}
