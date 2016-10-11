#YaraScanner
---

Microservice for scanning files with yara

Available functions are

- /scanner/v1/list/ - list all uploaded binaries

- /scanner/v1/{filename}/scan/ - scan specified file

- /scanner/v1/{filename}/remove/ - remove specified file

- /scanner/v1/{filename}/download/ - download a copy of specified file

- /scanner/v1/{filename}/upload/ - upload a binary

---
```
Usage of ./yarascanner:
  -address string
    	address to bind to (default "0.0.0.0")
  -port string
    	port to bind to (default "9999")
  -rules string
    	path to yara rules (default "rules")
  -uploads string
    	path to uploads directory (default "uploads")
```
