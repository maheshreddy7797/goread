# Goread
Read a text file and create a logs.txt file
## Usage:
```
  git clone https://github.com/maheshreddy7797/goread
  cd goread
  sudo docker build -t test .
  sudo docker run -v /path/to/file:/datadir test go run /datadir/filename.txt
```
Example in my case: sudo docker run -v /home/cloudbyte/testdir:/datadir test go run /datadir/read.txt
