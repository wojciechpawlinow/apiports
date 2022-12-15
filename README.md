# Ports domain service

## How to run  

### Container with the application

> _Make sure you have Docker installed_  

1. `make run` 

2. Try to get a port `curl localhost:8080/port/AEAJM`
```json
{
  "id": "AEAJM",
  "name": "Ajman",
  "city": "Ajman",
  "country": "United Arab Emirates",
  "alias": [],
  "regions": [],
  "coordinates": {},
  "province": "Ajman",
  "timezone": "Asia/Dubai",
  "unlocs": [
    "AEAJM"
  ],
  "code": "52000"
}
```

### Code linter and formatting

> _Make sure you have installed the [revive](https://github.com/mgechev/revive) package_.  
> `go install github.com/mgechev/revive@latest
`

`make lint`  

### Unit tests

`make test`