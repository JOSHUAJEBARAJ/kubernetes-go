## Custom controller

The usecase of this controller is whenever we can deployment the svc and ingress is deployed automatically.

### Building the controller

1. Informer 
2. Register function - when certain event happens, what function should be called
3. Queue - to maintain the group of objects
4. Routine - to process the queue or business logic