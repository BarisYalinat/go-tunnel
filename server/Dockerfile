# 1. Aşama: Go kodunu derlemek için resmi Go imajını kullanıyoruz
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Modül dosyalarını kopyala ve indir
COPY go.mod ./
RUN go mod download

# Sunucu kodunu kopyala
COPY server/main.go ./server/

# Kodu optimize ve hafif bir binary (çalıştırılabilir dosya) olarak derle
RUN CGO_ENABLED=0 GOOS=linux go build -o /server-app ./server/main.go

# 2. Aşama: Sadece derlenen binary'yi çalıştıracak çok hafif bir imaj oluşturuyoruz
FROM alpine:latest

WORKDIR /

# Derlenen dosyayı ilk aşamadan buraya kopyala
COPY --from=builder /server-app /server-app

# Tünel (9000) ve Ziyaretçi (8000) portlarını dışarıya aç
EXPOSE 9000
EXPOSE 8000

# Uygulamayı başlat
CMD ["/server-app"]