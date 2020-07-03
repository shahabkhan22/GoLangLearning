CONCURRENCY 

To understand concurrency, lets first understand how Scheduler works

1. OS Scheduler

OS scheduler take into account the layout and setup of the hardware they run on, including existence of multiple processors and cores.

OS creates a thread for each program. Its the job of the thread to account for sequential execution of the program.

Every program creates a Process and each PRocess is given an initial thread. Threads have the ability to create more threads.
All threads run independantly of each other and scheduling decisions are made at the thread level and not at the Process level.

Threads can run concurrently (each taking a turn on an individual core), or in parallel (each running at the same time on different cores). Threads also maintain their own state to allow for the safe, local, and independent execution of their instructions.

The OS scheduler is responsible for making sure cores are not idle if there are Threads that can be executing. It must also create the illusion that all the Threads that can execute are executing at the same time. In the process of creating this illusion, the scheduler needs to run Threads with a higher priority over lower priority Threads. However, Threads with a lower priority can’t be starved of execution time. The scheduler also needs to minimize scheduling latencies as much as possible by making quick and smart decisions.


Now RUN the OS_Test.go in this directory and check the output
Also run the following command from the terminal - 
go tool objdump -S -s "main.example" ./OS_Test


Read a little general idea about= 

Thread States
    1. Waiting  
    2. Runnable
    3. Executing

Type of WOrk- 
    1. CPU Bound
    2. I/O Bound


Context Switching- act of switching the threads on a core

When a scheduler pulls an Executing thread out off a core and replaces it with a Runnable Thread.
CSing is expensive as it is time consuming process.
Considering that hardware is able to execute 10 instructions per nanosecond, and context switching takes between ~1000 and ~1500 nanoseconds,
a CS can cost ~10k to ~15k instructions of latency.

I/O Bound program are advantageous to CS as other programs can use the executing time while program is in Waiting Queue.

CPU Bound program cause performance drops.


Scheduling Decision Scenario - 

1. Context-switch the main Thread off of core 1? Doing this could help performance, as the chances that this new Thread needs the same data that is already cached is pretty good. But the main Thread does not get its full time slice.

2. Have the Thread wait for core 1 to become available pending the completion of the main Thread’s time slice? The Thread is not running but latency on fetching data will be eliminated once it starts.

3. Have the Thread wait for the next available core? This would mean cache lines for the selected core would be flushed, retrieved, and duplicated, causing latency. However, the Thread would start more quickly and the main Thread could finish its time slice.

----------------------------------------------------------------------------------

2. Go Scheduler

