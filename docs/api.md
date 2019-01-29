# API Documentation

# API

The API listens on :6227 by default. This can be set by setting the `addr` ENV variable.

## Routes
PUT http://HOST:6227/apply
```json
{
	"type": "route",
	"list": [
		{
			"id": "1",
			"domainName": "team142.co.za",
			"address": "192.168.88.24:9000"
		}		
	]
}
```

## Services
PUT http://HOST:6227/apply
```json
{
	"type": "service",
	"list": [
		{
			"id": "1",
			"name": "nginx1",
			"image": "nginx:latest",
			"replicas": 1,
			"memoryLimit": 32
		}		
	]
}
```