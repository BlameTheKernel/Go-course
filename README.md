# Go-course
Installation in Linux:
Get the tar file with the last version (check the last one in https://golang.org/dl/) and substitute the number.
curl -O https://storage.googleapis.com/golang/go1.13.8.linux-amd64.tar.gz

Extract the file with:
tar -xvf go1.13.8.linux-amd64.tar.gz 

give the new dir root permissions and move it to /usr/local:
sudo chown -R root:root ./go
sudo mv go /usr/local

Add at the end of the ~/.profile file:
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

Execute the new ~/.profile:
source ~/.profile

Check if everything is all right with the version
go version
go version go1.13.8 linux/amd64

Each file has to have each own directory and everytime you modify something you must compile it with:
~/go/src/hello$go build (inside the dir your program is, in this case hello.go)
then:
~/go/src/hello$./hello
hello, world
