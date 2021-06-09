# go Best Practice Guides

# Perkenalan
Repositori ini dibuat oleh @rindibudiaramdhan dalam riwayat pembelajarannya mengenai GoLang. Repository ini semua deskripsi dan penjelasannya akan dibuat dalam bahasa Indonesia agar menjadi pembelajaran juga bagi teman-teman Indonesia saya yang kesulitan belajar Go yang literaturnya dan tutorialnya kebanyakan berbahasa asing (mayoritas Bahasa Inggris). Namun beberapa istilah akan tetap menggunakan bahasa Inggris agar pemahaman pembelajarannya sama secara universal.
Semoga bermanfaat.

# 12 Go *Best Practices*
Francesc Campoy Flores
Gopher dari Google

## *Best Practices*
dari Wikipedia:
> "Praktik terbaik adalah metode atau teknik yang secara konsisten menunjukkan hasil yang lebih unggul daripada yang dicapai dengan cara lain."

Teknik-teknik coding dalam Go antara lain
- Simpel
- Mudah Dibaca
- Mudah di-*maintenance*

## Contoh Code Kotor
```
type Gopher struct {
  Name string
  AgeYears int
}
```
```
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
  err = binary.Write(w, binary.LittleEndian, int32(g.Name))
  if err == nil {
    size += 4
    var n int
    n, err = w.Write([]byte(g.Name))
    size += int64(n)
    if err == nil {
      err = binary.Write(w, binary.LittleEndianm int64(g.AgeYears))
      if err == nil {
        size += 4
      }
      return
    }
    return
  }
  return
}
```

## 1. Bagaimana Menghindari *nesting* dengan *handling error* terlebih dahulu
```
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
 err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
 if err != nil {
   return
 }
 size += 4
 n, err := w.Write([]byte(g.Name))
 size += int64(n)
 if err != nil {
   return
 }
 err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
 if err = nil {
   size += 4
 }
 return
}
```
semakin sedikit nesting, maka semakin sedikit beban kognitif yang harus kita baca

## 2. Hindari Pengulangan Code Semaksimal Mungkin
Terapkan jenis utilitas satu kali untuk kode yang lebih sederhana

## 3. Code yang paling penting yang didahulukan
...
## 4. Dokumentasikan Code-mu
...
## 5. Lebih Pendek lebih baik
...
## 6. Gunakan Packages dengan multiple Files
...
## 7. Buatlah Seluruh Package-mu "go get"-able
...
## 8. Tanyakan apa yang sebenarnya kamu butuhkan
...
## 9. Jagalah independent package tetap independent
...
## 10. Hindari concurrency di API
...
## 11. Gunakan goroutines untuk me-manage state
...
## 12. Hindari goroutine leaks
...

## Some links
 Resources
- Go homepage http://golang.org
- Go interactive tour http://tour.golang.org
Other talks
- Lexical scanning with Go [video](http://www.youtube.com/watch?v=HxaD_trXwRE)
- Concurrency is not parallelism [video](http://vimeo.com/49718712)
- Go concurrency patterns [video](http://www.youtube.com/watch?v=f6kdp27TYZs)
- Advanced Go concurrency patterns [video](http://www.youtube.com/watch?v=QDDwwePbDtw)







