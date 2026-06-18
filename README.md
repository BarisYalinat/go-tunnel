# Go-Tunnel 🚀
Go (Golang) ile geliştirilmiş, ham TCP soket programlama ve Tersine Vekil (Reverse Proxy) mimarisine dayalı, lokal servisleri dış dünyaya açan tünelleme aracı (Ngrok / Playit.gg Klonu).

## 📌 Projenin Amacı ve Mimarisi
Bu proje; NAT veya Güvenlik Duvarı arkasında kalan lokal web sunucularının (örneğin `localhost:8080`), herhangi bir port yönlendirme (Port Forwarding) işlemine gerek kalmadan, dış dünyadan erişilebilir hale getirilmesini sağlar.

Sistem üç ana katmandan oluşur:
1. **Uzak Sunucu (Server):** Dış dünyadan gelen HTTP isteklerini dinler ve aktif tünel hattına yönlendirir.
2. **Lokal İstemci (Client):** Uzak sunucu ile sürekli açık bir TCP köprüsü (Tünel) kurar ve gelen istekleri yerel servise paslar.
3. **Yerel Servis (Target Web App):** Evde çalışan asıl web sitesi veya oyun sunucusu.

## 🛠️ Teknik Kazanımlar & Öğrenilen Konular
- **Soket Programlama:** `net.Listen` ve `net.Dial` ile ham TCP katmanında bağlantı yönetimi.
- **HTTP Protokol Akışı:** HTTP paketlerinin bütünlüğünü korumak adına `http.ReadRequest` ve `resp.Write` ile stream bazlı veri transferi.
- **Eşzamanlılık (Concurrency):** Go'nun `goroutine` altyapısı kullanılarak isteklerin kilitlenme (deadlock) yaşamadan asenkron olarak işlenmesi.
- **Konteynerleştirme (Docker):** Uygulamanın bağımlılıklardan arındırılarak her ortamda çalışabilmesi için `Dockerfile` optimizasyonu.

## 💻 Lokal Kurulum ve Çalıştırma

### 1. Sunucuyu Başlatın
```bash
cd server
go run main.go
