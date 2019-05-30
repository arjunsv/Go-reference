package main

import ("fmt")

func add(x float64, y float64) float64 {
    return x + y
}

func multipleOutput(a, b string) (string, string) {
    return a,b
}

func main() {
    // Implicitly typed to float64.
    num1, num2 := 4.2, 9.3

    fmt.Println(add(num1, num2))

    w1, w2 := "word1", "word2"

    fmt.Println(multipleOutput(w1, w2))
}

func typingExample() {

    var a int = 62
    var b float64 = float64(a)

    // x will be type int.
    x := a
}