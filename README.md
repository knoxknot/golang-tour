# golang-tour
Learning Go and cool tricks discovered journeying on this path

Please check this article for setting up a go development environment for an ubuntu distro. Or use this ansible role to set up if you already have ansible installed.

 # hello.go

Having written this snippet to print "Hello Go!" at our terminal, we have two ways to run this application.
  - ```go run src/github.com/knoxknot/golang-tour/hello.go```   this runs and returns an output for the hello.go program.
  - <code>go build github.com/knoxknot/golang-tour/</code> This builds several go packages into a single binary. Then we enter the binary name at the termninal and hit return to run.<aside> note that we omitted the "src" path name. </aside>