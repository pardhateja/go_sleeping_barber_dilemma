## This is an Git hub repository which helps in solving the problem of sleeper barber dilemma

# Question description:

Problem description: This is a simple demonstration of how to solve the Sleeping
Barber dilemma, a classic computer science problem which illustrates the
complexities that arise when there are multiple operating system processes. Here, we
have a finite number of barbers, a finite number of seats in a waiting room, a fixed
length of time the barbershop is open, and clients arriving at (roughly) regular
intervals. When a barber has nothing to do, he or she checks the waiting room for
new clients, and if one or more is there, a haircut takes place. Otherwise, the barber
goes to sleep until a new client arrives. So the rules are as follows:
• if there are no customers, the barber falls asleep in the chair
• a customer must wake the barber if he is asleep
• if a customer arrives while the barber is working, the customer leaves if all chairs
are occupied and sits in an empty chair if it's available
• when the barber finishes a haircut, he inspects the waiting room to see if there are
any waiting customers and falls asleep if there are none
• shop can stop accepting new clients at closing time, but the barbers cannot leave
until the waiting room is empty
• after the shop is closed and there are no clients left in the waiting area, the barber
goes home


## Run the file main.go and give the respective values to make the code run.

## Sample input and the output for the same:

    $ go run main.go
    enter the total available seats: 7 
    enter the avg time taken for each haircut: 200
    enter the incoming rate of the customers to the barber shop: 300
    enter the total time the shop will be open: 2
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Adding Client 1
    Barber is cutting client 1's hair
    Barber is finished cutting client 1's hair
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Adding Client 2
    Barber is cutting client 2's hair
    Barber is finished cutting client 2's hair
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Adding Client 3
    Barber is cutting client 3's hair
    Adding Client 4
    Barber is finished cutting client 3's hair
    Barber is cutting client 4's hair
    Adding Client 5
    Barber is finished cutting client 4's hair
    Barber is cutting client 5's hair
    Barber is finished cutting client 5's hair
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Adding Client 6
    Barber is cutting client 6's hair
    Barber is finished cutting client 6's hair
    Barber is sleeping: Zzzzzzzzzzzzz ...
    Closing shop
    Barber is going home
    Shop closed