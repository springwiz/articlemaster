# articlemaster
Catalog manager

Package implements a Catalog Manager which exposes Rest Apis for the Article/Item Management. It relies on Package gorilla/mux to manage routes and handle routing. The name mux stands for "HTTP request multiplexer". Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.The main features of the Catalog manager are:
  * Exposes the Rest Apis for Creating and Retrieving articles.
  * Exposes the Rest Apis for Tag Creation and Look up.
  * It is scalable and runs each HTTP Request in its own Go_Routine.
  * Employes Map based Indexes for efficient look up in the Data Storage.
  * Is thread safe and employes locking to ensure the data integrity.
  * The Data Repository is pluggable and could handle multiple database types.
  
# installation
1. Install the Go lang runtime on the local machine.
2. Create the path on the local machine at $GOPATH/src/github.com/springwiz. 
3. Change into articlemaster and Run go get github.com/springwiz/articlemaster.
4. Run go run main.go.

# assumptions
1. The package relies on in memory storage.
2. The Catalog Manager only implements the selected endpoints and leaves rest of the operations.
3. Throws back the errors encountered to the front end as JSON Messages.

# implementation details
  * The errors encountered on the backend are thrown from the REST Endpoints as JSON Messages.
  * The package relies on a mix of OOB Go Unit Testing and Postman for testing the solution.
  * The GoLang stack was chosen from the following:
      1. NodeJS 
      2. Java/SpringBoot
      3. GoLang/Gorilla Mux
    
    The following are the considered factors:
    The NodeJS stack is more suitable for business processes which are I/O bound. The NodeJS Event Loop is not meant to run longer       
    computations.
    The Java/SpringBoot based services are heavy weight and have much larger memory footprint. They have much higher boot time and 
    are not preferable for web scale applications. The Java's model of 1 thread per request fails for web scale applications as thread 
    context switches take most of the time/memory.
    The Golang services are light weight and offer rich support for concurrency via go routines and channels. The go routines are light     weight and offer rich support for parellelism.
