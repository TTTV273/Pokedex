# Tests

By now, you should be familiar enough with [unit testing](https://en.wikipedia.org/wiki/Unit_testing). Like anything, the more you do it, the better you will be at it. More importantly, you will begin to understand how to write code that is easier to test.

We'll use [TDD](https://en.wikipedia.org/wiki/Test-driven_development) for this part, and [Go's testing package](https://pkg.go.dev/testing).

If you're unsure how to write unit tests in Go, I highly recommend reading Dave Cheney's [excellent blog post](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) on the subject.

## Assignment

* Create a new `cleanInput(text string) []string` function. For now it should just return an empty slice of strings.

The purpose of this function will be to split the user's input into "words" based on whitespace. It should also lowercase the input and trim any leading or trailing whitespace. For example:

* `hello world` -> `["hello", "world"]`

* `Charmander Bulbasaur PIKACHU` -> `["charmander", "bulbasaur", "pikachu"]`

* Create a new file for some unit tests. I called mine `repl_test.go` since I put `cleanInput` in a new file, `repl.go` (but you can organize your project your way, the only requirement is that the test file ends in `_test.go`). Create a test suite for the `cleanInput` function. Here is the basic structure of the test file:

**All tests** go inside `TestXXX` functions that take a `*testing.T` argument:

  
```go
func TestCleanInput(t *testing.T) {
    // ...
}
```

  

Remember to import the `testing` package if it isn't imported already.

I like to start by creating a slice of test case structs, in this case:

  
```go
cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// add more cases here
}
```

  

Then I loop over the cases and run the tests:

  
```go
for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
}
```

  

* Once you have at least a few tests, run the tests using `go test ./...` from the root of the repo. We expect them to fail.

* Implement the `cleanInput` function to make the tests pass.

**Run and submit** the CLI tests.

---
*Bản dịch:*

# Tests (Kiểm thử)

Bây giờ thì anh đã khá quen thuộc với [unit testing](https://en.wikipedia.org/wiki/Unit_testing) rồi. Giống như mọi thứ khác, càng làm nhiều, anh sẽ càng giỏi. Quan trọng hơn, anh sẽ bắt đầu hiểu cách viết code sao cho dễ test hơn.

Chúng ta sẽ sử dụng [TDD (Test-driven development - Phát triển hướng kiểm thử)](https://en.wikipedia.org/wiki/Test-driven_development) cho phần này, và [package testing của Go](https://pkg.go.dev/testing).

Nếu anh chưa rõ cách viết unit test trong Go, tôi đặc biệt khuyên anh nên đọc [bài viết blog xuất sắc](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) của Dave Cheney về chủ đề này.

## Yêu cầu (Assignment)

* Tạo một hàm mới `cleanInput(text string) []string`. Hiện tại, nó chỉ cần trả về một slice string rỗng.

Mục đích của hàm này là tách input của người dùng thành các "từ" dựa trên khoảng trắng (whitespace). Nó cũng phải chuyển input thành chữ thường (lowercase) và loại bỏ khoảng trắng ở đầu hoặc cuối chuỗi (trim). Ví dụ:

* `hello world` -> `["hello", "world"]`

* `Charmander Bulbasaur PIKACHU` -> `["charmander", "bulbasaur", "pikachu"]`

* Tạo một file mới cho vài bài unit test. Tôi gọi file của mình là `repl_test.go` vì tôi đặt hàm `cleanInput` trong một file mới tên là `repl.go` (nhưng anh có thể tự tổ chức project theo cách của mình, yêu cầu duy nhất là file test phải có đuôi là `_test.go`). Tạo một bộ test (test suite) cho hàm `cleanInput`. Dưới đây là cấu trúc cơ bản của file test:

**Tất cả các test** đều nằm trong các hàm `TestXXX` nhận một tham số `*testing.T`:

```go
func TestCleanInput(t *testing.T) {
    // ...
}
```

Nhớ import package `testing` nếu nó chưa được import.

Tôi thích bắt đầu bằng cách tạo một slice chứa các struct test case, trong trường hợp này:

```go
cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// thêm các case khác ở đây
}
```

Sau đó tôi lặp qua các cases và chạy test:

```go
for _, c := range cases {
	actual := cleanInput(c.input)
	// Kiểm tra độ dài của slice 'actual' so với slice 'expected'
	// nếu chúng không khớp, dùng t.Errorf để in ra thông báo lỗi
	// và đánh rớt (fail) bài test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Kiểm tra từng từ trong slice
		// nếu chúng không khớp, dùng t.Errorf để in ra thông báo lỗi
		// và đánh rớt bài test
	}
}
```

* Sau khi anh có ít nhất vài test, hãy chạy test bằng lệnh `go test ./...` từ thư mục gốc của repo. Chúng ta kỳ vọng các bài test sẽ fail.

* Triển khai (implement) hàm `cleanInput` để các bài test pass (thành công).

**Chạy và nộp** (run and submit) các bài test CLI.
