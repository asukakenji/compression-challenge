# Compression Challenge

## Compiling the source code

```
cd /path/to/the/source
go build -o ffcompress .
```

## Compress a feed file using format 0

```
./ffcompress c 0 < feed_file.txt > compressed0.bin
```

## Decompress a feed file using format 0

```
./ffcompress d 0 < compressed0.bin > decompressed0.txt
```

## Analyzing a feed file

```
./ffcompress a < feed_file.txt
```
