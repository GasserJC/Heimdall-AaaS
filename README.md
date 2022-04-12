# Heimdall Authentication as a Service Go API
Create an light weight and performant AaaS API.  
This DOES NOT safe ANY user data, hashes provide a tool to authenticate users via a username and password without saving user data.  
Future functionality: Add User Removal  
  
Restrictions: persistant deployments.  
  
# Performance per Application on low Tier CPU  
User Count == 10,000  
Add User: ~130 Microseconds (.13 Milliseconds)  
Authenticate User: 1 Microsecond  
Database size: 328 KB, (32.8 Bytes per User of Disk Space)  
  
  
User Count == 100,000  
Add User: ~140 Microseconds (.14 Milliseconds)  
Authenticate User: 1 Microseconds (.001 Milliseconds)  
Database size: 4,648 KB, (46.48 Bytes per User of Disk Space)  
  
  
User Count == 1,000,000  
Add User: ~7000 Microseconds (7 Milliseconds)  
Authenticate User: ~ 7 Microseconds (.007 Milliseconds)  
Database size at 100,000 user count: 22,900 KB (24 Bytes per User of Disk Space)  

# Documentation

to be written...