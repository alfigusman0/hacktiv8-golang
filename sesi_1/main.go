package main

import (
	"fmt"
)

/* func main() {
	// fmt.Println("Hello, World!")

	// Call InitDB from db.go
	// InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
} */

func main() {
	/* var nama string = "Alfi"
	var nilai int = 90

	nilai = 100

	fmt.Println("Nama: ", nama)
	fmt.Println("Nilai: ", nilai)

	fmt.Printf("tipe data nama : %T\n", nama)
	fmt.Printf("tipe data nilai : %T\n", nilai)

	fmt.Printf("value dari nama : %v\n", nama)
	fmt.Printf("value dari nilai : %v\n", nilai) */

	/* nilai_total, err := strconv.Atoi("100")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Nilai total: %v\n", nilai_total) */

	const (
		c1 = iota
		c2
		c3
	)

	fmt.Println(c1, c2, c3)

	type Gender int

	const (
		Male Gender = iota
		Female
	)

	fmt.Println("Gender Male :", Male, "Gender Female :", Female)
}
