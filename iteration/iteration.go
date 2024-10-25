package iteration

func Repeat(character string, ammount int) (repeated string) {
	for i := 0; i < ammount; i++ {
		repeated += character
	}
	return
}

// otras formas de for
// i := 1
// for i <= 3 {
//     fmt.Println(i)
//     i = i + 1
// }
// for j := 0; j < 3; j++ {
//     fmt.Println(j)
// }
// for i := range 3 {
//     fmt.Println("range", i)
// }
// for {
//     fmt.Println("loop")
//     break
// }
// for n := range 6 {
//     if n%2 == 0 {
//         continue
//     }
//     fmt.Println(n)
// }
