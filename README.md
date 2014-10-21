gadgetWish
==========


###Required
* Go > 1.2.x
* MongoDB  > 2.5.x

###Adding Go dependencies
#####mgo
--------------
      go get gopkg.in/mgo.v2

#####gorilla
--------------
     go get github.com/gorilla/mux
     go get github.com/gorilla/sessions

###Running the App
####Bootstrapping
* Open bootstrap/main.go
* Modify the local constant
  * dbURI
  *  dbname
* Run bootstrap script like bellow
--------------
      cd app-path/bootstrap
      go run main.go

####Configure
Rename example.config.json to config.json and add mongodb connection string and db name

####Building app
* Simply start the server
--------------
     cd app-path/
     go run server.go
* Alternatively you can build and execute 
--------------
      go build server.go
      ./server
* Now go to you favourite browser and check http://127.0.0.1:8080
