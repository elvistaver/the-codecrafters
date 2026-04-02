# Strings Transformation CLI Tool
The strings Transformation Tool has about 6 functions which is been build individually using helper functions and called in the main function; the upper, lower, cap, title, snake, reverse

The string Transformer recieve's user input with a prompt for the user to enter an input choosing from the menu list then following with the text the user want to transform and it prints out to the console the output and requset user input again till the user types quit before the program ends.

## In writing the functions i used packages such as;

"bufio"- for allowing the user iput/ reading the user input , 

"fmt"- to print to the console the output,

"os"- this work inline with the bufio package allowing it to interact with the operating system to enable the user through standard input method in the os package to write to the terminal its input

"slices"- this package helped me to reverse strings with its built in method slice.Reverse

"strings" - this package also played a great role in my function for most of the major transformation such as turning strings to upper, title, cap, and also helping in spliting strings  for me to be able to do other manipulation like reverse, also for triming 

After building the functions i then now call them in my main function allocating them to specific command to responds to the user input in the menu list. 