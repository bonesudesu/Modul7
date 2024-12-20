nomer 1 :
Struktur Data
- Titik
Struktur data untuk merepresentasikan sebuah titik dengan koordinat (x, y):
type Titik struct {
    x, y int
}
- Lingkaran
Struktur data untuk merepresentasikan sebuah lingkaran dengan:

titikPusat: Sebuah titik sebagai pusat lingkaran.
radius: Jari-jari lingkaran.
type Lingkaran struct {
    titikPusat Titik
    radius     int
}


Fungsi
1.jarak Fungsi untuk menghitung jarak Euclidean antara dua titik (p, q):
func jarak(p, q Titik) float64 {
    return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}
- diDalam Fungsi untuk menentukan apakah suatu titik berada di dalam atau pada tepi sebuah lingkaran:
  func diDalam(c Lingkaran, p Titik) bool {
    return jarak(c.titikPusat, p) <= float64(c.radius)
}
Kondisi:
Jika jarak titik ke pusat lingkaran lebih kecil atau sama dengan jari-jari lingkaran, maka titik tersebut dianggap berada di dalam lingkaran.

Fungsi Utama (main)
- Input:

Pusat dan radius dari dua lingkaran: (cx1, cy1, r1) dan (cx2, cy2, r2).
Koordinat sebuah titik: (px, py).

0 0 5
4 4 3
1 1

- Membuat objek:

Lingkaran 1: Lingkaran{Titik{cx1, cy1}, r1}
Lingkaran 2: Lingkaran{Titik{cx2, cy2}, r2}
Titik: Titik{px, py}

- Logika:

Cek apakah titik berada di dalam:
Kedua lingkaran.
Salah satu lingkaran.
Di luar keduanya.

- Output:

Jika titik berada di kedua lingkaran:
"Titik di dalam lingkaran 1 dan 2"
Jika hanya di lingkaran 1:
"Titik di dalam lingkaran 1"
Jika hanya di lingkaran 2:
"Titik di dalam lingkaran 2"
Jika di luar keduanya:
"Titik di luar lingkaran 1 dan 2"

nomer 2:

1. Struktur Program
Input Awal:
Pengguna diminta memasukkan jumlah elemen array (n) dan elemen-elemen array tersebut.
Array dibuat menggunakan fungsi make() untuk mendukung panjang yang dinamis.
Menu Operasi:
Program berjalan dalam loop for untuk menyediakan menu pilihan. Pengguna dapat melakukan berbagai operasi pada array hingga memilih untuk keluar.

 2. Fitur-Fitur Program
Tampilkan keseluruhan isi array
Tampilkan elemen dengan indeks ganjil
Tampilkan elemen dengan indeks genap
Tampilkan elemen dengan indeks kelipatan x
Hapus elemen pada indeks tertentu
Tampilkan rata-rata
Tampilkan standar deviasi
Tampilkan frekuensi dari suatu bilangan
 Keluar

Nomer 3 :

1. Struktur Program
Variabel yang Digunakan
klubA, klubB: Nama dua klub yang bertanding.
skorA, skorB: Skor masing-masing klub dalam pertandingan.
pemenang: Slice untuk menyimpan daftar nama klub yang memenangkan pertandingan.

Fitur Program
1. Input Nama Klub
2. Input Skor dan Penentuan Pemenang
3. Menampilkan Daftar Pemenang


Nomer 4:

Struktur Program
1. Konstanta dan Tipe Data

NMAX: Konstanta yang menentukan kapasitas maksimum array, yaitu 127.
tabel: Tipe data array yang terdiri dari elemen bertipe rune (karakter Unicode), dengan kapasitas maksimum NMAX.

2. Fungsi-Fungsi

isiArray(t *tabel, n *int): Mengisi array dengan karakter masukan dari pengguna hingga tanda titik . atau hingga kapasitas maksimum tercapai.
cetakArray(t tabel, n int): Mencetak isi array karakter sebagai teks.
balikanArray(t *tabel, n int): Membalikkan urutan karakter dalam array.
palindrom(t tabel, n int): Mengecek apakah teks merupakan palindrom.

3. Fungsi Utama (main)

Membaca teks masukan.
Menampilkan teks asli, teks yang dibalik, dan status apakah teks tersebut palindrom.

Penjelasan Fungsi

1. Fungsi isiArray

Membaca masukan karakter satu per satu menggunakan fmt.Scanf("%c", &ch).
Karakter dimasukkan ke array hingga:
Tanda titik (.) ditemukan, atau
Kapasitas maksimum array tercapai.
Jumlah karakter yang dimasukkan disimpan dalam variabel n.

2. Fungsi cetakArray
Mencetak isi array sebagai teks hingga jumlah elemen n.
Digunakan untuk menampilkan teks asli atau teks yang sudah dibalik.

3.  Fungsi balikanArray
Membalikkan isi array dengan menukar elemen awal dan akhir secara iteratif:
Elemen indeks 0 ditukar dengan indeks n-1.
Elemen indeks 1 ditukar dengan indeks n-2, dan seterusnya hingga separuh array.
Proses ini dilakukan di tempat (in-place), tanpa memerlukan array tambahan.

4. Fungsi palindrom
Membuat salinan array t ke array sementara temp menggunakan copy.
Membalikkan array temp menggunakan balikanArray.
Membandingkan array asli t dengan array temp:
Jika semua elemen sama, teks adalah palindrom.
Jika ada elemen yang berbeda, teks bukan palindrom

Fungsi Utama (main)

1. Membaca Masukan:

Program meminta pengguna untuk memasukkan teks.
Teks dimasukkan ke array tab hingga tanda titik (.) ditemukan.

2. Menampilkan Teks Asli:

Menggunakan cetakArray untuk mencetak teks asli.

3. Membalikkan Teks:

Memanggil balikanArray untuk membalikkan teks.
Menampilkan teks yang sudah dibalik.

4. Mengecek Palindrom:

Memanggil fungsi palindrom untuk mengecek apakah teks adalah palindrom.
Menampilkan hasil pengecekan.

5. 



   





