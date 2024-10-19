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
    git clone https://github.com/muratbulut/yavuzlar-the-ripper.git
    cd yavuzlar-the-ripper
    ```

## Kullanım

### Sözlük Saldırısı

Bir wordlist dosyasını kullanarak hedef hashli parolayı bulmak için aşağıdaki komutu kullanın:

```sh
./yavuzlar-the-ripper -wordlist=wordlist.txt -hash=md5 -target=5d41402abc4b2a76b9719d911017c592 -workers=4
```

### Brute Force Saldırısı

```sh
./yavuzlar-the-ripper -bruteforce -hash=md5 -target=5d41402abc4b2a76b9719d911017c592 -workers=4
```