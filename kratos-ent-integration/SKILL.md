---
name: kratos-ent-integration
description: Use when integrating the Ent ORM with a Kratos application.
---

# Kratos Ent Integration

This skill provides a guide to integrating the Ent ORM with a Kratos application.

## Overview

Ent is a powerful entity framework for Go that can be easily integrated with Kratos.

## Usage

1.  Define your schema in the `internal/data` directory.
2.  Generate the Ent code:
    ```bash
    go generate ./...
    ```
3.  In your `data.go`, create a new Ent client and provide it to your repositories.

```go
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}
	return d, func() {
		log.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
```

## Examples

- [**data.go**](./examples/data.go): An example of how to initialize and use the Ent client in the data layer.

- **[kratos-layout](./../kratos-layout/SKILL.md)**: The Ent schema and generated code typically reside in the `internal/data` directory of your Kratos project.
