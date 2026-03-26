# Tugas Kecil 2 IF2211 Strategi Algoritma: Voxelization Objek 3D menggunakan Octree

## Deskripsi Program
Program ini merupakan aplikasi berbasis *Command Line Interface* (CLI) yang mengimplementasikan algoritma **Divide and Conquer** melalui struktur data **Octree** untuk melakukan *voxelization* pada model 3D berformat `.obj`. 

*Voxelization* adalah proses mengonversi model 3D yang tersusun dari poligon (segitiga halus) menjadi sekumpulan kubus-kubus kecil berukuran seragam (*voxel*), menyerupai gaya visual pada permainan Minecraft. 

**Fitur Utama & Optimasi:**
* **Surface-Only Voxelization:** Program secara presisi menggunakan kalkulasi *Cube-Triangle Intersection* (Separating Axis Theorem) untuk memastikan voxel hanya terbentuk pada permukaan objek, membiarkan bagian dalam tetap kosong (*hollow*).
* **Spatial Pruning:** Program dilengkapi dengan logika pemangkasan (*pruning*) yang secara otomatis menghentikan penelusuran pada ruang hampa atau ruang kosong di dalam model, sehingga mencegah pertumbuhan komputasi eksponensial yang tidak perlu.
* **Concurrency (Goroutine):** Algoritma pembagian ruang (*Divide*) dioptimasi menggunakan eksekusi paralel (Goroutine) pada level kedalaman awal (*early depths*) secara *thread-safe* (menggunakan `sync.Mutex`), memaksimalkan penggunaan CPU *multicore* untuk eksekusi yang lebih cepat (Pemenuhan Spesifikasi Bonus).

---

## Struktur Repositori
Repositori ini disusun berdasarkan spesifikasi tugas standar:
* `bin/` : Menyimpan file *executable* program hasil kompilasi.
* `doc/` : Menyimpan dokumen laporan Tugas Kecil dalam format PDF.
* `src/` : Menyimpan seluruh *source code* bahasa Go (`.go`).
* `test/` : Menyimpan kumpulan data uji berupa file `.obj` asli beserta file `.obj` hasil *voxelization*.

---

## Requirements & Instalasi
Untuk dapat mengkompilasi dan menjalankan program ini, sistem Anda harus memenuhi persyaratan berikut:
1. **Go Compiler:** Versi 1.16 atau yang lebih baru (dibutuhkan untuk dukungan *Go Modules*). Dapat diunduh di [golang.org/dl](https://golang.org/dl/).
2. **Sistem Operasi:** Windows, Linux, atau macOS.
3. **Aplikasi Penampil 3D (Opsional):** Untuk melihat visualisasi output, Anda dapat menggunakan aplikasi seperti **3D Viewer** (bawaan Windows 10/11), **Blender**, atau penampil berbasis web seperti [3DViewer.net](https://3dviewer.net/).

---

## Cara Mengkompilasi Program
Program ini harus dikompilasi terlebih dahulu menjadi file eksekusi (*executable*).
1. Buka *terminal* (Command Prompt / PowerShell / Bash).
2. Arahkan direktori aktif terminal ke *root folder* repositori ini (folder yang sejajar dengan file `go.mod`).
3. Jalankan perintah kompilasi berikut sesuai dengan Sistem Operasi Anda:

**Pengguna Windows:**
```bash
go build -o bin/voxelizer.exe ./src
```

**Pengguna Linux / macOS:**
```bash
go build -o bin/voxelizer ./src
```

## Cara Menjalankan Program
Program dijalankan melalui antarmuka baris perintah (CLI) dengan menyertakan dua argumen wajib:
<path_file.obj>: Jalur relatif atau absolut menuju file input 3D.
<max_depth>: Angka integer positif yang menentukan tingkat kedalaman maksimum pohon Octree (semakin tinggi nilainya, semakin kecil ukuran voxel dan semakin detail hasilnya).
Format Perintah:
```bash
# Untuk Windows
.\bin\voxelizer.exe <path_file.obj> <max_depth>

# Untuk Linux / macOS
./bin/voxelizer <path_file.obj> <max_depth>
```
Contoh Eksekusi:
```Bash
.\bin\voxelizer.exe test\pumpkin.obj 4
```

## Format Output
Setelah program selesai memproses data, program akan menghasilkan dua bentuk keluaran:
1. File Output .obj Program akan secara otomatis membuat file 3D baru yang berisi kumpulan voxel. File ini akan disimpan di dalam direktori yang sama dengan file input, dengan penambahan sufiks -voxelized pada namanya.
(Contoh: Jika input adalah test/pumpkin.obj, maka outputnya adalah test/pumpkin-voxelized.obj).

2. Laporan Statistik CLI Di layar terminal, program akan menampilkan laporan analitik proses, meliputi:
    - Total voxel, vertex, dan faces yang berhasil dibentuk.
    - Distribusi jumlah node yang terbentuk pada setiap level depth.
    - Distribusi jumlah node yang berhasil di-prune (diabaikan) pada setiap level depth.
    - Kedalaman aktual Octree.
    - Waktu eksekusi murni algoritma program (dalam milidetik/detik).
    - Jalur tempat file output disimpan.

## Author
13524076 - Fayyaz Akmal Lauda
