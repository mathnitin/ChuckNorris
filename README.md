# ChuckNorris Project.

## PROBLEM STATEMENT

Create a web service which combines two existing web services.

Fetch a random name from http://uinames.com/api/
Fetch a random Chuck Norris joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]
Combine the results and return them to the user.

Example
Fetching a name
$ curl http://uinames.com/api/
{"name":"Δαμέας","surname":"Γιάνναρης","gender":"male","region":"Greece"}

Fetching a joke
$ curl 'http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=\[nerdy\]'
{ "type": "success", "value": { "id": 181, "joke": "John Doe's OSI network model has only one layer - Physical.", "categories": [“nerdy”] } }

Using the new web service
$ curl ‘http://localhost:5000’
Δαμέας Γιάνναρης’s OSI network model has only one layer - Physical..


## INSTRUCTIONS FOR RUNNING WEB SERVICE

#### Prep
mkdir -p $GOPATH/src/github.com/mathnitin
cd $GOPATH/src/github.com/mathnitin

#### Checkout git
git clone https://github.com/mathnitin/ChuckNorris.git
cd ChuckNorris

#### Instructions to run binary
###### Run web-service.
NITINMAT-M-K12N:ChuckNorris nitinmat$ go run main.go


#### Instructions to build and run Docker container 
###### Build Docker image
NITINMAT-M-K12N:ChuckNorris nitinmat$ docker build -t chucknorris .
###### Run Docker image
NITINMAT-M-K12N:ChuckNorris nitinmat$ docker run -p 5000:5000 chucknorris:latest
( The above command will start a webserver on localhost:5000 port )
