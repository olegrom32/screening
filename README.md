# screening

## Q1

Please find the solution with comments in [/q1/main.go](/q1/main.go)

## Q2

Depends on what is considered a "module".

Generally, I would depend on at least unit tests covering all the business logic. Ideally, functional or e2e tests as well.
If there's no tests, I would start with planning tech debt tickets to cover the module with tests.

Then, depending on the size of the module, refactoring may be done as one task/pull request or as many. Some or maybe all the unit tests will have
to be rewritten but the test cases must stay almost the same. Functional tests are more likely to stay the same,
indicating if the already existing functionality is working the same way / the contracts are preserved.

E2e tests will have to stay the same and will be the best indicator that no contracts were violated.

In case the module is more the size of the microservice, I would probably try to go with a
[Strangler Fig](https://martinfowler.com/bliki/StranglerFigApplication.html) approach/pattern.
Just create a new microservice and start rewriting the exposed endpoints one by one, while proxying the not yet
refactored code to the original module/service. This gives us more flexibility in terms of time we can spend on each
endpoint. Also, the refactored endpoints can be released one-by-one without having to wait for all the module to get refactored.

## Q3

There are two general approaches to coordinate multiple processes / services.
1. Messaging
2. Shared memory

I will have to say right away that I myself in the real projects used only the messaging mechanisms. I never had to 
run multiple go processes on the same OS, as the services are usually dockerized. However, if there ever was a need for real time
communication speed and/or strong consistency, the shared memory approach would be the way to go. So, data transfer speed is the first benefit of shared
memory.

Since I didn't have exposure to this approach on real projects, the other benefits or disadvantages I can think of -
are just my gut feeling, as I may not have some deep understanding of this. The downside of shared memory, as I see it,
could be the need to deploy both processes to the same VM/OS, which means it's not cloud, it's not managed by the cloud provider,
which can bring some challenges, like the need to have people/resources to maintain that VM. VMs/bare-metal also generally introduce
complications when undergoing certifications like PCI DSS.
The advantage of shared memory, especially in go, could be the possibility to use go channels to communicate. Channels are thread-safe
and do not require any additional synchronization other than that it itself provides.

Messaging, in its turn, introduces increase in infrastructure complexity, but also a nice decoupling of the services, they
can be deployed to different VMs, serverless instances, networks, datacenters or cloud providers. The messaging bus is
usually already available as a managed cloud service (SQS, SNS, Pub/Sub, etc) or as an on-premise service (Kafka, RabbitMQ, NATS, etc).
Messaging buses are equipped with retry mechanisms, message routing, audit log, granular access control, dead-lettering, etc.
Messages are eventually consistent, though. Will not be suitable for all the cases.

So to sum up, the approach would be dictated by the business needs. In my experience, messaging buses were good enough
in all the cases. I would love to have an opportunity to work with other mechanisms of IPC, too.
