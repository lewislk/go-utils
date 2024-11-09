# 在golang中，各种进制的数值
- 以`292`为例
```
	0b100100100 // 二进制表示法 Ob 开头
	0444        // 8进制表示法 0 开头
	292         // 10进制表示法
	0x124       // 16进制表示法 0x开头
```

# os.OpenFile 参数说明
```
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
```
- name 文件名
- flag
  - os.O_RDONLY‌：只读模式。
  - os.O_WRONLY‌：只写模式。
  - os.O_RDWR‌：读写模式。
  - os.O_CREATE‌：如果文件不存在，则创建文件。
  - os.O_TRUNC‌：如果可能，打开时清空文件。
  - os.O_APPEND‌：写入时追加到文件末尾。
- perm
  - 表示文件的模式和权限。这个参数是一个 Unix 样式的权限模式，通常是一个整数，表示为八进制数，如 0666 或 0755
  - 这个参数是一个uint32的数，其中0~9位，用来表示 Unix 文件权限(`rwxrwxrwx`)
  - 例如，希望创建的文件对所有用户是可读可写的(rw_rw_rw_)，那二进制表示法(`110110110`),对应8进制(`0666`)