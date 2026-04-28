# Ưu tiên kiểm thử hướng bảng (Table Driven Tests)

Tôi là một người rất hâm mộ việc kiểm thử, đặc biệt là [unit testing](https://dave.cheney.net/2019/04/03/absolute-unit-test) và TDD ([tất nhiên là phải thực hiện đúng cách](https://www.youtube.com/watch?v=EZ05e7EMOLM)). Một thực hành phổ biến trong các dự án Go là ý tưởng về kiểm thử hướng bảng (table driven test). Bài viết này sẽ khám phá cách thức và lý do tại sao nên viết kiểm thử hướng bảng.

Giả sử chúng ta có một hàm dùng để tách chuỗi:

```go
// Split tách s thành các chuỗi con cách nhau bởi sep và
// trả về một slice chứa các chuỗi con nằm giữa các dấu ngăn cách đó.
func Split(s, sep string) []string {
    var result []string
    i := strings.Index(s, sep)
    for i > -1 {
        result = append(result, s[:i])
        s = s[i+len(sep):]
        i = strings.Index(s, sep)
    }
    return append(result, s)
}
```

Trong Go, các đơn vị kiểm thử (unit tests) chỉ là các hàm Go thông thường (với một vài quy tắc), vì vậy chúng ta viết một unit test cho hàm này bắt đầu bằng một tệp trong cùng thư mục, cùng tên package là `strings`.

```go
package split

import (
    "reflect"
    "testing"
)

func TestSplit(t *testing.T) {
    got := Split("a/b/c", "/")
    want := []string{"a", "b", "c"}
    if !reflect.DeepEqual(want, got) {
         t.Fatalf("expected: %v, got: %v", want, got)
    }
}
```

Kiểm thử chỉ là các hàm Go thông thường với một vài quy tắc:
1. Tên hàm kiểm thử phải bắt đầu bằng `Test`.
2. Hàm kiểm thử phải nhận một đối số kiểu `*testing.T`. `*testing.T` là một kiểu được inject bởi chính package testing, nhằm cung cấp các phương thức để in kết quả, bỏ qua (skip) hoặc đánh dấu kiểm thử thất bại.

Trong bài kiểm tra này, chúng ta gọi `Split` với một số đầu vào, sau đó so sánh kết quả trả về với kết quả mong đợi.

## Độ bao phủ mã nguồn (Code coverage)

Câu hỏi tiếp theo là: độ bao phủ của package này là bao nhiêu? May mắn thay, công cụ `go` có sẵn tính năng đo độ bao phủ nhánh. Chúng ta có thể thực thi như sau:

```bash
% go test -coverprofile=c.out
PASS
coverage: 100.0% of statements
ok      split   0.010s
```

Kết quả cho thấy chúng ta có độ bao phủ 100%, điều này không gây ngạc nhiên vì chỉ có một nhánh duy nhất trong đoạn mã này.

Nếu muốn tìm hiểu sâu hơn, công cụ `go` có vài tùy chọn để in báo cáo. Chúng ta có thể dùng `go tool cover -func` để phân tích độ bao phủ theo từng hàm:

```bash
% go tool cover -func=c.out
split/split.go:8:       Split          100.0%
total:                  (statements)   100.0%
```

## Vượt xa mức bao phủ 100%

Chúng ta đã viết một trường hợp kiểm thử và đạt 100% độ bao phủ, nhưng câu chuyện chưa dừng lại ở đó. Chúng ta có độ bao phủ nhánh tốt nhưng cần kiểm tra các điều kiện biên. Ví dụ, điều gì xảy ra nếu tách chuỗi bằng dấu phẩy?

```go
func TestSplitWrongSep(t *testing.T) {
    got := Split("a/b/c", ",")
    want := []string{"a/b/c"}
    if !reflect.DeepEqual(want, got) {
        t.Fatalf("expected: %v, got: %v", want, got)
    }
}
```

Hoặc nếu không có dấu ngăn cách nào trong chuỗi nguồn?

```go
func TestSplitNoSep(t *testing.T) {
    got := Split("abc", "/")
    want := []string{"abc"}
    if !reflect.DeepEqual(want, got) {
        t.Fatalf("expected: %v, got: %v", want, got)
    }
}
```

Chúng ta đang bắt đầu xây dựng một tập hợp các trường hợp kiểm thử để vận hành các điều kiện biên. Điều này rất tốt.

## Giới thiệu kiểm thử hướng bảng (Table driven tests)

Tuy nhiên, có rất nhiều sự lặp lại trong các bài kiểm tra. Với mỗi trường hợp, chỉ có đầu vào, kết quả mong đợi và tên bài kiểm tra là thay đổi. Mọi thứ khác đều là mã rập khuôn (boilerplate). Điều chúng ta muốn là thiết lập tất cả đầu vào, đầu ra mong đợi và đưa chúng vào một bộ khung kiểm thử duy nhất. Đây là lúc thích hợp để áp dụng kiểm thử hướng bảng.

```go
func TestSplit(t *testing.T) {
    type test struct {
        input string
        sep   string
        want  []string
    }

    tests := []test{
        {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
        {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
        {input: "abc", sep: "/", want: []string{"abc"}},
    }

    for _, tc := range tests {
        got := Split(tc.input, tc.sep)
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
    }
}
```

Bây giờ, việc thêm một bài kiểm tra mới trở nên cực kỳ đơn giản; chỉ cần thêm một dòng vào cấu trúc `tests`. Ví dụ, điều gì sẽ xảy ra nếu chuỗi đầu vào có dấu ngăn cách ở cuối?

```go
{input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}}, // dấu ngăn cách ở cuối
```

Nhưng khi chạy `go test`, chúng ta nhận được:
`--- FAIL: TestSplit (0.00s)`
`split_test.go:24: expected: [a b c], got: [a b c ]`

Tạm gác lại thất bại của bài kiểm tra, có vài vấn đề cần bàn luận: bằng cách chuyển từ các hàm riêng lẻ sang hàng trong bảng, chúng ta đã mất đi tên của trường hợp kiểm thử bị lỗi.

## Đặt tên cho các trường hợp kiểm thử

Một mô hình phổ biến khác là thêm trường `name` vào cấu trúc kiểm thử.

```go
func TestSplit(t *testing.T) {
    tests := []struct {
        name  string
        input string
        sep   string
        want  []string
    }{
        {name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
        {name: "trailing sep", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
    }

    for _, tc := range tests {
        got := Split(tc.input, tc.sep)
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
        }
    }
}
```

Bây giờ khi thất bại, chúng ta có một cái tên mô tả rõ ràng. Chúng ta thậm chí có thể sử dụng cú pháp `map` để định nghĩa: sử dụng map giúp thứ tự thực thi là *không xác định*, điều này cực kỳ hữu ích để phát hiện các trường hợp kiểm thử phụ thuộc vào trạng thái toàn cục bị thay đổi bởi các bài kiểm tra trước đó.

## Giới thiệu kiểm thử con (Sub tests)

Kể từ Go 1.7, chúng ta có tính năng `sub tests`.

```go
for name, tc := range tests {
    t.Run(name, func(t *testing.T) {
        got := Split(tc.input, tc.sep)
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
    })
}
```

Mỗi subtest là một hàm ẩn danh riêng biệt, cho phép sử dụng `t.Fatalf`, `t.Skipf` mà vẫn giữ được sự gọn gàng của kiểm thử hướng bảng. Bạn có thể chạy riêng lẻ từng subtest bằng flag `-run`:
`go test -run=.*/trailing -v`

## So sánh kết quả bằng thư viện go-cmp

Thay vì dùng `reflect.DeepEqual` đôi khi khó nhìn ra sự khác biệt trong các cấu trúc phức tạp, Google cung cấp thư viện [go-cmp](https://github.com/google/go-cmp). Hàm `cmp.Diff` sẽ tạo ra một mô tả văn bản về sự khác biệt giữa hai giá trị một cách đệ quy.

```go
diff := cmp.Diff(tc.want, got)
if diff != "" {
    t.Fatalf(diff)
}
```

Kết quả báo lỗi sẽ chi tiết hơn:
`split_test.go:27: {[]string}[?->3]:`
`    -: <non-existent>`
`    +: ""`

Nó cho chúng ta biết chuỗi có độ dài khác nhau, tại index thứ 3 đáng lẽ không tồn tại nhưng thực tế lại nhận được một chuỗi rỗng "". Từ đây, việc sửa lỗi hàm `Split` trở nên rõ ràng.

---
*Bản dịch tuân thủ các nguyên tắc logic và kỹ thuật của văn bản gốc.*
