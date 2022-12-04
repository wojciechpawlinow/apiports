# Ports domain service

## How to run  

### Container with the application

> ⚠️ _Make sure you have Docker installed_  

`make run` 

### Code linter and formatting

> ⚠️ _Make sure you have installed the [revive](https://github.com/mgechev/revive) package_.  
> `go install github.com/mgechev/revive@latest
`

`make lint`  

### Unit tests

`make test`

## Notes

1. Missing `Update` method. There are few samples with aliases which cover some of ports names.  
The question is which object is the newest version? API design could have been polished more in terms of the `Update` method 
2. That `Update` method could be fulfilled by a single method `Upsert()` at the persistence layer.
3. I struggled too much with Docker config (go + MongoDB) so eventually the application runs from host and mongo inside a container.  
I had docker-compose.yml, however, I don't want to keep any dead or not used code in the code base. The issue was that from the app container  
I wasn't able to ping the DB although I used proper name of the container, they were in the same bridget network and so on...
4. I added some metrics to show the performance of decoding JSON object by object in order to avoid loading all to a memory.
5. No HTTP server provided as the description doesn't explicitly tell about it.
6. Domain entity types have been assumed from that provided JSON file data.