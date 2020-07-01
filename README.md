# Product Images

## Uploading 

Note: need to use `--data-binary` to ensure file is not converted to text


## POST
```
curl -vvv localhost:9090/images/1/car.png -X POST --data-binary @car.png
```

## GET

```
curl -vvv localhost:9090/images/1/car.png --output car2.png 
```