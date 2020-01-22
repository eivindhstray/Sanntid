# Choosing a language

In Exercise 3, 4 (Network exercises) and the project, you will be using a language of your own choice. You are of course free to change your mind at any time, but to help avoid this situation (and all its associated costs) it is worth doing some research already now.

You should start by looking at the [Programming Language part of the Project](https://github.com/TTK4145/Project/tree/master#programming-language). Send in your suggestions if you find more and/or better resources.

Here are a few things you should consider:
 - Think about how want to move data around (reading buttons, network, setting motor & lights, state machines, etc). Do you think in a shared-variable way or a message-passing way? Will you be using concurrency at all?
 - How will you split into modules? Functions, objects, threads? Think about what modules you need, and how they need to interact. This is an iterative design process that will take you many tries to get "right" (if such a thing even exists!).
 - The networking part is often difficult. Can you find anything useful in the standard libraries, or other libraries?
 - While working on new sections on the project you'll want to avoid introducing bugts to the parts that already work properly. Does the language have a framework for making and running tests, or can you create one? Testing multithreaded code is especially difficult.
 - Code analysis/debugging/IDE support?
