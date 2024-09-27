**Hands-on Week 1**

**Muhammad Grandiv Lava Putra**

**22/493242/TK/54023**

**Penjelasan Alur Kode**

Program di dalam folder 'week1' berfungsi untuk mengonversi suhu dari satuan Celcius menjadi 3 satuan lainnya: Fahrenheit, Kelvin, dan Reamur. 
Struktur `switch-case` digunakan untuk memudahkan user memilih konversi yang ingin dilakukan.

1. Program akan meminta pengguna untuk memasukkan suhu dalam satuan Celcius yang kemudian akan disimpan ke dalam variabel `celcius` dengan tipe data `float64` supaya dapat menerima angka desimal.

2. Program menampilkan tiga pilih untuk mengonversi suhu tersebut ke dalam satuan lain dan dengan memasukkan pilihan (1, 2, atau 3), pilihan tersebut akan disimpan ke variabel `choice` bertipe `int`

3. `switch-case` digunakan untuk menentukan satuan konversi berdasarkan pilihan pengguna. Jika pilihan 1 maka suhu Celcius dikonversi ke Fahrenheit dengan rumus `(9*c/5)+32` dan hasilnya diformat ke Fahrenheit.
Jika pilihan 2 maka suhu Celcius dikonversi menjadi Kelvin dengan rumus `C+273.15` dan diformat ke Kelvin. Jika pilihan 3 maka suhu Celcius dikonversi ke Reamur dengan rumus `4/5*C` dan diformat ke Reamur.

4. Jika pengguna memasukkan pilihan yang tidak valid, program akan menampilkan pesan bahwa pilihan tersebut tidak valid.

