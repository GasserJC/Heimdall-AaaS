# Heimdall Authentication as a Service Go API
Create a light weight and performant AaaS API.  
This DOES NOT save ANY raw user data, hashes provide a tool to authenticate users via a username and password without saving any raw or reversable data.  
Future functionality: Add User Removal 
  
Restrictions: persistant deployments.  
  
# Performance per Application on low Tier CPU  
User Count == 10,000  
Add User: ~840 Microseconds
Authenticate User: 99 Microseconds 
  
  
User Count == 100,000  
Add User: ~930 Microseconds
Authenticate User: ~110 Microseconds  

# Documentation

to be written...