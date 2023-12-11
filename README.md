# iprv(1) - calculate electrical current, power, resistance and/or voltage

Kozmix Go, 11-JUL-2023

```
iprv [I=val] [V=val] [Vp=val] [P=val] [R=val]
```


<a name="description"></a>

# Description

**iprv,**
given any two out of electrical current, power, resistance or
voltage, will calculate the other two values. You must call
**iprv**
with exactly two arguments.


<a name="options"></a>

# Options


**iprv**
accepts SI prefixes (k, M, μ etc.) as part of its command line
options, which are:


* **I=val**  
  specify current in amperes
* **V=val**  
  specify voltage in volts (DC or AC RMS)
* **Vp=val**  
  specify voltage in volts (AC peak for sine wave AC)
* **P=val**  
  specify power in watts
* **R=val**  
  specify resistance in ohms
  
* All arguments are case insensitive. The SI prefixes are not.  
  

<a name="example"></a>

# Example

.EX
$ iprv r=8 vp=7.4
resistance R = 8 Ω
voltage    V = 5.23 V
current    I = 654.07 mA
power      P = 3.42 W
.EE


<a name="author"></a>

# Author

svm

