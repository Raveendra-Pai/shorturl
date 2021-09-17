# shorturl
Sample GoLang code to generate short url from long url.

This is my first attempt to learn the GoLang and exercise the tools the language offers. 

Basically this app hosts REST API to generate shorturl by receiving long url in HTTP Post, with the path "/Posturl". Long url will be sent as POST request body content.
On succesful POST, Application generates short urls and return shorturl as Response body. 
User can use that response shorturl and make a HTTP GET request through browser or from any http client application,
this application will receives the GET request and redirect the url to the original long url website.

The Application maintains map internally, which will store the short url as the key and long url as the value. Alternatively this key/values can be stored on NoSQL or SQL Databases if user wants persistent data. 


