## Mapped

Shorten that URL...

### Motivations

There are 2 apps almost every developer has built. A todo app and A url
shortner.. I realized that I haven't built a url shortnening application and I
been looking forward to learning devops. So here, I am building
a ___convoluted, sophisticated, web-scale___ url shortner with modern practices.


### Stack

- Deployment
  - Docker
- Services
  - Consul
  - Redis
  - MongoDB (an hosted service, maybe Amazon RDS or google cloud sql.. Hell, a database in a container ???? )
  - CI with Drone
  - CD with Drone
  - Monitoring ( Whatever can go wrong will eventually go wrong a.k.a Javascript
    rule)
    - Prometheus
    - Grafana
  - Oklog

- Languages
  - Golang
  - JS (React)


### Tests

-  `MAPPED_CONSUL_ADDR` for consul server to use during tests.


### RoadMap

- [ ] Dashboard
  - [ ] List all links
  - [ ] Link sorting
  - [ ] Link statistics / analytics

