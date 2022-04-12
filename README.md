# Heimdall Authentication as a Service Go API
Current Goal: Create an easy and simple to implement authentication for a single application with the standard library.  
Next Goal: Improve Performance via concurrent infrastructure.  
  
Restrictions: persistant deployments.

# Performance per Application on low Tier CPU
User Count == 10,000  
Add User: ~130 Microseconds  
Authenticate User: 1 Microsecond
Database size: 328 KB, (32.8 Bytes per User of Disk Space)  
  
  
User Count == 100,000  
Add User: ~140 Microseconds
Authenticate User: 1 Microseconds
Database size: 4,648 KB, (46.48 Bytes per User of Disk Space)
  

User Count == 1,000,000
Add User: ~7000 Microseconds 
Authenticate User: ~ 7 Microseconds
Database size at 100,000 user count: 22,900 KB (24 Bytes per User of Disk Space)