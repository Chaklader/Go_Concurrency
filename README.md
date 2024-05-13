
# GO CONCURRENCY

<br>

## Programs Created For Learning

<br>

- Simple Channel: The provided Go program creates a simple communication channel between the main goroutine and a worker goroutine. The 
worker goroutine converts the strings received from the main goroutine to uppercase and appends exclamation marks, sending 
the modified strings back to the main goroutine.

<br>

- The provided Go program simulates a pizzeria that takes orders and makes pizzas using goroutines and channels. The producer 
goroutine generates pizza orders, while the consumer goroutine processes the orders and reports the status of each order. The 
program aims to demonstrate the use of channels for communication between goroutines and the concurrent execution of tasks.


- This is the famous "Dining Philosophers" problem proposed by Edsger Dijkstra, a Dutch computer scientist, in 1965. It is 
a classic synchronization problem in computer science and is used to illustrate the challenges of avoiding deadlock and starvation 
in concurrent systems. The problem is stated as follows:

There are five philosophers sitting around a circular table, with one fork placed between each pair of adjacent philosophers. 
Each philosopher must alternate between thinking and eating. However, a philosopher can only eat when they have acquired both 
the fork to their left and the fork to their right. In other words, two adjacent philosophers cannot eat simultaneously since 
they share a fork between them.

The challenge is to design a solution that allows the philosophers to eat and think without violating the following constraints:

<br>

1. No two adjacent philosophers can eat simultaneously (mutual exclusion).
2. Philosophers should not starve (deadlock-free).
3. The solution should avoid unnecessary waiting (livelock-free).
4. The solution should be fair, allowing each philosopher to eat infinitely often.

<br>

The provided Go program implements Dijkstra's solution to the Dining Philosophers problem using Go's concurrency primitives, 
such as goroutines and mutexes. The program simulates the philosophers' behavior by creating a goroutine for each philosopher, 
where each goroutine alternates between thinking and eating. The forks are represented as mutexes, and a philosopher must 
acquire the mutexes for both their left and right forks before eating.

The program demonstrates how to use synchronization mechanisms to achieve mutual exclusion and avoid deadlocks, ensuring 
that the philosophers can eat and think without violating the problem's constraints.


- The provided Go program demonstrates a solution to the classic "Sleeping Barber" problem, which is a concurrency problem 
in computer science. The problem involves a barbershop with a waiting room, barbers, and customers arriving at random 
intervals.

The program simulates the following scenario:

1. The barbershop has a finite number of seats in the waiting room.
2. If there are no customers, the barber falls asleep in their chair.
3. A customer must wake the barber if they are asleep.
4. If a customer arrives while the barber is working, and all chairs are occupied, the customer leaves.
5. If a chair is available, the customer takes a seat in the waiting room.
6. When the barber finishes a haircut, they inspect the waiting room for waiting customers. If there are none, the barber falls asleep.
7. The shop stops accepting new customers at closing time, but the barbers cannot leave until the waiting room is empty.
8. After the shop is closed and there are no customers left in the waiting area, the barbers go home.

The program creates a barbershop struct with properties like shop capacity, haircut duration, number of barbers, and channels 
for communication. It adds barbers to the shop and starts a goroutine for each barber to handle customers. Customers arrive 
at random intervals, and if the waiting room is not full, they are added to the customers' channel. If the waiting room is 
full, the customer leaves.

When a barber is available, they take a customer from the channel, cut their hair, and then check for more customers. If 
no customers are waiting, the barber falls asleep until a new customer arrives and wakes them up. The program also simulates 
the closing of the shop after a specified time. When the shop closes, it stops accepting new customers, and the barbers go 
home once all the remaining customers have been served. The program uses Go's concurrency primitives, such as goroutines 
and channels, to coordinate the interaction between barbers and customers, ensuring that the problem's constraints are met.



GO CONCURRENCY TOPICS  
–––––––––––––––––––––

* Goroutines
* Channels
* Select statement
* Buffered channels
* Range and close
* WaitGroups
* Mutex
* RWMutex
* Atomicity
* Race conditions
* Deadlocks
* Resource starvation
* Channel synchronization
* Channel direction
* Context package
* Cancellation and timeouts
* Worker pools
* Pipelines
* Fan-out, fan-in patterns
* Generators
* Multiplexing
* Concurrency patterns (e.g., producer-consumer, publish-subscribe)
* Graceful shutdown
* Parallelism vs. concurrency
* Debugging concurrent programs
* Performance optimization in concurrent programs
* Concurrency-safe data structures
* Goroutine leaks
* Concurrency best practices
* Concurrency testing techniques
