# Christmas Coding Challenge 2017

*Deadline for submissions - 15/12/17*

# The Challenge

Leeds Dock has had a super productive year with many big wins, to reward the staff Renee has comissioned Santa's Elves to deliver gifts by hand to each desk.

For the sake of this exercise assume there are an infinite number of employees at Leeds Dock who sit at desks numbered sequentially: 1, 2, 3, 4, 5, and so on.

Each Elf is assigned a number, too, and delivers presents to desks based on that number:

- The first Elf (number 1) delivers presents to every desk: 1, 2, 3, 4, 5, ....
- The second Elf (number 2) delivers presents to every second desk: 2, 4, 6, 8, 10, ....
- Elf number 3 delivers presents to every third desk: 3, 6, 9, 12, 15, ....

There are infinitely many Elves, numbered starting with 1. Each Elf delivers presents equal to ten times his or her number at the desks they make deliveries to.

So, the first nine desks end up like this:

    Desk 1 got 10 presents.
    Desk 2 got 30 presents.
    Desk 3 got 40 presents.
    Desk 4 got 70 presents.
    Desk 5 got 60 presents.
    Desk 6 got 120 presents.
    Desk 7 got 80 presents.
    Desk 8 got 150 presents.
    Desk 9 got 130 presents.

The first desk gets 10 presents: it is visited only by Elf 1, which delivers 1 * 10 = 10 presents. The fourth desk gets 70 presents, because it is visited by Elves 1, 2, and 4, for a total of 10 + 20 + 40 = 70 presents.

What is the lowest desk number of the desk to get at least 50000000 presents?

# How to win

Fork this repo into your own namespace, *make it private*, and grant me access (_do not_ create a merge request, as previously requested, the diff will be publicly visible even if the repo isn't). I will verify the solution as it appears on the `master` branch.

Your fork should contain a simple Dockerfile containing all the required build steps and an appropriate CMD instruction (see the example [Dockerfile](Dockerfile)).

It must be possible to be called like this (see the [run-script](run-script.sh)):

    docker build -t christmas_comp .
    docker run --memory=1G christmas_comp

(note the `1G` memory limit on the `run`)

The output should be written to `stdout` on 2 lines, the first line being the desk number. And the second being the time taken to work it out in microseconds:

    123456
    1022μs

To make it fair across languages with a slower startup time (JVM for example), you should measure the execution time within your code. For example in pseudocode:

    start_time = getTime()
    ...
    // compute result
    ...
    execution_time = getTime() - start_time

The winning entry will be the fastest solution to deliver the correct answer.

# The prize

💰 £50 Amazon voucher for the winning entry 💰
