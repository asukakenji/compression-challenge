#!/bin/bash

# Result: The integral part of the 3rd field only contains 108 or 109.
cut -f3 -d' ' feed_file.txt | cut -d'.' -f1 | sort | uniq

# Result: The integral part of the 4th field only contains 108 or 109.
cut -f4 -d' ' feed_file.txt | cut -d'.' -f1 | sort | uniq

# Result: The integral part of the 5th field only contains 108 or 109.
cut -f5 -d' ' feed_file.txt | cut -d'.' -f1 | sort | uniq

# Result: The 6th contains only these choices:
# DUBA
# FFS
# FXCM
# FXDC
# FXDD
# GAIN
# KZ
# PEP
# PFD
# SBD
# SEBS
# TDFX
cut -f6 -d' ' feed_file.txt | sort | uniq

# Download the feed file
# I tried curl. I have no idea why it didn't work!
wget 'https://www.dropbox.com/s/1mcc2demrs9nbc0/feed_file%20%285%29.txt?dl=0'
