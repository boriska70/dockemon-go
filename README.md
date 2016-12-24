[![CircleCI](https://circleci.com/gh/boriska70/dockermon-go.svg?style=svg)](https://circleci.com/gh/boriska70/dockermon-go)
[![](https://badge.imagelayers.io/boriska70/dockermon:latest.svg)](https://imagelayers.io/?images=boriska70/dockermon:latest 'Get your own badge on imagelayers.io')

# dockermon-go


### Build
 - docker build -t boriska70/dockermon-builder -f Dockerfile.build .
 - docker run -it --name=dockermon-builder boriska70/dockermon-builder
 - mkdir .dist
 - docker cp  dockermon-builder:/go/src/github.com/boriska70/dockermon-go/.dist/dockermon ./.dist/
 - docker build -t boriska70/dockermon .

### Run in docker
  - Assuming that we link to elasticsearch running as another docker named es:
  `docker run --rm --name dmg --log-driver=json-file -v /var/run/docker.sock:/var/run/docker.sock --link es:es boriska70/dockermon-go -esurl=http://es:9200`
  - Elasticsearch can be started as
  `docker run -d --name es -p 9200:9200 -p 9300:9300 elasticsearch elasticsearch -Des.network.host=0.0.0.0 -Des.network.bind_host=0.0.0.0 -Des.cluster.name=elasticlaster -Des.node.name=$(hostname)`
  - Kibana run: docker run --link es:elasticsearch -d -p5601:5601 --name kibana kibana

Useful:
  - Get containers: curl --unix-socket /var/run/docker.sock http:/containers/json (see https://docs.docker.com/engine/reference/api/docker_remote_api/ for details)
  - Stream events: curl --unix-socket /var/run/docker.sock http://localhost/events (see https://docs.docker.com/engine/reference/api/docker_remote_api/ for details)
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
