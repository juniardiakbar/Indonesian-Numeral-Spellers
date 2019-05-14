# Indonesian Numeral Speller

## Latar Belakang
Mengeja angka merupakan salah satu kegiatan dasar yang dilakukan setiap harinya. Contoh kegiatan tersebut yaitu mengeja harga barang, nilai data, tanggal dan tahun, serta masih banyak lagi. Meskipun terkesan hal sepele, berdasarkan penelitian dari para dokter di Indonesia, seorang anak baru bisa membaca dan mengeja angka pada umur 4-6 tahun. Rentang usia tersebut tentunya terasa kurang cepat. Padahal, semakin cepat seorang bisa membaca dan mengeja angka, maka semakin cepat pula anak tersebut dapat belajar berhitung dan mempelajari hal-hal lainnya, bahkan termasuk belajar pemrograman.

Dari permasalah di atas, maka diperlukanlah suatu sarana pembelajaran yang dapat membantu anak-anak balita di Indonesia untuk membaca dan mengeja angka. Dengan adanya solusi tersebut, diharapkan anak-anak dapat membaca dan mengeja angka lebih cepat sehingga mampu segera mempelajari hal-hal lebih besar lainnya dan tentunya meningkatkan tingkat pendidikan di Indonesia.

## Spesifikasi
1. Terdapat 2 buah endpoint API yaitu '**GET** /spell' yang menerima parameter angka, serta '**POST** /read' yang menerima body/payload berupa text/ejaan. 

2. API dibuat dengan menggunakan **Go** dan memanfaatkan **github.com/gin-gonic/gin** serta **github.com/gin-contrib/cors**

3. Response berupa JSON dengan format
**GET**
Response :
```
{
    "text" : string
}
```
**POST**
Response : (jika tidak terdapat error)
```
{
    "number" : int
}
```
Response : (jika terdapat error)
```
{
    "number" : string
}
```

4. Aplikasi website dibuat dengan teknologi **React.js** dan memanfaatkan **axios** untuk melakukan request ke API

## Contoh Kasus Uji
Semua kasus uji dijalankan dengan **POSTMAN** 
### Contoh Kasus Uji 1 : Pengejaan
Request :
```
GET '/spell?number=605004321'
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/1.png)

### Contoh Kasus Uji 2 : Pengejaan

Request :
```
GET '/spell?number=-32'
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/2.png)

### Contoh Kasus Uji 3 : Pembacaan
Request :
```
POST '/read'
{
    "text" : enam puluh juta tujuh ratus lima puluh ribu tiga ratus tiga belas 
}
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/3.png)

### Contoh Kasus Uji 4 : Pengejaan

Request :
```
POST '/read'
{
    "text" : enam puluh lima tujuh ribu
}
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/4.png)

## Penggunaan
- Install gin dan cors pada go
- Install npm dan axios pada react.js
- Jalankan api.go dengan cara mengetikkan "go run api.go" di terminal
- Rest API akan berjaan di 0.0.0.0:8080/
- Jalankan App React.js dengan cara mengetikkan npm start di direktori numeral_spellers pada terminal
- App akan berjalan di localhost:3000/

## Uji Coba App React JS
### Contoh Kasus Uji 1 : Pengejaan
![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/5.png)
### Contoh Kasus Uji 2 : Pengejaan
![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/6.png)
### Contoh Kasus Uji 3 : Pembacaan
![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/7.png)
### Contoh Kasus Uji 4 : Pebacaan
![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/8.png)