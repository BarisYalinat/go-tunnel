package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

var tunnelConn net.Conn

func main() {
	// A. TÜNEL PORTU (9000)
	tunnelListener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Tünel portu başlatılamadı:", err)
		os.Exit(1)
	}
	fmt.Println("1. TÜNEL HATTI: Port 9000 üzerinden evdeki client bekleniyor...")

	go func() {
		for {
			conn, err := tunnelListener.Accept()
			if err != nil {
				continue
			}
			fmt.Println("-> Evdeki bilgisayar tünel hattını başarıyla açtı!");
			tunnelConn = conn
		}
	}()

	// B. ZİYARETÇİ PORTU (8000) - Tam HTTP Sunucusu
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if tunnelConn == nil {
			http.Error(w, "Evdeki bilgisayarla tünel kurulu değil.", http.StatusServiceUnavailable)
			return
		}

		fmt.Println("-> Tarayıcıdan HTTP isteği geldi, tünele yazılıyor...")

		// 1. Tarayıcıdan gelen isteği tünel üzerinden eve gönderiyoruz
		err := r.Write(tunnelConn)
		if err != nil {
			fmt.Println("Tünel hattına istek yazılırken hata:", err)
			return
		}

		// 2. Evden (client) gelecek yanıtı HTTP standartlarında oku
		respReader := bufio.NewReader(tunnelConn)
		resp, err := http.ReadResponse(respReader, r)
		if err != nil {
			fmt.Println("Evden gelen HTTP yanıtı okunamadı:", err)
			return
		}
		defer resp.Body.Close()

		// 3. Evden gelen HTTP başlıklarını (Headers) tarayıcıya kopyala
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// 4. Evden gelen asıl HTML içeriğini tarayıcıya aktar ve bağlantıyı temizce kapat
		io.Copy(w, resp.Body)
	})

	fmt.Println("2. ZİYARETÇİ PORTU: Port 8000 üzerinden istekler dinleniyor...")
	_ = http.ListenAndServe(":8000", nil)
}