set -x

sudo go build -o /usr/local/bin/uv main.go run.go
sudo mkdir /etc/uv
sudo cp uv.json /etc/uv/
