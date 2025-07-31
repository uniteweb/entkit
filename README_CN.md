
# EntKit

> A lightweight extension library for [Ent](https://entgo.io), providing essential features like soft delete, optimistic locking, filters, and common mixins.

**EntKit** 是一个面向 [Ent ORM](https://entgo.io) 的轻量级扩展库，内置常用的结构增强、方法生成、查询构造器等功能，帮助开发者更高效地构建实体层逻辑。

---

## 📦 功能模块

- 🧩 **软删除支持**：自动生成软删除方法
- 🔒 **乐观锁更新机制**：支持版本控制及可配置的重试逻辑
- 🔍 **查询过滤器**：标准化结构体查询过滤器构造
- 🧬 **常用 Mixin 集合**：内置时间戳、软删除、版本号字段支持

---

## 🔧 安装

```bash
go get github.com/your-org/entkit@latest
```

---

## 🚀 使用示例

### 1️⃣ 启用扩展并生成代码

```go
func main() {
	softdelete, err := softdelete.NewExtension()
	if err != nil {
		log.Fatalf("failed to create soft delete extension: %v", err)
	}

	optimisticlock := optimisticlock.NewExtension(
		optimisticlock.WithRetry(), // 启用重试机制
	)

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept, // 必须启用以支持 Interceptor（如软删除过滤）
		},
	}, entc.Extensions(softdelete, optimisticlock))

	if err != nil {
		panic(err)
	}
}
```

使用 `go generate` 快速生成代码：

```bash
go generate ./...
```

---

### 2️⃣ 使用 SoftDeleteMixin 实现软删除能力

```go
// ent/schema/user.go
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin(),
	}
}
```

生成软删除方法后，即可调用：

```go
err := client.Debug().User.SoftDelete(ctx, userID)
```

#### 自动过滤已删除数据

```go
client = client.Intercept(softdelete.Interceptor())

client.Debug().User.Query().
	Where(user.NameEQ("Alice")).
	AllX(ctx)

// 生成 SQL:
// SELECT ... FROM users WHERE name = ? AND deleted_at IS NULL
```

#### 显式跳过软删除过滤

```go
client.Debug().User.Query().
	Where(user.NameEQ("Alice")).
	AllX(softdelete.Skip(ctx))

// 生成 SQL:
// SELECT ... FROM users WHERE name = ?
```

---

### 3️⃣ 使用 OptimisticLockMixin 实现乐观锁机制

```go
// ent/schema/user.go
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		optimisticlock.OptimisticLockMixin{},
	}
}
```

#### 乐观锁更新

```go
err := client.User.UpdateOneWithLock(ctx, userID, oldVersion, func(u *ent.UserUpdateOne) *ent.UserUpdateOne {
	return u.SetName("Alice").SetAge(18)
})

// 若版本冲突，err == ent.ErrOptimisticLock
```

#### 自动重试更新

```go
retryCount := 5
interval := 1 * time.Second

err := client.User.UpdateOneWithLockAndRetry(ctx, userID, oldVersion, func(u *ent.UserUpdateOne) *ent.UserUpdateOne {
	return u.SetName("Alice").SetAge(18)
}, retryCount, interval)
```

---

## 🧬 常用 Mixin

### 🕒 TimeMixin

为 Schema 自动添加标准时间戳字段：

- `created_at`：记录实体创建时间
- `updated_at`：每次更新时自动刷新

适用于需要审计或时间追踪的实体模型。

**示例：**

```go
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entkit.TimeMixin{},
	}
}
```

---

## 🔍 查询工具

### 📄 WithPagination

为 Ent 查询添加分页功能（`limit` + `offset`）：

```go
qry := client.Debug().User.Query()

entkit.WithPagination(qry, 1, 10).AllX(ctx)
```

---

## 🤝 参与贡献

本项目基于 [Ent](https://entgo.io) 构建，欢迎贡献者提交 PR 或 Issue！

---

## 📝 License

MIT © Your Name
