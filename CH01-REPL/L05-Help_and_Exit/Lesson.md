# Help and Exit

A REPL is only useful if it does something! Our REPL will work using the concept of "commands". A command is a single word that maps to an action.

We're going to support two commands in this step:

* `help`: prints a help message describing how to use the REPL

* `exit`: exits the program

## Assignment

* Remove your logic that prints the first word (the command) back to the user

* Add a callback for the `exit` command. Commands in our REPL are just callback functions with no arguments, but that return an `error`. For example:

  
```go
func commandExit() error
```

  

This function should print `Closing the Pokedex... Goodbye!` then immediately exit the program. I used [`os.Exit(0)`](https://pkg.go.dev/os#Exit).

* Create a "registry" of commands. This will give us a nice abstraction for managing the many commands we'll be adding. I created a struct type that describes a command:

  
```go
type cliCommand struct {
	name        string
	description string
	callback    func() error
}
```

  

Then I created a `map` of supported commands:

  
```go
map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}
```

  

* Register the `exit` command. Update your REPL loop to use the "command" the user typed in to look up the callback function in the registry. If the command is found, call the callback (and print any errors that are returned). If there isn't a handler, just print `Unknown command`.

* Test your program (obviously).

* Add a `help` command, its callback, and register it. It should print:

  
```
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
```

  

You can dynamically generate the "usage" section by iterating over my registry of commands. That way the help command will always be up-to-date with the available commands.

* Test your code again manually.

**Run and submit** the CLI tests from the **root of the repo**.

---
*Bản dịch:*

# Help and Exit (Trợ giúp và Thoát)

Một REPL chỉ hữu ích khi nó có thể thực hiện được việc gì đó! REPL của chúng ta sẽ hoạt động dựa trên khái niệm "các câu lệnh" (commands). Một lệnh là một từ đơn lẻ tương ứng với một hành động.

Trong bước này, chúng ta sẽ hỗ trợ hai lệnh sau:

* `help`: in ra một thông báo trợ giúp mô tả cách sử dụng REPL

* `exit`: thoát khỏi chương trình

## Yêu cầu (Assignment)

* Xóa logic in ra từ đầu tiên (câu lệnh) trả lại cho người dùng mà anh đã viết ở bài trước.

* Thêm một callback (hàm gọi lại) cho lệnh `exit`. Các lệnh trong REPL của chúng ta chỉ là các hàm callback không nhận tham số nào, nhưng trả về một `error` (lỗi). Ví dụ:

```go
func commandExit() error
```

Hàm này nên in ra `Closing the Pokedex... Goodbye!` sau đó ngay lập tức thoát khỏi chương trình. Tôi đã sử dụng hàm [`os.Exit(0)`](https://pkg.go.dev/os#Exit).

* Tạo một "registry" (nơi đăng ký) các lệnh. Việc này sẽ cung cấp cho chúng ta một tính trừu tượng (abstraction) gọn gàng để quản lý nhiều lệnh mà chúng ta sẽ thêm vào sau này. Tôi đã tạo một kiểu struct mô tả một lệnh như sau:

```go
type cliCommand struct {
	name        string
	description string
	callback    func() error
}
```

Sau đó tôi tạo một `map` chứa các lệnh được hỗ trợ:

```go
map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}
```

* Đăng ký lệnh `exit`. Cập nhật vòng lặp REPL của anh để dùng "lệnh" mà người dùng vừa gõ vào để tra cứu (look up) hàm callback tương ứng trong registry (bảng map). Nếu tìm thấy lệnh, hãy gọi hàm callback đó (và in ra bất kỳ lỗi nào được trả về nếu có). Nếu không có trình xử lý (handler) nào cho lệnh đó, chỉ cần in ra `Unknown command`.

* Kiểm thử chương trình (tất nhiên rồi).

* Thêm lệnh `help`, hàm callback của nó, và đăng ký nó vào map. Lệnh này nên in ra:

```
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
```

Anh có thể tạo phần "Usage:" (Cách sử dụng) một cách linh động (dynamically) bằng cách lặp qua registry (map) các lệnh của chúng ta. Bằng cách đó, lệnh help sẽ luôn được cập nhật chính xác với các lệnh hiện có.

* Kiểm thử lại code thủ công một lần nữa.

**Chạy và nộp** (run and submit) các bài test CLI từ **thư mục gốc của repo**.

---

## 📝 Ghi chú từ buổi học

### Hai tầng code trong Go

Trong Go có hai "tầng" để đặt code:

**Tầng 1 — Package level** (ngoài mọi hàm):
- Chỉ dùng để **định nghĩa** — `func`, `type`, `struct`, `const`
- Không thể chạy code hay tạo biến ở đây

**Tầng 2 — Function level** (bên trong hàm):
- Dùng để **thực thi** — tạo biến (`:=`), gọi hàm, tính toán

```go
// Tầng 1: định nghĩa hàm ✅
func commandExit() error { ... }

// Tầng 1: tạo biến ❌ — KHÔNG được phép!
commands := map[string]cliCommand{ ... }

func main() {
    // Tầng 2: tạo biến ✅
    commands := map[string]cliCommand{ ... }
}
```

**Quy tắc nhớ nhanh:**
- Thấy `func`, `type`, `struct` → tầng 1 (ngoài hàm)
- Thấy `:=` hay `var` → tầng 2 (trong hàm)

### Tại sao map nằm ngoài vòng lặp for?

`commands` map chỉ cần tạo **một lần** trước vòng lặp. Nếu đặt bên trong `for`, map sẽ được tạo lại mỗi lần người dùng gõ lệnh — lãng phí không cần thiết vì nội dung map không thay đổi.

```go
func main() {
    commands := map[string]cliCommand{ ... }  // ✅ tạo một lần
    for {
        // dùng commands ở đây
    }
}
```

### Tại sao main.go dùng được hàm từ repl.go mà không cần import?

Vì cả hai file đều có `package main` ở đầu. Trong Go, tất cả file cùng package được **gộp lại thành một** khi compile — không cần import nhau.

`main.go` thấy `commandExit`, `cleanInput`, `cliCommand` vì chúng cùng package `main` với `repl.go`.