FROM golang:1.22-alpine

# CGO'yu etkinleştir
ENV CGO_ENABLED=1

# Çalışma dizinini belirleyin
WORKDIR /app

# Bağımlılıkları indirmek için go mod dosyalarını kopyalayın
COPY go.mod go.sum ./
RUN go mod tidy

# GCC ve diğer gerekli bağımlılıkları yükleyin
RUN apk add --no-cache gcc musl-dev

# Uygulama kaynaklarını kopyalayın
COPY . .

# Uygulamayı derleyin
RUN go build -o inventory-services .

# Uygulamayı başlatın
CMD ["./inventory-services"]
