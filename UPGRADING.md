# Upgrade guide

This document describes the notable breaking changes, if any, in each version of
the Go client. See [CHANGELOG.md](./CHANGELOG.md) for a comprehensive list of
changes.

## v2.0

- The module path has been bumped to `github.com/turbopuffer/turbopuffer-go/v2`.
  Update your imports accordingly:

  ```go
  import (
      "github.com/turbopuffer/turbopuffer-go/v2"
      "github.com/turbopuffer/turbopuffer-go/v2/option"
  )
  ```

- The `CopyFromNamespace` field on `NamespaceWriteParams` has been removed in
  favor of a dedicated `CopyFrom` method.

  Old:

  ```go
  ns.Write(ctx, turbopuffer.NamespaceWriteParams{
      CopyFromNamespace: turbopuffer.CopyFromNamespaceParams{
          SourceNamespace: "src",
          SourceRegion:    turbopuffer.String("gcp-us-central1"),
      },
  })
  ```

  New:

  ```go
  ns.CopyFrom(ctx, turbopuffer.NamespaceCopyFromParams{
      SourceNamespace: "src",
      SourceRegion:    turbopuffer.String("gcp-us-central1"),
  })
  ```

- The `BranchFromNamespace` field on `NamespaceWriteParams` has been removed in
  favor of a dedicated `BranchFrom` method.

  Old:

  ```go
  ns.Write(ctx, turbopuffer.NamespaceWriteParams{
      BranchFromNamespace: turbopuffer.BranchFromNamespaceParams{
          SourceNamespace: "src",
      },
  })
  ```

  New:

  ```go
  ns.BranchFrom(ctx, turbopuffer.NamespaceBranchFromParams{
      SourceNamespace: "src",
  })
  ```

- The `encryption` parameter has been restructured.

  Old:

  ```go
  ns.Write(ctx, turbopuffer.NamespaceWriteParams{
      UpsertRows: []turbopuffer.RowParam{ /* ... */ },
      Encryption: turbopuffer.NamespaceWriteParamsEncryption{
          Cmek: turbopuffer.NamespaceWriteParamsEncryptionCmek{KeyName: "..."},
      },
  })
  ```

  New:

  ```go
  ns.Write(ctx, turbopuffer.NamespaceWriteParams{
      UpsertRows: []turbopuffer.RowParam{ /* ... */ },
      Encryption: turbopuffer.EncryptionParamCustomerManaged("..."),
  })
  ```

  A new `Default` variant lets you migrate a namespace from CMEK to default
  encryption.

- The `NewRankByVector` and `NewRankBySparseVector` constructors have been
  renamed to `NewRankByAnn` and `NewRankBySparseKnn`.

- `NamespaceQueryParams.GroupBy` is now `[]turbopuffer.GroupBy` instead of
  `[]string`. Wrap plain attribute names with `turbopuffer.NewGroupByAttr`.

  Old:

  ```go
  ns.Query(ctx, turbopuffer.NamespaceQueryParams{
      GroupBy: []string{"color", "size"},
  })
  ```

  New:

  ```go
  ns.Query(ctx, turbopuffer.NamespaceQueryParams{
      GroupBy: []turbopuffer.GroupBy{
          turbopuffer.NewGroupByAttr("color"),
          turbopuffer.NewGroupByAttr("size"),
      },
  })
  ```
