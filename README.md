# WebscrappingWithGo
Scrapping weather APIs with Go to store them in a local MariaDB instance


## Installation

```
# Download
wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz

# Extract
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz

# setup Environment variables
nano ~/.profile
export PATH=$PATH:/usr/local/go/bin
# reload
source ~/.profile

# verify
go version
```


## Run
```
go run <file_name.go>
```