# Go UserRegistration Uygulaması

## Genel Bakış

**GoUserAuth**, kullanıcı kimlik doğrulama işlevleri sunan Go ile yazılmış bir web uygulamasıdır. Uygulama, `kullanıcı kaydı` ve `girişi` işlemlerini içerir. `SQLite3` veritabanı ve `bcrypt` ile şifre karma kullanır. Kullanıcı kayıt ve giriş işlemleri için HTTP yollarını, girdi doğrulaması ve hata yönetimi ile birlikte sağlar.

## Kullanılan Teknolojiler

- **Go**: Uygulamayı oluşturmak için kullanılan programlama dili.
- **SQLite**: Kullanıcı bilgilerini depolamak için kullanılan veritabanı.
- **bcrypt**: Şifreleri karmak için kullanılan kütüphane.
- **Postman**: Uç noktalarla etkileşim kurmak için kullanılan API test aracı.

## Projenin Kurulumu

1. **Depoyu klonlayın:**

    ```sh
    git clone https://github.com/erenerdogmus/GoUserAuth.git
    cd GoUserAuth
    ```

2. **Bağımlılıkları yükleyin:**

    ```sh
    go mod tidy
    ```

3. **Uygulamayı çalıştırın:**

    ```sh
    go run ./cmd/web/
    ```

## Veritabanı Başlatma

Uygulama başladığında veritabanını otomatik olarak başlatır ve gerekli tabloları oluşturur. `SQLite3` veritabanı dosyasının (`user.db`) erişilebilir olduğundan emin olun.

## API Uç Noktaları

### Kullanıcı Kaydı

- **Uç Nokta:** `/user/signup`
- **Yöntem:** `POST`
- **İstek Parametreleri:**
  - `username` (string): Kullanıcının kullanıcı adı.
  - `email` (string): Kullanıcının e-posta adresi.
  - `password` (string): Kullanıcı için şifre.
  - `confirm` (string): Şifre onayı.
- **Yanıt:**
  - `id` (int64): Oluşturulan kullanıcının ID'si.
  - `username` (string): Oluşturulan kullanıcının kullanıcı adı.
  - `email` (string): Oluşturulan kullanıcının e-posta adresi.
  - `message` (string): Başarı mesajı.

### Kullanıcı Girişi

- **Uç Nokta:** `/user/login`
- **Yöntem:** `POST`
- **İstek Parametreleri:**
  - `email` (string): Kullanıcının e-posta adresi.
  - `password` (string): Kullanıcının şifresi.
- **Yanıt:**
  - `id` (int64): Doğrulanan kullanıcının ID'si.
  - `username` (string): Doğrulanan kullanıcının kullanıcı adı.
  - `email` (string): Doğrulanan kullanıcının e-posta adresi.
  - `message` (string): Başarı mesajı.

## Postman Kullanımı

1. **Postman'i açın** ve projeniz için yeni bir koleksiyon oluşturun.

2. **Kullanıcı kaydı için yeni bir istek ekleyin:**

    - **Yöntem:** `POST`
    - **URL:** `http://localhost:8080/user/signup`
    - **Body:** 
      - `form-data` seçin.
      - Aşağıdaki alanları ekleyin:
        - `username`: İstediğiniz kullanıcı adı.
        - `email`: E-posta adresiniz.
        - `password`: Şifreniz.
        - `confirm`: Şifre onayı.

3. **Kullanıcı girişi için yeni bir istek ekleyin:**

    - **Yöntem:** `POST`
    - **URL:** `http://localhost:8080/user/login`
    - **Body:** 
      - `form-data` seçin.
      - Aşağıdaki alanları ekleyin:
        - `email`: E-posta adresiniz.
        - `password`: Şifreniz.

4. **İstekleri gönderin** ve yanıtları gözlemleyin.

## Hata Yönetimi

Uygulama, geçersiz girdi veya işlem hataları için detaylı hata yanıtları sağlar. Her hata yanıtı, bir `durum kodu` ve `açıklayıcı bir mesaj` içerir.

## Sonuç

Bu Go web uygulaması, kullanıcı kaydı ve girişi için temel kimlik doğrulama işlevleri sunar. `Postman` kullanarak API ile etkileşim kurabilir ve kullanıcı kaydı ve giriş özelliklerini test edebilirsiniz.
