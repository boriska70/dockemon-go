# dockermon-go


### Build
  - Prepare image for building. Run with **--build-arg http_proxy**, when running behind proxy, assuming that http_proxy and https_proxy environment variables are set. For example, `docker build -t boriska70/dockermon -f Dockerfile.build --build-arg http_proxy=http://1.2.3.4:5678 --build-arg https_proxy=http://1.2.3.4:5678 .`
  - Build the project: `docker run --rm -v "$PWD":/go/src/github.com/boriska70/dockermon-go boriska70/dockermon script/go_build.sh`
  - Create docker runtime image: `docker build -t boriska70/dockermon-go .`

### Run in docker
  - Assuming that we link to elasticsearch running as another docker named es:
  `docker run --rm --name dmg --log-driver=json-file -v /var/run/docker.sock:/var/run/docker.sock --link es:es boriska70/dockermon-go -esurl=http://es:9200`
  - Elasticsearch can be started as
  `docker run -d --name es -p 9200:9200 -p 9300:9300 elasticsearch:2.3.4 elasticsearch -Des.network.host=0.0.0.0 -Des.network.bind_host=0.0.0.0 -Des.cluster.name=elasticlaster -Des.node.name=$(hostname)`
  - Kibana run: docker run --link es:elasticsearch -d kibana

Useful:
  - curl --unix-socket /var/run/docker.sock http:/containers/json (see https://docs.docker.com/engine/reference/api/docker_remote_api/ for details)
  - start/stop container events:
  `[
      {"status":"start","id":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","from":"alpine","Type":"container","Action":"start","Actor":{"ID":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","Attributes":{"image":"alpine","name":"alpine"}},"time":1473106693,"timeNano":1473106693262908400},
      {"Type":"network","Action":"connect","Actor":{"ID":"e893d978e108d8ac175fae938ed02d12f9f3570843586b43606e4c083a62facc","Attributes":{"container":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","name":"bridge","type":"bridge"}},"time":1473106692,"timeNano":1473106692841554700},
      {"status":"kill","id":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","from":"alpine","Type":"container","Action":"kill","Actor":{"ID":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","Attributes":{"image":"alpine","name":"alpine","signal":"15"}},"time":1473107891,"timeNano":1473107891869814300},
      {"status":"kill","id":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","from":"alpine","Type":"container","Action":"kill","Actor":{"ID":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","Attributes":{"image":"alpine","name":"alpine","signal":"9"}},"time":1473107901,"timeNano":1473107901871658000},
      {"status":"die","id":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","from":"alpine","Type":"container","Action":"die","Actor":{"ID":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","Attributes":{"exitCode":"137","image":"alpine","name":"alpine"}},"time":1473107901,"timeNano":1473107901922774300},
      {"Type":"network","Action":"disconnect","Actor":{"ID":"e893d978e108d8ac175fae938ed02d12f9f3570843586b43606e4c083a62facc","Attributes":{"container":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","name":"bridge","type":"bridge"}},"time":1473107902,"timeNano":1473107902267808500},
      {"status":"stop","id":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","from":"alpine","Type":"container","Action":"stop","Actor":{"ID":"405efae3b420464a9da7be7cd9de8d2ff160ffcfdac01517d9b686e8137f9053","Attributes":{"image":"alpine","name":"alpine"}},"time":1473107902,"timeNano":1473107902406099500}
      ]`
  - Event structure format: https://godoc.org/github.com/fsouza/go-dockerclient#APIEvents
  - Possible events list: https://docs.docker.com/engine/reference/commandline/events/
