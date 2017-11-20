# Christmas Coding Challenge 2017

*Deadline for submissions - 15/12/17*

# The Challenge

By random draw you have been selected to be Santa's helper this year, your job is to deliver presents to 10 lucky Sky employees scattered over Leeds Dock.

Use your skills to find the most efficient way of delivering all the presents by making as few steps as possible.

You are given a seating layout in the form of an ascii file, it details the locations and distances between each employee. `x` is your current location, from where you start; the other numbers are (in no particular order) the deliveries you need to make. Walls are marked as #, and open passages are marked as `.`. Numbers behave like open passages.

For example, suppose you have a map like the following (simplified to 4 employees):

    ###########
    #x.0.....1#
    #.#######.#
    #3.......2#
    ###########

To make all the deliveries as efficiently as possible, you would have to take the following path:

- x to 3 (2 steps)
- 3 to 0 (4 steps; diagonal moves are not allowed)
- 0 to 1 (6 steps)
- 1 to 2 (2 steps)

Which gives a total of 14 steps

[Given the actual map](map.txt), and starting from location x, what is the minimum number of steps required to visit every number marked on the map (in any order) at least once?

# How to win

Fork this repo into your own namespace and grant me access (_do not_ create a merge request, as previously requested, the diff will be publicly visible even if the repo isn't). I will verify the solution as it appears on the `master` branch.

Your fork should contain a simple Dockerfile containing all the required build steps and an appropriate CMD instruction (see the example [Dockerfile](Dockerfile)).

It must be possible to be called like this (see the [run-script](run-script.sh)):

    docker build -t christmas_comp .
    docker run --memory=1G christmas_comp

(note the `1G` memory limit on the `conatiner run`)

The output should be written to `stdout` on 2 lines, the first line being a description of the shortest path taken in terms of "nsew" characters and the total number of steps. And the second being the execution time in milliseconds. For example (for the above sample map):

    eewwsseeeeeeeenn 16
    1022ms

The submission that computes the shortest valid path will be the winner.

If more than one solution agrees on the number of steps, the execution time will decide.

Submissions that take longer than a reasonable amount of time to execute will be excluded.

To make it fair accross languages with a slower startup time (JVM for example), you should measure the execution time within your code. For example in pseudocode:

    start_time = getTime()
    ...
    // compute result
    ...
    execution_time = getTime() - start_time

# The prize

💰 £50 Amazon voucher for the winning entry 💰
