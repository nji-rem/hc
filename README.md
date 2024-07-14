![img.png](assets/logo.png)

HC (short for Habbo Club) is a WIP Habbo emulator targeting the v9 client. HC is written in Go, an easy to learn 
language. 

### Architecture
HC is a modular monolith, meaning that each module is isolated. HC tries to be as fine-grained as possible in terms of 
module granularity. Components - or packages in Go - communicate through orchestration **and** choreography. Components 
that need orchestration, as you'll regularly see with microservices that use e.g. REST or gRPC, use dependency inversion 
(SOLI**D**); the latter uses an event-driven approach.

#### Project structure
**cmd/v9** - Presentation layer and entrypoint. Here lies all messaging (incoming and outgoing), route definitions, and dependency
injection spaghetti.

**pkg** - Shared code. This is mostly 'framework code', e.g. encoding or useful utilities.

**internal** - Contains modules (such as networking or packet resolvers). This code is **not** accessible by the 
presentation layer (`cmd` directory), because the presentation layer relies on events or interfaces.

**api** - Contains all interfaces, structures, and other things. This is accessible by every root directory, including 
`pkg`. 