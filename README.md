# go-do
A simple CLI Todo application made in Go

<a id="readme-top"></a>


<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Godo Screen Shot][product-screenshot]]

By doing this small cli project, i wanted to learn a bit of Go. The todo application is a simple cli that i will add simple features to.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Golang]]


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started



### Prerequisites

* go 1.23.2 >=
  ```
  find the download on https://go.dev/dl/
  ```

### Installation

1. Download Go 
2. Clone the repo
   ```sh
   git clone https://github.com/Pjoess/go-do.git
   ```
3. Install using the install.sh file
   ```sh
   cd ./godo
   ./install.sh
   ```
4. Get started with go-do!
   ```
   godo -help
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

To see all commands available:
```sh
godo -help
```

Add a new todo:
```sh
godo -add "String"
```

List all todo's:
```sh
godo -list
```

Delete a todo:
```sh
godo -delete index
```

Complete toggle a todo:
```sh
godo -complete index
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


[product-screenshot]: picture.png
[Golang]: https://go.dev