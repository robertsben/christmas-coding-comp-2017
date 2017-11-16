# Christmas Coding Challenge 2017

Deadline for submissions - 15/12/17

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

- x to 0 (2 steps)
- 0 to 3 (4 steps; diagonal moves are not allowed)
- 3 to 2 (8 steps)
- 2 to 3 (2 steps)

Which gives a total of 16 steps

[Given the actual map](map.txt), and starting from location x, what is the minimum number of steps required to visit every number marked on the map at least once?

# How to win

Submission should be in the form of a merge request into `master` of this repo from your own _private_ fork. With it being a private fork you'll need to explicitly grant me access.

Your fork should contain a simple Dockerfile containing all the required build steps and an appropriate CMD instruction (see the example [Dockerfile](Dockerfile)).

It must be possible to be called like this (see the `run-script` in `master`):

    docker build -t christmas_comp .
    docker run christmas_comp

The output should be written to `stdout` on 2 lines, the first line being the minimum number of steps that are needed, and the second being the execution time in microseconds. For example:

    14
    1022µs

The fastest time with the correct answer wins. In the event of a tie the cleanliness of the code will be the deciding factor.

Submissions can be made by linking to a git repo or a zip file and you should state an expected execution time.

To make it fair accross languages with a slower startup time (JVM for example), you should measure the execution time within your code. For example in pseudocode:

    start_time = getTime()
    ...
    // compute result
    ...
    execution_time = getTime() - start_time

# The prize

TBC
