package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	// 1. Sunucuya tünel portundan bağlan
	tunnelConn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Uzak sunucuya bağlanılamadı:", err)
		os.Exit(1)
	}
	defer tunnelConn.Close()
	fmt.Println("Uzak sunucuyla tünel hattı kuruldu. İstekler dinleniyor...")

	reader := bufio.NewReader(tunnelConn)
	for {
		// Tünelden gelen veriyi tam bir HTTP isteği olarak oku
		req, err := http.ReadRequest(reader)
		if err != nil {
			fmt.Println("Tünel hattından istek okunurken hata (bağlantı kopmuş olabilir):", err)
			break
		}

		fmt.Println("-> Tünelden bir HTTP isteği alındı, yerel siteye (8080) yönlendiriliyor...")

		// İstek adresini bizim yerel web sitemize (8080) çeviriyoruz
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:8080"
		req.RequestURI = "" 

		// Gerçek bir HTTP istemcisi gibi yerel sitemize isteği atıyoruz
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[Hata] Yerel web sitesinden (8080) yanıt alınamadı. Açık mı?")
			continue
		}

		// Yerel siteden dönen HTTP cevabını tünel üzerinden sunucuya geri yazıyoruz
		err = resp.Write(tunnelConn)
		resp.Body.Close()
		if err != nil {
			fmt.Println("Tünel hattına cevap yazılırken hata:", err)
			break
		}
	}
}