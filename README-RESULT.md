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

|               | File Size (bytes) | Obtained by                                         |
|---------------|------------------:|-----------------------------------------------------|
| **1st Trial** |         `564,152` | `./ffcompress c 0 < feed_file.txt > compressed0.bin` |
| **2nd Trial** |         `670,582` | `./ffcompress c 1 < feed_file.txt > compressed1.bin` |
| **3rd Trial** |         `561,906` | `./ffcompress c 2 < feed_file.txt > compressed2.bin` |
| **4th Trial** |         `439,681` | `./ffcompress c 3 < feed_file.txt > compressed3.bin` |

### First Trial

By using the fixed-length scheme to store the time, each record could be stored in:

    17 + 10*3 + 4 = 51 bits

To ease implementation, a 64-bit integer is used.

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

    6 + 11*3 + 4 = 43 bits

To ease implementation, a 64-bit integer is used.
