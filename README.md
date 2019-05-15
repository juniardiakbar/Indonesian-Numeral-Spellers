# Indonesian Numeral Speller

## Spesification
1. There are 2 API endpoint '**GET** /spell' and '**POST** /read'. GET receive number as parameter and POST recieve body/payload as text spell

2. API is made with **Go** and using **github.com/gin-gonic/gin**, **github.com/gin-contrib/cors**

3. Response given in JSON with format:

**GET**
Response :
```
{
    "text" : string
}
```
**POST**
Response : (if there is no error)
```
{
    "number" : int
}
```
Response : (if there is an error)
```
{
    "number" : string
}
```


4. Website App is made with **React.js** technology and using **axios** for doing request to API

## Test Case 
### Test Case 1 : Spelling
Request :
```
GET '/spell?number=605004321'
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/1.jpg)

### Test Case 2 : Spelling

Request :
```
GET '/spell?number=-32'
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/2.jpg)

### Test Case 3 : Reading
Request :
```
POST '/read'
{
    "text" : enam puluh juta tujuh ratus lima puluh ribu tiga ratus tiga belas 
}
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/3.jpg)

### Test Case 4 : Reading

Request :
```
POST '/read'
{
    "text" : enam puluh lima tujuh ribu
}
```
Response :

![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/4.jpg)

## How To Use
- Install gin and cors in go
- Install npm and axios in react.js
- Run api.go by typing **"go run api.go"** in terminal
- Rest API will running in 0.0.0.0:8080/
- Run React.js App by typing **"npm start"** in numeral_spellers directory in terminal
- App will running in localhost:3000/

## React JS App Testing
### Test Case 1 : Spell
![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/5.jpg)
### Test Case 2 : Read
![alt text](https://raw.githubusercontent.com/juniardiakbar/Indonesian-Numeral-Spellers/master/pict/6.jpg)
