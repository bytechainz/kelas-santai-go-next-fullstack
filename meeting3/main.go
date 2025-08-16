package main

import "fmt"

func main() {

	var nilai = [3]int{}

	fmt.Println(nilai)
	nilai[0] = 10
	nilai[1] = 20
	nilai[2] = 30

	fmt.Println(nilai)

	// membuat array dengan inisialisai langsung
	var nilai2 = [3]int{10, 20, 30}
	fmt.Println(nilai2)

	// ambil nilai array
	fmt.Println("mengambil nilai 30 :", nilai[2])
	fmt.Println("jumlah array : ", len(nilai2))

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(nilai2); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	for i, fruit := range fruits {
		fmt.Printf("element from for range %d : %s : ", i, fruit)
	}

	fmt.Println("----------  slice -------")
	nilaiSlice := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// slice low to high
	slice1 := nilaiSlice[2:5]
	fmt.Println("slice 1 : ", slice1)

	// slice low
	slice2 := nilaiSlice[6:]
	fmt.Println("slice 2 : ", slice2)

	// slice high
	slice3 := nilaiSlice[:4]
	fmt.Println("slice 3 : ", slice3)

	// slice low:high
	slice4 := nilaiSlice[:]
	fmt.Println("slice 4 : ", slice4)

	// mengubah nilai slice
	nilaiSlice[2] = 99
	fmt.Println("setelah diubah : ", slice4)

	//membuat slice
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice 5 : ", slice5)

	//menambahkan data pada slice
	slice5 = append(slice5, 11, 12, 13)
	fmt.Println("setelah penambahan data pada slice ", slice5)

	// menghapus

	// MAP : tipe data yang berbentuk key dan value
}
