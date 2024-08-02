The project has next components' architecture:
```
          Infrastructure layer                              Services layer                             Infrastructure layer              
+-------------------------------------------+  +-------------------------------------+  +-----------------------------------------------+
|                                           |  |                                     |  |                                               |
| +----------------+                        |  |          +----------------+         |  |                                               |
| |       DB       <---+                    |  |          |    Services    |         |  |                                               |
| +----------------+   |                    |  |          +-+------------+-+         |  |                                               |
|                      |                    |  |            |            |           |  |                                               |
| +----------------+   |   +------------+   |  | +----------v-+       +--v---------+ |  |   +-------------+        +-----------------+  |
| |  Currency API  <---+---+  Adapters  +---+--+->  Gateways  |       |  Services  <-+--+---+ Controllers +-------->  WEB Framework  |  |
| +----------------+   |   +------------+   |  | | interfaces |       | interfaces | |  |   +-------------+        +-----------------+  |
|                      |                    |  | +------------+       +------------+ |  |                                               |
| +----------------+   |                    |  |                                     |  |                                               |
| |     Emails     <---+                    |  +-------------------------------------+  |                                               |
| +----------------+                        |                                           |                                               |
|                                           +-------------------------------------------+                                               |
|                                                                                                                                       |
+---------------------------------------------------------------------------------------------------------------------------------------+
```


The service layer consists of processes in accordance with business logic. The infrastructure layer includes external tools such as
database, email functionality, APIs, and others. Business logic should not depend on infrastructure details, so
there are interfaces that are needed to apply the Dependency Inversion Principle. 
