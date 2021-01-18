# project-jc

- This project is written in Go using only the standard library.
- When launched, it monitors port 8080 and waits for http connections.
- It supports four endpoints:

  * A POST to /hash accepts a password in the body, in the format "password=MySecretPasswordSample",
    and returns a unique jobId immediately. It then, in the background (via a go routine), waits 5 seconds
    and then computes the SHA512 password hash for the specified password and associates the result
    with the jobId that was returned for this hash request.
    
  * A GET to /hash/{jobId} returns the jobId's corresponding base64-encoded SHA512 password hash.
      NOTE: If a valid jobId is specified before the SHA512 password hash has been computed, an "" is returned.
            If no jobId is specified, then the error "missing job ID" is returned.
            If an invalid jobId is specified, i.e. "xyz4", then the error "invalid job ID" is returned.
            If an unassigned jobId is specified, i.e. "6000000", then the error "job ID not found" is returned.
      
  * A GET to /stats returns JSON data structure {"total":1000,"average":127} where the total value 
    is the total hash request count since the server was started and average value is the average
    time of a hash request in milliseconds.
    
  * A GET to /shutdown triggers a graceful shutdown. That is, no new requests are allowed, and all 
    in-progress hash computations are completed before the server executable exits.
    
- The software supports multiple simultaneous connections.
- The software supports a graceful shutdown when a ctrl-C is invoked at the cmd-line. That is, no 
  new requests are allowed, and all in-progress hash computations are completed before the server
  executable exits.
  
  Building and Running the Server:
  - cd to dir project-jc
  - invoke: go build
  - run the server executable by invoking: project-jc (or project-jc.exe in Windows)
  - to gracefully stop the server, invoke a ctrl-C at the cmd-line
  
  Building and Running the testclient:
  - cd to dir project-jc/testclient
  - invoke: go build
  - run the testclient by invoking: testclient (or testclient.exe in Windows)
  - you may run multiple instances of the testclient at the same time
  - to stop the testclient, invoke a ctrl-C at the cmd-line
  
  Building and Running the data unit tests:
  - cd to dir project-jc/data
  - invoke: go test
  
Please contact Lino Pereira at 303-408-5369, or at lino.pereira@comcast.net, with any questions or issues.
