# Example Age Implementation

### Encryption

1. Go to **cmd/encrypt** directory
2. Remove example encrypted file at `cmd/encrypt/encfiles/a.txt.age`
3. Run `go run cmd.go` from **cmd/encrypt** directory
4. You'll see a new file appears at `cmd/encrypt/encfiles/a.txt.age`

### Decryption

1. Go to **cmd/decrypt** directory
2. Remove example decrypted file at `cmd/decrypt/files/a.txt`
3. Copy encrypted file from `cmd/encrypt/encfiles/a.txt.age` to `cmd/decrypt/encfiles/a.txt.age`
4. Run `go run cmd.go` from **cmd/decrypt** directory
5. You'll see a new file appears at `cmd/decrypt/files/a.txt`

## Generate Key

1. Go to **cmd/generate-key**
2. Remove example generated key file at `cmd/generate-key/key.txt`
3. Run `go run cmd.go`
4. You'll see a new generated key stored at `key.txt`