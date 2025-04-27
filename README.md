# Intro

This repo contains my implementation of [Huffman encoding and decoding](https://en.wikipedia.org/wiki/Huffman_coding).
I recently encountered this problem and thougth it would be a good way to practice bit manipulation
in go.

I initially just implemented string encoding/decoding to make verifying easier, after which I
converted to just encoding/decoding byte strings to support images, video or just any general blob
of bytes. I ran into the problem where decoding bits back to text would sometimes produce extra
characters at the end of my input. I had not implemented a way for decoding implementation to know
where to stop decoding. Since the smallest single chunk of memory is a byte which might only contain
1 significant bit, there was a possible 7 bits of extra information in the last byte.

To solve this problem I initially added a header to my encoded byte string. This header was a 64 bit
int that contained the length of the entire encoded message. After talking to [Devan Benz](https://github.com/devanbenz) about my implementation,
he suggested just passing along the amount of padding bits, a maximum of 7, instead. This solved my
initial constraint where an encoded message could only be 2^64 bytes long, and it makes the encode
bytestring 7 bytes shorter.

~~# Todo~~

~~What I still need to work on is a way to read and write the tree to/from disk. This means coming up
with a binary format for the tree. The way the tree is created now is just counting all occurrences
of a byte in a byte string and storing that in a hashmap as the weights. My initial implementation
is to just store the byte and it's weight as byte pairs and writing that to disk. If the weight of
any byte is more than the max an u8 int can store I will have to look at other solutions.~~

I ended up encoding the tree and an array of <byte><weight as u8> and just wrote that to a file.
This made it possible to store and load back in the tree when running the program separately meaning
I was also able to turn this into a cli quickly.
