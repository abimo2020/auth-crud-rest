# Auth CRUD Rest
### Sebuah program RESTful API untuk melakukan CRUD dan Authentication pengguna.
###  Tech stack yang digunakan :
- Go versi 1.22.3
    - Go digunakan sebagai bahasa pemrograman dalam mengembangkan tugas ini.
- Echo v4
    - Echo merupakan sebuah framework dari Go yang memudahkan dalam pengembangan RESTful API ini.
- Mysql
    - Mysql digunakan sebagai database SQL untuk menyimpan data user
- GORM
    - GORM merupakan sebuah library ORM, digunakan untuk database connection pada mysql serta melakukan query.
- Docker & Docker Compose
    - Docker digunakan sebagai container untuk membungkus aplikasi-aplikasi yang digunakan agar mempunyai lingkungan yang sama.
    - Docker Compose digunakan untuk penyederhanaan dalam mengeksekusi beberapa container sekaligus 
- Redis
    - Redis merupakan databse NoSQL berupa pasangan key dan value. Pada tugas ini penggunaan redis ditujukan untuk caching, yaitu menyimpak token jwt auth agar mempercepat aplikasi dalam mengambil data, dimana tidak melakukan pembuatan token berulang-ulang ketika melakukan login. Hal tersebut terjadi karena saat token masih valid atau tidak expired maka token akan diambil melalui redis.
- RabbitMQ
    - RabbitMQ digunakan sebagai message broker, pada tugas ini digunakan untuk proses pertukaran pesan yaitu publish dan consume untuk melakukan pembuatan dan perubahan user.
  
### Postman Documentation Link : https://documenter.getpostman.com/view/27529365/2sA3Qy7VVW
