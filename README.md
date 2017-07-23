# Compression Challenge

## Compiling the source code

```
cd /path/to/the/source
go build -o ffcompress .
```

## Compressing a feed file using format 3

```
./ffcompress c 3 < feed_file.txt > compressed3.bin
```

## Decompressing a feed file using format 3

```
./ffcompress d 3 < compressed3.bin > decompressed3.txt
```

## Analyzing a feed file

```
./ffcompress a < feed_file.txt
```
