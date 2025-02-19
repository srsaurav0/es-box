# Product Search Auto-Suggestion Beego


## Setup and Installation

1. Clone the repository
   ```
   git clone https://github.com/srsaurav0/es-box.git
   cd es-box
   ```

2. Create config file in conf/app.conf
   ```
   touch conf/app.conf
   ```

3. Copy configurations from app.conf and add API key.
   ```
   appname = es-box
   httpport = 8080
   runmode = dev
   ES_LOCAL_API_KEY=QnZnMjdwUUJfZXVoNWRBbE1MaTg6c19PM0hWUVFRay1QM0QyLXNuWE1fZw==
   ES_LOCAL_URL=http://elasticsearch:9200
   ```

4. Run docker container
   ```
   docker compose up --build
   ```
   This command initializes the app in **localhost:8080**


## App Usage

- Go to ***[localhost:5601](http://localhost:5601/app/home#/tutorial_directory/sampleData)*** and add **Sample eCommerce orders**

- Go to **localhost:8080** and search for products in the search bar.