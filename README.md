## Checking IMDB copycat for movies

This is a POC for a backend written in Go. I added to the db reqeust an cache.

If you want to tried it out, the simply downlaod the code and set your own variables into .env and try the backend
## Libs used

+ backend:
    * fiber - server
    * gorequest - client
+ database:
    * gorm - ORM
    * sqlite - database technology
+ util:
    * regex
    * godotenv - asking project variables from .env file
    * air - hot reloading
    
