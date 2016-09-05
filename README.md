# dockermon-go


# Build
  - Prepare image for building. Run with ** --build-arg ** http_proxy, when running behind proxy. For example, `docker build -t boriska70/dockermon -f Dockerfile.build --build-arg http_proxy=http://1.2.3.4:5678 --build-arg https_proxy=http://1.2.3.4:5678 .`
  - Build the project: `docker run --rm -v "$PWD":/go/src/github.com/boriska70/dockermon-go boriska70/dockermon script/go_build.sh`
  - Create docker runtime image: docker build -t boriska70/dockermon-go .

# Run in docker
  - Assuming that we link to elasticsearch running as another docker:
  docker run --rm --name dmg --log-driver=json-file -v /var/run/docker.sock:/var/run/docker.sock --link es:es boriska70/dockermon-go -esurl=http://es:9200