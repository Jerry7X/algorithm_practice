package main

import "fmt"
//import "algorithm_practice"
import "algorithm_practice/quick_sort"
//import "quick_sort"

func main() {
   var example1 [20]int = [20]int{1,134,23,34,653,2,65,101,102,321,3,98,71,1001,9,72,723,11,82,7}
   var result = quick_sort.Quick_sort(example1)
   var i int
   for i=0;i<20;i++ {
     fmt.Printf("%d  ", result[i])
   }

}
