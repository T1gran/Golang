package main
import "fmt"
 
func main() {
     
    intCh := make(chan int, 3) 
    intCh <- 10
    intCh <- 3
    intCh <- 24
    fmt.Println(<-intCh)     // 10
    fmt.Println(<-intCh)     // 3
    fmt.Println(<-intCh)     //24
	
	fmt.Println(cap(intCh))     // 3
    fmt.Println(len(intCh))     // 0
}