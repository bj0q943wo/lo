module github.com/myusername/lo

go 1.18

require (
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17
)

// Personal fork of samber/lo for learning generics in Go 1.18.
// Original: https://github.com/samber/lo
//
// Notes:
// - Experimenting with additional slice helpers (Compact, Flatten variants)
// - Tracking upstream: last synced 2023-04-01
// - TODO: look into adding a ChunkBy helper (group consecutive equal elements)
