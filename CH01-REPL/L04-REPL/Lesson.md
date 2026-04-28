# REPL

[REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) stands for "read-eval-print loop". It's a common type of program that allows for interactivity. You type in a command, and the program evaluates it and prints the result. You can then type in another command, and so on.

Your native terminal's command line is a REPL! Our Pokedex will _also_ be a REPL. In this step, we're just going to build the bones of the REPL: a loop that parses and cleans an input and prints the first word back to the user. Here is what your program should be able to do after this step:

  
```bash
> go run .
Pokedex > help me
help
Pokedex > Exit
exit
Pokedex > WHAT even is a pokeman?
what
```

  

_`Pokedex > ` is our own custom command line prompt. The program waited and recorded my input (in the first instance, `help me`) and didn't continue until I pressed `enter`_.

## Assignment

* Remove your "Hello, World!" logic.

* Create support for a simple REPL

* Create a [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) that reads from `os.Stdin`, for example: `scanner := bufio.NewScanner(os.Stdin)`. When you later call [`scanner.Scan`](https://pkg.go.dev/bufio#Scanner.Scan) it will block and wait for input until the user presses enter.

* Start an infinite `for` loop. This loop will execute once for every command the user types in (we don't want to exit the program after just one command)

* Use [fmt.Print](https://pkg.go.dev/fmt#Print) to print the prompt `Pokedex > ` **without** a newline character

* Use the scanner's [.Scan](https://pkg.go.dev/bufio#Scanner.Scan) and [.Text](https://pkg.go.dev/bufio#Scanner.Text) methods to get the user's input as a string

* Clean the user's input string

* Capture the first "word" of the input and use it to print: `Your command was: <first word>`

* Test your program. Here's my example session:

  
```
wagslane@MacBook-Pro-2 pokedexcli % go run .
Pokedex > well hello there
Your command was: well
Pokedex > Hello there
Your command was: hello
Pokedex > POKEMON was underrated
Your command was: pokemon
```

  

You can terminate the program by pressing `ctrl+c`.

* Run the CLI again and [`tee`](https://en.wikipedia.org/wiki/Tee_(command)) the output (copies the stdout) to a new file called `repl.log` (and `.gitignore` the log).

  
```bash
go run . | tee repl.log
```

  

* Use this as the first input: `CHARMANDER is better than bulbasaur`.

* Use this as the second input: `Pikachu is kinda mean to ash`.

* Terminate the program by pressing `ctrl+c`.

**Run and submit** the CLI tests.

---
*Bản dịch:*

# REPL

[REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) là viết tắt của "read-eval-print loop" (vòng lặp đọc-đánh giá-in). Đây là một loại chương trình phổ biến cho phép tương tác. Anh gõ vào một lệnh, chương trình sẽ đánh giá (thực thi) nó và in ra kết quả. Sau đó anh có thể gõ một lệnh khác, và cứ tiếp tục như vậy.

Command line (dòng lệnh) trên terminal của anh chính là một REPL! Pokedex của chúng ta *cũng sẽ* là một REPL. Trong bước này, chúng ta sẽ chỉ xây dựng bộ khung cơ bản của REPL: một vòng lặp dùng để parse (phân tích) và làm sạch input, sau đó in từ đầu tiên trả lại cho người dùng. Đây là những gì chương trình của anh có thể làm được sau bước này:

```bash
> go run .
Pokedex > help me
help
Pokedex > Exit
exit
Pokedex > WHAT even is a pokeman?
what
```

_`Pokedex > ` là dấu nhắc lệnh (prompt) tùy chỉnh của riêng chúng ta. Chương trình đã chờ và ghi lại input của tôi (trong trường hợp đầu tiên là `help me`) và không tiếp tục cho đến khi tôi nhấn `enter`_.

## Yêu cầu (Assignment)

* Xóa logic "Hello, World!" của anh đi.

* Xây dựng hỗ trợ cho một REPL đơn giản.

* Tạo một [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) để đọc từ `os.Stdin`, ví dụ: `scanner := bufio.NewScanner(os.Stdin)`. Sau đó, khi anh gọi [`scanner.Scan`](https://pkg.go.dev/bufio#Scanner.Scan), nó sẽ block (chặn) và chờ input cho đến khi người dùng nhấn enter.

* Bắt đầu một vòng lặp `for` vô hạn (infinite loop). Vòng lặp này sẽ thực thi một lần cho mỗi lệnh mà người dùng gõ vào (chúng ta không muốn thoát chương trình chỉ sau một lệnh).

* Sử dụng [fmt.Print](https://pkg.go.dev/fmt#Print) để in ra dấu nhắc `Pokedex > ` **mà không có** ký tự xuống dòng (newline character).

* Sử dụng các phương thức [.Scan](https://pkg.go.dev/bufio#Scanner.Scan) và [.Text](https://pkg.go.dev/bufio#Scanner.Text) của scanner để lấy input của người dùng dưới dạng một chuỗi (string).

* Làm sạch chuỗi input của người dùng (dùng hàm `cleanInput` đã viết ở bài trước).

* Lấy "từ" đầu tiên của input và sử dụng nó để in ra: `Your command was: <first word>`

* Kiểm thử chương trình. Đây là ví dụ phiên chạy (session) của tôi:

```
wagslane@MacBook-Pro-2 pokedexcli % go run .
Pokedex > well hello there
Your command was: well
Pokedex > Hello there
Your command was: hello
Pokedex > POKEMON was underrated
Your command was: pokemon
```

Anh có thể kết thúc chương trình bằng cách nhấn `ctrl+c`.

* Chạy CLI một lần nữa và dùng lệnh [`tee`](https://en.wikipedia.org/wiki/Tee_(command)) để xuất kết quả (copy stdout) ra một file mới tên là `repl.log` (và thêm file log này vào `.gitignore`).

```bash
go run . | tee repl.log
```

* Nhập nội dung sau làm input đầu tiên: `CHARMANDER is better than bulbasaur`.

* Nhập nội dung sau làm input thứ hai: `Pikachu is kinda mean to ash`.

* Kết thúc chương trình bằng cách nhấn `ctrl+c`.

**Chạy và nộp** (run and submit) các bài test CLI.