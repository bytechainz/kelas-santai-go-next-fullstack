package main // package main merupakan package utama yang akan digunakan
import (
	"fmt"
	"time"
)

func main() { // function utama yang akan di eksekusi
	fmt.Println("hello mas adam")

	// penulisan komentar
	/*
		ini komentar multiline
	*/

	// -------- variable

	// cara 1
	var nama string
	nama = "mas budi"

	fmt.Println("nama saya adalah ", nama)

	// cara 2
	var nama2 = "budi"
	fmt.Println(nama2)

	namaBelakang := "jumadi"
	fmt.Println(namaBelakang)

	namaBelakang = "bigdata"
	fmt.Println("data terbaru : ", namaBelakang)

	// ---------- tipe data --------
	umur := 24

	fmt.Println("nama saya adalah", umur)

	berat := 78.5
	fmt.Println("berat saya adalah", berat)

	makan := false
	fmt.Println("apakah saya sudah makan : ", makan)

	// penjumlahan
	a := 10
	b := 20
	fmt.Println("hasil dari penjulahan ", a+b)

	fmt.Println("-----------konstanta----------")
	const pi = 3.14
	fmt.Println("nilai pi adalah : ", pi)

	fmt.Println("-----------operator----------")
	// artimatika
	x := 10
	z := 5
	fmt.Println("-----------aritmatika----------")
	fmt.Println(x + z)
	fmt.Println(x - z)
	fmt.Println(x / z)
	fmt.Println(x * z)
	fmt.Println(x % z)

	// perbandingan
	fmt.Println("-----------perbandingan----------")
	fmt.Println(x > z)
	fmt.Println(x < z)
	fmt.Println(x == z)
	fmt.Println(x != z)
	fmt.Println(x >= z)
	fmt.Println(x <= z)

	fmt.Println("-----------logika----------")
	var left = false
	var right = true

	fmt.Println(left && right)
	fmt.Println(left || right)
	fmt.Println(!left)

	fmt.Println("-----------seleksi kondisi----------")

	nilai := 81
	if nilai > 80 {
		fmt.Println("nilai A")
	} else if nilai > 60 {
		fmt.Println("nilai b")
	} else {
		fmt.Println("nilai c")
	}

	fullname := "sukardi"

	if (fullname == "sukardi" || fullname == "suradi") && nilai > 80 {
		fmt.Println("selamat datang")
	} else {
		fmt.Println("anda tidak dikenal ")
	}

	fmt.Println("-----------switch case----------")
	nilaiUjian := 75

	switch nilaiUjian {
	case 10:
		fmt.Println("nilai e")
	case 20:
		fmt.Println("nilai d")
	default:
		fmt.Println("nilai anda tidak valid")
	}

	fmt.Println("----------- perulangan ----------")
	for i := 1; i <= 5; i++ {
		fmt.Println("ini adalah nilai ke : ", i)
	}

	var txt = "123333333333333"
	for i, v := range txt {
		fmt.Println("index", i, "value : ", v)
	}

	var txt2 = [5]int{10, 20, 30, 40, 50}
	for i, v := range txt2 {
		fmt.Println("index : ", i, "value : ", v)
	}

	fmt.Println("----------- fungsi ----------")
	TimeNowJakarta()
	tampung := TimeNowJakartaReturn()

	fmt.Println(tampung)

	luas := AlasSegitiga(2.0, 2.0)
	fmt.Println(luas)

}

func TimeNowJakarta() {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		panic(err)
	}

	now := time.Now().In(loc)

	timeString := now.Format("02 January 2006, 03:04 PM")
	fmt.Println("current time in jakarta : ", timeString)
}

func TimeNowJakartaReturn() string {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		panic(err)
	}

	now := time.Now().In(loc)

	timeString := now.Format("02 January 2006, 03:04 PM")
	return timeString
}

func AlasSegitiga(alas, tinggi float64) float64 {
	luas := 0.5 * alas * tinggi
	return luas
}
