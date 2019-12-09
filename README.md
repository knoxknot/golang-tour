# golang-tour
Learning Go and cool tricks discovered journeying on this path

Please check this [article](https://medium.com/@nwoyesamuelc/a-neater-go-local-development-environment-hello-go-9ba77598f7b4) for setting up a go development environment on an ubuntu distro. Or use the [go ansible role](https://github.com/knoxknot/ansible-roles.git) from this repository to set up if you already have ansible installed.

I have also added 1 to every go filename to avoid conflicting with the current file worked on at any given time. Thus to run that source file as a go program simply remove the 1 behind the .go extension. Please do note also that since the intention was a successive tutorial of showing the design and concept of go language, each of these files have a main function which will conflict with another if both files are in .go extensions.  

**Tips and Tricks**
 git request-pull master origin dev


 # hello.go
Having written this snippet to print "Hello Go!" at our terminal, we have two ways to run this application.
  - ```go run src/github.com/knoxknot/golang-tour/hello.go```   this runs and returns an output for the hello.go program.
  - <code>go build github.com/knoxknot/golang-tour/</code> This builds several go packages into a single binary. Then we enter the binary name at the termninal and hit return to run.<aside> note that we omitted the "src" path name. </aside>
