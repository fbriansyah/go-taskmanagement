# Setup Project
Pada bagian ini kita memulai projek dengan menyiapkan device kita dengan install aplikasi penunjang, seperti berikut ini.
- [Make](https://gnuwin32.sourceforge.net/packages/make.htm) digunakan untuk menyederhanakan command.
- [Protoc](https://github.com/protocolbuffers/protobuf/releases) digunakan untuk generate code dari file `proto`.
- [Buf](https://buf.build/docs/cli/installation/) digunakan untuk generate grpc code dari file `proto`.

## Additional Info
Project ini menggunakan go versi 1.24.1 dimana terdapat perbedaan dalam config tools.
- Untuk versi Go di bawah 1.24.1 un-comment code pada `internal/tools.go`
- Untuk Versi go 1.24.1 keatas bisa menggunakan config tools pada `go.mod`