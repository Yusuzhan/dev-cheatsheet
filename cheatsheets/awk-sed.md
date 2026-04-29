---
title: Awk / Sed
icon: fa-file-lines
primary: "#C74634"
lang: bash
---

## fa-pen sed Basic Substitution

```bash
sed 's/old/new/' file.txt              # replace first occurrence per line
sed 's/old/new/g' file.txt             # replace all occurrences
sed 's/old/new/2' file.txt             # replace 2nd occurrence per line
sed -i 's/old/new/g' file.txt          # in-place edit
echo "hello123" | sed 's/[0-9]//g'     # remove all digits
```

## fa-eraser sed Delete & Insert

```bash
sed '/pattern/d' file.txt              # delete matching lines
sed '3d' file.txt                      # delete line 3
sed '1,5d' file.txt                    # delete lines 1-5
sed '/^$/d' file.txt                   # delete blank lines
sed '2i\inserted line' file.txt        # insert before line 2
sed '2a\appended line' file.txt        # append after line 2
sed '2c\replaced line' file.txt        # replace line 2
```

## fa-layer-group sed Multi-line

```bash
sed 'N;s/\n/ /' file.txt              # join pairs of lines
sed ':a;N;$!ba;s/\n/ /g' file.txt     # join all lines into one
sed '/start/,/end/d' file.txt         # delete range between patterns
sed '/start/,/end/s/old/new/g' file.txt
```

## fa-floppy-disk sed In-place Edit

```bash
sed -i.bak 's/old/new/g' file.txt     # in-place with backup
sed -i '' 's/old/new/g' file.txt      # macOS in-place (no backup)
sed -i -e 's/old/new/g' file.txt
```

## fa-map-pin sed Addresses & Ranges

```bash
sed -n '5p' file.txt                   # print only line 5
sed -n '5,10p' file.txt                # print lines 5-10
sed -n '$p' file.txt                   # print last line
sed -n '/pattern/p' file.txt           # print matching lines
sed -n '/start/,/end/p' file.txt       # print range between patterns
sed '1~2d' file.txt                    # delete every other line starting from 1
```

## fa-terminal awk Basic Usage

```bash
awk '{print}' file.txt                 # print all lines
awk '{print $1}' file.txt              # print first field
awk '{print $1, $3}' file.txt          # print field 1 and 3
awk '/pattern/{print}' file.txt        # print matching lines
awk '{print NR, $0}' file.txt          # print with line numbers
```

## fa-table-columns awk Fields & Separators

```bash
awk -F',' '{print $1, $2}' file.csv   # comma separator
awk -F'\t' '{print $1}' file.tsv      # tab separator
awk -F':' '{print $1}' /etc/passwd
awk 'BEGIN{FS=","; OFS="|"} {print $1, $2}' file.csv
awk '{print NF}' file.txt              # number of fields per line
awk '{print $NF}' file.txt             # last field
awk '{print $(NF-1)}' file.txt         # second to last field
```

## fa-filter awk Patterns & Conditions

```bash
awk '$3 > 100' file.txt                # field 3 greater than 100
awk '$1 == "root"' /etc/passwd
awk '/start/,/end/' file.txt           # range pattern
awk 'NR > 1' file.txt                  # skip header line
awk 'NR >= 5 && NR <= 10' file.txt     # lines 5 to 10
awk '$3 ~ /^[0-9]+$/' file.txt         # regex match field 3
awk '$3 !~ /pattern/' file.txt         # regex NOT match
```

## fa-boxes-stacked awk Variables

```bash
awk '{sum += $1} END {print sum}' file.txt
awk '{sum += $1; count++} END {print sum/count}' file.txt
awk 'BEGIN {pi=3.14159; print pi}'
awk '{arr[NR] = $0} END {for(i=NR;i>0;i--) print arr[i]}' file.txt
awk '{len = length($0); if(len > max) max = len} END {print max}' file.txt
```

## fa-code-branch awk Control Flow

```bash
awk '{if ($3 > 50) print $1, "high"; else print $1, "low"}' file.txt
awk '{
  for (i = 1; i <= NF; i++) {
    if ($i > 100) print $i
  }
}' file.txt
awk '{
  i = 1
  while (i <= NF) {
    print $i; i++
  }
}' file.txt
```

## fa-database awk Arrays

```bash
awk '{count[$1]++} END {for(k in count) print k, count[k]}' file.txt
awk '{sum[$1] += $2} END {for(k in sum) print k, sum[k]}' file.txt
awk '!seen[$1]++' file.txt             # deduplicate by field 1
awk '{a[$1] = a[$1] ? a[$1]","$2 : $2} END {for(k in a) print k, a[k]}' file.txt
```

## fa-wand-magic-sparkles awk Functions

```bash
awk '{print length($0)}' file.txt      # string length
awk '{print substr($0, 1, 10)}' file.txt
awk '{print toupper($1)}' file.txt
awk '{print tolower($1)}' file.txt
awk '{print index($0, "pattern")}' file.txt
awk '{print split($0, arr, ",")}' file.txt
awk '{gsub(/old/, "new"); print}' file.txt
```

## fa-bolt awk One-liners

```bash
awk 'END{print NR}' file.txt           # count lines
awk '{s+=$1} END{print s}' file.txt    # sum column 1
awk 'NR%2==0' file.txt                 # print even lines
awk '{gsub(/\r/,""); print}' file.txt  # remove carriage returns
awk '{$1=$1; print}' file.txt          # normalize whitespace
awk '{print length, $0}' file.txt | sort -rn | head -5  # longest lines
```

## fa-lightbulb Practical Examples

```bash
# Extract usernames from passwd
awk -F':' '{print $1}' /etc/passwd

# Sum file sizes in directory listing
ls -l | awk '{sum += $5} END {print sum " bytes"}'

# CSV column extraction with header skip
awk -F',' 'NR > 1 {print $2, $5}' data.csv

# Find top 10 IPs in access log
awk '{print $1}' access.log | sort | uniq -c | sort -rn | head -10

# Swap two columns
awk '{print $2, $1, $3}' file.txt

# Count occurrences of unique values in column 1
awk -F',' '{count[$1]++} END {for(k in count) print k, count[k]}' data.csv

# Remove HTML tags
sed 's/<[^>]*>//g' file.html

# Add line numbers to file
sed '=' file.txt | sed 'N;s/\n/\t/'
```
