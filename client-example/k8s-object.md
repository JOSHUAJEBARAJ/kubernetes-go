## 

If any struct implments the runtime.Object interface, it can be used as a Kubernetes object. The runtime.Object interface is defined as follows:


Pod implements the gettypekind and other method using the typemeta interface and its implement the deep copy directly 


Each kubernetes object has the following things 
- TypeMeta
    - Kind 
    - apiversion 
- ObjectMeta - name , namespace , labels , annotations
 spec - for example how many replicas
 status - done 
