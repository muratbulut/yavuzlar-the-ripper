package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {
	// Komut satırı bayrakları
	help := flag.Bool("help", false, "Kullanılabilir parametreleri gösterir")
	wordlistPath := flag.String("wordlist", "", "Wordlist dosyasının yolu")
	bruteForce := flag.Bool("bruteforce", false, "Brute force saldırısı yap")
	hashType := flag.String("hash", "md5", "Hash türü (md5, sha1, sha256)")
	workers := flag.Int("workers", runtime.NumCPU(), "İşçi sayısı")
	targetHash := flag.String("target", "", "Hedef hashli parola")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Yavuzlar The Ripper'a hoşgeldiniz, tool'u kullanmak için -help komutunu çalıştırın.")
		return
	}

	if *help {
		fmt.Println("Kullanılabilir parametreler:")
		fmt.Println("-help: Kullanılabilir parametreleri gösterir")
		fmt.Println("-wordlist: Wordlist dosyasının yolu")
		fmt.Println("-bruteforce: Brute force saldırısı yap")
		fmt.Println("-hash: Hash türü (md5, sha1, sha256)")
		fmt.Println("-workers: İşçi sayısı")
		fmt.Println("-target: Hedef hashli parola")
		return
	}

	if *targetHash == "" {
		fmt.Println("Lütfen bir hedef hashli parola belirtin.")
		return
	}

	if *bruteForce {
		fmt.Println("Brute force saldırısı başlatılıyor...")
		bruteForceAttack(*targetHash, *hashType, *workers)
		return
	}

	if *wordlistPath == "" {
		fmt.Println("Lütfen bir wordlist dosyası yolu belirtin veya brute force saldırısı yapın.")
		return
	}

	file, err := os.Open(*wordlistPath)
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	wordChan := make(chan string)

	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for word := range wordChan {
				hashedWord := hashPassword(word, *hashType)
				if hashedWord == *targetHash {
					fmt.Println("Parola bulundu:", word)
					os.Exit(0)
				}
			}
		}()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		wordChan <- word
	}
	close(wordChan)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Dosya okuma hatası:", err)
	} else {
		fmt.Println("Parola bulunamadı.")
	}
}

func hashPassword(password, hashType string) string {
	var hash []byte
	switch hashType {
	case "md5":
		h := md5.Sum([]byte(password))
		hash = h[:]
	case "sha1":
		h := sha1.Sum([]byte(password))
		hash = h[:]
	case "sha256":
		h := sha256.Sum256([]byte(password))
		hash = h[:]
	default:
		fmt.Println("Desteklenmeyen hash türü:", hashType)
		os.Exit(1)
	}
	return hex.EncodeToString(hash)
}

func bruteForceAttack(targetHash, hashType string, workers int) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var wg sync.WaitGroup
	passwordChan := make(chan string)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for password := range passwordChan {
				hashedPassword := hashPassword(password, hashType)
				if hashedPassword == targetHash {
					fmt.Println("Parola bulundu:", password)
					os.Exit(0)
				}
			}
		}()
	}

	generatePasswords("", charset, targetHash, passwordChan)
	close(passwordChan)

	wg.Wait()
	fmt.Println("Parola bulunamadı.")
}

func generatePasswords(current, charset, targetHash string, passwordChan chan string) {
	if len(current) > 6 { //
		return
	}
	passwordChan <- current
	for _, char := range charset {
		generatePasswords(current+string(char), charset, targetHash, passwordChan)
	}
}
