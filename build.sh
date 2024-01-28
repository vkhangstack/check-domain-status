env GOOS=linux GOARCH=amd64 go build -o build/check-domain-status main.go
env GOOS=windows GOARCH=amd64 go build -o build/check-domain-status.exe main.go
cp -r domains.txt .env.example build/