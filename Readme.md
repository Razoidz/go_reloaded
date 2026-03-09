pseudocode
create 2 files(sample and result)txt as argument, write on the file.

functions to replace
(hex) to hexadecimal

(Bin) to decimal
(up) to uppercase
(low) to lowercase
(cap) to capitalized
if number appears beside (low), (up) and (cap), it turns the corresponding function based on the number :Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING"

punctuations should be close to the previous word and with space apart from the next one.

group punctuations should be close to pevious word 

punctuation mark ' should be place left and right of the word in the middle of them without space

Every instance of (a) should be turned into (an) if the next word begins with a vowel (a, e, i, o, u) or a h 


(hex) to hexadecimal:
convert the string to slice 
loop through slice to get index to replace[i-1]
convert index[i-1] from hex to to int
convert the int back to string due to it being a slice
remove (hex) from the slice via append function
i-- to adjust slicefter removing (hex)

loop
if worrd[index] == up && index > o{
    word[index-1] = string.ToLower(word[index-1])
    word = append(word[:i],word[i+1:]) 
}
