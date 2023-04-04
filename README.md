<p align="center">
  <img src="https://raw.githubusercontent.com/Eklavya-11/Hidu/main/assets/Hidu.jpeg">
</p>
<h1 align="center">Hidu - Hidden Input Parameters Finder</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>
<h3 align="center">Hidu is a tool designed to find hidden input parameters in web applications.</h3>


## - Installation & Requirements:
```
> git clone https://github.com/Eklavya-11/hidu.git

> cd hidu

> go build main.go

> mv main hidu

> chmod +x hidu

> ./hidu -h
```
<br>


## - Usage & Explanation:

In order to ensure stability, some web applications use forms, which may result in the application handling hidden input parameters within its source code.<br> A typical example of such an input parameter is: `<input type="hidden" name="validate" value="test">`.

This is where Hidu comes in handy. When a web environment has an input parameter like this: <br> `<input type="hidden" name="test" value="">` <br> with a value of 0, it's highly probable that the parameter is reflected in the front-end, allowing for the exploitation of reflected XSS.

For example, executing the command cat index.html yields the following output: `<input type="hidden" name="testing" value="">`.
Hidu reads from standard input.

You can use a file containing a list of targets as well: <br>
cat targets | hidu

**Hidu only brings to us the url to be tested, so, to test if parameter is reflecting, you can use other tools such as: httpx, kxss or manual analisys.**

## This project is solely intended for educational and bug bounty purposes. I do not endorse any illegal activities.
If any error in the program, talk to me immediatly.
