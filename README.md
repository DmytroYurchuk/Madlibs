# Madlibs

# Requirements:
A small service that serves a single REST endpoint (/madlib). This endpoint should return a templated "madlib" sentence:

`It was a {adjective} day. I went downstairs to see if I could {verb} dinner. I asked, "Does the stew need fresh {noun}?"`

The randomized words should be fetched from https://reminiscent-steady-albertosaurus.glitch.me/. There are three separate routes per part of speech: /adjective, /verb, and /noun.

# Implementation
There are two implementations:
- Python
- Go 