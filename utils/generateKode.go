package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateKode() string {
	t := time.Now().Format("02012006")
	rand.Seed(time.Now().UnixNano())

	// Menghasilkan angka bulat acak antara 0 hingga 99
	randomInt := rand.Intn(1000000)

	return fmt.Sprintf("TRX-JUAL%s-%05d", t, randomInt)
}
