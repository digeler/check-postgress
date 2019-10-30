# check-postgress
check postgress connection

run the file after you edit the settings pstestdep.yaml

 env:
         - name: host
           value: ""  # yourdbnam
         - name: user
           value: "" # your username
         - name: password
           value: "" #password 
         - name: dbname
           value: "" #dbname  
           
           once done it will run connectivity tests for every 2 sec 
           
           the output will be in the pv file :
           
           root@psgress-test-c56b94c89-7lxgg:/var/logs/pslogs# cat ps.log 
Successfully connected!   2019-10-30 16:52:42.540544271 +0000 UTC
Successfully connected!   2019-10-30 16:52:45.343449938 +0000 UTC
Successfully connected!   2019-10-30 16:52:48.101259126 +0000 UTC
Successfully connected!   2019-10-30 16:52:50.941542407 +0000 UTC
Successfully connected!   2019-10-30 16:52:53.736711009 +0000 UTC
Successfully connected!   2019-10-30 16:52:56.474371327 +0000 UTC
Successfully connected!   2019-10-30 16:52:59.214423666 +0000 UTC
Successfully connected!   2019-10-30 16:53:01.997851969 +0000 UTC
Successfully connected!   2019-10-30 16:53:04.798042013 +0000 UTC
Successfully connected!   2019-10-30 16:53:07.553737583 +0000 UTC
Successfully connected!   2019-10-30 16:53:10.321969859 +0000 UTC
Successfully connected!   2019-10-30 16:53:13.051206306 +0000 UTC
Successfully connected!   2019-10-30 16:53:15.898775549 +0000 UTC
Successfully connected!   2019-10-30 16:53:18.65694974 +0000 UTC
Successfully connected!   2019-10-30 16:53:21.386101086 +0000 UTC
Successfully connected!   2019-10-30 16:53:24.13984064 +0000 UTC
Successfully connected!   2019-10-30 16:53:26.909439227 +0000 UTC
Successfully connected!   2019-10-30 16:53:29.664608793 +0000 UTC
Successfully connected!   2019-10-30 16:53:32.410120377 +0000 UTC


