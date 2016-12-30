#YaraScanner
---

Microservice for scanning files with yara

Available functions are

- GET /scanner/v1/files/ - list all uploaded binaries

- PUT /scanner/v1/files/{filename} - upload a binary

- GET /scanner/v1/files/{filename} - download a copy of specified file

- DELETE /scanner/v1/files/{filename} - remove specified file

- GET /scanner/v1/files/{filename}/scan/ - scan specified file

- GET /scanner/v1/ruleset/ - list all loaded ruleset

- GET /scanner/v1/ruleset/{ruleset} - list all rules from a loaded ruleset
---
```
Usage of ./yarascanner:
  -address string
    	address to bind to (default "0.0.0.0")
  -port string
    	port to bind to (default "9999")
  -i string
    	path to yara rules index file
  -uploads string
    	path to uploads directory (default "uploads")
```
---
Requirements
ubuntu/debian
libyara-dev
