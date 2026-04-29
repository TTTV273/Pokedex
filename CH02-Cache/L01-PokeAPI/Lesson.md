# PokeAPI

Now we're going to use the [PokeAPI](https://pokeapi.co/) to get some real data from the Pokemon world!

## Assignment

* Add the `map` command. It displays the names of `20` location areas in the Pokemon world. Each subsequent call to `map` should display the next 20 locations, and so on. This will be how we explore the Pokemon world. Example usage:

  
```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

  

Here are some pointers for implementing this command:

* You'll need to use the [PokeAPI location-area endpoint](https://pokeapi.co/docs/v2#location-areas) to get the **location areas**. Note that this is a different endpoint than the "location" endpoint. Calling the endpoint without an `id` will return a batch of location areas.

* Update all commands (e.g. `help`, `exit`, `map`) to now accept a pointer to a "config" struct as a parameter. This struct will contain the `Next` and `Previous` URLs that you'll need to paginate through location areas.

* [Here's an example](https://pkg.go.dev/net/http#example-Get) of how to make a GET request in Go.

* [Here's how to unmarshal](https://www.boot.dev/blog/golang/json-golang/#example-unmarshal-json-to-struct-decode) a slice of bytes into a Go struct.

* You can make `GET` requests in your browser or by using [`curl`](https://curl.se/)! It's convenient for testing and debugging.

* Add the `mapb` (map back) command. It's similar to the `map` command, however, instead of displaying the _next_ 20 locations, it displays the _previous_ 20 locations. It's a way to go back.

If you're on the first "page" of results, this command should just print "you're on the first page". Example usage:

  
```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
Pokedex > map
mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c
Pokedex > mapb
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

  

**Run and submit** the CLI tests from the **root of the repo**.

## Tips

* [JSON lint](https://jsonlint.com/) is a useful tool for debugging JSON, it makes it easier to read.

* [JSON to Go](https://mholt.github.io/json-to-go/) a useful tool for converting JSON to Go structs. You can use it to generate the structs you'll need to parse the PokeAPI response. Keep in mind it sometimes can't know the exact type of a field that you want, because there are multiple valid options. For nullable strings, use `*string`.

* I recommend creating an [internal package](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface) that manages your PokeAPI interactions. It's not required, but it's a good organizational and architectural pattern.

---
*Bản dịch:*

# PokeAPI

Bây giờ chúng ta sẽ sử dụng [PokeAPI](https://pokeapi.co/) để lấy một vài dữ liệu thực tế từ thế giới Pokemon!

## Yêu cầu (Assignment)

* Thêm lệnh `map`. Lệnh này hiển thị tên của `20` khu vực địa điểm (location areas) trong thế giới Pokemon. Mỗi lần gọi lệnh `map` tiếp theo sẽ hiển thị 20 địa điểm kế tiếp, và cứ thế. Đây sẽ là cách chúng ta khám phá thế giới Pokemon. Ví dụ sử dụng:

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

Dưới đây là một vài gợi ý để triển khai lệnh này:

* Anh sẽ cần sử dụng [endpoint location-area của PokeAPI](https://pokeapi.co/docs/v2#location-areas) để lấy các **location areas**. Lưu ý rằng endpoint này khác với endpoint "location". Gọi endpoint mà không truyền `id` sẽ trả về một danh sách (batch) các location areas.

* Cập nhật tất cả các lệnh (ví dụ: `help`, `exit`, `map`) để bây giờ chúng nhận vào một con trỏ (pointer) trỏ tới struct "config" làm tham số. Struct này sẽ chứa các URL `Next` (tiếp theo) và `Previous` (trước đó) mà anh sẽ cần để phân trang (paginate) qua các location areas.

* [Đây là một ví dụ](https://pkg.go.dev/net/http#example-Get) về cách thực hiện một HTTP GET request trong Go.

* [Đây là cách unmarshal](https://www.boot.dev/blog/golang/json-golang/#example-unmarshal-json-to-struct-decode) (giải mã) một slice of bytes thành một Go struct.

* Anh có thể thực hiện các `GET` requests trên trình duyệt hoặc bằng lệnh [`curl`](https://curl.se/)! Việc này rất tiện lợi để kiểm thử (testing) và gỡ lỗi (debugging).

* Thêm lệnh `mapb` (map back). Nó tương tự như lệnh `map`, tuy nhiên, thay vì hiển thị 20 địa điểm *kế tiếp*, nó hiển thị 20 địa điểm *trước đó*. Đây là cách để quay lại.

Nếu anh đang ở "trang" (page) kết quả đầu tiên, lệnh này chỉ cần in ra "you're on the first page" (bạn đang ở trang đầu tiên). Ví dụ sử dụng:

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
Pokedex > map
mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c
Pokedex > mapb
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

**Chạy và nộp** (run and submit) các bài test CLI từ **thư mục gốc của repo**.

## Mẹo (Tips)

* [JSON lint](https://jsonlint.com/) là một công cụ hữu ích để debug JSON, nó làm cho JSON dễ đọc hơn.

* [JSON to Go](https://mholt.github.io/json-to-go/) là một công cụ hữu ích để chuyển đổi từ chuỗi JSON sang các Go structs. Anh có thể dùng nó để tạo (generate) các structs cần thiết cho việc parse (phân tích) kết quả trả về từ PokeAPI. Hãy nhớ rằng đôi khi nó không thể biết chính xác kiểu dữ liệu của một trường mà anh muốn, vì có nhiều lựa chọn hợp lệ. Với các chuỗi có thể null (nullable strings), hãy dùng `*string`.

* Tôi khuyên anh nên tạo một [internal package](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface) (package nội bộ) để quản lý các tương tác với PokeAPI của anh. Điều này không bắt buộc, nhưng nó là một pattern cấu trúc và kiến trúc tốt.
