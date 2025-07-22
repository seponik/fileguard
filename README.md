<h1 align="center">🔐 FileGuard</h1>

**FileGuard** is a minimal and secure file encryption and decryption library written in Go. It also includes an optional command-line tool 🖥️.

## ✨ Features

- 🔒 AES-based file encryption and decryption
- 📘 Simple, clean public API
- 🖥️ Optional CLI for quick usage
- 🛠️ Easy to integrate into any Go project

## 📦 Installation

### 📚 Library

```
go get github.com/seponik/fileguard/pkg/fileguard
```

🧰 CLI Tool


1. Go to the [Releases](https://github.com/seponik/fileguard/releases) section of this repository.
2. Download the latest release for your operating system.
3. Follow the instructions provided in the Usage section or run the downloaded file.


## 🚀 Usage

🧩 In Go code

```go
package main

import (
	"fmt"

	"github.com/seponik/fileguard/pkg/fileguard"
)

func main() {
  // Encrypting file using fileguard
  err := fileguard.EncryptFile("file.txt", "secretkey")
  if err != nil {
    fmt.Println(err)
  }

  // Decrypting file using fileguard
  err = fileguard.DecryptFile("file.txt.fg", "secretkey")
  if err != nil {
    fmt.Println(err)
  }
}
```

🖥️ From the CLI

**Encrypting**
```bash
./fileguard -e example.txt
```

**Decrypting**
```bash
./fileguard -d example.txt.fg
```
##
 ⚠️ Security Note

Always use strong passwords.
