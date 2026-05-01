# Explore

After a user uses the `map` commands to find a location area, we want them to be able to see a **list of all the Pokémon** located there.

## Assignment

Add an `explore` command. It takes the name of a location area as an argument.

**Run and submit** the CLI tests.

## Tips

* Use the same [PokeAPI location-area endpoint](https://pokeapi.co/docs/v2#location-areas), but this time you'll need to pass the `name` of the location area being explored. By adding a `name` or `id`, the API will return _a lot_ more information about the location area.

* Feel free to use tools like JSON lint and JSON to Go to help you parse the response.

* Parse the Pokemon's names from the response and display them to the user.

* Make sure to use the caching layer again! Re-exploring an area should be blazingly fast.

* You'll need to alter the function signature of _all_ your commands to allow them to allow parameters. E.g. `explore <area_name>`

**Example usage:**

  
```bash
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados
 - remoraid
 - octillery
 - wingull
 - pelipper
 - shellos
 - gastrodon
Pokedex >
```

---
*Bản dịch:*

# Explore (Khám phá)

Sau khi người dùng sử dụng lệnh `map` để tìm kiếm một khu vực địa điểm (location area), chúng ta muốn họ có thể xem được **danh sách tất cả các Pokémon** có mặt ở đó.

## Yêu cầu (Assignment)

Thêm một lệnh `explore`. Nó nhận vào tên của một khu vực địa điểm (location area) làm tham số.

**Chạy và nộp** (run and submit) các bài test CLI.

## Mẹo (Tips)

* Sử dụng cùng một [endpoint location-area của PokeAPI](https://pokeapi.co/docs/v2#location-areas), nhưng lần này anh sẽ cần truyền vào `name` (tên) của khu vực địa điểm đang được khám phá. Bằng cách thêm `name` hoặc `id`, API sẽ trả về *rất nhiều* thông tin hơn về khu vực đó.

* Cứ tự nhiên sử dụng các công cụ như JSON lint và JSON to Go để giúp anh parse (phân tích) kết quả trả về.

* Lấy (parse) tên của các Pokemon từ response và hiển thị chúng cho người dùng.

* Đảm bảo rằng anh sử dụng lại lớp caching của mình nhé! Việc khám phá lại (re-exploring) một khu vực đáng ra phải nhanh như chớp (blazingly fast).

* Anh sẽ cần phải sửa đổi function signature (chữ ký hàm/định dạng tham số của hàm) của *tất cả* các lệnh để cho phép chúng nhận các tham số (parameters). Ví dụ: `explore <area_name>`.

**Ví dụ sử dụng:**

```bash
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados
 - remoraid
 - octillery
 - wingull
 - pelipper
 - shellos
 - gastrodon
Pokedex >
```

---

## 📝 Ghi chú từ buổi học

### Tên tham số và tên biến có thể khác nhau (Scope)

Khi **định nghĩa** hàm, anh đặt tên tham số tùy ý. Khi **gọi** hàm, anh truyền vào một biến — tên biến không cần trùng với tên tham số. Go sao chép giá trị từ biến vào tham số.

**Ví dụ:**

```go
func cong(a int, b int) int {
    return a + b
}

func main() {
    x := 3
    y := 5
    ketqua := cong(x, y) // x → a, y → b
    fmt.Println(ketqua)  // 8
}
```

**Trong bài này:**

```go
// main.go — biến tên là words
words := cleanInput(input)
command.callback(&cfg, words[1:])

// repl.go — tham số tên là args
func commandExit(cfg *config, args []string) error {
    // args nhận giá trị từ words[1:]
}
```

`words[1:]` là giá trị được truyền vào → Go gán vào `args`. Tên khác nhau, cùng giá trị.

> Giống như: ở nhà mẹ gọi là Vũ, ở trường bạn gọi là Thiên Vũ — nhưng vẫn là một người.

---

### Giải thích 2 lệnh curl khám phá API

**Câu 1 — dùng grep:**

```bash
curl "https://pokeapi.co/api/v2/location-area/canalave-city-area" | grep -o '"pokemon":{[^}]*}' | head -c 300
```

- `curl "..."` — tải JSON về, in ra stdout
- `|` — pipe: kết quả curl → làm input cho lệnh tiếp theo
- `grep -o '"pokemon":{[^}]*}'` — tìm và in ra **chỉ phần khớp** với pattern. `[^}]*` nghĩa là "bất kỳ ký tự nào trừ `}`"
- `| head -c 300` — chỉ lấy 300 ký tự đầu tiên

**Câu 2 — dùng jq:**

```bash
curl "https://pokeapi.co/api/v2/location-area/canalave-city-area" | jq '.pokemon_encounters[0]'
```

- `jq` — tool parse và format JSON
- `'.pokemon_encounters[0]'` — lấy phần tử đầu tiên (index 0) của array `pokemon_encounters`

> Câu 1 dùng regex thô — nhanh nhưng dễ sai với JSON phức tạp. Câu 2 dùng `jq` — hiểu đúng cấu trúc JSON, đáng tin hơn.

---

### Khi nào cần JSON tag, khi nào không?

`encoding/json` match field name **case-insensitive** — nghĩa là `Next` tự khớp với `"next"`, `Previous` khớp với `"previous"` mà không cần tag.

**Nhưng bắt buộc phải có tag khi:**
- Tên JSON có **dấu gạch dưới** `_`: `pokemon_encounters` → không thể tự map sang `PokemonEncounters`
- Tên JSON viết **hoàn toàn khác** với tên field trong Go

**Thói quen tốt:** Luôn thêm JSON tag cho tất cả fields — tránh phụ thuộc vào hành vi case-insensitive ngầm định, code rõ ràng hơn khi đọc lại.

```go
// Không có tag — vẫn chạy được nhờ case-insensitive
type locationAreaResponse struct {
    Next *string
}

// Có tag — rõ ràng, an toàn, convention chuẩn
type locationAreaResponse struct {
    Next *string `json:"next"`
}
```