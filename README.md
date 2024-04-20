![Açıklama](https://github.com/omerfdev/turkishCitizenIDCheckAPI/blob/master/tckId.png)
# TCK No Doğrulama Servisi

Bu basit Go programı, Türkiye Cumhuriyeti Kimlik Numarası (TCKN) doğrulaması için bir HTTP API sağlar.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun ve Gin paketinin yüklü olduğundan emin olun.
2. **Servisi Başlatın**: `go run main.go` komutuyla servisi başlatın.
3. **TCK No Doğrulama**: Tarayıcıdan veya bir API test aracından `/validateTCK?tck=YOUR_TCK_NUMBER` endpoint'ine bir GET isteği yaparak TCK numarasını doğrulayın.

## API Dökümantasyonu

### `/validateTCK` Endpoint'i

- **Method**: GET
- **Parametreler**:
  - `tck`: Doğrulanacak TCK numarası.
- **Cevap**:
  - Başarılı doğrulama durumunda HTTP status kodu 200 ve bir "message" yanıtı döner: `{"message": "TCK is valid."}`
  - Hatalı doğrulama durumunda HTTP status kodu 400 ve bir "error" yanıtı döner: `{"error": "HATA_METNİ"}`

Örnek istek:
GET http://localhost:8080/validateTCK?tck=12345678901
