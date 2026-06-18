# 1. Aşama: Derleme
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Şu an server klasörünün içinde olduğumuz için direkt yanındaki dosyaları kopyalıyoruz
COPY main.go ./

# Modül olmadan direkt tek dosyayı derle diyoruz (go.mod bağımlılığını kaldırdık)
RUN CGO_ENABLED=0 GOOS=linux go build -o /server-app main.go

# 2. Aşama: Çalıştırma imajı
FROM alpine:latest

WORKDIR /

COPY --from=builder /server-app /server-app

EXPOSE 9000
EXPOSE 8000

CMD ["/server-app"]