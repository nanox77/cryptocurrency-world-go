## Cryptocurrency Go example using Gin

System has 3 layers:
1. Handlers     -> Its responsibility is getting request and parse data. 
2. Actions      -> Its responsibility is use models to build business logic.
3. Repositories -> Its responsibility is return data from internal (memory) or external (databases, apis, etc) services. 

Base url:
```
https://glistening-invention-production.up.railway.app/
```

### Endpoints available
```
GET    /users
POST   /users/:userId/deposit
POST   /users/:userId/withdraw          
POST   /users/:userId/buy
POST   /users/:userId/sell
POST   /users/:userId/transfer
```

### Test api
You can use `resources/test_api.json` file to test api from Postman.

### Notes
- In order to reduce complexity of exercise, BTC and ETH prices are set via environment variables.
The best option will be getting from external service like Coinbase API or Blockchain API.


- In order to avoid double-spending I implemented a lock system. 
When client use the API, user is locked until request ends and then is unlocked.


- In order to reduce complexity of exercise, system has 3 users with its respective wallet.


- I used [railway](https://railway.app/) to deploy