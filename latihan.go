// package main
// import "fmt"
// const NMAX int = 100
// func main(){
// 	var A[NMAX]int
// 	var n, i int
// 	fmt.Scan(&n)
// 	for i = 0; i < n; i++{
// 		fmt.Scan(&A[i])
// 	}
// 	hasil := avaregValue(A[:n], n)
// 	fmt.Print(hasil)
// }
// func avaregValue(nums [] int, n int)int{
// 	var i, jum, c int
// 	for i = 0; i < n; i++{
// 		if nums[i]%2 == 0 && nums[i]%3 == 0{
// 			jum+= nums[i]
// 			c++
// 		}
// 	}
// 	if c == 0{
// 		return 0
// 	}else{
// 		return jum/c
// 	}
// }


// package main
// import "fmt"

// func subtractProductAndSum(n int) int {
//     x := n
//     kali := 1
//     tambah := 0
//     for n > 0 {
//         hasil := n % 10
//         kali = kali * hasil
//         n = n / 10
//     }
//     for x > 0 {
//         hasil := x % 10
//         tambah = tambah + hasil
//         x = x / 10
//     }

//     return kali - tambah
// }
// func main() {
// 	var n int
// 	fmt.Scan(&n)
//     fmt.Println(subtractProductAndSum(n))
// }



// package main
// import "fmt"
// const NMAX int = 100
// func main(){
// 	var A[NMAX]int
// 	var n, i int
// 	fmt.Scan(&n)
// 	for i = 0; i <n; i++{
// 		fmt.Scan(&A[i])
// 	}
// 	hasil := findNumber(n, A[:n])
// 	fmt.Print(hasil)
// }
// func findNumber(n int, A[] int)int{
// 	var i, genap int
// 	for i = 0; i<n; i++{
// 		if A[i]%2 == 0{
// 			genap++
// 		}
// 	}
// 	return genap
// }



package main
import "fmt"
type arrBalita struct{
	beratBalita float64
}
const NMAX int = 100
func main(){
	var A[NMAX]arrBalita
	var n , i int
	var min, max, rata float64
	fmt.Scan(&n)
	for i =0; i < n; i++{
		fmt.Scan(&A[i].beratBalita)
	}
	hitungMinMax(&min, &max, A, n)
	rata = rerata(A, n)
	fmt.Println(min, max, rata)
}
func hitungMinMax(min, max *float64, A[NMAX]arrBalita, n int){
	var i, j int
	*min = A[0].beratBalita
	*max = A[0].beratBalita
	for i = 0; i <n; i++{
		if *min > A[j].beratBalita{
			*min = (A[i].beratBalita)
		}else if *max < A[i].beratBalita{
			*max = A[i].beratBalita
		}
	}
}
func rerata(arrBerat [NMAX] arrBalita, n int) float64{
	var i int
	var rata2 float64
	for i = 0; i < n; i++{
		rata2 += arrBerat[i].beratBalita
	}
	return rata2/float64(i)
}
