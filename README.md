# Snipbox

<a href="https://golang.org/doc/go1.23"><img alt="Go 1.23" src="https://img.shields.io/badge/golang-1.23-blue?logo=go&color=5EC9E3"></a>

Repo for snippetbox project as part of Let's Go! book by Alex Edwards 
(found [here](https://lets-go.alexedwards.net))

## ✨ My Changes
This section highlights the modifications made to the original project as described in the Let's Go book:

- Creating a Docker-driven Go development workflow
- Add live realod using [air-verse/air](https://github.com/air-verse/air)
- Add Docker image (production)

---
## Setup Mysql docker container 

### Docker Compose File

```yml
services:
  web:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: snipbox-dev
    ports:
      - "4000:4000" # default app port
      - "4001:4001" # air proxy port (Optional)
    environment:
      - PORT=4000 # default app port 
    volumes:
      - ./:/app
    networks:
      - snipbox-net
    depends_on:
      db:
        condition: service_healthy

  db:
    image: mysql:lts 
    container_name: snipbox-db
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: snippetbox
      MYSQL_USER: web
      MYSQL_PASSWORD: pass
    ports:
      - "3306:3306" # Optinal
    volumes:
      - mysql-data:/var/lib/mysql
    networks:
      - snipbox-net
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 5s
      timeout: 5s
      retries: 2
  
volumes:
  mysql-data:
    external: true # Use external vluume for data persistence

networks:
  snipbox-net:
    driver: bridge
```

### Setup User

>  [!IMPORTANT]  
>  **We use `%` instead of `localhost` because**:
> - `'localhost'` restricts the user to connect **only from the same machine** where the MySQL server is running.
> - In this workflow, MySQL runs in a separate container, so `localhost` doesn't work for external connections.
> - Using `%` allows the user to connect from any host, enabling communication between containers.


1. Create User
```sql
CREATE USER 'web'@'%';
```

2. Grant access
```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';
```

```sql
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'%' IDENTIFIED BY 'pass';
```

3. Check users
```sql
SELECT user, host FROM mysql.user;
```

4. Connection String
```go
dsn := flag.String("dsn", "web:pass@tcp(mysql:3306)/snippetbox?parseTime=true", "MySQL data source name")
```

## Docker Compose Production

> [!TIP]
> Create self signed tls certificate & create `.env` before run docker compose
> the generate_cert.go file should be located under  `/usr/local/go/src/crypto/tls` or `/usr/lib/go/src/crypto/tls` (Manjaro/Arch)
> 
> #### Generate Certificate
> ```bash
> mkdir -p tls
> cd tls
> go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
> ``` 

```yml
services:
  web:
    image: ghcr.io/nmdra/snipbox:latest
    container_name: snipbox-prod
    ports:
      - "4000:4000" # default app port
    env_file:
      - .env
    volumes:
      - ./tls/:/tls/ # TLS cert directory
    networks:
      - snipbox-net
    depends_on:
      db:
        condition: service_healthy

  db:
    image: mysql:lts 
    # container_name: snipbox-db
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: snippetbox
      MYSQL_USER: web
      MYSQL_PASSWORD: pass
    ports:
      - "3306:3306" # Expose MySQL port (Optional)
    volumes:
      - mysql-data:/var/lib/mysql
    networks:
      - snipbox-net
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 5s
      timeout: 5s
      retries: 2
  
volumes:
  mysql-data:
    external: true

networks:
  snipbox-net:
    driver: bridge
```

<div align="center">
  <a href="blog.nimendra.xyz"> 🌎 nmdra.xyz</a> |
  <a href="https://github.com/nmdra"> 👨‍💻 Github</a> |
  <a href="https://twitter.com/nimendra_"> 🐦 Twitter</a>
</div>
