# golang-tour
Learning Go and cool tricks discovered journeying on this path

Please check this article for setting up a go development environment for an ubuntu distro. Or use this ansible role to set up if you already have ansible installed.

 # hello.go

Now that we have written this snippet to print "Hello Go!" at our terminal. We have two ways to run this application
  - ```go run src/github.com/knoxknot/golang-tour/hello.go```   this runs the hello.go program and you see output on the fly
  - <code>go build github.com/knoxknot/golang-tour/</code> we usually use to to build several go packages into a single binary file. Then enter the file name and hit the return button to run.<quote> note we omitted the "src" path name. </quote>