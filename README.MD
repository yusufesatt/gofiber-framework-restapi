# GoFiber Framework REST API

GoFiber framework kullanılarak Udemy kursunda izlemiş olduğum videolarla geliştirilmiş bir Todo REST API uygulaması.

## 🚀 Özellikler

- ✅ CRUD işlemleri (Create, Read, Update, Delete)
- ✅ Go Fiber v3 framework
- ✅ GORM ORM ile veritabanı yönetimi
- ✅ SQLite veritabanı
- ✅ RESTful API tasarımı
- ✅ Health check endpoint



## 📚 API Endpoints

### Health Check
```
GET /
```
Uygulamanın çalışıp çalışmadığını kontrol eder.

**Response:**
```json
{
  "success": true
}
```

### Todo İşlemleri

#### Tüm Todoları Listele
```
GET /todos
```

#### Belirli Bir Todo Getir
```
GET /todos/:todoID
```

#### Yeni Todo Oluştur
```
POST /todos
```

#### Todo Güncelle
```
PUT /todos/:todoID
```

#### Todo Sil
```
DELETE /todos/:todoID
```

## 🏗️ Proje Yapısı

```
.
├── dal/           # Data Access Layer - Veri modelleri
├── database/      # Veritabanı bağlantı ve yapılandırması
├── services/      # İş mantığı ve handler fonksiyonları
├── types/         # Tip tanımlamaları
├── main.go        # Uygulama giriş noktası
└── database.db    # SQLite veritabanı
```

## 🔧 Kullanılan Teknolojiler

- [Go Fiber v3](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM kütüphanesi
- [SQLite](https://www.sqlite.org/) - Veritabanı
- [Go Playground Validator](https://github.com/go-playground/validator) - Veri doğrulama
