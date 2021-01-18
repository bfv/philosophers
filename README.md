# Philosophers dilemma

5 Philosophers go out to eat for sushi. They sit down at a round table. They have 5 plates and between the plates is a chopstick, so 5 plates, 5 chopsticks. To eat a philosopher has to have 2 chopstick, one of the left of the plate, one on the right.

## assignment
Each philosopher has its own number, 1 to 5. 
Create a threaded solution where each philosopher has it's own thread. The eating happens in its own thread. The thread has to make sure that it is allowed to eat (i.e. that both the left and right chopstick are available). In order to eat the philosopher (in the thread) has to ask the waiter first. The waiter runs in its own thread. The waiter should let 2 philosophers eat concurrently. Each philosopher eats 3 times. When a philosopher starts to eat, print the number of the philosopher and which meal it is. When the philosopher is done eating a meal, print a line stating the philosopher (number) has finished his nth meal. Eating takes 1 second (philosophers are fast eaters, they got more important things to do)

What's important:
- locking of the chopsticks
- preventing deadlock
- concurrent eating (so threads or whatever your technology calls it)
- each of the 5 philosophers eats 3 meals

Bon appetite!
