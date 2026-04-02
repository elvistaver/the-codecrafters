# File-pipeline CLI Tool
This is a file system that takes user input, save it in the input.txt and transform it base on the command choosen and writes back the transformed text to the output.txt file

In this project there are five functions and the first is;
 The title- it turns the first character of each word to capital letter and the rest to small letters, here i used the for range loop to interate over each character and take the current index of each character of word in the string and change the first character to capital and the rest to lowercase

 The upper - it turns all strings to capital letters and i used the strings package for it also

 The TrimSpace- it removes the leading and trailing spaces in a string, thats the space before and after a string aslo used strings package for it

 The Reverse- it reverse the text from right to left, i used the slices package for it

 The Remove- it removes space and remove dashes in a string

After the helper functions, there is the main function which calls all the helper functions and also is in the main function that the file sysmtem api is built for reading and writing files using the os package