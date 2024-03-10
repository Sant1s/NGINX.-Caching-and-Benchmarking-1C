# Замер производительности
## Производительность сервиса на Flask
### Тестирование в однопоточном режиме
```bash
ab -n 5000 http://localhost:5000/api/date
```
Получаем такой результат
```bash
Server Software:        Werkzeug/3.0.1
Server Hostname:        localhost
Server Port:            5000

Document Path:          /api/date
Document Length:        640003 bytes

Concurrency Level:      1
Time taken for tests:   161.560 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      3200865000 bytes
HTML transferred:       3200015000 bytes
Requests per second:    30.95 [#/sec] (mean)
Time per request:       32.312 [ms] (mean)
Time per request:       32.312 [ms] (mean, across all concurrent requests)
Transfer rate:          19347.86 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:    30   32   1.4     32      39
Waiting:       19   22   1.4     21      28
Total:         30   32   1.4     32      39

Percentage of the requests served within a certain time (ms)
  50%     32
  66%     33
  75%     33
  80%     34
  90%     34
  95%     35
  98%     35
  99%     35
 100%     39 (longest request)
```
### Тестирование в многопоточном режиме
```bash
ab -n 5000 -c 16 http://localhost:5000/api/date
```
Получаем такой результат
```bash
Server Software:        Werkzeug/3.0.1
Server Hostname:        localhost
Server Port:            5000

Document Path:          /api/date
Document Length:        640003 bytes

Concurrency Level:      16
Time taken for tests:   123.374 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      3200865000 bytes
HTML transferred:       3200015000 bytes
Requests per second:    40.53 [#/sec] (mean)
Time per request:       394.797 [ms] (mean)
Time per request:       24.675 [ms] (mean, across all concurrent requests)
Transfer rate:          25336.34 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:    66  394  68.7    390     746
Waiting:       24  330  55.4    328     642
Total:         67  394  68.6    390     746

Percentage of the requests served within a certain time (ms)
  50%    390
  66%    418
  75%    436
  80%    449
  90%    482
  95%    513
  98%    552
  99%    576
 100%    746 (longest request)
 ```

## Производительность сервиса на Golang
### Тестирование в однопоточном режиме
```bash
ab -n 5000 http://localhost:8080/api/date
```
Получаем такие результаты
```bash
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /api/date
Document Length:        330011 bytes

Concurrency Level:      1
Time taken for tests:   7.712 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      1650495000 bytes
HTML transferred:       1650055000 bytes
Requests per second:    648.37 [#/sec] (mean)
Time per request:       1.542 [ms] (mean)
Time per request:       1.542 [ms] (mean, across all concurrent requests)
Transfer rate:          209010.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    1   0.3      1       3
Waiting:        1    1   0.3      1       3
Total:          1    2   0.3      1       3
ERROR: The median and mean for the total time are more than twice the standard
       deviation apart. These results are NOT reliable.

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      2
  75%      2
  80%      2
  90%      2
  95%      2
  98%      2
  99%      2
 100%      3 (longest request)
 ```

### Тестирование в многопоточном режиме
```bash
ab -n 5000 -c 16 http://localhost:8080/api/date
```
Получаем вот такие результаты
```bash
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /api/date
Document Length:        330011 bytes

Concurrency Level:      16
Time taken for tests:   1.037 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      1650495000 bytes
HTML transferred:       1650055000 bytes
Requests per second:    4821.34 [#/sec] (mean)
Time per request:       3.319 [ms] (mean)
Time per request:       0.207 [ms] (mean, across all concurrent requests)
Transfer rate:          1554218.41 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     1    3   0.7      3       7
Waiting:        1    2   0.6      2       5
Total:          1    3   0.7      3       7

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      4
  80%      4
  90%      4
  95%      4
  98%      5
  99%      5
 100%      7 (longest request)
 ```

 Из чего можно сделать вывод о перимуществе конкурентного Golang

## Производительность с проксированием
### Только сервис на Flask
```bash
ab -n 5000 -c 16 http://localhost/api/date
```
Получаем вот такой результат
```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/date
Document Length:        640003 bytes

Concurrency Level:      16
Time taken for tests:   1.191 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      3200825000 bytes
HTML transferred:       3200015000 bytes
Requests per second:    4199.36 [#/sec] (mean)
Time per request:       3.810 [ms] (mean)
Time per request:       0.238 [ms] (mean, across all concurrent requests)
Transfer rate:          2625273.62 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    4   0.4      4      29
Waiting:        0    0   0.4      0      27
Total:          1    4   0.4      4      29

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      4
  80%      4
  90%      4
  95%      4
  98%      4
  99%      4
 100%     29 (longest request)
 ```

 ### Только сервис на Golang
 ```bash
 ab -n 5000 -c 16 http://localhost/api/date
 ```
 Получаем вот такие результаты
 ```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/date
Document Length:        330011 bytes

Concurrency Level:      16
Time taken for tests:   0.748 seconds
Complete requests:      5000
Failed requests:        5
   (Connect: 0, Receive: 0, Length: 5, Exceptions: 0)
Total transferred:      1652295080 bytes
HTML transferred:       1651604960 bytes
Requests per second:    6681.23 [#/sec] (mean)
Time per request:       2.395 [ms] (mean)
Time per request:       0.150 [ms] (mean, across all concurrent requests)
Transfer rate:          2156126.24 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     1    2   0.1      2       4
Waiting:        0    0   0.1      0       2
Total:          1    2   0.1      2       4

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      2
  80%      2
  90%      2
  95%      2
  98%      3
  99%      3
 100%      4 (longest request)
 ```
 ### Балансировка на оба сервиса
