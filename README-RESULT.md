# Result

## Commercial Products

|                   | File Size (bytes) | Obtained By                                    |
|-------------------|------------------:|------------------------------------------------|
| **Original File** |       `3,065,633` | (N/A)                                          |
| **Tar + GZIP**    |         `496,430` | `tar zcvf feed_file.txt.tar.gz feed_file.txt`  |
| **Tar + BZ2**     |         `368,845` | `tar jcvf feed_file.txt.tar.bz2 feed_file.txt` |
| **GZIP**          |         `486,865` | `gzip -9 -k feed_file.txt`                     |
| **BZ2**           |         `369,678` | `bzip2 -9 -k feed_file.txt`                    |
| **XZ**            |         `302,640` | `xz -9 -k feed_file.txt`                       |
| **XZ (Extreme)**  |         `304,080` | `xz -9 -e -k feed_file.txt`                    |

## My Trials

|               | File Size (bytes) | Obtained by                                          |
|---------------|------------------:|------------------------------------------------------|
| **1st Trial** |         `564,152` | `./ffcompress c 0 < feed_file.txt > compressed0.bin` |
| **2nd Trial** |         `670,582` | `./ffcompress c 1 < feed_file.txt > compressed1.bin` |
| **3rd Trial** |         `531,191` | `./ffcompress c 2 < feed_file.txt > compressed2.bin` |
| **4th Trial** |         `418,454` | `./ffcompress c 3 < feed_file.txt > compressed3.bin` |
| **5th Trial** |         `352,598` | `./ffcompress c 4 < feed_file.txt > compressed4.bin` |

### First Trial

By using the fixed-length scheme to store the time, each record could be stored in:

    17 + 10*3 + 4 = 51 bits

The resultant is encoded into an unsigned 64-bit integer with "gaps". The gaps
are suppose to ease human to read the code and speed up bit arithmetics (alignment):

    20 + 12*3 + 4 = 60 bits

### Second Trial

Each field is encoded into a variable-length integer with the built-in
`"encoding/binary"` package in Golang.

The result is in fact larger than the first trial.

### Thrid Trial

Like the first trial, each record is encoded into an unsigned 64-bit integer.
Unlike the first trial, there is no gaps this time.
The integer is then encoded using the same method as in the second trial
(i.e. using `"encoding/binary"`).

### Fourth Trial

The fourth trial is like the third trial, but instead of storing the absolute
timestamp, the number of seconds elapsed counting from previous line (time diff)
is stored.

Therefore, for each record, it could be stored in:

    6 + 10*3 + 4 = 40 bits

The integer is then encoded using the same method as in the second trial
(i.e. using `"encoding/binary"`).

### Fifth Trial

The fifth trial is like the fourth trial, but instead of using `"encoding/binary"`,
the fields are stored in a continuous stream of bits, using the BitWriter type
(see `bitwriter.go`) I created for this challenge.

## Summary

### Imperfectness

The above compression formats do not restore the original file losslessly.
For fields 2/3/4, the values `"123.400"`, `"123.40"`, and `"123.4"` are treated
as if they are the same. After decompression, the trailing zeros are removed.
I asked the team at the beginning of the challenge whether I could make such an
assumption, and the feedback was positive (as far as I could understand).

### Further Compression

#### Frequency Distribution Analysis

The file could be further compressed by analyzing the frequency distribution of
the values. This is the same principle as Shannon Entropy, and applying Huffman
Code. For instance, if there are a lot of `1`s and just a few `39`s in field 0
(time diff), instead of using a fixed number of bits to store them (like
`000001` and `100111`), a variable-length prefix code could be used so that
shorter codes are used for more frequent values.

Some code has already been written in `analysis.go` and `analyze.go`. However,
time is not enough to integrate these into the algorithm.

#### Second Pass of Compression

The file could be further compressed by applying a second pass of compression
using commercial / generic compression methods. Here are the results of
compressing `compressed4.bin`:

|                   | File Size (bytes) | Obtained By                                          |
|-------------------|------------------:|------------------------------------------------------|
| **Original File** |         `352,598` | `./ffcompress c 4 < feed_file.txt > compressed4.bin` |
| **Tar + GZIP**    |         `204,979` | `tar zcvf compressed4.bin.tar.gz compressed4.bin`    |
| **Tar + BZ2**     |         `155,336` | `tar jcvf compressed4.bin.tar.bz2 compressed4.bin`   |
| **GZIP**          |         `204,816` | `gzip -9 -k compressed4.bin`                         |
| **BZ2**           |         `155,164` | `bzip2 -9 -k compressed4.bin`                        |
| **XZ**            |         `161,728` | `xz -9 -k compressed4.bin`                           |
| **XZ (Extreme)**  |         `162,124` | `xz -9 -e -k compressed4.bin`                        |
