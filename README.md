# Product Images

## Uploading 

Note: need to use `--data-binary` to ensure file is not converted to text


## POST
```
curl -vvv localhost:3333/images/1/car.png -X POST --data-binary @car.png
```

## GET

```
curl -vvv localhost:3333/images/1/car.png --output car2.png 
```
**Now support __--compressed__ as well**
```
curl -vvv localhost:3333/images/1/car.png --compressed --output car3.png 
```

## POST Multi-Part

````http request
POST / HTTP/1.1
Host: localhost:3333
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="id"

5
------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file"; filename="gopher.jpeg"
Content-Type: file

------WebKitFormBoundary7MA4YWxkTrZu0gW--
````

```shell script
curl --location --request POST 'http://localhost:3333/' \
--form 'id=5' \
--form 'file=@gopher.jpeg'
```