```bash
ab -n 5000 -c 16 http://localhost/api/date
```
Получаем вот такие результаты
```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/date
Document Length:        330011 bytes

Concurrency Level:      16
Time taken for tests:   1.189 seconds
Complete requests:      5000
Failed requests:        4999
   (Connect: 0, Receive: 0, Length: 4999, Exceptions: 0)
Total transferred:      3200514984 bytes
HTML transferred:       3199705008 bytes
Requests per second:    4203.91 [#/sec] (mean)
Time per request:       3.806 [ms] (mean)
Time per request:       0.238 [ms] (mean, across all concurrent requests)
Transfer rate:          2627864.26 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    4   0.4      4      30
Waiting:        0    0   0.4      0      26
Total:          2    4   0.4      4      30

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      4
  80%      4
  90%      4
  95%      4
  98%      4
  99%      4
 100%     30 (longest request)
 ```

 ```bash
 ab -n 5000 -c 10 -T 'application/x-www-form-urlencoded' -p post_data.txt http://localhost/api/name
 ```
 Получаем вот такой результат
 ```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/name
Document Length:        290003 bytes

Concurrency Level:      10
Time taken for tests:   0.674 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      1450825000 bytes
Total body sent:        815000
HTML transferred:       1450015000 bytes
Requests per second:    7423.40 [#/sec] (mean)
Time per request:       1.347 [ms] (mean)
Time per request:       0.135 [ms] (mean, across all concurrent requests)
Transfer rate:          2103525.65 [Kbytes/sec] received
                        1181.65 kb/s sent
                        2104707.31 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    1   0.2      1      18
Waiting:        0    0   0.2      0      16
Total:          1    1   0.2      1      18

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      1
 100%     18 (longest request)
```
## Производительность с разным колличеством одновременных подключений
```bash
ab -n 5000 -T 'application/x-www-form-urlencoded' -p post_data.txt http://localhost/api/name
```
Получаем вот такой результат
```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/name
Document Length:        290003 bytes

Concurrency Level:      1
Time taken for tests:   0.895 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      1450825000 bytes
Total body sent:        815000
HTML transferred:       1450015000 bytes
Requests per second:    5585.74 [#/sec] (mean)
Time per request:       0.179 [ms] (mean)
Time per request:       0.179 [ms] (mean, across all concurrent requests)
Transfer rate:          1582800.03 [Kbytes/sec] received
                        889.14 kb/s sent
                        1583689.17 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    0   0.0      0       0
Waiting:        0    0   0.0      0       0
Total:          0    0   0.0      0       0

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      0
  75%      0
  80%      0
  90%      0
  95%      0
  98%      0
  99%      0
 100%      0 (longest request)
 ```

```bash
ab -n 5000 -c 100 -T 'application/x-www-form-urlencoded' -p post_data.txt http://localhost/api/name
```
Получаем вот такой результат
```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/name
Document Length:        290003 bytes

Concurrency Level:      100
Time taken for tests:   0.661 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      1450825000 bytes
Total body sent:        815000
HTML transferred:       1450015000 bytes
Requests per second:    7559.29 [#/sec] (mean)
Time per request:       13.229 [ms] (mean)
Time per request:       0.132 [ms] (mean, across all concurrent requests)
Transfer rate:          2142031.89 [Kbytes/sec] received
                        1203.29 kb/s sent
                        2143235.17 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.2      1       3
Processing:     5   12   0.7     12      27
Waiting:        0    1   0.3      1      16
Total:          6   13   0.6     13      27

Percentage of the requests served within a certain time (ms)
  50%     13
  66%     13
  75%     13
  80%     13
  90%     13
  95%     13
  98%     14
  99%     14
 100%     27 (longest request)
 ```
```bash
ab -n 5000 -c 8 -T 'application/x-www-form-urlencoded' -p post_data.txt http://localhost/api/name
```
Получаем вот такой результат
```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/name
Document Length:        290003 bytes

Concurrency Level:      8
Time taken for tests:   0.677 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      1450825000 bytes
Total body sent:        815000
HTML transferred:       1450015000 bytes
Requests per second:    7382.93 [#/sec] (mean)
Time per request:       1.084 [ms] (mean)
Time per request:       0.135 [ms] (mean, across all concurrent requests)
Transfer rate:          2092058.17 [Kbytes/sec] received
                        1175.21 kb/s sent
                        2093233.38 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:     0    1   0.0      1       2
Waiting:        0    0   0.0      0       1
Total:          1    1   0.0      1       2

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      1
 100%      2 (longest request)
 ```
```bash
ab -n 5000 -c 16 http://localhost/api/date
```
Получаем вот такой результат
```bash
Server Software:        nginx/1.18.0
Server Hostname:        localhost
Server Port:            80

Document Path:          /api/date
Document Length:        640003 bytes

Concurrency Level:      16
Time taken for tests:   1.177 seconds
Complete requests:      5000
Failed requests:        0
Total transferred:      3200825000 bytes
HTML transferred:       3200015000 bytes
Requests per second:    4247.64 [#/sec] (mean)
Time per request:       3.767 [ms] (mean)
Time per request:       0.235 [ms] (mean, across all concurrent requests)
Transfer rate:          2655457.72 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       1
Processing:     1    4   1.8      4      76
Waiting:        0    0   1.6      0      72
Total:          1    4   1.8      4      77

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      4
  80%      4
  90%      4
  95%      4
  98%      4
  99%      4
 100%     77 (longest request)
 ```
 