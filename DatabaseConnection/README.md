# Struktur Project
> ├── config                                                               
> │   └── config.go                                                        
> ├── mahasiswa                                                            
> │   └── repository_mysql.go                                              
> ├── main.go                                                              
> ├── models                                                               
> │   └── mahasiswa.go                                                     
> └── utils                                                                
> └── res.go  

# Penjelasan:
- **main.go**, digunakan untuk menjalankan melakukan aksi terhadap data, bisa dikatakan sebuath controller
- **utils/res.go**, digunakan untuk mencetak data dengan format JSON
- **config/config.go**, digunakan untuk melakukan konfigurasi MySQL
- **mahasiswa/repository_mysql.go**, digunakan untuk melakukan query ke database
- **models/mahasiswa.go**, digunakan untuk membuat struct. struktur