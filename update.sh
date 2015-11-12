docker kill stalks && docker rm stalks && docker build -t stalks-build . && docker run -p 15000:8000 -it -d --name stalks stalks-build
