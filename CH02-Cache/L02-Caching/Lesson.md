# Caching

It's time to implement [caching](https://en.wikipedia.org/wiki/Cache_(computing))! This will make moving around the map feel a lot snappier. We'll be building a _flexible_ caching system to help with performance in future steps.

## What Is a Cache?

A cache temporarily stores data so that future requests for that data can be served faster.

In our case, we'll be caching responses from the PokeAPI so that when we need that same data again, we can grab it from memory instead of making another network request.

## Assignment

* Create a new internal package called `pokecache` in your `internal` directory (if you haven't already created an `internal` directory in your project, do so now). This package will be responsible for all of our caching logic.

I used a `Cache` struct to hold a `map[string]cacheEntry` and a mutex to protect the map across goroutines. A `cacheEntry` should be a struct with two fields:

* `createdAt` - A [time.Time](https://pkg.go.dev/time#Time) that represents when the entry was created.

* `val` - A `[]byte` that represents the raw data we're caching.

You'll probably want to expose a `NewCache()` function that creates a new cache with a configurable `interval` ([time.Duration](https://pkg.go.dev/time#Duration)).

* Create a `cache.Add()` method that adds a new entry to the cache. It should take a `key` (a `string`) and a `val` (a `[]byte`).

* Create a `cache.Get()` method that gets an entry from the cache. It should take a `key` (a `string`) and return a `[]byte` and a `bool`. The `bool` should be `true` if the entry was found and `false` if it wasn't.

* Create a `cache.reapLoop()` method that is called when the cache is created (by the `NewCache` function). Each time an `interval` (the `time.Duration` passed to `NewCache`) passes it should remove any entries that are older than the `interval`. This makes sure that the cache doesn't grow too large over time. For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.

I used a [time.Ticker](https://pkg.go.dev/time#Ticker) to make this happen. If you want additional help, see the `Tips` section below.

Maps are _not_ thread-safe in Go. You should use a [sync.Mutex](https://www.boot.dev/blog/golang/golang-mutex/) to lock access to the map when you're adding, getting entries or reaping entries. It's unlikely that you'll have issues because reaping only happens every ~5 seconds, but it's still _possible_, so you should make your cache package safe for concurrent use.

* Update your code that makes requests to the PokeAPI to use the cache. If you already have the data for a given URL (which is our cache key) in the cache, you should use that instead of making a new request. Whenever you do make a request, you should add the response to the cache.

* **Write at least 1 [test](https://go.dev/doc/tutorial/add-a-test) for your cache package**! The tip below should help you get started.

* Test your application manually to make sure that the cache works as expected. When you use the `map` command to get data for the first time there should be a noticeable waiting time. However, when you use `mapb` it should be instantaneous because the data for that page is already in the cache. Feel free to add some logging that informs you in the command line when the cache is being used.

**Run and submit** the CLI tests from the **root of the repo**.

## Tips

### Clearing the Cache

You can use a [time.Ticker](https://pkg.go.dev/time#Ticker) inside a goroutine started by `NewCache`. In a loop like `for range ticker.C { ... }`, check the entries and remove any whose `createdAt` is older than the cache's `interval`.

### Running Tests

You can run tests for all packages in a Go module by running `go test ./...` from the root of the module.

  
```go
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
```

---

## 📝 Ghi chú từ buổi học

### Tổng quan — Mình đã làm gì?

Bài này xây dựng hệ thống cache theo 3 lớp:
1. **Package `pokecache`** — logic cache độc lập (struct, Add, Get, reapLoop)
2. **Tích hợp vào `config`** — thêm `Cache` vào struct dùng chung
3. **Dùng trong `commandMap`/`commandMapb`** — check cache trước khi gọi API

---

### Bước 1 — Tạo package `internal/pokecache`

```bash
mkdir -p internal/pokecache && touch internal/pokecache/cache.go
```

**Tại sao `internal/`?** — Quy ước đặc biệt của Go: code bên trong chỉ import được trong cùng module. Package bên ngoài không dùng được, giúp ẩn implementation detail.

```
github.com/TTTV273/Pokedex/internal/pokecache  ← chỉ dùng được trong module này
```

---

### Bước 2 — Định nghĩa structs

```go
type cacheEntry struct {
    createdAt time.Time  // thời điểm tạo entry
    val       []byte     // data thô (raw JSON response)
}

type Cache struct {
    entries  map[string]cacheEntry  // kho lưu data, key là URL
    mu       sync.Mutex             // khóa bảo vệ map khi đa luồng
    interval time.Duration          // thời gian sống của entry
}
```

**Tại sao `cacheEntry` fields dùng lowercase?** — Chỉ dùng bên trong package `pokecache`, không package nào khác đọc trực tiếp. Quy tắc: chỉ viết hoa khi package khác cần dùng.

**Tại sao cần `mu sync.Mutex`?** — Go map không thread-safe. `reapLoop` chạy ngầm trong goroutine, có thể xóa entries cùng lúc `Add` đang ghi → crash. Mutex đảm bảo chỉ 1 goroutine dùng map tại một thời điểm.

---

### Bước 3 — Hàm `NewCache`

```go
func NewCache(interval time.Duration) Cache {
    c := Cache{}
    c.entries = make(map[string]cacheEntry)  // map nil không dùng được, cần make
    c.interval = interval
    go c.reapLoop()  // chạy reapLoop trong background goroutine
    return c
}
```

**Tại sao cần `make`?** — `Cache{}` tạo struct với `entries = nil`. Map nil không ghi được, sẽ panic. `make(map[string]cacheEntry)` tạo map rỗng nhưng sẵn sàng dùng.

**`go c.reapLoop()`** — `go` khởi chạy goroutine (lightweight thread) chạy song song. Trong khi REPL đang chờ input, `reapLoop` chạy ngầm dọn dẹp cache định kỳ.

---

### Bước 4 — Method `Add`

```go
func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.entries[key] = cacheEntry{time.Now(), val}
}
```

**`defer c.mu.Unlock()`** — Đảm bảo Unlock luôn được gọi khi hàm kết thúc, dù có lỗi hay không. Tránh deadlock.

---

### Bước 5 — Method `Get`

```go
func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    entry, ok := c.entries[key]
    if ok {
        return entry.val, true
    }
    return nil, false
}
```

**Pattern `entry, ok := map[key]`** — Cách chuẩn của Go để kiểm tra key có trong map không. `ok = true` → có, `ok = false` → không có.

---

### Bước 6 — Method `reapLoop`

```go
func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    for range ticker.C {
        c.mu.Lock()
        for key, entry := range c.entries {
            if time.Since(entry.createdAt) > c.interval {
                delete(c.entries, key)
            }
        }
        c.mu.Unlock()
    }
}
```

**`time.NewTicker` và `ticker.C`** — `ticker.C` là một **channel** (kênh giao tiếp giữa các goroutine). `time.NewTicker` tạo ticker và bên trong nó tự động gửi tín hiệu vào channel `C` mỗi `interval` — mình không cần làm gì thêm.

`for range ticker.C` **không chạy liên tục** — nó **block** (đứng im chờ) cho đến khi có tín hiệu. Flow thực tế:
1. Chờ... chờ... chờ... (5 phút)
2. Nhận tín hiệu → dọn cache
3. Chờ... chờ... chờ... (5 phút)
4. Nhận tín hiệu → dọn cache
5. ...mãi mãi

`for` ở đây không phải "chạy nhanh liên tục" mà là "lặp lại hành động sau mỗi lần nhận tín hiệu".

**Không dùng `defer` trong loop** — `defer` chạy khi **hàm kết thúc**, không phải khi lần lặp kết thúc. Nếu dùng `defer Unlock()` trong `for range ticker.C`, mutex sẽ không bao giờ unlock cho đến khi hàm return (không bao giờ) → deadlock.

**`time.Since(entry.createdAt)`** — trả về thời gian đã trôi qua kể từ lúc entry được tạo. Nếu lớn hơn `interval` → entry quá cũ, xóa đi.

---

### Bước 7 — Tích hợp cache vào `commandMap`

```go
// Thêm vào config struct
type config struct {
    Next     *string
    Previous *string
    Cache    pokecache.Cache  // dùng chung giữa các commands
}

// Khởi tạo trong main.go
cfg := config{
    Cache: pokecache.NewCache(5 * time.Minute),
}
```

**Pattern check cache trước:**

```go
var body []byte
var err error
var res *http.Response
if cachedBody, ok := cfg.Cache.Get(url); ok {
    body = cachedBody  // dùng data từ cache, bỏ qua HTTP request
} else {
    res, err = http.Get(url)   // lưu ý: = không phải :=
    // ...
    body, err = io.ReadAll(res.Body)
    // ...
    cfg.Cache.Add(url, body)   // lưu vào cache cho lần sau
}
// dùng body cho json.Unmarshal
```

**Lỗi hay gặp — Shadow variable:** Nếu dùng `:=` trong else block, Go tạo **biến mới** cùng tên thay vì gán vào biến bên ngoài:

```go
var body []byte
if ... {
    body, err := io.ReadAll(...)  // ❌ tạo body MỚI trong block, body ngoài vẫn nil
    body, err = io.ReadAll(...)   // ✅ gán vào body đã khai báo bên ngoài
}
```

**Fix:** Khai báo `var body []byte` và `var err error` trước if/else, dùng `=` thay vì `:=` bên trong.

---
*Bản dịch:*

# Caching (Lưu trữ đệm)

Đã đến lúc triển khai [caching](https://en.wikipedia.org/wiki/Cache_(computing)) rồi! Việc này sẽ làm cho việc di chuyển xung quanh bản đồ cảm thấy nhanh và mượt mà hơn nhiều. Chúng ta sẽ xây dựng một hệ thống caching *linh hoạt* (flexible) để giúp ích cho hiệu suất ở các bước sau.

## Cache là gì?

Cache lưu trữ tạm thời dữ liệu để các yêu cầu (requests) cho dữ liệu đó trong tương lai có thể được phục vụ nhanh hơn.

Trong trường hợp của chúng ta, chúng ta sẽ cache (lưu đệm) các phản hồi (responses) từ PokeAPI để khi cần lại cùng một dữ liệu đó, chúng ta có thể lấy nó từ bộ nhớ (memory) thay vì phải tạo một request mạng (network request) mới.

## Yêu cầu (Assignment)

* Tạo một internal package mới tên là `pokecache` trong thư mục `internal` của anh (nếu anh chưa tạo thư mục `internal` trong project của mình, hãy làm ngay bây giờ). Package này sẽ chịu trách nhiệm cho toàn bộ logic caching của chúng ta.

Tôi đã sử dụng một struct `Cache` để chứa một `map[string]cacheEntry` và một mutex để bảo vệ map đó khi chạy trên nhiều goroutines. Một `cacheEntry` nên là một struct có hai trường (fields):

* `createdAt` - Một [time.Time](https://pkg.go.dev/time#Time) đại diện cho thời điểm entry (mục) này được tạo ra.

* `val` - Một `[]byte` đại diện cho dữ liệu thô (raw data) mà chúng ta đang cache.

Anh có thể sẽ muốn viết (expose) một hàm `NewCache()` để tạo một cache mới với một tham số `interval` (khoảng thời gian) có thể cấu hình được ([time.Duration](https://pkg.go.dev/time#Duration)).

* Tạo một phương thức `cache.Add()` để thêm một entry mới vào cache. Nó nên nhận một `key` (là `string`) và một `val` (là `[]byte`).

* Tạo một phương thức `cache.Get()` để lấy một entry từ cache. Nó nên nhận một `key` (là `string`) và trả về một `[]byte` cùng với một `bool`. Giá trị `bool` nên là `true` nếu entry được tìm thấy và `false` nếu không tìm thấy.

* Tạo một phương thức `cache.reapLoop()` được gọi khi cache được tạo (bởi hàm `NewCache`). Mỗi khi một khoảng thời gian `interval` (cái `time.Duration` được truyền vào `NewCache`) trôi qua, nó sẽ xóa bất kỳ entry nào cũ hơn `interval` đó. Điều này đảm bảo rằng cache không phình to quá mức theo thời gian. Ví dụ, nếu interval là 5 giây, và một entry được thêm vào 7 giây trước, entry đó sẽ bị xóa.

Tôi đã dùng một [time.Ticker](https://pkg.go.dev/time#Ticker) để thực hiện việc này. Nếu anh cần thêm sự trợ giúp, hãy xem phần `Mẹo (Tips)` bên dưới.

Trong Go, Maps là cấu trúc dữ liệu *không* an toàn luồng (not thread-safe). Anh nên sử dụng một [sync.Mutex](https://www.boot.dev/blog/golang/golang-mutex/) để khóa (lock) quyền truy cập vào map khi anh đang thêm (adding), lấy (getting) entries hoặc đang dọn dẹp (reaping) các entries cũ. Rất khó có khả năng anh gặp lỗi vì việc dọn dẹp (reaping) chỉ xảy ra khoảng mỗi 5 giây, nhưng nó vẫn *có khả năng* xảy ra, vì vậy anh nên làm cho package cache của mình an toàn khi sử dụng đồng thời (concurrent use).

* Cập nhật phần code gửi request tới PokeAPI của anh để sử dụng cache. Nếu anh đã có dữ liệu cho một URL nhất định (đóng vai trò là cache key của chúng ta) trong cache, anh nên dùng dữ liệu đó thay vì tạo một request mới. Bất cứ khi nào anh tạo một request mới, anh nên thêm response của nó vào cache.

* **Viết ít nhất 1 [test](https://go.dev/doc/tutorial/add-a-test) cho package cache của anh**! Cái mẹo bên dưới sẽ giúp anh bắt đầu.

* Kiểm thử ứng dụng của anh thủ công để đảm bảo rằng cache hoạt động như mong đợi. Khi anh dùng lệnh `map` để lấy dữ liệu lần đầu tiên, sẽ có một khoảng thời gian chờ đáng kể. Tuy nhiên, khi anh dùng `mapb` nó sẽ hiển thị ngay lập tức vì dữ liệu của trang đó đã nằm sẵn trong cache rồi. Anh cứ tự do thêm một vài dòng log (in ra terminal) để báo cho anh biết khi nào cache đang được sử dụng nhé.

**Chạy và nộp** (run and submit) các bài test CLI từ **thư mục gốc của repo**.

## Mẹo (Tips)

### Dọn dẹp Cache (Clearing the Cache)

Anh có thể sử dụng [time.Ticker](https://pkg.go.dev/time#Ticker) bên trong một goroutine được khởi chạy bởi `NewCache`. Trong một vòng lặp như `for range ticker.C { ... }`, hãy kiểm tra các entries và xóa bất kỳ entry nào có `createdAt` cũ hơn `interval` của cache.

### Chạy Tests

Anh có thể chạy tests cho tất cả các packages trong một Go module bằng cách chạy lệnh `go test ./...` từ thư mục gốc của module đó.

```go
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
```