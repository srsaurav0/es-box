## Setup and Installation

1. Create config file in conf/app.conf
```
touch conf/app.conf
```

2. Copy configurations from app.conf and add API key.
```
appname = es-box
httpport = 8080
runmode = dev
ES_LOCAL_API_KEY=QnZnMjdwUUJfZXVoNWRBbE1MaTg6c19PM0hWUVFRay1QM0QyLXNuWE1fZw==
ES_LOCAL_URL=http://elasticsearch:9200
```

3. Run docker container
```
docker compose up --build
```