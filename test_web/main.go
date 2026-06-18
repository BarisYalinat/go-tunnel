package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Tarayıcıya düz bir HTML sayfası basacak basit bir fonksiyon
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Tebrikler Baris! Tünel Başarıyla Çalıştı 🚀</h1><p>Bu yazı evindeki bilgisayardan tünel vasıtasıyla geldi.</p>")
	})

	fmt.Println("Test Web Sitesi Port 8080 üzerinde çalışıyor...")
	_ = http.ListenAndServe(":8080", nil)
}