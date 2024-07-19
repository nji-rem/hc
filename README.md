![img.png](assets/logo.png)

HC (short for Habbo Club) is a WIP Habbo emulator targeting the v9 client. HC is written in Go, an easy to learn 
language. 

### Architecture
HC is a modular monolith, meaning that each module is isolated. HC tries to be as fine-grained as possible in terms of 
module granularity. Components - or packages in Go - communicate through orchestration **and** choreography. Components 
that need orchestration, as you'll regularly see with microservices that use e.g. REST or gRPC, use dependency inversion 
(SOLI**D**); the latter uses an event-driven approach.

Architecture chapter is still WIP, because it's in an experimental phase.