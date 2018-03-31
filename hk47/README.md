
# hk47

*"It's like if C-3PO had the personality of Deadpool." - scottski02 on [YouTube](https://www.youtube.com/watch?v=Vg1gTas7OAA)*

hk47 is a local text-to-speech. it sounds like this: https://soundcloud.com/juuso-haavisto/example/s-CZzxn

**how this works** is that each alphabet (in latin) is mapped to a sound file, which is then played in sequence

notes:
- requires `sox` on the vacuum cleaner
- works only for text in lowercase alphabet
- i think is cool even with this functionality

basically, when no TTS engine is available (`say` of macOS or whatever cloud providers you have), the vacuum cleaner should do Star Wars droid bleeps instead

**why** makes more sense if you look at the other folders and figure out the purpose of the whole project -- if your vacuum is able to analyze the room, how could would it to ask it to speak itself out loud?

endgame would ability to translate beeps back to text. but, it seems rather hard problem

there are some example sound files generated with a synthesizer on the `/sounds` folder.

ill try to find the time to generate more upbeat sounds and find a tune palette which sounds nice overall.

name is a reference to the best droid in Star Wars universe: https://www.youtube.com/watch?v=Vg1gTas7OAA
