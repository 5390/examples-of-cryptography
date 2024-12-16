# examples-of-cryptography
**Cryptography in Golang**

Cryptography is the practice of securing communication and data from unauthorized access. It includes:

**Encryption/Decryption**: Converting plaintext ↔ ciphertext.
**Integrity**: Ensuring no data tampering.
**Authentication**: Verifying sender/recipient identity.
**Types of Cryptographic Algorithms**
**Symmetric Key Cryptography**: Single key for encryption and decryption.
  Examples: AES, ChaCha20.
**Asymmetric Key Cryptography**: Public/Private key pair.
  Examples: RSA, ECC.
**Hashing Algorithm**s: Converts input to fixed-length output.
Examples: SHA-256, MD5.
**Key Exchange**: Securely exchange cryptographic keys.
  Example: ECDH.
**Post-Quantum Cryptography**: Algorithms secure against quantum attacks.
  Example: Dilithium, Kyber.
  
**Golang Implementations**

1. Symmetric Key Cryptography: AES

func encryptAES(key []byte, text string) (string, error) { /* AES Encryption Logic */ }
func decryptAES(key []byte, cipherText string) (string, error) { /* AES Decryption Logic */ }

2. Asymmetric Key Cryptography: RSA
//RSA Key Pair Generation and Encryption.
privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
cipherText, _ := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, message)

3. Hashing Algorithms: SHA-256
hash := sha256.Sum256([]byte("Hello, World!"))
fmt.Printf("SHA-256 Hash: %x\n", hash)

4. Key Exchange: ECDH
privateKeyA, _ := ecdh.P256().GenerateKey(rand.Reader)
sharedSecret, _ := privateKeyA.ECDH(publicKeyB)
Post-Quantum Cryptography
Post-quantum algorithms include Dilithium (signatures) and Kyber (key exchange). Golang libraries like CIRCL support implementation.

Hyperledger Fabric and Solana
Hyperledger Fabric:
Consensus: Raft (Leader Election).
Smart Contracts: Chaincode in Go.
Storage: CouchDB/LevelDB.
Solana:
Consensus: Proof of History (PoH) + Tower BFT.
Parallel Transaction Processing via Sealevel.
Accounts Model for State Storage.
Code Examples
Raft Consensus Simulation

Simplified leader election example using Goroutines.
Hyperledger Fabric Chaincode

Example of asset creation and query.
Solana-like Account Transfer Simulation

Demonstrates parallel processing.
Storage Management
Hyperledger Fabric: CouchDB for World State.
Solana: Account-based ledger with snapshotting.
Summary
This document explains cryptographic algorithms and their implementation in Golang, alongside blockchain frameworks like Hyperledger Fabric and Solana.

References:
Golang Crypto Package Documentation
Cloudflare CIRCL
