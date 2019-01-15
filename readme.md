
# RESTFUL API IN GO

This project contains four RESTful APIs method: GET, POST, PUT, and DELETE that sends corresponding requests to the requested url.
To use these APIs, download [Postman](https://www.getpostman.com/).
In terminal, enter folowing lines and replace some_path with your main folder path.
```
cd some_path
go run main.go handlers.go cookie_handler.go

``` 

## Check connection 

After running, open a browser page and navigate to [localhoat:4040/](localhoat:4040/). You should see "Hello there!" 

Then try API in Postman.

## POST

POST method takes an JSON input of an array of name and value pairs and sets them as cookies to the requested url.

### Example


#### Url:
```
localhost:4040/cookie
```
#### Input: 
```
[
    {
        "name": "Cookie1",
        "value": "This is cookie 1"
    },
    {
        "name": "Cookie2",
        "value": "This is cookie 2"
    }

]
```

#### Method: POST

#### Output:

```
[
    {
        "Name": "Cookie1",
        "Value": "This is cookie 1",
        "Path": "",
        "Domain": "",
        "Expires": "0001-01-01T00:00:00Z",
        "RawExpires": "",
        "MaxAge": 0,
        "Secure": false,
        "HttpOnly": false,
        "SameSite": 0,
        "Raw": "",
        "Unparsed": null
    },
    {
        "Name": "Cookie2",
        "Value": "This is cookie 2",
        "Path": "",
        "Domain": "",
        "Expires": "0001-01-01T00:00:00Z",
        "RawExpires": "",
        "MaxAge": 0,
        "Secure": false,
        "HttpOnly": false,
        "SameSite": 0,
        "Raw": "",
        "Unparsed": null
    }
]

```
#### Errors:

Invalid Json input and Duplicate input.
```
[
    {
        "err": "true",
        "msg": "Duplicate! Please check input"
    }
]
```

## GET
GET method returns all the cookies under the requested url in JSON format.


### Example

#### Url:
```
localhost:4040/cookie
```
#### Input: nil

#### Method: GET

#### Output:
```
[] //if there is no cookie.

```

or

```
//cookies
[
    {
        "name": "Cookie1",
        "value": "This is cookie 1"
    },
    {
        "name": "Cookie2",
        "value": "This is cookie 2"
    }
]
```


## PUT

PUT method takes an JSON input, updates the existing cookies and post any new cookies in the input.

### Example

#### Url:
```
localhost:4040/cookie
```
#### Input: 
```
//Suppse cookie1 is posted as above
[
    {
        "name": "Cookie1",
        "value": "This is cookie 100"
    }
]
```

#### Method: PUT

#### Output:

```
[
    {
        "Name": "Cookie1",
        "Value": "This is cookie 100",  //UPDATED
        "Path": "",
        "Domain": "",
        "Expires": "0001-01-01T00:00:00Z",
        "RawExpires": "",
        "MaxAge": 0,
        "Secure": false,
        "HttpOnly": false,
        "SameSite": 0,
        "Raw": "",
        "Unparsed": null
    }
]

```
#### Errors:

Invalid Json input and Duplicate input.
```
[
    {
        "err": "true",
        "msg": "Invalid input! Please check input"
    }
]
```

## DELETE

DELETE method deletes all cookies of the requested url

### Example

#### Url:
```
localhost:4040/cookie
```
#### Input: nil

#### Method: DELETE

#### Output:

```
//The cookie(s) that has just been deleted
[
    {
        "Name": "Cookie1",
        "Value": "This is cookie 100",  //UPDATED
        "Path": "",
        "Domain": "",
        "Expires": "0001-01-01T00:00:00Z",
        "RawExpires": "",
        "MaxAge": 0,
        "Secure": false,
        "HttpOnly": false,
        "SameSite": 0,
        "Raw": "",
        "Unparsed": null
    }
]

```
Or
```
//No cookies deleted.
[] 
```

## Tests
To run all the tests, make sure server is running.
Then in a new terminal window, run
```
go test -v
```




