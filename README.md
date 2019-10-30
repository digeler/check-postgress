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
           root@psgress-test-74jgq:/var/logs/pslogs# tail -f  aks-agentpool-35064155-0.txt 
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:01:55.46939178 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:01:57.474532921 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:01:59.476610046 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:01.481702787 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:03.483729612 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:05.488764252 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:07.496066804 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:09.498097029 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:11.500541356 +0000 UTC
dial tcp: lookup dinor.test on 10.0.0.10:53: no such host
real error  dial tcp: lookup dinor.test on 10.0.0.10:53: no such host 2019-10-30 18:02:13.503282984 +0000 UTC
           

