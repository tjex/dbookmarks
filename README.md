# dbookmarks

This is a very simple go program that extracts urls (using [xurls](https://github.com/mvdan/xurls) from a text file, pipes them into [dmenu-mac](https://github.com/oNaiPs/dmenu-mac) 
and opens them in browser upon selection. It would've been simpler to do as a bash script, but I'm teaching myself go (and programming 
in general - so this is why).

For context, it's in following with Luke Smith's "bookmark management" approach as seen [here](https://www.youtube.com/watch?v=d_11QaTlf1I). 

URLS in a text file are read line by line. If there are two URLS on one line, only the first one will be picked.

The program requires one argument, the location of the bookmark file. 

```bash
dbookmarks ~/path/to/bookmarks.txt 
```

bookmarks.txt
```
black hole https://google.com
shameless plug https://tjex.net
https://github.com
```
Non-url text can be on the sameline. It will be stripped out, and you will only be presented the URLS in the dmenu-mac UI. 

## Installation
```
git clone https://github.com/tjex/dbookmarks
cd dbookmarks
go build .
```
