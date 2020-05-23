## Ignite
Ignite is a tool to help start/stop/suspend/reset vmware fusion's virtual machine from a config file.

#### Install
```bash
% go get github.com/alimy/ignite@latest
```
#### Ignitefile
Ignitefile is a [HCL](https://github.com/hashicorp/hcl) format file.A sample file is [here](assets/Ignitefile).

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
INFO[0001] start tier: mysql node 2.                    
INFO[0001] start tier: mysql node 1.                    
INFO[0004] start tier: mysql router. 

# suspend workspace that named mysql-cluster
% ignite suspend -f Ignitefile mysql-cluster
INFO[0000] suspend workspace: mysql-cluster             
INFO[0000] suspend tier: mysql router.                  
INFO[0000] suspend tier: mysql node 2.                  
INFO[0000] suspend tier: mysql node 1.                  
INFO[0002] suspend tier: mysql master.

# ssh to tier
% ignite ssh -u alimy mysql-master
INFO[0000] try ssh to alimy@192.168.117.138 on port 22
Activate the web console with: systemctl enable --now cockpit.socket

Last login: Sat May 23 18:17:04 2020 from 192.168.117.1
```
