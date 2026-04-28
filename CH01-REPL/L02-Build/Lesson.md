# Build

By the end of this step, you'll have a simple working program in Go!

## Assignment

* Create a `main.go` file. It should be part of `package main` in the root of your project and have a `main()` function that just prints the text "Hello, World!".

* [Create a Go module](https://golang.org/doc/tutorial/create-module) in the root of your project. Here's the command:

  
```
go mod init MODULE_NAME
```

  

I recommend naming the module by its remote Git location (you should store all your projects in Git!). For example, my GitHub name is `wagslane` so my module name might be `github.com/wagslane/pokedexcli`.

* Build your program:

  
```bash
go build
```

  

The executable will be named after the directory containing the `main.go` file, hence mine was `pokedexcli`. You should `.gitignore` the executable.

* Run your program:

  
```bash
./pokedexcli
```

  

It should print "Hello, World!" to the console!

**Run and submit** the CLI tests from the **root of the repo**.

---
*Bản dịch:*

# Build (Biên dịch)

Đến cuối bước này, anh sẽ có một chương trình Go đơn giản hoạt động được!

## Yêu cầu (Assignment)

* Tạo một file `main.go`. File này cần thuộc `package main` nằm ở thư mục gốc (root) của dự án và có một hàm `main()` chỉ in ra dòng chữ "Hello, World!".

* [Tạo một Go module](https://golang.org/doc/tutorial/create-module) ở thư mục gốc của dự án. Lệnh như sau:

```bash
go mod init MODULE_NAME
```

Khuyến nghị đặt tên module theo đường dẫn Git remote (anh nên lưu trữ tất cả các dự án trên Git!). Ví dụ: tên GitHub là `wagslane` thì tên module có thể là `github.com/wagslane/pokedexcli`. Tương tự với anh thì có thể là `github.com/TTTV273/pokedexcli`.

* Build (biên dịch) chương trình:

```bash
go build
```

File thực thi (executable) sẽ tự động lấy tên của thư mục chứa file `main.go`. Anh nên thêm file thực thi này vào `.gitignore`.

* Chạy chương trình:

```bash
./pokedexcli
```

Chương trình sẽ in ra "Hello, World!" trên màn hình console!

**Chạy và nộp** (run and submit) các bài test CLI từ **thư mục gốc của repo**.