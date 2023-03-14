To run this file you need to:

`docker compose up`

Here are the urls for you to run with postman or thunderclient (check out thunderclient, its postman but an extension for VS code) or what have you.

POST:
http://localhost:8000/store
    body: 
    {"Value": "WHATEVER_VALUE",
    "Key":"WHATEVER_KEY"
    }


GET: http://localhost:8000/store/WHATEVER_KEY

DELETE: http://localhost:8000/store/WHATEVER_KEY

As you can see, I followed the REST api practices. This handles GET, POST and DELETE HTTP keywords. 
I also have the system overwrite by simply making an overwrite with a key. 


I moved around the routers into "services" folder, that way it's outside of main. I used DB.client() in a single instance so it's not constantly creating new instances of the redis client. I did this for performance reasons.
I also moved the client into services as it's own service to be able to move this around for whatever reason. To make changes to the DB.client(), I would probably have to remove the instancing I did with the client to just every Post





Tests:
This covers deletion and overwriting:

`docker-compose exec app sh -c "cd tests && go test -v endpoint_test.go"`

In here, I use requests to create a key-value and then either delete them with another request or I overwrite them with another post request.

