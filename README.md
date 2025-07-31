# EntKit

> A lightweight extension library for [Ent](https://entgo.io), offering essential features like soft delete, optimistic locking, filters, and common mixins.

**EntKit** is a utility toolkit for the Ent ORM, providing common schema enhancements, method extensions, and query helpers to help you build entity logic more efficiently.

---

## 📦 Features

- 🧩 Soft Delete: Auto-generate soft delete methods
- 🔒 Optimistic Locking: Version-based update with retry support
- 🔍 Generic field filters
- 🧬 Useful Mixins: Timestamps, soft delete, optimistic lock, etc.

---

## 🔧 Installation

```bash
go get github.com/your-org/entkit@latest
```

---

## 🚀 Usage Example

### 1. Register Extensions

```go
func main() {
    softdelete, err := softdelete.NewExstension()  // initialize SoftDelete extension
    if err != nil {
        log.Fatalf("failed to create softdelete extension: %v", err)
    }

    optimisticlock := optimisticlock.NewExtension(optimisticlock.WithRetry())  // initialize OptimisticLock extension

    err = entc.Generate("./schema", &gen.Config{
        Features: []gen.Feature{
            gen.FeatureIntercept, // required for soft delete
        },
    }, entc.Extensions(softdelete, optimisticlock))

    if err != nil {
        panic(err)
    }
}
```

Generate code with:

```bash
go generate ./...
```

### 2. Soft Delete Mixin & Method

Add `SoftDeleteMixin` to your schema:

```go
// ent/schema/user.go
func (User) Mixin() ent.Mixin {
    return []ent.Mixin{
        softdelete.SoftDeleteMixin(),
    }
}
```

Then generate the code:

```bash
go generate ./...
```

Call soft delete:

```go
err := client.Debug().User.SoftDelete(ctx, userID)
```

Intercept query to auto-filter soft-deleted records:

```go
client.Intercept(softdelete.Interceptor())

client.Debug().User.Query().Where(user.NameEQ("test")).AllX(ctx)

// SQL:
// SELECT `users`.`id`, `users`.`deleted_at`, `users`.`name`, `users`.`age` 
// FROM `users` WHERE `users`.`name` = ? AND `users`.`deleted_at` IS NULL 
// args=[Alice]
```

Skip soft delete filter:

```go
client.Debug().User.Query().Where(user.NameEQ("test")).AllX(
    softdelete.Skip(ctx),
)

// SQL:
// SELECT `users`.`id`, `users`.`deleted_at`, `users`.`name`, `users`.`age` 
// FROM `users` WHERE `users`.`name` = ?
// args=[Alice]
```

---

### 3. Optimistic Locking with VersionMixin

```go
// ent/schema/user.go
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        optimisticlock.OptimisticLockMixin{},
    }
}
```

Generate code:

```bash
go generate ./...
```

Update with version check:

```go
err := client.User.UpdateOneWithLock(ctx, userID, oldVersion, func(uvuo *ent.UserUpdateOne) *ent.UserUpdateOne {
    return uvuo.SetName("Alice").SetAge(18)
})

// If version conflict occurs, err will be ent.ErrOptimisticLock
```

Update with retry:

```go
retryCount := 5
interval := 1 * time.Second

err := client.User.UpdateOneWithLockAndRetry(ctx, userID, oldVersion, func(uvuo *ent.UserUpdateOne) *ent.UserUpdateOne {
    return uvuo.SetName("Alice").SetAge(18)
})

// Retries on version conflict
```

---

## 🧬 Mixins

### TimeMixin

Adds standard timestamp fields to schema:

- `created_at`: Entity creation time
- `updated_at`: Updated automatically on each update

Example:

```go
func (User) Mixin() []ent.Mixin {
    return []ent.Mixin{
        entkit.TimeMixin{},
    }
}
```

---

## 🔍 Filters

### WithPagination

Adds pagination support to query chain (limit + offset):

```go
qry := client.Debug().User.Query()

entkit.WithPagination(qry, 1, 10).AllX(ctx)
```

---

## 🤝 Contribution

Based on [Ent](https://entgo.io) ORM.
