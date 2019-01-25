# Planck's Cloud

<img src="https://travis-ci.org/plancks-cloud/plancks-cloud.svg?branch=master" />&nbsp;
<img src="https://goreportcard.com/badge/github.com/plancks-cloud/plancks-cloud"> &nbsp;
<a href="https://codeclimate.com/github/plancks-cloud/plancks-cloud/maintainability"><img src="https://api.codeclimate.com/v1/badges/81aff827de3938808c2d/maintainability" /></a>&nbsp;
[![codebeat badge](https://codebeat.co/badges/25407218-e856-4f5e-ac7c-9d045dc0fe5a)](https://codebeat.co/projects/github-com-plancks-cloud-plancks-cloud-master)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)

# Running

```
git clone https://github.com/plancks-cloud/plancks-cloud.git
go mod download
go mod vendor
go run application.go

```

# Building for Arm

```
./build-arm.sh
```

# API

The API listens on :6227 by default. This can be set by setting the `addr` ENV variable.

## Routes
PUT http://HOST:6227/apply
```
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
```
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

# Proxy

The API listens on :6228 by default. This can be set by setting the `proxy` ENV variable.

