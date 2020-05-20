## Ignite
Ignite is a tool to help start/stop/suspend/reset vmware fusion's virtual machine from a config file.

#### Install
```bash
% go get github.com/alimy/ignite@latest
```
#### Ignitefile
Ignitefile is a [HCL](https://github.com/hasicorp/hcl) format file.A sample file is [here](assets/Ignitefile).

#### Usage
```bash
% cat Ignitefile
version = "1"
...
...

# start workspace that named mysql-cluster
% ignite start mysql-cluster
INFO[0000] start workspace: mysql-cluster               
INFO[0000] start tier: mysql master.                    
2020-05-20T21:37:45.040| ServiceImpl_Opener: PID 8351
INFO[0001] start tier: mysql node 2.                    
INFO[0001] start tier: mysql node 1.                    
2020-05-20T21:37:46.540| ServiceImpl_Opener: PID 8451
2020-05-20T21:37:46.618| ServiceImpl_Opener: PID 8500
INFO[0003] start tier: mysql router.                    
2020-05-20T21:37:48.387| ServiceImpl_Opener: PID 8504

# suspend workspace that named mysql-cluster
% ignite suspend -f Ignitefile mysql-cluster
INFO[0000] suspend workspace: mysql-cluster             
INFO[0000] suspend tier: mysql router.                  
INFO[0000] suspend tier: mysql node 2.                  
INFO[0000] suspend tier: mysql node 1.                  
INFO[0002] suspend tier: mysql master.  

```
