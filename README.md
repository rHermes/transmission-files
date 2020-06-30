# transmission-files

I wanted to get a list of the directories and files that my `transmission-gtk` had.
Didn't find a nice way to do this, so ended up parsing the `.resume` files, normally
located in `~/.config/transmission/resume`.

The tool either reads from stdin or from the list of resume files given as args.
It prints out the current destination for each of them
