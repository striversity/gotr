# Section 04 - Lab 01 : Word Count 1

## TODO 1 - Count the occurrence of each _word_ in a text file

### Requirements

1. The filename is passed as an argument to the program
2. Use *input.NewFileReader()* to create an *input.FileReader* object.
    - The *input.FileReader* object has the method *ReadLine()*.
3. You will also need *strings.Split()* to break lines into words.
    - TIP: I also used *strings.TrimSpace()* in my solution.