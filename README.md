# Heimdall Authentication as a Service Go API
Create a light weight and performant AaaS API.  
This DOES NOT save ANY raw user data, hashes provide a tool to authenticate users via a username and password without saving any raw or reversable data.  
Future functionality: Add User Removal 
  
Restrictions: persistant deployments.  
  
# Performance per Application on low Tier CPU Ubuntu 20.04 Container  
User Count == 10,000  
Add User: ~75 Microseconds
Authenticate User: 43 Microseconds 
  
  
User Count == 100,000  
Add User: ~95 Microseconds
Authenticate User: ~85 Microseconds  

# Documentation

to be written...