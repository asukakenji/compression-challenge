# Compression Challenge

## Compiling the source code

```
cd /path/to/the/source
go build -o ffcompress .
```

## Compressing a feed file using format 4

```
./ffcompress c 4 < feed_file.txt > compressed4.bin
```

## Decompressing an format 4 archive

```
./ffcompress d 4 < compressed4.bin > decompressed4.txt
```

## Analyzing a feed file

```
./ffcompress a < feed_file.txt
```
