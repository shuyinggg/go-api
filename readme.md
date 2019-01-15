1. Add logrus package and use for logging
2. Implement 3 REST APIs
 2.1 GET /cookie will return a Json of all cookies or an empty json if none are present
 2.2 POST /cookie Take a JSON input of an array of name value pairs and sets those cookies
 2.3 DELETE /cookie Will delete all the cookies in the request
3. Write a unit tests to make sure that 2.1, 2.2 and 2.3 works correctly.
4. write a package called cookie_handler that will do the core logic for all cookie parsing.

Guidelines
1. Make sure all the error cases are handled. 
  e.g is client sends invalid json input or duplicate json input

   invalid json input
   respond with message {"error": true, "message": "invalid json"}

   For duplicate json input, there is no error, the response just sets the last value of the cookie.

Definition of Done.

1. There should be a good documentation for how someone can use these APIS
2. The test cases should run and cover all usecases
3. Code is neatly organized

Godoc
readme
t.log()
error

      


#RESTFUL API IN GO
This project contains four RESTful APIs method: GET, POST, PUT, and DELETE that sends corresponding requests to the requested url.
To use these APIs, download Postman and run main.go, handlers.go, cookie_handler.go.


## Check connection /
Open a browser page and navigate to [localhoat:4040/](localhoat:4040/). You should see "Hello there!" 

Then use Postman and navigate to /cookie

## GET /cookie
GET method returns all the cookies under localhost:4040/cookie in JSON format.
### To use
To use get, open Postman and put localhost:4040/cookie as in the url address box. Next,
select GET in the drop down box to the left of the url. 
If there is no cookie available, the program will return an empty JSON as [].
There should also be a "GET SUCCESS" message appearing in the terminal.

3. POST
POST method takes an JSON input of an array of name and value pairs and sets them as
cookies to localhost:4040/cookie.
Similarly, to use get, open Postman and put localhost:4040/cookie as in the url address box. Next,
select POST in the drop down box to the left of the url. Then select the body label down below and mark
"raw", the put the input in the big blank box.
The input must be a valid JSON input without any duplicated pairs. The input can not be empty.
Duplicated pairs mean that the cookies can't have the same name, and this is only true for current input.
Sample input can be found in samplecookies.text.
If the user enters a not JSON input or an empty input, the program should exit and return "Invalid JSON! Please check input".
If the user enters duplicated pairs, the program should exit and return "Duplicate! Please check input".
If either of the case happens, Postman will return nothing and no cookies will be posted.
If success, in terminal, there should be a "POST SUCCESS" message.

4. PUT
PUT method POST method takes an JSON input of an array of name and value pairs which names already exist in the current cookie,
and update their values to localhost:4040/cookie.
Similarly, to use get, open Postman and put localhost:4040/cookie as in the url address box. Next,
select PUT in the drop down box to the left of the url. Then select the body label down below and mark
"raw", the type the input in the big blank box.
JSON input rules are the same as method POST. Note that if the user "PUT"s a cookie that doesn't appear before,
the program will treate it as a "POST" request and update other cookies.
If success, in terminal, there should be a "PUT SUCCESS" message.

5. DELETE
DELETE method deletes all cookies under localhost:4040/cookie. 
To use delete, open Postman and put localhost:4040/cookie as in the url address box. Next, 
select DELETE in the drop down box to the left of the url. The program will return the cookies deleted. 
If success, in terminal, there should be a "DELETE SUCCESS" message.
If there are no cookies, the program will still run and return an empty JSON.




