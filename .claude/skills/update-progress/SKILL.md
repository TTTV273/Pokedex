---
name: update-progress
description: Cập nhật tiến độ học tập vào Share_Memory.md khi hoàn thành bài học. Use when user says "cập nhật tiến độ", "ghi nhận bài học", "update progress", or completes a lesson in Learn HTTP Clients in Go course.
disable-model-invocation: true
allowed-tools: Read, Write, Edit, Glob
---

# /update-progress - Cập nhật tiến độ học tập

Cập nhật `Share_Memory.md` khi hoàn thành một bài học với format chuẩn và terminology đúng.

**Cú pháp:** `/update-progress <CHAPTER>/<LESSON>`

**Ví dụ:** `/update-progress CH11/L02`

---

## Terminology Guidelines

### ✅ Được phép:
- "Completed successfully", "Successfully implemented"
- "Can apply", "Can implement", "Can create"
- "Understands", "Solid understanding", "Good grasp"
- "Passed X/X tests"

### ❌ CẤM DÙNG:
- "Mastery", "Mastered", "Expert", "Advanced mastery"
- "Perfect", "Excellence", "Flawless"

---

## Quy trình

### Bước 1: Thu thập thông tin

Đọc các file sau:
1. `$ARGUMENTS/Lesson.md` - Nội dung bài học
2. `$ARGUMENTS/*.go` - Code đã implement
3. Kết quả test từ user (nếu có)

### Bước 2: Chuẩn bị data

```yaml
lesson_id: "L02 - Pros and Cons"
chapter: "Ch11 - Mark and Sweep GC"
date: "DD/MM/YYYY"
lesson_focus: "Mô tả 1-2 câu về mục tiêu bài học"
key_insights:
  - "Insight 1: Giải thích ngắn gọn"
  - "Insight 2: Giải thích ngắn gọn"
capabilities:
  - "**Can implement** [hành động cụ thể]"
  - "**Can apply** [pattern/concept cụ thể]"
  - "**Understands** [khái niệm]"
test_result: "X/X tests passed"
```

### Bước 3: Ghi vào Share_Memory.md

#### 3.1. Cập nhật "Bài học mới nhất" (đầu file)

```markdown
*   **Chương hiện tại:** [Chapter Name] (Progress: X/x lessons completed)
*   **Bài học mới nhất:** [Lesson ID] ✅ **HOÀN THÀNH** (DD/MM/YYYY)
    *   ✅ **Skill Gained:** [Tóm tắt kỹ năng chính]
    *   ✅ **Key Insight:**
        *   **[Insight 1]:** [Giải thích]
        *   **[Insight 2]:** [Giải thích]
    *   **Capabilities:**
        *   **Can implement** [hành động]
        *   **Can apply** [pattern]
```

#### 3.2. Thêm entry chi tiết

```markdown
#### ✅ [Lesson ID] (DD/MM/YYYY) - COMPLETED
* **Lesson Focus:** [Mô tả mục tiêu bài học]
* **Key Insight:**
    *   **[Insight Name 1]:** [Giải thích chi tiết]
    *   **[Insight Name 2]:** [Giải thích chi tiết]
* **Capabilities:**
    *   **Can implement** [hành động cụ thể]
    *   **Can apply** [pattern/concept]
    *   **Understands** [khái niệm]
* **Trạng thái:** ✅ Hoàn thành thành công.
```

---

## Checklist trước khi lưu

- [ ] Không có từ "mastery", "master", "mastered"
- [ ] Không có từ "expert", "expertise"
- [ ] Tất cả claims đều có bằng chứng (test passed, code working)
- [ ] Dùng đúng format với **bold** cho Capability types
- [ ] Ngày tháng đúng format DD/MM/YYYY
- [ ] Progress counter được cập nhật

---

## Lưu ý

1. **Không thay đổi các entry cũ** - Chỉ thêm entry mới
2. **Giữ nguyên thứ tự chronological** - Entry mới nhất ở trên
3. **Dùng tiếng Việt** - Cho Key Insights và explanations
4. **Technical terms giữ tiếng Anh** - malloc, free, pointer, struct, etc.
