# Analysis

Initialization
```
set b 99
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000           << b is initially 109900 and increased with 17 at the end of every outer loop
set c b
sub c -17000            << c is initially 126900 and not changed
```

Outer loop
```
---
    set f 1             << f determines whether h is increased at the end of the middle loop
    set d 2
    ---
        set e 2
        ---
            set g d
            mul g e
            sub g b
            jnz g 2     << if g == 0 -> we set f to 0, g == 0 if b==d*e
              set f 0
            sub e -1    << e is increased on each inner loop
            set g e
            sub g b
        jnz g -8        << if g != 0 -> g==0 iff b==e
        ---
        sub d -1        << d is increased on each middle loop
        set g d
        sub g b
    jnz g -13           << if g != 0 -> we loop again, g==0 iif b==d
    ---
    jnz f 2             << if f == 0 -> we increase h, f==0 iif ???
      sub h -1          << only modification to h  
    set g b
    sub g c
    jnz g 2             << if g == 0 -> we are done, g==0 iff b==c
      jnz 1 3           << only exit point
    sub b -17           << must happen 1000x to make b==c
jnz 1 -23
---
```
