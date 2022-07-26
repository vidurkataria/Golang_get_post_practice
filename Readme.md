# Post Method call
It will add the data to a userInfo struct on the server.
curl command - curl -v -X POST -d '{"name":"NewUser", "email": "newuser@gmail.com"}' http://127.0.0.1:9342/add_data/


# Get Method call
It will return all the data in json format
curl command -  curl -v -X GET http://127.0.0.1:9342/add_data/