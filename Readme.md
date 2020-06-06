## 简介
查看系统状态

## Test

```
MB1:mystat yezi$ make
go build -o Mystat main.go
```
```
MB1:mystat yezi$ ./Mystat -a
 9:58  up 21 days, 14:34, 4 users, load averages: 3.78 4.04 4.39
USER     TTY      FROM              LOGIN@  IDLE WHAT
yezi     console  -                15 520  21days -
yezi     s000     -                15 520  21days -bash
yezi     s001     -                15 520      - w
yezi     s002     -                15 520  20days -bash



Processes: 389 total, 2 running, 387 sleeping, 1662 threads 
2020/06/06 09:58:12
Load Avg: 3.78, 4.04, 4.39 
CPU usage: 10.92% user, 21.84% sys, 67.22% idle 
SharedLibs: 164M resident, 43M data, 26M linkedit.
```
