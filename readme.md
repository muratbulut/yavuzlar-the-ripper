# Yavuzlar The Ripper

Yavuzlar The Ripper, parolaları kırmak için kullanılan bir araçtır. Hem sözlük saldırısı hem de brute force saldırısı yapabilir. Bu araç, MD5, SHA-1 ve SHA-256 hash türlerini destekler.

## Özellikler

- Sözlük saldırısı
- Brute force saldırısı
- MD5, SHA-1 ve SHA-256 hash türlerini destekler
- Paralel işleme desteği

## Kurulum

1. Go programlama dilini kurun: [Go Kurulumu](https://golang.org/doc/install)
2. Bu projeyi klonlayın:
    ```sh
    git clone https://github.com/kullaniciadi/yavuzlar-the-ripper.git
    cd yavuzlar-the-ripper
    ```

## Kullanım

### Sözlük Saldırısı

Bir wordlist dosyasını kullanarak hedef hashli parolayı bulmak için aşağıdaki komutu kullanın:

```sh
go run main.go -wordlist=wordlist.txt -hash=md5 -target=5d41402abc4b2a76b9719d911017c592 -workers=4