# Inspect

Just like in the Pokemon _game_, our Pokedex will only allow players to see details about a Pokemon if they have seen it before (or in our case, caught it)

## Assignment

Add an `inspect` command. It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon. **Example usage**:

  
```bash
Pokedex > inspect pidgey
you have not caught that pokemon
Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
```

  

You should _not_ need to make an API call to get this information, since you should have already stored it when the user caught the Pokemon.

If the user has not caught the Pokemon, just print a message saying so.

_Test this one manually! Move on when you're satisfied with your implementation._

---

# Kiểm tra (Inspect)

Giống như trong _trò chơi_ Pokemon, Pokedex của chúng ta sẽ chỉ cho phép người chơi xem thông tin chi tiết về một Pokemon nếu họ đã từng thấy nó trước đây (hoặc trong trường hợp của chúng ta là đã bắt được nó).

## Bài tập

Thêm lệnh `inspect`. Lệnh này nhận tên của một Pokemon và in ra tên, chiều cao, cân nặng, các chỉ số (stats) và (các) hệ (type) của Pokemon đó. **Ví dụ cách dùng**:

  
```bash
Pokedex > inspect pidgey
you have not caught that pokemon
Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
```

  

Bạn _không_ cần thực hiện gọi API để lấy thông tin này, vì bạn đã lưu trữ nó khi người dùng bắt được Pokemon rồi.

Nếu người dùng chưa bắt được Pokemon, chỉ cần in ra thông báo cho biết điều đó.

_Hãy tự test chức năng này thủ công nhé! Bạn có thể chuyển sang bài tiếp theo khi đã hài lòng với phần code của mình._
