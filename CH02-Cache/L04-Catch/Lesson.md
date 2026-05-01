# Catch

It's time to catch some pokemon! Catching Pokemon adds them to the user's Pokedex.

## Assignment

* **Add a `catch` command**. It takes the name of a Pokemon as an argument. _Example usage_:

  
```bash
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu escaped!
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
```

  

* Be sure to print the `Throwing a Pokeball at <pokemon>...` message before determining if the Pokemon was caught or not.

* Use the [Pokemon endpoint](https://pokeapi.co/docs/v2#pokemon) to get information about a Pokemon by name.

* Give the user a _chance_ to catch the Pokemon using the [math/rand package](https://pkg.go.dev/math/rand#Rand.Intn).

* You can use the pokemon's "base experience" to determine the chance of catching it. The higher the base experience, the harder it should be to catch.

* Once the Pokemon is caught, add it to the user's Pokedex. I used a `map[string]Pokemon` to keep track of caught Pokemon.

* Test the `catch` command manually - make sure you can actually catch a Pokemon within a reasonable number of tries.

**Run and submit** the CLI tests.

---

# Bắt Pokemon (Catch)

Đến lúc bắt vài con Pokemon rồi! Bắt được Pokemon sẽ thêm chúng vào Pokedex của người dùng.

## Bài tập

* **Thêm lệnh `catch`**. Lệnh này nhận tên của một Pokemon làm tham số. _Ví dụ cách dùng_:

  
```bash
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu escaped!
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
```

  

* Đảm bảo in ra thông báo `Throwing a Pokeball at <pokemon>...` trước khi xác định xem Pokemon có bị bắt hay không.

* Sử dụng [Pokemon endpoint](https://pokeapi.co/docs/v2#pokemon) để lấy thông tin về một Pokemon dựa vào tên.

* Cho người dùng một _cơ hội (tỉ lệ)_ để bắt Pokemon bằng cách sử dụng [package math/rand](https://pkg.go.dev/math/rand#Rand.Intn).

* Bạn có thể sử dụng "base experience" (kinh nghiệm cơ bản) của pokemon để xác định tỉ lệ bắt nó. Kinh nghiệm cơ bản càng cao thì càng khó bắt.

* Một khi Pokemon bị bắt, hãy thêm nó vào Pokedex của người dùng. Mình đã dùng một `map[string]Pokemon` để theo dõi các Pokemon đã bắt được.

* Test lệnh `catch` thủ công - đảm bảo rằng bạn thực sự có thể bắt được một Pokemon trong một số lần thử hợp lý.

**Chạy và nộp** các bài test CLI.
